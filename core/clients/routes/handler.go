package routes

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/util"
	"Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/logs"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/tools"

	"log"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

//this will properly handle the incoming connection
//performs all the valid checks without issues happening
//this will make sure its not ignored without errors happening
func NewHandlerFunction(conn *ssh.ServerConn, channel ssh.Channel, created time.Time) error {

	//redeems the plan properly and safely
	//this will ensure its done without errors
	isRedeem, _ := tools.NeedleHaystack(toml.ConfigurationToml.AppSettings.Redeem, conn.User())

	//checks for the action type
	//this will ensure its done without errors
	if json.ConfigSettings.Masters.Server.DynamicAuth && isRedeem {
		//trigger the redeem route properly
		//this will ensure its done without issues
		return RedeemRoute(channel, conn) //executes redeem
	}

	//writes to the channel properly without issues
	//this will ensure its done without issues happening
	if _, err := channel.Write([]byte("\033]0;"+strings.ReplaceAll(toml.ConfigurationToml.Login.Title, "<<$cnc>>", toml.ConfigurationToml.AppSettings.AppName)+"\007")); err != nil {
		return err //returns the error correctly
	}
	//stores the future user information
	//this will allow for better control without issues
	var usernameFromLogin string = "" //stored in type string properly
	if !json.ConfigSettings.Masters.Server.DynamicAuth {
		usernameFromLogin = conn.User() //sets the static username properly
	} else {
		//renders the header properly
		//this will ensure its done without any errors happening on reqeust
		usersFromReq, err := LoginRequest(channel, conn) //properly tries to work
		//error handles the login reqeust information
		//this will ensure its done without any errors happening
		if err != nil || usersFromReq == nil { //error handles the statement properly without issues
			channel.Close() //closes the channel properly
			return err //returns the error
		}

		//tries to properly insert the login reqeust from the remote conn
		//this will ensure its done without any errors happening on reqeust
		database.Conn.LoginAttempt(conn.RemoteAddr().String(), string(conn.ClientVersion()), usersFromReq.Username, err == nil)

		//prints to the terminal about the login properly
		//this will ensure its done without any errors happening
		log.Printf("[SSH SERVER] NEW SSH CONNECTION (%s) (%s) (%s)\r\n", conn.RemoteAddr().String(), usersFromReq.Username, string(conn.ClientVersion()))
		
		//tries to correctly write the log into the file
		//this will ensure its done without any errors happening
		if err := logs.WriteLog(filepath.Join(deployment.Assets, "logs", "connections.json"), logs.ConnectionLog{Type: "ssh", Address: conn.RemoteAddr().String(), Username: usersFromReq.Username, Time: time.Now()}); err != nil {
			log.Printf("logging fault: %s\r\n", err.Error()) //alerts the main terminal properly
		}

		//sets the username properly without issues
		//this will ensure its complete without errors
		usernameFromLogin = usersFromReq.Username
	}

	//tries to grab the username from the database
	//this will ensure its done correctly without issues happening
	user, err := database.Conn.FindUser(usernameFromLogin)
	if err != nil { //correctly error handles the statement
		return err //returns the error correctly
	}


	//stores the default path without issues
	//this will ensure its done without any errors
	themePath, colours := make([]string, 0), toml.DecorationToml.Gradient.Colours //sets the default
	//adds support for theme swapping
	//this will try to properly render without issues
	if user.Theme != "default" { //checks for abnormal theme defaults
		//tries to find the theme properly
		//this will ensure its done without any errors
		thm := toml.ThemeConfig.Theme[user.Theme]
		if thm == nil { //checks if theme was found properly
			log.Printf("[UserTheme] [invalid theme has been found] [%s]\r\n", user.Username)
		} else { //else tries to resolve theme without issues
			//this will ensure its properly found without errors
			themePath = strings.Split(thm.Branding, "/") //updates properly
			colours = thm.Decor.Colours
		}
	}

	//starts the ranking system
	//this will ensure its done without any errors
	sys := ranks.MakeRank(user.Username) //starts
	if err := sys.SyncWithString(user.Ranks); err != nil {
		return err //returns the error correctly and properly
	}

	//stores the future address properly and safely
	//this will ensure its done without any errors happening
	var address string = strings.Join(strings.Split(conn.RemoteAddr().String(), ":")[:len(strings.Split(conn.RemoteAddr().String(), ":"))-1], ":") //stores the remote host
	if sys.CanAccessArray(toml.IPRewriteToml.Ranks) { //checks for rewrite
		address = toml.IPRewriteToml.Rewritten //sets rewrite
	}

	//runs the util tracer
	//this will get the matrix for the user
	Matrix, err := util.ParentTracer(user.Parent, make([]int, 0))
	if err != nil { //error handles properly
		return err //error handles
	}

	r := ranks.MakeRank(user.Username)
	//tries to sync the ranks
	//this will allow for proper handling
	if err := r.SyncWithString(user.Ranks); err != nil {
		return err //returns the error properly
	}
	//makes the session correctly and properly
	//this will make sure its done correctly without errors
	session := sessions.MakeCorrectiveSession(user, conn, channel, created.Unix(), themePath, colours, address, Matrix, *r)

	if session.User.Locked { //checks for locked accounts properly and safely
		return language.ExecuteLanguage([]string{"sessions", "account_locked.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	}

	//checks for banned user properly
	//this will ensure its done without any errors
	if r.CanAccess("banned") { //checks for banned user
		delete(sessions.Sessions, session.ID)
		return language.ExecuteLanguage([]string{"views", "banned", "user-banned.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	}

	//checks if the plan has expired
	//this will try to properly handle without issues
	if user.Expiry < time.Now().Unix() && !sys.CanAccessArray(toml.ConfigurationToml.AppSettings.ByPlan) { //returns the function properly
		delete(sessions.Sessions, session.ID) //deletes the session properly without issues happening
		if strv, err := language.MakeTermFX([]string{"views", "expired", "title.tfx"}, session); err != nil { //error
			session.Write("\033]0;You have been banned!\007") //writes to the session properly
		} else { //renders the default value properly without issues happening on reqeust
			session.Write("\033]0;"+strv+"\007") //writes to the session properly
		}
		return language.ExecuteLanguage([]string{"views", "expired", "plan_expired.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	}
	
	//tries to properly remove the session
	//this will perform the catpcha seq without issues
	if err := CaptchaRoute(session); err != nil { //errors
		delete(sessions.Sessions, session.ID) //deletes
		return err //returns the error properly
	}

	var forced bool = false
	//checks for mfa being wanted on the user in forced
	//this will ensure its done without errors happening on purposr
	if len(user.MFA_secret) <= 0 && toml.CatpchaToml.Mfa.Forced && !r.CanAccessArray(toml.CatpchaToml.Mfa.BypassForced) || user.MFA_secret == "FORCE" {
		if err := WantMFA(session); err != nil { //err handles properly
			delete(sessions.Sessions, session.ID) //deletes
			return err //returns the error properly
		}

		forced = true //forced system here
	}

	//mfa verfiy checker
	//only runs if mfa is enabled
	if len(user.MFA_secret) > 0 && !forced{

		//verifys the maps properly
		//ensures its done without errors
		if err := VerifyMFA(*session); err != nil {
			delete(sessions.Sessions, session.ID) //deletes
			return err //returns the error properly
		}
	}

	//checks if the user has new user enabled
	//this will enforce them to change there password on login without errors
	if user.NewUser { //checks if the user is new user properly without issue
		//error handles the request without issues
		//this will ensure its done without any errors
		if err := NewLogin(session); err != nil { //error handles
			delete(sessions.Sessions, created.Unix()) //deletes properly
			session.Channel.Close() //closes the channel properly without issues
			return err //returns the error correctly and properly without errors
		}
	}

	//checks if there is more sessions open then allowed
	//this will stop any other sessions opened properly without issues
	if len(session.CountOpenSessions()) - 1 >= user.MaxSessions { //checks for more
		//prints/writers the banner to the main system without issues happening on purpose
		//this iwll make sure the viewer knows what has happened with the system properly without issues
		if err := language.ExecuteLanguage([]string{"sessions", "max_sessions_reached.itl"}, session.Channel, deployment.Engine, session, make(map[string]string)); err != nil {
			return err //returns the error properly and safely
		}
		//closes the channel once that has rendered properly
		//this will ensure its done without issues happening on purpose
		return channel.Close() //closes the channel properly without issues 
	}

	//enables the newprompt session
	//this will ensure its correctly and properly started without issues
	return NewPrompt(session)
}