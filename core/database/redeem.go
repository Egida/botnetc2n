package database

import (
	"Nosviak2/core/sources/tools"
	"encoding/json"
)

//stores the information properly
//this will allow for better control
type Redeem struct { //stored in structure
	//stores the target string
	Token string  //stored in string
	//stores the object properly
	//this will allow for better control
	Bundle string  
}

//runs the redeem function properly
//this will ensure its done without any errors
func (d *Connection) Redeem(src string) (*TargetRedeem, error) {
	src = tools.SanatizeTool(src)

	//querys the database for the redeem properly
	//this will allow for better control without issues
	row := d.Conn.QueryRow("SELECT `bundle` FROM `tokens` WHERE `token` = ?", src)
	if row.Err() != nil { //error handles the statement properly
		return nil, row.Err() //returns the error properly
	}

	var dst string
	//stores the dst properly
	//stores the string properly without issues
	if err := row.Scan(&dst); err != nil {
		return nil, err //returns the error properly
	}
	var promptly TargetRedeem //redeem storage
	//tries to properly parse the redeem statement
	//this will try to enforce the rules without issues
	if err := json.Unmarshal([]byte(dst), &promptly); err != nil {
		return nil, err //returns the error properly
	}
	//stores the structure properly
	//this will ensure its done without any errors
	return &promptly, nil
}

const (
	//stores the number for user redeem
	//this will allow me to add different redeem statements
	RedeemUser int = 1
)

//stores our target from the redeem
//this will be exported from the statement properly
type TargetRedeem struct { //stored in type structure
	//stores the statement type properly
	//this will store the correct type properly
	// # 1 - user account creation
	Type int `json:"type"`
	//stores the user constructure properly
	//this will ensure its done without any errors
	User *User //leaves the username, password & ranks blank
}


//allows the redeem statement properly
//this will ensure its done without any errors
func (d *Connection) MakeRedeem(tgt *TargetRedeem) (string, error) {
	//converts the value into string properly
	//this will try to properly handle without issues
	bytVal, err := json.Marshal(tgt) //performs the statement
	if err != nil { //performs the system properly without issues
		return "", err //returns the error properly
	}

	//generates the redeem token
	//this will be used properly to store the information
	token := tools.CreateStrongPassword(10) + "-" + tools.CreateStrongPassword(10) + "-" + tools.CreateStrongPassword(10)//creates the perfection


	//tries to insert properly
	//this will ensure its done without errors
	//this will return the token and the information properly
	_, err = d.Conn.Exec("INSERT INTO `tokens` (`token`, `bundle`) VALUES (?, ?)", string(token), string(bytVal))
	return token, err
}

//redeems the token properly
//this will ensure its done without any errors
func (d *Connection) RemoveRedeem(rdm string) error {
	rdm = tools.SanatizeTool(rdm)
	//returns the error properly without issues
	_, err := d.Conn.Exec("DELETE FROM `tokens` WHERE `token` = ?", HashProduct(rdm))
	return err
}

type SystemTokens struct {
	Token  string
	Bundle TargetRedeem
}

//grabs all possible tokens properly and safely
//this will allow us to capture all possible tokens
func (d *Connection) GrabTokens() ([]SystemTokens, error) {

	//runs the query properly and safely
	//this will ensure its done without issues
	conRow, err := d.Conn.Query("SELECT `token`, `bundle` from `tokens`")
	if err != nil { //err handles
		return make([]SystemTokens, 0), err
	}

	defer conRow.Close()
	//stores the array properly
	tokens := make([]SystemTokens, 0)

	for conRow.Next() {
		system, token := "", "" //scans the row properly
		if err := conRow.Scan(&token, &system); err != nil {
			return make([]SystemTokens, 0), err
		}

		var pure TargetRedeem //takes the system properly
		if err := json.Unmarshal([]byte(system), &pure); err != nil {
			return make([]SystemTokens, 0), err
		}

		//saves into the system properly and safley
		//ensures its done without issues happening
		tokens = append(tokens, SystemTokens{Token: token, Bundle: pure})
	}

	return tokens, nil
}