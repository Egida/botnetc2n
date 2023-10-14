package database

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/logs"
	"Nosviak2/core/sources/tools"
	"log"
	"path/filepath"
	"strings"
	"time"
)

//stores information about the attack
//this will allow for features like ongoing etc
type AttackLinear struct { //stored inside the structure
	Method, Target, Username string //stores the method, target and username used
	ID, Duration, Port int //stores the duration and the port properly
	Created, Finish int64 //stores the created and finish field properly
	SentViaAPI bool
}

//correctly inserts the attack information without issues
//this will ensure its done without errors happening on request
func (c *Connection) PushAttack(attk *AttackLinear) error { //inserts into the database properly without issues happening
	//tries to correctly write the log into the file
	//this will ensure its done without any errors happening
	if err := logs.WriteLog(filepath.Join(deployment.Assets, "logs", "attacks.json"), logs.AttackLog{Target: attk.Target, Duration: attk.Duration, Port: attk.Port, Method: attk.Method, Created: attk.Created, End: attk.Finish, Username: attk.Username}); err != nil {
		log.Printf("logging fault: %s\r\n", err.Error()) //alerts the main terminal properly
	}
	
	//starts the query without issues happening
	//this will follow from the logging reqeust without issues
	_, err := c.Conn.Exec("INSERT INTO `attacks` (`id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?, ?)", attk.Target, attk.Username, attk.Duration, attk.Port, attk.Created, attk.Finish, attk.SentViaAPI, attk.Method)
	return err
}


//gets a users all ongoing attacks properly
//this will ensure we can display them in different areas
func (c *Connection) Attacking(user string) ([]AttackLinear, error) {
	user = tools.SanatizeTool(user)
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE `username` = ? AND `Finish` > ?", user, time.Now().Unix())
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}

//gets all the ongoing attacks properly
//this will ensure its done properly without errors
func (c *Connection) GlobalRunning() ([]AttackLinear, error) {
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE `Finish` > ?", time.Now().Unix())
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}

//gets a users all ongoing attacks properly
//this will ensure we can display them in different areas
func (c *Connection) AttackingWithMethod(user string, method string) ([]AttackLinear, error) {
	user = tools.SanatizeTool(user); method = tools.SanatizeTool(method)
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE `username` = ? AND `Finish` > ? AND `method` = ?", user, time.Now().Unix(), method)
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}

//gets a users all ongoing attacks properly
//this will ensure we can display them in different areas
func (c *Connection) AttackingTarget(target string) ([]AttackLinear, error) {
	target = tools.SanatizeTool(target)
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE `Finish` > ? AND `target` = ?", time.Now().Unix(), target)
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}

//gets a users all ongoing attacks properly
//this will ensure we can display them in different areas
func (c *Connection) UserSent(user string) ([]AttackLinear, error) {
	user = tools.SanatizeTool(user)
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE `username` = ?", user)
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}

//gets a users all attacks properly
//this will ensure we can display them in different areas
func (c *Connection) GlobalSent() ([]AttackLinear, error) {
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks`")
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}

//get the most used method from inside the database
//this will allow you to access the most used method properly
func (C *Connection) MostUsedMethod() (string, int, error) {
	//performs the query properly
	//this will ensure its done without any errors
	qrow, err := C.Conn.Query("SELECT `method`, count(*) as theCount FROM `attacks` AND GROUP BY `method` ORDER BY theCount DESC")
	if err != nil { //err handles the request properly
		return "", 0, err //returns the error properly
	}

	defer qrow.Close()
	//stores the information properly
	//this will be used to scan into properly
	method, used := "[EOF]", 0 //stores the data
	//scans the query output properly
	//this will produce the query output
	if err := qrow.Scan(&method, &used); err != nil {
		return "", 0, err //returns the error
	}

	//returns the values properly
	//this will ensure its done without any errors
	return method, used, nil
}

//gets a methods sent properly
//this will ensure we can display them in different areas
func (c *Connection) MethodSent(method string) ([]AttackLinear, error) {
	method = tools.SanatizeTool(method)
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE `method` = ?", method)
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		Attacking = append(Attacking, AttackLinear{})
		err := rows.Scan(&Attacking[len(Attacking)-1].ID, &Attacking[len(Attacking)-1].Target, &Attacking[len(Attacking)-1].Username, &Attacking[len(Attacking)-1].Duration, &Attacking[len(Attacking)-1].Port, &Attacking[len(Attacking)-1].Created, &Attacking[len(Attacking)-1].Finish, &Attacking[len(Attacking)-1].SentViaAPI, &Attacking[len(Attacking)-1].Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}


// GrabMultiplyRunning will get all running attacks with them methods
func (c *Connection) GrabMultiplyRunning(methods []string) ([]AttackLinear, error) {
	//correctly querys the sql database
	//this will ensure its done properly
	rows, err := c.Conn.Query("SELECT `id`, `target`, `username`, `duration`, `port`, `created`, `finish`, `api`, `method` FROM `attacks` WHERE "+"`method` = '"+ strings.Join(methods, "' OR `method` = '")+"'")
	if err != nil { //correctly error handles the database properly
		return make([]AttackLinear, 0), err //returns the error properly
	}

	defer rows.Close()
	//stores all the attack requests we will store
	//this will ensure its done properly without issues
	var Attacking []AttackLinear = make([]AttackLinear, 0) //stored in array

	//for loops through all properly
	//this will ensure its done correctly
	for rows.Next() { //loops through each row
		var line AttackLinear = AttackLinear{}
		err := rows.Scan(&line.ID, &line.Target, &line.Username, &line.Duration, &line.Port, &line.Created, &line.Finish, &line.SentViaAPI, &line.Method)
		if err != nil { //error handling properly without issues
			return make([]AttackLinear, 0), err //stores inside array
		}

		if line.Finish < time.Now().Unix() {
			continue
		}

		Attacking = append(Attacking, line)
	}

	//returns the array of structs properly
	//this will ensure its done correctly and properly
	return Attacking, nil
}