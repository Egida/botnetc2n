package attacks

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/toml"
	"net"
	"strconv"
	"strings"
)

//performs the buffer callback for the attacks safely
//this will ensure its done without issues happening on purpose
//BufferCallback allows for us to parse the incoming transcript and post messages
func BufferCallback(session *sessions.Session, buf string, cleanup bool) (bool, error) {

	//invalid method detection here properly
	//this will allow for live buffer invalid attack command detection
	if QueryMethod(strings.Replace(strings.Split(buf, " ")[0], toml.AttacksToml.Attacks.Prefix, "", 1)) == nil {
		return true,language.ExecuteLanguage([]string{"attacks", "alerts", "invalid_method_alert.itl"}, session.Channel, deployment.Engine, session, map[string]string{"method":strings.Split(buf, " ")[0]})
	} else if cleanup { //cleans the method cleanup system properly
		session.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[0m\x1b8"); return false, nil //remove possible invalid method ,essage
	}

	//valid ip address has been tracked here
	if net.ParseIP(strings.Split(buf, " ")[1]) != nil && buf == strings.Join(strings.Split(buf, " ")[:2], " ") {

		//runs the api request
		//this will ensure its looked up
		go func() {
			out, err := Lookup(strings.Split(buf, " ")[1])
			if err != nil || out == nil {
				return //ends routine
			}

			//tries to suggest the method properly
			entry, suggested := MatchSuggestion(out.Isp, out.As, strings.Replace(strings.Split(buf, " ")[0], toml.AttacksToml.Attacks.Prefix, "", -1))
			if suggested == nil { //validates
				session.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[0m\x1b8"); return
			}

			//ignores the system suggestion properly
			if session.CanAccessArray(suggested.AccessBypass) {
				return //ends routine
			}

			if !cleanup { //suggests the method properly and safely
				language.ExecuteLanguage([]string{"attacks", "alerts", "method-suggestion.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":strings.Split(buf, " ")[1], "method":entry}); return
			}
		}()
	}

	
	//checks if the user can bypass blacklist
	//ensures we dont display messages about blacklist if so
	if !session.CanAccess("bypass-bl") && buf == strings.Join(strings.Split(buf, " ")[:2], " ") { //can access rank properly
		if blacklisted, err := CheckBlacklist(strings.Split(buf, " ")[1]); blacklisted || err != nil {
			return true,language.ExecuteLanguage([]string{"attacks", "alerts", "error_target_blacklisted.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":strings.Split(buf, " ")[1]})
		} else if cleanup { //detects non blacklisted target if cleanup mode is enabled
			session.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[0m\x1b8"); return false, nil //remove possible blacklisted message
		}
	}


	//checks if the user is trying to attack themself properly
	//this will ensure its not ignored properly and its enforced properly
	if strings.ToLower(strings.Split(buf, " ")[1]) == strings.Split(session.Conn.RemoteAddr().String(), ":")[0] {
		return true, language.ExecuteLanguage([]string{"attacks", "alerts", "warning_self.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":strings.Split(buf, " ")[1]})
	} else { //checks for no self ip and cleanup mode being enabled
		session.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[0m\x1b8"); return false, nil //removes self message
	}

	//checks the length properly
	//this will ensure its done properly
	if len(strings.Split(buf, " ")) < 3 {
		return false, nil
	}

	//tries to validate the length properly
	//this will ensure its done safely and properly
	if len(strings.Split(buf, " ")[2]) > 0 { //leng
		//tries to convert the duration into an int
		//this will ensure its done properly and safely
		duration, err := strconv.Atoi(strings.Split(buf, " ")[2])
		if err != nil || duration == 0 || duration > session.User.MaxTime { //checks within the guidelines properly
			return true, language.ExecuteLanguage([]string{"attacks", "alerts", "error_invalid_duration.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":strings.Split(buf, " ")[1], "duration":strings.Split(buf, " ")[2]})
		} else if cleanup { //checks the cleanup properly and safely
			session.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K\x1b[0m\x1b8"); return false, nil //removes self message
		}
	}

	//this will ensure its done without issues
	//allows for proper control without issues happening
	return false, nil //returns no values properly and safely
}