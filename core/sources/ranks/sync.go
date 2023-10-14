package ranks

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

//syncs the ranks into the structure properly
//this will make sure its done without errors happening
func (r *Ranks) SyncWithString(usr string) error {
	if len(usr) <= 0 { //checks if there is anything inside the string
		return nil //returns nothing as the string still matters but ignores it properly
	}
	//properly decodes the string into byte base64
	//this will ensure its done without issues happening
	database, err := base64.RawStdEncoding.DecodeString(usr)
	if err != nil { //error handles the statement properly
		return err //returns the error correctly
	}

	var context UserRanked //stores the structure properly
	//correctly tries to decode without issues
	//this will ensure its without errors happening
	if err := json.Unmarshal(database, &context); err != nil {
		return err //returns the error properly without issues
	}

	//checks if the usernames match properly
	//this will ensure its properly done without errors
	if r.username != context.Username { //returns the error properly
		return errors.New("username doesnt match with one inside structure")
	}

	//saves into the array correctly
	//this will make sure its done without issues happening
	r.ranks = append(r.ranks, context.Ranks...); return nil
}