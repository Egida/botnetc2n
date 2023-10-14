package database

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/layouts/json"
)

//stores the SQL connection information
//this will allow for proper handling without issues
type Connection struct { //stored inside a structure
	//stores the target host properly
	//this will allow for proper handling without issues
	Conn *sql.DB
	//stores the authentication string used properly
	//this will be used if we need to recreate any connections
	AuthString string
	//stores the time the connection was made
	//this will allow for better handling without issues
	Time time.Time
}

//stores the information properly
//this will literally properly connect
var Conn *Connection = nil

//pushes the connection towards the database
//this will allow for better handling without issues
func MakeConnection() error { //returns error if one happens
	//gets the auth string properly
	//this will store if for future refs
	conn := makeString(json.ConfigSettings)
	
	//tries to connect to the database properly
	//this will ensure its been done properly without issues
	Connect, err := sql.Open("mysql", conn) //opens the conn
	if err != nil { //error handles the new connection
		return err //returns the error
	}

	//tries to ping the remote connection
	//this will ensure its done properly without issues
	if err := Connect.Ping(); err != nil {
		return err //returns the error
	}

	//creates the structure properly
	//this will store all information for future refs
	Conn = &Connection{
		Conn: Connect, //sets the conn properly
		AuthString: conn, //sets the authstring properly
		Time: time.Now(), //defaults the time properly
	}; return nil //returns nil as nothing went wrong
}

//creates the auth string properly
//this will take the details without issues happening
func makeString(s *models.ConfigurationJson) string { //returns the formatted string properly
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", s.Database.Username, s.Database.Password, s.Database.Host, s.Database.Name)
}

//properly hashes the reactant for comparing
//this will be used in areas like password handling
func HashProduct(reactant string) (string) {
	//encodes the string properly and safely
	//this will ensure its done properly and safely
	return hex.EncodeToString(sha256.New().Sum([]byte(reactant)))
}