package database

import (
	"Nosviak2/core/sources/tools"
	"time"
)

//stores the login information
//this will help with management properly and better handling
type Login struct {
	//stores the remote address used for the login
	//this ensures its done properly without issues happening
	Address string //stores the IPv4 address

	//stores the date which address the date for the login
	//allows the user to view when the login attempt was properly
	TimeStore int64 //stored in unix format

	//stores the SSHClient banner which was used
	//this will help with narrowing the request down properly
	Banner string //stores the sshBanner used

	//stores the username who made the login attempt
	//this will help with better system handling without issues
	Username string //stores the usernames

	//stores if the login attempt was complete
	//this will store if the login reqeust gained access
	Success bool
}

//tries to log the attempt into the database
//this ensures its done properly without issues happening
func (c *Connection) LoginAttempt(address, banner, user string, success bool) error {
	user = tools.SanatizeTool(user)
	banner = tools.SanatizeTool(banner)
	//tries to query input the information
	//this will ensure its done properly without errors happening
	_, row := c.Conn.Exec("INSERT INTO `logins` (`addressIPv4`, `timeDate`, `sshBanner`, `username`, `status`) VALUES (?, ?, ?, ?, ?)", address, time.Now().Unix(), banner, user, success)
	if row != nil { //error handles the statement properly
		return row //returns the error correctly
	}

	return nil
}


//gets all the login attempts for an account
//this will ensure its done without issues happening
func (c *Connection) GetLogins(who string) ([]Login, error) {
	who = tools.SanatizeTool(who)
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `addressIPv4`, `timeDate`, `sshBanner`, `username`, `status` FROM `logins` WHERE `username` = ?", who)
	if err != nil { //correctly error handles the database properly
		return make([]Login, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the login requests we will store
	//this will ensure its done properly without issues
	var LoginsTillNow []Login = make([]Login, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		LoginsTillNow = append(LoginsTillNow, Login{})
		err := rows.Scan(&LoginsTillNow[len(LoginsTillNow)-1].Address, &LoginsTillNow[len(LoginsTillNow)-1].TimeStore, &LoginsTillNow[len(LoginsTillNow)-1].Banner, &LoginsTillNow[len(LoginsTillNow)-1].Username, &LoginsTillNow[len(LoginsTillNow)-1].Success)
		if err != nil { //error handling properly without issues
			return make([]Login, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return LoginsTillNow, nil
}

//gets all the login attempts for an account
//this will ensure its done without issues happening
func (c *Connection) AllLogins() ([]Login, error) {
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `addressIPv4`, `timeDate`, `sshBanner`, `username`, `status` FROM `logins`")
	if err != nil { //correctly error handles the database properly
		return make([]Login, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the login requests we will store
	//this will ensure its done properly without issues
	var LoginsTillNow []Login = make([]Login, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		LoginsTillNow = append(LoginsTillNow, Login{})
		err := rows.Scan(&LoginsTillNow[len(LoginsTillNow)-1].Address, &LoginsTillNow[len(LoginsTillNow)-1].TimeStore, &LoginsTillNow[len(LoginsTillNow)-1].Banner, &LoginsTillNow[len(LoginsTillNow)-1].Username, &LoginsTillNow[len(LoginsTillNow)-1].Success)
		if err != nil { //error handling properly without issues
			return make([]Login, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return LoginsTillNow, nil
}