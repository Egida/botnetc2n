package interference

import "Nosviak2/core/database"

//checks if the tables exist properly
//this will ensure its done without any errors happening
//this will insert into the database so the user doesnt have to properly
func AppearDatabase() (bool, error) { //checks if the database tables exist properly
	//querys the database properly and safely
	//this will ensure its done without any errors happening
	_, row := database.Conn.Conn.Query("SELECT * FROM `logins`,`users`,`attacks`,`tokens`;")

	//checks if the information is there
	//this will ensure its done without any errors
	if row == nil { //checks properly to see if they exist
		return true, nil //returns the information properly
	}

	//returns that they dont exist properly
	//this will ensure its done without any errors
	return false, nil
}