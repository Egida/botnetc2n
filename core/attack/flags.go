package attacks

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/language"
	"strconv"
	"strings"
	"unicode/utf8"
)





//stores information about the keyvalue parsed
//this will ensure its done without any errors happening
type KeyValue struct { //stored in structure properly and safely
	Value string //stores the value properly
	ID int //stores the keyvalue id properly
	Name string //stores the name properly without issues
	PositionString int //stores the position inside the array
}


//veifys the flags properly without issues
//this will ensure its done without any errors happening
func VerifyInformation(current map[string]*KeyValue, method *Method, session *sessions.Session) (bool, error) {
	//ranges through all the of the keyValues properly
	//this will ensure its done without any errors happening
	for name, information := range method.Options.KeyValues { //ranges through
		//checks if it has already been registered
		//this will ensure its done without any errors happening
		if _, ok := current[name]; ok {continue} //continues if found

		//checks if its forced properly and safely
		//this will ensure its done without any errors happening
		if information.Forced { //checks if there is a path properly
			if session == nil {continue} //ignores the session properly
			if len(information.Views["forced"]) <= 0 { //checks properly
				return false, language.ExecuteLanguage(strings.Split(information.Views["forced"], "/"), session.Channel, deployment.Engine, session, map[string]string{"kv":name})
			} else { //executes the default type properly without issues happening
				return false, session.Write(" missing forced key value is missing inside your statement\r\n") //renders default
			}
		}

		//saves properly into without issues happening
		//this will ensure its done without any errors happening
		current[name] = &KeyValue{Value: information.Default, ID: information.ID, Name: name}
	}; return true, nil
}

//tries to resolve the keyvalue properly
//this will ensure its done without any errors
func ResolveFlag(incoming string, method *Method, s *sessions.Session) (*KeyValue, bool, error) {
	//tries to validate the keyValue properly
	//this will ensure its done without any errors
	if strings.Count(incoming, "=") <= 0 { //checks the syntax properly without issues
		return nil, false, language.ExecuteLanguage([]string{"attacks", "invalid_kv_syntax.itl"}, s.Channel, deployment.Engine, s, map[string]string{"kv":incoming})
	}
	//tries to parse the value properly
	//this will ensure its done without any errors
	key, value := strings.Split(incoming, "=")[0], strings.Split(incoming, "=")[1]
	//tries to find the keyvalue properly
	//this will ensure its done without any errors
	keyHeader := method.Options.KeyValues[key] //tries to fetch properly
	if keyHeader == nil { //checks to see if it was found properly safely
		return nil, false, language.ExecuteLanguage([]string{"attacks", "invalid_kv_flag.itl"}, s.Channel, deployment.Engine, s, map[string]string{"kv":incoming})
	}

	//tries to verify the flag type properly
	//this will ensure its done without any errors properly
	if !VerifyType(value, keyHeader) { //verifys properly safely
		if len(keyHeader.Views["invalid_type"]) <= 0 { //checks properly
			return nil, false, language.ExecuteLanguage(strings.Split(keyHeader.Views["invalid_type"], "/"), s.Channel, deployment.Engine, s, map[string]string{"kv":incoming})
		} else { //executes the default type properly without issues happening
			return nil, false, s.Write(" invalid type given for a keyValue option\r\n") //renders default
		}
	}

	//detects the its type int properly and safely
	//this will ensure its done without any errors happening
	if strings.ToLower(keyHeader.Type) == "int" { //int detection
		value, err := strconv.Atoi(value) //converts to int
		if err != nil { //error handles properly without issues
			if len(keyHeader.Views["invalid_type"]) <= 0 { //checks properly
				return nil, false, language.ExecuteLanguage(strings.Split(keyHeader.Views["invalid_type"], "/"), s.Channel, deployment.Engine, s, map[string]string{"kv":incoming})
			} else { //executes the default type properly without issues happening
				return nil, false, s.Write(" invalid type given for a keyValue option\r\n") //renders default
			}
		}

		//checks the key value size properly
		//this will ensure its done without any errors
		if value > keyHeader.MaxIntSize { //checks the size properly
			if len(keyHeader.Views["too_big"]) <= 0 { //checks properly
				return nil, false, language.ExecuteLanguage(strings.Split(keyHeader.Views["too_big"], "/"), s.Channel, deployment.Engine, s, map[string]string{"kv":incoming})
			} else { //executes the default type properly without issues happening
				return nil, false, s.Write(" value given surpassed the option given properly\r\n") //renders default
			}
		}
	//checks the max length inside the string type properly
	//this will ensure its done without any errors happening on purpose
	} else if strings.ToLower(keyHeader.Type) == "string" { //string detection
		if utf8.RuneCountInString(value) > keyHeader.MaxLength { //over length detection
			if len(keyHeader.Views["too_big"]) <= 0 { //checks properly without issues for the branding
				return nil, false, language.ExecuteLanguage(strings.Split(keyHeader.Views["too_big"], "/"), s.Channel, deployment.Engine, s, map[string]string{"kv":incoming})
			} else { //executes the default type properly without issues happening
				return nil, false, s.Write(" value given surpassed the option given properly\r\n") //renders default
			}
		}
	}

	//returns the values properly and safely
	//this will ensure its done without any errors happening
	return &KeyValue{ID: keyHeader.ID, Value: value, Name: key}, true, nil
}

//tries to verify the keyvalue type properly
//this will ensure its done without any errors happening
func VerifyType(incoming string, flag *models.KeyValue) bool {
	//lowers the flag type properly
	//this will ensure its done without errors
	switch strings.ToLower(flag.Type) {

	//tries to validate that its int
	//this will ensure its done without errors
	case "int": //tries to verify int protocol properly
		if _, err := strconv.Atoi(incoming); err != nil {
			return false //error handles properly
		}; return true //confirms the success properly
	case "string":
		return true
	};return false
}