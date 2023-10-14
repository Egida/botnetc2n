package apis

import (
	"Nosviak2/core/database"
	"Nosviak2/core/sources/events"
	"net/http"
)


type OngoingAttacks struct {

	Success			bool						`json:"success"`
	Error			string						`json:"error"`
	Username		string						`json:"username"`
	Running 		[]database.AttackLinear		`json:"ongoing"`
}

//Ongoing will allow you to view all your running attacks
func Ongoing(w http.ResponseWriter, r *http.Request) {
	

	authenticationMethod := 0 //default is set to 0
	if r.URL.Query().Has("username") && r.URL.Query().Has("password") {
		authenticationMethod = 1 //username & password auth method
		events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{"USER&PASS-AUTH",r.RemoteAddr, r.URL.String(), r.URL.Path})
	} else if r.URL.Query().Has("key") {
		authenticationMethod = 2 //token auth method
		events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{"TOKEN-AUTH",r.RemoteAddr, r.URL.String(), r.URL.Path})
	} else { //the unknown authentication method will be written here
		events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{"UNWANTED-AUTH",r.RemoteAddr, r.URL.String(), r.URL.Path})
		EncodeAndReturn(OngoingAttacks{Success: false, Error: ErrUnwantedAuthMethod.Error()}, w); return
	}

	var user *database.User = nil
	if authenticationMethod == 1 { //username & password based authentication method
		//tries to get the username which was given from the database
		has, err := database.Conn.FindUser(r.URL.Query().Get("username"))
		if err != nil || has == nil { //unknown username has been given or database returned error
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-USER-GIVEN",r.RemoteAddr, r.URL.String(), r.URL.Path})
			EncodeAndReturn(OngoingAttacks{Success: false, Error: ErrUnknownUserGiven.Error()}, w); return
		}

		user = has //updates the option
	} else if authenticationMethod == 2 { //auth via token
		//tries to get the user via there api token
		has, err := database.Conn.GetUserViaToken(r.URL.Query().Get("key"))
		if err != nil || has == nil { //unknown token has been given or database returned error
			events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-TOKEN-GIVEN",r.RemoteAddr, r.URL.String(), r.URL.Path})
			EncodeAndReturn(OngoingAttacks{Success: false, Error: ErrUnknownToken.Error()}, w); return
		}

		user = has //updates the option
	}

	//tries to validate the authentication of the remote system
	if authenticationMethod == 1 && user.Password != database.HashProduct(r.URL.Query().Get("password")) { //password authentication method
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-PASSWORD-GIVEN",r.RemoteAddr, r.URL.String(), r.URL.Path})
		EncodeAndReturn(OngoingAttacks{Success: false, Error: ErrUnknownPassword.Error(), Username: user.Username}, w); return
	} else if authenticationMethod == 2 && user.Token != database.HashProduct(r.URL.Query().Get("key")) {
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"UNKNOWN-TOKEN-GIVEN",r.RemoteAddr, r.URL.String(), r.URL.Path})
		EncodeAndReturn(OngoingAttacks{Success: false, Error: ErrUnknownToken.Error(), Username: user.Username}, w); return
	}

	//Attacking will fetch all the users ongoing attacks
	OngoingAttack, err := database.Conn.Attacking(user.Username)
	if err != nil { //err handles the system
		events.DebugLaunch(events.FunctionWithError, "API", "REQUEST", []string{"DATABASE-ERROR-RETURNED",r.RemoteAddr, r.URL.String(), r.URL.Path})
		EncodeAndReturn(OngoingAttacks{Success: false, Error: ErrUnknownError.Error(), Username: user.Username}, w); return
	}


	events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{"OBJECT-RETURNED",r.RemoteAddr, r.URL.String(), r.URL.Path})
	//writes the commands output from the api properly and safely
	EncodeAndReturn(OngoingAttacks{Success: true, Username: user.Username, Running: OngoingAttack}, w)
}