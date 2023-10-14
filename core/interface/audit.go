package interference

import (
	"Nosviak2/core/database"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"encoding/json"
	"fmt"
	"os"
)

//this will properly run the SQL Audit safely
//this will insert the tables without errors happening
//makes sure that on setup the tables are inserted without issues
func RunTerminalAudit() error { //returns an error if one happened properly
	//tries to insert the users table properly
	//this will ensure its done without any errors happening
	if _, err := database.Conn.Conn.Query(database.Conn.User()); err != nil {
		return err //returns the error if one happened properly
	}

	//tries to insert the attacks table properly
	//this will ensure its done without any errors happening
	if _, err := database.Conn.Conn.Query(database.Conn.Attacks()); err != nil {
		return err //returns the error if one happened properly
	}

	//tries to insert the logins table properly
	//this will ensure its done without any errors happening
	if _, err := database.Conn.Conn.Query(database.Conn.Logins()); err != nil {
		return err //returns the error if one happened properly
	}

	//tries to insert the tokens table properly
	//this will ensure its done without any errors happening
	if _, err := database.Conn.Conn.Query(database.Conn.Tokens()); err != nil {
		return err //returns the error if one happened properly
	}


	//creates the strong password properly
	//this will ensure its done without any errors
	Password := tools.CreateStrongPassword(32) //password filter
	APIKey   := tools.CreateStrongPassword(toml.ApiToml.API.KeyLen) //password filter

	//tries to insert the user properly
	//this will ensure its done without any errors happening
	if _, err := database.Conn.Conn.Query(database.Conn.UserInsert("root", Password, APIKey)); err != nil {
		return err //returns the error properly if one happened
	}

	//renders the information properly for the user
	//this will ensure they know about the audit going correc
	fmt.Printf("\t- user: root\r\n\t- password: %s\r\n\t- api: %s\r\n", Password, APIKey)

	//stores if we want to write into the file
	//this will ensure its done without errors happening
	Object, err := json.MarshalIndent(&Created{Username: "root", Password: Password, APIUsers: APIKey}, "", "\t")
	if err == nil { //renders into the file properly
		newest, err := os.Create("new_default.json")
		if err != nil { //err handles proeprly
			return nil
		}

		newest.Write(Object) //writes properly
	}

	return nil
}

//stores the created username
//this will be formatted into the json file
type Created struct {
	Username string `json:"username"`
	Password string `json:"password"`
	APIUsers string `json:"api"`
}