package ranks

import (
	"encoding/base64"
	"encoding/json"
)

//tries to properly compress into a string
//this will ensure its done without issues happening
func (r *Ranks) MakeString() (string, error) {

	//correctly tries to marshal the structure 
	//this will ensure its done without issues happening
	Values, err := json.Marshal(&UserRanked{Username: r.username, Ranks: r.ranks})
	if err != nil { //properly tries to error handle without issues
		return "", err //returns the error correctly
	}

	return base64.RawStdEncoding.EncodeToString(Values), nil
}