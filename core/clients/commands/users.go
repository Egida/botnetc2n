package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	"Nosviak2/core/clients/views/util"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"fmt"
	"time"

	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
	"golang.org/x/term"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "users",
		Aliases:            []string{"usrs"},
		CommandPermissions: []string{"admin", "moderator", "reseller"},
		CommandDescription: "moderation and management for users",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			// Stores all the user accounts
			var users []database.User = make([]database.User, 0)
			var err error = nil

			// Checks for reseller and hasn't got admin or moderator
			if s.CanAccess("reseller") && !s.CanAccessArray([]string{"admin", "moderator"}) {
				users, err = database.Conn.ParentTracer(s.User.Identity) // shows all the resellers users
			} else {
				users, err = database.Conn.GetUsers() // shows all the users inside the database

			}

			if err != nil { //basic error handling without issues happening on request and returns the error statement
				return language.ExecuteLanguage([]string{"users", "databaseErr.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure

			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-id.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-username.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-maxtime.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-concurrents.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-cooldown.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-ranks.txt").Containing, lexer.Escapes)},
				},
			}

			//ranges through all the users properly
			//this will insert each user into the database without issues happening on request
			for _, usr := range users {
				//creates the rank information
				//this will ensure its done without issues
				r := ranks.MakeRank(usr.Username) //makes the rank
				//tries to correctly sync the render without issues happening
				//this will make sure its done without issues happening on request
				if err := r.SyncWithString(usr.Ranks); err != nil {
					return err //returns the error correctly and properly
				}
				//properly deploys the rank into string
				//this will ensure its done without errors happening
				ranks, err := r.DeployRanks(true) //deploys into string format
				if err != nil {                   //error handles the syntax properly without issues
					return err //returns the error properly
				}

				var user string = strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-username.txt").Containing, lexer.Escapes), "<<$username>>", usr.Username)
				if len(s.GetSessions(usr.Username)) > 0 {
					user = strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-user-online.txt").Containing, lexer.Escapes), "<<$username>>", usr.Username)
				}

				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.Identity))}, //id
					{Align: simpletable.AlignLeft, Text: user}, //username
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-maxtime.txt").Containing, lexer.Escapes), "<<$maxtime>>", HandleTime(usr.MaxTime))},             //maxtime
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-concurrents.txt").Containing, lexer.Escapes), "<<$concurrents>>", HandleTime(usr.Concurrents))}, //conns
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-cooldown.txt").Containing, lexer.Escapes), "<<$cooldown>>", strconv.Itoa(usr.Cooldown))},        //cooldown
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))},                  //ranks
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("users", table, s).TextureTable()
		},

		//stores the invalid subcommand function
		//this will ensure its done without errors happening
		InvalidSubCommand: func(s *sessions.Session, cmd []string) error {
			//tries to cache the rank level properly
			//this will check if the rank exists without issues
			_, ok := ranks.PresetRanks[strings.ToLower(strings.Split(cmd[1], "=")[0])] //tries to find the rank name
			if !ok {                                                                   //properly allows us to try to find the information without issues happening
				return language.ExecuteLanguage([]string{"errors", "subcommand404.itl"}, s.Channel, deployment.Engine, s, map[string]string{"command": cmd[0], "subcommand": cmd[1]})
			}
			//error checks the length without issues happening
			if len(cmd) <= 2 || len(strings.Split(cmd[1], "=")) <= 1 { //checks for the invalid language path

				//filter detection here
				//this will now filter for that rank
				if !strings.Contains(cmd[1], "=") {
					return FilterRank(s, "users", cmd[1:]...)
				}
				return language.ExecuteLanguage([]string{"users", "syntax-ranks.itl"}, s.Channel, deployment.Engine, s, map[string]string{"rank": strings.Split(cmd[1], "=")[0]})
			}

			//checks the permissions properly
			//this will ensure its done without errors
			if len(ranks.PresetRanks[strings.ToLower(strings.Split(cmd[1], "=")[0])].Manage_ranks) > 0 && !s.CanAccessArray(ranks.PresetRanks[strings.ToLower(strings.Split(cmd[1], "=")[0])].Manage_ranks) {
				return language.ExecuteLanguage([]string{"errors", "command403.itl"}, s.Channel, deployment.Engine, s, map[string]string{"command": cmd[0]}) //executes properly
			}

			//tries to parse the boolean option
			//this will ensure its done without errors
			Boolean, err := strconv.ParseBool(strings.Split(cmd[1], "=")[1]) //parses the boolean
			if err != nil {                                                  //tries to correctly error handle without issues happening
				return language.ExecuteLanguage([]string{"users", "syntax-ranks.itl"}, s.Channel, deployment.Engine, s, map[string]string{"rank": strings.Split(cmd[1], "=")[0]})
			}

			//ranges through all the users through args
			//this will ensure each user is given admin permissions
			for proc := 2; proc < len(cmd); proc++ { //loops through the args
				//tries to get the username properly
				//this will ensure its done without errors
				user, err := database.Conn.FindUser(cmd[proc]) //tries to find the user
				if err != nil {                                //error handles the statement without issues happening properly
					if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "user-EOF.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
						s.Write(fmt.Sprintf("%s is an unclassified username which wasnt found inside the database\r\n", cmd[proc]))
						continue //writes to the session properly without issues
					}
					continue //continues the for loop properly without issues
				}

				//this will properly make the rank
				//this will ensure its done without errors happening
				rs := ranks.MakeRank(user.Username) //stores the ranks properly
				rs.SyncWithString(user.Ranks)       //syncs the ranks properly without errors
				//this will properly check if they can access without errors happening on reqeust
				if rs.CanAccess(strings.ToLower(strings.Split(cmd[1], "=")[0])) == Boolean {
					if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), strconv.FormatBool(Boolean) + "-has.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
						s.Write(fmt.Sprintf("the user %s can either already access or not access this rank\r\n", cmd[proc]))
						continue //writes to the session properly without issues
					}
					continue //continues the for loop properly without issues
				}

				//checks to see if they can modify
				//if this statement returns true they cant modify
				if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
					if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
						s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
					}
					continue //continues looping properly without issues
				}

				if Boolean { //checks the boolean properly
					//this will properly try to error handle without issues
					if err := rs.GiveRank(strings.ToLower(strings.Split(cmd[1], "=")[0])); err != nil { //error handles the statement properly and makes sure its safe
						if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "giveFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
							s.Write(fmt.Sprintf("failed to give %s the rank "+strings.Split(cmd[1], "=")[0]+"\r\n", cmd[proc]))
							continue //writes to the session properly without issues
						}
						continue //continues the for loop properly without issues
					}
				} else {
					//this will properly try to error handle without issues
					if err := rs.RemoveRank(strings.ToLower(strings.Split(cmd[1], "=")[0])); err != nil { //error handles the statement properly and makes sure its safe
						if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "takeFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
							s.Write(fmt.Sprintf("failed to take %s's the rank "+strings.Split(cmd[1], "=")[0]+"\r\n", cmd[proc]))
							continue //writes to the session properly without issues
						}
						continue //continues the for loop properly without issues
					}
				}
				//properly tries to make the string without issues
				//this will ensure its done without errors happening
				str, err := rs.MakeString() //makes into string properly
				if err != nil {             //error handles the statement properly
					if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "controlFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
						s.Write(fmt.Sprintf("failed to control %s's rank "+strings.Split(cmd[1], "=")[0]+"\r\n", cmd[proc]))
						continue //writes to the session properly without issues
					}
					continue //continues the for loop properly without issues
				}
				//correctly tries to update the rank
				//this will ensure its done without errors happening
				if err := database.Conn.EditRanks(str, user.Username); err != nil { //error handles the statement
					if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "controlFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
						s.Write(fmt.Sprintf("failed to control %s's rank "+strings.Split(cmd[1], "=")[0]+"\r\n", cmd[proc]))
						continue //writes to the session properly without issues
					}
					continue //continues the for loop properly without issues
				}

				//adds support for live session updating
				//makes sure its done without any errors
				s.FunctionRemote(user.Username, func(t *sessions.Session) {
					t.Ranks = rs                             //syncs with the rs string
					t.User.Ranks = str                       //syncs with both properly
					sessions.Sessions[t.ID].Ranks = rs       //format
					sessions.Sessions[t.ID].User.Ranks = str //string
					//tries to broadcast the correct message properly
					//this will broadcast the message from alerts properly
					err := language.ExecuteLanguage([]string{"alerts", strings.ToLower(strings.Split(cmd[1], "=")[0]) + strconv.FormatBool(Boolean) + ".itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username})
					if err != nil { //error handles properly without issues happening
						//renders the default style properly without issues happening on reqeust
						t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Your " + strings.ToLower(strings.Split(cmd[1], "=")[0]) + " status has been changed to " + strconv.FormatBool(Boolean) + " by " + s.User.Username + "\x1b[0m\x1b8")
					}

					//checks for the close when awarded
					//this will forcefully close session when awarded!
					if Boolean && ranks.PresetRanks[strings.Split(cmd[1], "=")[0]].CloseWhenAwarded {
						t.Channel.Close() //closes the channel properly
						return            //ends the function proeprly
					}
				})

				//correctly executes the language without issues
				//this will ensure its done without issues happening on request
				if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "success-" + strconv.FormatBool(Boolean) + ".itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[proc]}); err != nil {
					s.Write(fmt.Sprintf("correctly edited %s with the rank "+strings.Split(cmd[1], "=")[0], cmd[proc]))
					continue //writes to the session properly without issues
				}
				continue //continues the for loop properly without issues

			}

			//tries to format the option given
			//this will ensure its done without issues happening
			return nil

		},

		//executes most subcommands
		//this will store all the subcommands for user
		SubCommands: []SubCommand{
			{
				SubcommandName:     "list",
				Description:        "list all users",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//tries to correctly get all the users
					//this will ensure its done without issues happening on reqeust
					users, err := database.Conn.GetUsers() //gets the users properly without issues
					if err != nil {                        //basic error handling without issues happening on request and returns the error statement
						return language.ExecuteLanguage([]string{"users", "list", "databaseErr.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//creates the simpletable
					//this will store our information
					table := simpletable.New() //makes the structure

					//sets table header
					//this will be used to define information
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "list", "headers", "header-id.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "list", "headers", "header-username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "list", "headers", "header-maxtime.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "list", "headers", "header-concurrents.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "list", "headers", "header-cooldown.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "list", "headers", "header-ranks.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through all the users properly
					//this will insert each user into the database without issues happening on request
					for _, usr := range users {
						//creates the rank information
						//this will ensure its done without issues
						r := ranks.MakeRank(usr.Username) //makes the rank
						//tries to correctly sync the render without issues happening
						//this will make sure its done without issues happening on request
						if err := r.SyncWithString(usr.Ranks); err != nil {
							return err //returns the error correctly and properly
						}
						//properly deploys the rank into string
						//this will ensure its done without errors happening
						ranks, err := r.DeployRanks(true) //deploys into string format
						if err != nil {                   //error handles the syntax properly without issues
							return err //returns the error properly
						}

						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "list", "values", "value-id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.Identity))},                    //id
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "list", "values", "value-username.txt").Containing, lexer.Escapes), "<<$username>>", usr.Username)},                        //username
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "list", "values", "value-maxtime.txt").Containing, lexer.Escapes), "<<$maxtime>>", strconv.Itoa(usr.MaxTime))},             //maxtime
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "list", "values", "value-concurrents.txt").Containing, lexer.Escapes), "<<$concurrents>>", strconv.Itoa(usr.Concurrents))}, //conns
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "list", "values", "value-cooldown.txt").Containing, lexer.Escapes), "<<$cooldown>>", strconv.Itoa(usr.Cooldown))},          //cooldown
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "list", "values", "value-ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))},                  //ranks
						}

						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}

					//renders the table properly
					//this will ensure its done without issues
					return pager.MakeTable("users", table, s).TextureTable()
				},
			},
			{
				SubcommandName:     "createtoken",
				Description:        "create redeemable tokens",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//stores the order of operations safely
					//this will allow for better handling without issues happening
					var order []string = []string{"maxtime", "cooldown", "concurrents"}

					//stores the values inside a map without issues happening
					//this will make sure its done without errors happening on reqeust
					var args map[string]string = make(map[string]string)

					//ranges through all the args given without issues happening
					//this ensures its done without any errors happening on request
					for pos := 2; pos < len(cmd); pos++ { //loops through all the args
						if pos-2 >= len(order) { //makes sure we dont ignore without issues
							break //breaks from the loop
						}
						//inserts into the map without issues
						//this will make sure its done without errors happening
						args[order[pos-2]] = cmd[pos]
						continue //and continues looping
					}

					//ranges through the order properly
					//this will ensure its done without errors happening on request
					for setting := range order {
						//this will ensure that the option hasnt been made inside map
						//stops items from repeating without issues taking place on purpose etc
						if _, ok := args[order[setting]]; ok { //checks if the value is true or false
							continue //continues the loop properly
						}
						//tries to correctly deploy the prompt without issues happening
						//this will try to stop without issues happening on request without errors
						if err := language.ExecuteLanguage([]string{"users", "createtoken", order[setting] + ".itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
							return err //returns the error correctly and properly
						}
						//prepares and tries to take the input with out issues
						//this will make sure when its executed it doesnt cause issues
						value, err := term.NewTerminal(s.Channel, "").ReadLine() //reads the input
						if err != nil {                                          //properly error handles without issues happening on request
							return err //returns the error properly on purpose
						}

						//saves into the value properly
						//this makes sure its done without issues happening
						args[order[setting]] = value //saves into the socket properly
					}

					//this will properly try to validate the information
					//makes sure the fields are valid without issues happening
					MaxTime, err := strconv.Atoi(args["maxtime"]) //uses the maxtime
					if err != nil {                               //error handles the information without issues happening
						return language.ExecuteLanguage([]string{"users", "createtoken", "maxtime-atoi.itl"}, s.Channel, deployment.Engine, s, map[string]string{"maxtime": args["maxtime"]})
					}

					//this will properly try to validate the information
					//makes sure the fields are valid without issues happening
					Cooldown, err := strconv.Atoi(args["cooldown"]) //uses the cooldown
					if err != nil {                                 //error handles the information without issues happening
						return language.ExecuteLanguage([]string{"users", "createtoken", "cooldown-atoi.itl"}, s.Channel, deployment.Engine, s, map[string]string{"cooldown": args["cooldown"]})
					}

					//this will properly try to validate the information
					//makes sure the fields are valid without issues happening
					Concurrents, err := strconv.Atoi(args["concurrents"]) //uses the concurrents
					if err != nil {                                       //error handles the information without issues happening
						return language.ExecuteLanguage([]string{"users", "createtoken", "concurrents-atoi.itl"}, s.Channel, deployment.Engine, s, map[string]string{"concurrents": args["concurrents"]})
					}

					//tries to properly create the token for the database
					//this will ensure its done without any errors happening
					token, err := database.Conn.MakeRedeem(&database.TargetRedeem{Type: database.RedeemUser, User: &database.User{Theme: "default", MaxTime: MaxTime, Cooldown: Cooldown, Concurrents: Concurrents, MaxSessions: json.ConfigSettings.Masters.Accounts.MaxSessions, Expiry: time.Now().Add((time.Hour * 24) * time.Duration(json.ConfigSettings.Masters.Accounts.DaysExpiry)).Unix(), NewUser: true, Parent: s.User.Identity}})
					if err != nil {
						return language.ExecuteLanguage([]string{"users", "createtoken", "creation-error.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//returns the success message properly
					//this will ensure its completed without errors happening on request
					return language.ExecuteLanguage([]string{"users", "createtoken", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"concurrents": args["concurrents"], "cooldown": args["cooldown"], "maxtime": args["maxtime"], "maxsessions": strconv.Itoa(json.ConfigSettings.Masters.Accounts.MaxSessions), "token": token})
				},
			},

			{
				//this will insert the brand new user into the database
				//allows for better management without errors happening on reqeust
				SubcommandName:     "create",
				Description:        "create users",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//stores the order of operations safely
					//this will allow for better handling without issues happening
					var order []string = []string{"username", "maxtime", "cooldown", "concurrents"}

					//stores the values inside a map without issues happening
					//this will make sure its done without errors happening on reqeust
					var args map[string]string = make(map[string]string)

					//ranges through all the args given without issues happening
					//this ensures its done without any errors happening on request
					for pos := 2; pos < len(cmd); pos++ { //loops through all the args
						if pos-2 >= len(order) { //makes sure we dont ignore without issues
							break //breaks from the loop
						}
						//inserts into the map without issues
						//this will make sure its done without errors happening
						args[order[pos-2]] = cmd[pos]
						continue //and continues looping
					}

					//ranges through the order properly
					//this will ensure its done without errors happening on request
					for setting := range order {
						//this will ensure that the option hasnt been made inside map
						//stops items from repeating without issues taking place on purpose etc
						if _, ok := args[order[setting]]; ok { //checks if the value is true or false
							continue //continues the loop properly
						}
						//tries to correctly deploy the prompt without issues happening
						//this will try to stop without issues happening on request without errors
						if err := language.ExecuteLanguage([]string{"users", "create", order[setting] + ".itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
							return err //returns the error correctly and properly
						}
						//prepares and tries to take the input with out issues
						//this will make sure when its executed it doesnt cause issues
						value, err := term.NewTerminal(s.Channel, "").ReadLine() //reads the input
						if err != nil {                                          //properly error handles without issues happening on request
							return err //returns the error properly on purpose
						}

						//saves into the value properly
						//this makes sure its done without issues happening
						args[order[setting]] = value //saves into the socket properly
					}

					if len(args["username"]) <= 0 { //checks the username properly and safely without issues happening
						return language.ExecuteLanguage([]string{"users", "create", "invalid-username.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//generates a properly secure password
					//this will make sure its done without errors happening
					SecurePassword := tools.CreateStrongPassword(json.ConfigSettings.Masters.Accounts.PasswordLength)
					ApiKey := tools.CreateStrongPassword(toml.ApiToml.API.KeyLen)

					//this will properly deploy the randomly generated password
					//ensures its safer without issues happening on request making it safer
					if err := language.ExecuteLanguage([]string{"users", "create", "password.itl"}, s.Channel, deployment.Engine, s, map[string]string{"password": SecurePassword}); err != nil {
						return err //returns the error correctly and properly
					}

					//this will properly deploy the randomly generated apikey
					//ensures its safer without issues happening on request making it safer
					if err := language.ExecuteLanguage([]string{"users", "create", "api-key.itl"}, s.Channel, deployment.Engine, s, map[string]string{"apikey": ApiKey}); err != nil {
						return err //returns the error correctly and properly
					}

					//checks if the username already exists
					//makes sure its not ignored without issues happening
					User, err := database.Conn.FindUser(args["username"])
					if User != nil && err == nil { //basic error handling properly and safely
						return language.ExecuteLanguage([]string{"users", "create", "usr-already.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": args["username"]})
					}

					//this will properly try to validate the information
					//makes sure the fields are valid without issues happening
					MaxTime, err := strconv.Atoi(args["maxtime"]) //uses the maxtime
					if err != nil {                               //error handles the information without issues happening
						return language.ExecuteLanguage([]string{"users", "create", "maxtime-atoi.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": args["username"], "maxtime": args["maxtime"]})
					}

					//this will properly try to validate the information
					//makes sure the fields are valid without issues happening
					Cooldown, err := strconv.Atoi(args["cooldown"]) //uses the cooldown
					if err != nil {                                 //error handles the information without issues happening
						return language.ExecuteLanguage([]string{"users", "create", "cooldown-atoi.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": args["username"], "cooldown": args["cooldown"]})
					}

					//this will properly try to validate the information
					//makes sure the fields are valid without issues happening
					Concurrents, err := strconv.Atoi(args["concurrents"]) //uses the concurrents
					if err != nil {                                       //error handles the information without issues happening
						return language.ExecuteLanguage([]string{"users", "create", "concurrents-atoi.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": args["username"], "concurrents": args["concurrents"]})
					}

					//properly tries to create the user
					//this will ensure its done without error happening on request
					err = database.Conn.MakeUser(&database.User{Username: args["username"], Password: SecurePassword, Ranks: "", MaxTime: MaxTime, Concurrents: Concurrents, Cooldown: Cooldown, MaxSessions: json.ConfigSettings.Masters.Accounts.MaxSessions, NewUser: true, Theme: "default", Expiry: time.Now().Add((time.Hour * 24) * time.Duration(json.ConfigSettings.Masters.Accounts.DaysExpiry)).Unix(), Parent: s.User.Identity, Token: ApiKey})
					if err != nil { //error handles the request properly without issues happening
						return language.ExecuteLanguage([]string{"users", "create", "creation-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": args["username"]})
					}

					//returns the success message properly
					//this will ensure its completed without errors happening on request
					return language.ExecuteLanguage([]string{"users", "create", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": args["username"], "concurrents": args["concurrents"], "cooldown": args["cooldown"], "maxtime": args["maxtime"], "maxsessions": strconv.Itoa(json.ConfigSettings.Masters.Accounts.MaxSessions)})
				},
			},
			{
				SubcommandName:     "admin",
				Description:        "view users with powersaving",
				CommandPermissions: []string{"admin"}, CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					return FilterRank(s, "users", "admin")
				},
			},
			{
				SubcommandName:     "admin=",
				Description:        "control admin privileges",
				CommandPermissions: []string{"admin"},
				CommandSplit:       "=", //allows for splits to be added properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "admin", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//parses the unit properly into boolean
					//this will allow many different factors to happen
					decider, err := strconv.ParseBool(strings.Split(cmd[1], "=")[1])
					if err != nil { //returns the error properly statement and safely without issues happening
						return language.ExecuteLanguage([]string{"users", "admin", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to promote already can access it
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "admin", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						users := ranks.MakeRank(usr.Username)
						//properly tries to sync the ranks
						//this will ensure its done without issues happening
						if err := users.SyncWithString(usr.Ranks); err != nil { //basic error handles the statement without issues happening
							language.ExecuteLanguage([]string{"users", "admin", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//checks if the rank system already has it
						//this will ensure we dont dounble rank a user with issues
						if users.CanAccess("admin") == decider { //checks for it properly
							if decider { //returns already has the rank properly
								language.ExecuteLanguage([]string{"users", "admin", "has.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							} else { //already has not got permissions alert
								language.ExecuteLanguage([]string{"users", "admin", "not.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							}
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//decides what we will do
						//this will ensure we properly do the correct action
						if decider { //tries to properly give them the rank without issues
							if err := users.GiveRank("admin"); err != nil { //error handles the give
								language.ExecuteLanguage([]string{"users", "admin", "giveFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							}
						} else { //tries to properly remove the rank without issues happening
							if err := users.RemoveRank("admin"); err != nil { //error handles the take statement properly
								language.ExecuteLanguage([]string{"users", "admin", "takeFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							}
						}

						//compresses the ranks into string format
						//this will ensure they are done properly without issues
						format, err := users.MakeString() //converts into string format
						if err != nil {                   //error handles and tries to deploy the database error if needed
							language.ExecuteLanguage([]string{"users", "admin", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//correctly formats the information
						//this will ensure its done without issues happening
						if err := database.Conn.EditRanks(format, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "admin", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							sessions.Sessions[t.ID].User.Ranks = format //updates the string and next we will render the alert branding properly
							sessions.Sessions[t.ID].Ranks = users       //ranks structure inside session
							t.Ranks = users
							t.User.Ranks = format //formats without issues happening
							if err := language.ExecuteLanguage([]string{"alerts", "admin-" + strconv.FormatBool(decider) + ".itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Your admin status has been changed to " + strconv.FormatBool(decider) + " by " + s.User.Username + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})

						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "admin", "success-" + strconv.FormatBool(decider) + ".itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
						continue
					}

					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {

					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "moderator",
				Description:        "view users with moderator",
				CommandPermissions: []string{"admin", "moderator"}, CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					return FilterRank(s, "users", "moderator")
				},
			},
			{
				SubcommandName:     "moderator=",
				Description:        "control moderator privileges",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       "=", //allows for splits to be added properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "moderator", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//parses the unit properly into boolean
					//this will allow many different factors to happen
					decider, err := strconv.ParseBool(strings.Split(cmd[1], "=")[1])
					if err != nil { //returns the error properly statement and safely without issues happening
						return language.ExecuteLanguage([]string{"users", "moderator", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to promote already can access it
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "moderator", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						users := ranks.MakeRank(usr.Username)
						//properly tries to sync the ranks
						//this will ensure its done without issues happening
						if err := users.SyncWithString(usr.Ranks); err != nil { //basic error handles the statement without issues happening
							language.ExecuteLanguage([]string{"users", "moderator", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//checks if the rank system already has it
						//this will ensure we dont dounble rank a user with issues
						if users.CanAccess("moderator") == decider { //checks for it properly
							if decider { //returns already has the rank properly
								language.ExecuteLanguage([]string{"users", "moderator", "has.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							} else { //already has not got permissions alert
								language.ExecuteLanguage([]string{"users", "moderator", "not.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							}
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//decides what we will do
						//this will ensure we properly do the correct action
						if decider { //tries to properly give them the rank without issues
							if err := users.GiveRank("moderator"); err != nil { //error handles the give
								language.ExecuteLanguage([]string{"users", "moderator", "giveFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							}
						} else { //tries to properly remove the rank without issues happening
							if err := users.RemoveRank("moderator"); err != nil { //error handles the take statement properly
								language.ExecuteLanguage([]string{"users", "moderator", "takeFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
								continue
							}
						}

						//compresses the ranks into string format
						//this will ensure they are done properly without issues
						format, err := users.MakeString() //converts into string format
						if err != nil {                   //error handles and tries to deploy the database error if needed
							language.ExecuteLanguage([]string{"users", "moderator", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//correctly formats the information
						//this will ensure its done without issues happening
						if err := database.Conn.EditRanks(format, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "moderator", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							sessions.Sessions[t.ID].User.Ranks = format //updates the string and next we will render the alert branding properly
							sessions.Sessions[t.ID].Ranks = users       //the ranks structure properly
							t.Ranks = users                             //ranks inside session
							t.User.Ranks = format                       //ranks string inside session
							if err := language.ExecuteLanguage([]string{"alerts", "moderator-" + strconv.FormatBool(decider) + ".itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Your moderator status has been changed to " + strconv.FormatBool(decider) + " by " + s.User.Username + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})

						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "moderator", "success-" + strconv.FormatBool(decider) + ".itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
						continue
					}

					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "remove",
				Description:        "remove users from database",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//checks the cmd length properly
					//this will ensure the args needed are given properly
					if len(cmd) <= 2 { //checks the length correctly and returns otherwise properly
						return language.ExecuteLanguage([]string{"users", "remove", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the users given
					//this will ensure this is properly done without issues
					for arg := 2; arg < len(cmd); arg++ { //allows us to access each one
						//tries to get them from the database properly
						//this will ensure its done without errors happening
						usr, err := database.Conn.FindUser(cmd[arg]) //properly tries to find the user
						if err != nil {                              //error handles properly without issues happening this will also continue the looping
							language.ExecuteLanguage([]string{"users", "remove", "user-EOF.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly remove the user
						//this will ensure its done without issues happening
						err = database.Conn.RemoveUser(usr.Username) //removes the user
						if err != nil {                              //error handles the statement correctly and also continues the looping
							language.ExecuteLanguage([]string{"users", "remove", "database-EOF.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							if err := language.ExecuteLanguage([]string{"alerts", "user-removed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Your account has been removed by " + s.User.Username + "\x1b[0m\x1b8")
								return //kills the function properly
							}
							//closes the channel
							t.Channel.Close()               //stops ssh connection
							delete(sessions.Sessions, t.ID) //removes from the sessions
						})
						//tries to correctly execute the pointer without issues
						//this will ensure its done without issues happening on request
						language.ExecuteLanguage([]string{"users", "remove", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
						continue
					}

					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "password=",
				Description:        "update a users password",
				CommandPermissions: []string{"admin"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "password", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Password := strings.Split(cmd[1], "=")[1]
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "password", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.Password(Password, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "password", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}

						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.Password = database.HashProduct(Password) //updates the sessions password correctly and properly
							sessions.Sessions[t.ID].User.Password = database.HashProduct(Password)
							if err := language.ExecuteLanguage([]string{"alerts", "password-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Your account password has been changed by " + s.User.Username + ", please log back in\x1b[0m\x1b8")
								return //kills the function properly
							}
							//closes the channel
							t.Channel.Close()               //stops ssh connection
							delete(sessions.Sessions, t.ID) //removes from the sessions
						})

						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "password", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "cooldown=",
				Description:        "change a users cooldown",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "cooldown", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Cooldown, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "cooldown", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "cooldown", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.EditCooldown(Cooldown, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "cooldown", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.Cooldown = Cooldown //updates the users cooldown properly
							sessions.Sessions[t.ID].User.Cooldown = Cooldown
							if err := language.ExecuteLanguage([]string{"alerts", "cooldown-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldcooldown": strconv.Itoa(usr.Cooldown), "newcooldown": strconv.Itoa(Cooldown)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Cooldown changed from " + strconv.Itoa(usr.Cooldown) + " -> " + strconv.Itoa(Cooldown) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "cooldown", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldcooldown": strconv.Itoa(usr.Cooldown), "newcooldown": strconv.Itoa(Cooldown)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "cooldown_add=",
				Description:        "add onto a users cooldown",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "cooldown", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Cooldown, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "cooldown", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "cooldown", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.EditCooldown(Cooldown, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "cooldown", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.Cooldown = usr.Cooldown + Cooldown //updates the users cooldown properly
							sessions.Sessions[t.ID].User.Cooldown = usr.Cooldown + Cooldown
							if err := language.ExecuteLanguage([]string{"alerts", "cooldown-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldcooldown": strconv.Itoa(usr.Cooldown), "newcooldown": strconv.Itoa(usr.Cooldown + Cooldown)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Cooldown changed from " + strconv.Itoa(usr.Cooldown) + " -> " + strconv.Itoa(usr.Cooldown+Cooldown) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "cooldown", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldcooldown": strconv.Itoa(usr.Cooldown), "newcooldown": strconv.Itoa(usr.Cooldown + Cooldown)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "maxtime=",
				Description:        "change a users maxtime",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "maxtime", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Maxtime, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "maxtime", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "maxtime", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.EditMaxTime(Maxtime, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "maxtime", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.MaxTime = Maxtime //updates the users cooldown properly
							sessions.Sessions[t.ID].User.MaxTime = Maxtime
							if err := language.ExecuteLanguage([]string{"alerts", "maxtime-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldmaxtime": strconv.Itoa(usr.MaxTime), "newmaxtime": strconv.Itoa(Maxtime)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Maxtime changed from " + strconv.Itoa(usr.MaxTime) + " -> " + strconv.Itoa(Maxtime) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "maxtime", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldmaxtime": strconv.Itoa(usr.MaxTime), "newmaxtime": strconv.Itoa(Maxtime)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "maxtime_add=",
				Description:        "add onto a users maxtime",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "maxtime_add", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Maxtime, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "maxtime_add", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "maxtime_add", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.EditMaxTime(usr.MaxTime+Maxtime, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "maxtime_add", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.MaxTime = usr.MaxTime + Maxtime //updates the users cooldown properly
							sessions.Sessions[t.ID].User.MaxTime = usr.MaxTime + Maxtime
							if err := language.ExecuteLanguage([]string{"alerts", "maxtime-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldmaxtime": strconv.Itoa(usr.MaxTime), "newmaxtime": strconv.Itoa(usr.MaxTime + Maxtime)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Maxtime changed from " + strconv.Itoa(usr.MaxTime) + " -> " + strconv.Itoa(usr.MaxTime+Maxtime) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "maxtime_add", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldmaxtime": strconv.Itoa(usr.MaxTime), "newmaxtime": strconv.Itoa(usr.MaxTime + Maxtime)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "concurrents=",
				Description:        "change a users conns",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "concurrents", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Conns, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "concurrents", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "concurrents", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}
						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.EditConcurrents(Conns, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "concurrents", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.Concurrents = Conns //updates the users cooldown properly
							sessions.Sessions[t.ID].User.Concurrents = Conns
							if err := language.ExecuteLanguage([]string{"alerts", "concurrents-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldconcurrents": strconv.Itoa(usr.Concurrents), "newconcurrents": strconv.Itoa(Conns)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K concurrents changed from " + strconv.Itoa(usr.Concurrents) + " -> " + strconv.Itoa(Conns) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "concurrents", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldconcurrents": strconv.Itoa(usr.Concurrents), "newconcurrents": strconv.Itoa(Conns)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "concurrents_add=",
				Description:        "add onto a users conns",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "concurrents_add", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//gets the password properly and safely
					//this will be properly used inside this function without issues
					Conns, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "concurrents_add", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "concurrents_add", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}
						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.EditConcurrents(usr.Concurrents+Conns, usr.Username); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "concurrents_add", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.Concurrents = usr.Concurrents + Conns //updates the users cooldown properly
							sessions.Sessions[t.ID].User.Concurrents = usr.Concurrents + Conns
							if err := language.ExecuteLanguage([]string{"alerts", "concurrents-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldconcurrents": strconv.Itoa(usr.Concurrents), "newconcurrents": strconv.Itoa(usr.Concurrents + Conns)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K concurrents changed from " + strconv.Itoa(usr.Concurrents) + " -> " + strconv.Itoa(usr.Concurrents+Conns) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "concurrents_add", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldconcurrents": strconv.Itoa(usr.Concurrents), "newconcurrents": strconv.Itoa(usr.Concurrents + Conns)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "add-days=", //sets the subcommand name
				Description:        "add days onto plans",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "add_days", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "add_days", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "add_days", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", "add_days", "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//syncs the users time properly
						//this will ensure its done without any issues
						currentExpiry := time.Unix(user.Expiry, 0) //syncs properly

						//works the new expiry properly
						//this will ensure its done without any errors
						new := currentExpiry.Add((time.Hour * 24) * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "add_days", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "days-added.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "days": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "add_days", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "days": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "add-hours=", //sets the subcommand name
				Description:        "adds hours onto plans",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "add_hours", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "add_hours", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "add_hours", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//syncs the users time properly
						//this will ensure its done without any issues
						currentExpiry := time.Unix(user.Expiry, 0) //syncs properly

						//works the new expiry properly
						//this will ensure its done without any errors
						new := currentExpiry.Add(time.Hour * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "add_hours", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "hours-added.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "days": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "add_hours", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "days": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "add-minutes=", //sets the subcommand name
				Description:        "adds minutes onto plans",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "add_minutes", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "add_minutes", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "add_minutes", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//syncs the users time properly
						//this will ensure its done without any issues
						currentExpiry := time.Unix(user.Expiry, 0) //syncs properly

						//works the new expiry properly
						//this will ensure its done without any errors
						new := currentExpiry.Add(time.Minute * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "add_minutes", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "minutes-added.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "minutes": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "add_minutes", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "minutes": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "add-seconds=", //sets the subcommand name
				Description:        "adds seconds onto plans",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "add_seconds", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "add_seconds", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "add_seconds", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//syncs the users time properly
						//this will ensure its done without any issues
						currentExpiry := time.Unix(user.Expiry, 0) //syncs properly

						//works the new expiry properly
						//this will ensure its done without any errors
						new := currentExpiry.Add(time.Second * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "add_seconds", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "seconds-added.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "seconds": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "add_seconds", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "seconds": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "set-days=", //sets the subcommand name
				Description:        "set days for a plans",
				CommandPermissions: []string{"admin"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "set_days", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "set_days", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "set_days", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", "set_days", "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//works the new expiry properly
						//this will ensure its done without any errors
						new := time.Now().Add((time.Hour * 24) * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "set_days", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "days-set.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "days": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "set_days", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "days": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "set-hours=", //sets the subcommand name
				Description:        "set hours for a plans",
				CommandPermissions: []string{"admin"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "set_hours", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "set_hours", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "set_hours", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", "set_hours", "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//works the new expiry properly
						//this will ensure its done without any errors
						new := time.Now().Add((time.Hour) * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "set_hours", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "hours-set.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "hours": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "set_hours", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "hours": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "set-minutes=", //sets the subcommand name
				Description:        "set minutes for a plans",
				CommandPermissions: []string{"admin"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "set_minutes", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "set_minutes", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "set_minutes", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", "set_minutes", "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//works the new expiry properly
						//this will ensure its done without any errors
						new := time.Now().Add((time.Minute) * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "set_minutes", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "minutes-set.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "minutes": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "set_minutes", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "minutes": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "set-seconds=", //sets the subcommand name
				Description:        "set seconds for a plans",
				CommandPermissions: []string{"admin"},
				CommandSplit:       "=", //sets the command split properly
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "set_seconds", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Duration, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "set_seconds", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "set_seconds", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", "set_seconds", "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						if user.Expiry <= 0 { //checks for the time properly
							//properly tries to format without issues happening
							//allows for better settings without issues happening on reqeust
							user.Expiry = time.Now().Unix() //adds the time properly
						}

						//works the new expiry properly
						//this will ensure its done without any errors
						new := time.Now().Add((time.Second) * time.Duration(Duration)).Unix()

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Expiry(user.Username, new); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "set_seconds", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.Expiry = new //updates the new expiry properly
							sessions.Sessions[t.ID].User.Expiry = new
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "seconds-set.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "seconds": strconv.Itoa(Duration)})
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "set_seconds", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "seconds": strconv.Itoa(Duration)})
						continue
					}
					return nil
				},
			},
			{
				SubcommandName:     "filter=",
				Description:        "list users with privileges",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=", //splits the first command name by the first object
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//checks the amount inside the string properly
					//this will ensure its done without any errors
					if strings.Count(cmd[1], "=") < 1 { //len checks
						return language.ExecuteLanguage([]string{"users", "filter", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//stores all of our filter targets properly
					//this will ensure its done without any errors happening
					access := strings.Split(strings.Join(strings.Split(cmd[1], "=")[1:], "="), ",")

					//tries to access all of the users properly
					//this will ensure its done without any errors happening
					users, err := database.Conn.GetUsers() //gets the users properly
					if err != nil {                        //error handles properly without issues happening
						return err //returns the error which happened on purpose properly
					}

					//creates our new simpletable
					//this will be what we perform our split with
					table := simpletable.New() //stored inside a simpletable

					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-id.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-maxtime.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-concurrents.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-cooldown.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-ranks.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through the users properly
					//this will ensure its done without any errors
					for _, user := range users { //ranges through properly

						//creates our newer system preset
						//this will ensure its done without errors
						system := ranks.MakeRank(user.Username) //err checks sync
						if err := system.SyncWithString(user.Ranks); err != nil {
							continue //continues looping properly without issues
						}

						//checks if the user can access the objects
						//this will ensure its done without any errors
						if !system.CanAccessArray(access) { //checks properly
							continue
						}

						//tries to deploy the ranks properly
						//this will form a pretty rank schema
						Ranks, err := system.DeployRanks(true) //deploys
						if err != nil {                        //error handles without issues
							continue //continues looping properly without errors
						}

						var ur string = strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-username.txt").Containing, lexer.Escapes), "<<$username>>", user.Username)
						if len(s.GetSessions(user.Username)) > 0 {
							ur = strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-user-online.txt").Containing, lexer.Escapes), "<<$username>>", user.Username)
						}

						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(user.Identity))}, //id
							{Align: simpletable.AlignLeft, Text: ur}, //username
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-maxtime.txt").Containing, lexer.Escapes), "<<$maxtime>>", HandleTime(user.MaxTime))},             //maxtime
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-concurrents.txt").Containing, lexer.Escapes), "<<$concurrents>>", HandleTime(user.Concurrents))}, //conns
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-cooldown.txt").Containing, lexer.Escapes), "<<$cooldown>>", strconv.Itoa(user.Cooldown))},        //cooldown
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(Ranks, " "))},                 //ranks
						}

						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}

					//runs our pager system properly and safely
					//this will ensure its done without any errors happening
					return pager.MakeTable("users", table, s).TextureTable()
				},
			},
			{
				SubcommandName:     "slaves=",
				Description:        "control slave access",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit:       "=", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "slaves", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the duration into int
					//this will properly try without issues happening
					Amount, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "slaves", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "slaves", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}
						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to update the max slaves for the users
						//this will ensure its done without errors happening
						if err := database.Conn.EditMaxSlaves(user.Username, Amount); err != nil { //err handles properly
							language.ExecuteLanguage([]string{"users", "slaves", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							t.User.MaxSlaves = Amount                       //updates the internal sessions
							sessions.Sessions[t.ID].User.MaxSlaves = Amount //updates the external
							//tries to render the alert properly
							//this will ensure its done without any errors
							language.ExecuteLanguage([]string{"alert", "slaves_updated.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username, "old": strconv.Itoa(user.MaxSlaves)})
						})

						//renders the information properly and safely
						//this will make sure its done without errors happening
						language.ExecuteLanguage([]string{"users", "slaves", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username, "old": strconv.Itoa(user.MaxSlaves)})
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "newuser=",
				Description:        "control newuser privileges",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       "=", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "newuser", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to convert the status into boolean
					//this will properly try without issues happening
					Status, err := strconv.ParseBool(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "newuser", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "newuser", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//checks if they already have it properly
						if user.NewUser == Status { //renders the information properly and safely
							language.ExecuteLanguage([]string{"users", "newuser", "already_" + strconv.FormatBool(Status) + ".itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						var errDB error = nil
						switch Status { //performs the reqeust without errors
						case true: //enables the new user option properly
							errDB = database.Conn.EnableNewUser(user.Username)
						default: //disables the new user option properly
							errDB = database.Conn.DisableNewUser(user.Username)
						}

						if errDB != nil { //error handles the statement properly
							//this will ensure its done without errors happenign on purpose
							language.ExecuteLanguage([]string{"users", "newuser", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//renders the information properly and safely
						//this will make sure its done without errors happening
						language.ExecuteLanguage([]string{"users", "newuser", "success_" + strconv.FormatBool(Status) + ".itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username})
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "view",
				Description:        "view user account information",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//checks the length within the arguments properly
					//this will ensure its done without any errors happening
					if len(cmd) <= 2 { //checks the length given properly safely
						return language.ExecuteLanguage([]string{"users", "view", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to get the username properly
					//this will ensure we have gotton it properly
					user, err := database.Conn.FindUser(cmd[2]) //find user
					if err != nil || user == nil {              //tries to find the user properly
						return language.ExecuteLanguage([]string{"users", "view", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"user": cmd[2]})
					}

					//tries to get all of the user logins properly
					//this will ensure its done properly without issues
					logins, err := database.Conn.GetLogins(user.Username)
					if err != nil { //error handles properly without issues
						return language.ExecuteLanguage([]string{"users", "view", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"user": cmd[2]})
					}

					//tries to fetch all of the users attack ever sent
					//this will ensure its done without any errors happening
					attacks, err := database.Conn.UserSent(user.Username)
					if err != nil { //error handles properly without issues
						return language.ExecuteLanguage([]string{"users", "view", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"user": cmd[2]})
					}

					//stores the last attack ever sent properly
					//this will ensure its done without any errors happening
					var LastTarget *database.AttackLinear = &database.AttackLinear{Method: "EOF", Target: "EOF", Username: user.Username, ID: -1, Duration: 0, Port: 0, Created: time.Now().Unix(), Finish: time.Now().Unix(), SentViaAPI: false}
					if len(attacks) > 0 { //checks length proeprly without issues
						LastTarget = &attacks[len(attacks)-1]
					}
					//auto fills a default login section properly
					//this will ensure its done without any errors happening
					var LastLogin *database.Login = &database.Login{Address: "EOF", TimeStore: time.Now().Unix(), Banner: "EOF", Username: user.Username, Success: false}
					if len(attacks) > 0 { //checks the length properly without issues happening
						LastLogin = &logins[len(logins)-1]
					}

					//combines the ranks properly without issues
					//this will ensure its done without any errors happening
					System := ranks.MakeRank(user.Username)                   //makes the new instance properly
					if err := System.SyncWithString(user.Ranks); err != nil { //error handles properly
						return language.ExecuteLanguage([]string{"users", "view", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"user": cmd[2]})
					}

					//deploys the system properly
					//this will ensure its done without any errors
					Ranks, err := System.DeployRanks(true) //compiles the ranks properly
					if err != nil {                        //compiles into a string properly without issues happening
						return language.ExecuteLanguage([]string{"users", "view", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"user": cmd[2]})
					}

					//writes information about the users information properly
					//this will ensure its done without any errors happening on purpose
					s.Write(Centre("Header", 20, "") + Centre("Value", 20, "") + Centre("Header", 20, "") + Centre("Value", 20, "") + "\r\n")

					var parent string = "EOF"
					//tries to get the username via the parent properly
					//this will ensure its done without any errors happening
					parentUser, err := database.Conn.GetUserViaParent(user.Parent)
					if err == nil && parentUser != nil { //checks to see if we have got the parent username
						if parentUser.Identity == user.Identity { //tries to get the id properly without issues
							parent = "<original>" //checks properly
						} else { //else statement properly
							parent = parentUser.Username //gets the parents username properly
						}
					}

					var MFA bool = false
					if len(user.MFA_secret) > 0 {
						MFA = true
					}
					//first line about the login information properly
					//this will ensure its done without any errors happening on purpose
					s.Write("\x1b[38;5;15m" + BuildLine("Logins:", strconv.Itoa(len(logins)), "Last Login:", tools.ResolveTimeStamp(time.Unix(LastLogin.TimeStore, 0), true), s))
					s.Write("\x1b[38;5;15m" + BuildLine("maxTime:", strconv.Itoa(user.MaxTime), "cooldown:", strconv.Itoa(user.Cooldown), s))                                                              //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("concurrents:", strconv.Itoa(user.Concurrents), "Ranks:", strings.Join(Ranks, " "), s))                                                            //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Expires:", tools.ResolveTimeStampUnix(time.Unix(user.Expiry, 0), true), "Created:", tools.ResolveTimeStamp(time.Unix(user.Created, 0), true), s)) //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Updated:", tools.ResolveTimeStamp(time.Unix(user.Updated, 0), true), "Parent:", parent, s))                                                       //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Attacks:", strconv.Itoa(len(attacks)), "NewUser:", MakeBoolean(user.NewUser), s))                                                                 //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Last Target:", LastTarget.Target, "Banned:", MakeBoolean(System.CanAccess("banned")), s))                                                         //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Last Method:", LastTarget.Method, "Admin:", MakeBoolean(System.CanAccess("admin")), s))                                                           //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Attack finished:", tools.ResolveTimeStamp(time.Unix(LastTarget.Finish, 0), true), "Moderator:", MakeBoolean(System.CanAccess("moderator")), s))   //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Theme:", user.Theme, "User ID:", strconv.Itoa(user.Identity), s))                                                                                 //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Max Sessions:", strconv.Itoa(user.MaxSessions), "APIUser:", MakeBoolean(System.CanAccess("api")), s))                                             //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Max Slaves:", strconv.Itoa(user.MaxSlaves), "MFA:", MakeBoolean(MFA), s))                                                                         //displays information properly
					s.Write("\x1b[38;5;15m" + BuildLine("Plan:", user.Plan, "Locked:", MakeBoolean(user.Locked), s))                                                                                       //displays information properly
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "subordinates=",
				Description:        "store all user subordinates",
				CommandPermissions: []string{"admin"}, CommandSplit: "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					if strings.Count(cmd[1], "=") <= 0 { //counts the ranges information properly and safely
						return language.ExecuteLanguage([]string{"users", "subordinates", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//this will ensure its done without issues
					user := strings.Split(cmd[1], "=")[1] //user storage

					//tries to grab the user properly
					//this will ensure we can track the tracer
					target, err := database.Conn.FindUser(user)
					if err != nil { //err handles
						return language.ExecuteLanguage([]string{"users", "subordinates", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user})
					}

					//grabs all users which the user created properly
					//this will ensure its done without errors happening
					subordinate, err := database.Conn.ParentTracer(target.Identity)
					if err != nil { //err handles properly
						return language.ExecuteLanguage([]string{"users", "subordinates", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user})
					}

					//creates the new table properly
					//this will be used to render properly
					Table := simpletable.New() //new table

					//fills the table headers properly
					Table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "subordinates", "header", "id.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "subordinates", "header", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "subordinates", "header", "ranks.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through the user subordinates properly
					//this will display everything who the user created
					for _, username := range subordinate { //ranges through

						systemRank := ranks.MakeRank(username.Username)
						if err := systemRank.SyncWithString(username.Ranks); err != nil {
							continue
						}

						//deploys the system safely
						//this will ensure its done without errors
						pure, err := systemRank.DeployRanks(true)
						if err != nil { //err handles properly
							continue
						}

						row := []*simpletable.Cell{ //stores the information properly``
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(strings.ReplaceAll(views.GetView("users", "subordinates", "value", "id.txt").Containing, "<<$id>>", strconv.Itoa(username.Identity)), lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(strings.ReplaceAll(views.GetView("users", "subordinates", "value", "username.txt").Containing, "<<$username>>", username.Username), lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(strings.ReplaceAll(views.GetView("users", "subordinates", "value", "ranks.txt").Containing, "<<$ranks>>", strings.Join(pure, " ")), lexer.Escapes)},
						}

						//saves into table properly
						Table.Body.Cells = append(Table.Body.Cells, row)
					}

					//creates the table properly and safely
					//this will render the information into the terminal
					return pager.MakeTable("users", Table, s).TextureTable()
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "lock",
				Description:        "lock a users account down",
				CommandPermissions: []string{"admin"}, CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "lock", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg])                      //tries to find
						if err != nil || user == nil || user.Username == s.User.Username { //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "lock", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Lock(user.Username); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "lock", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							language.ExecuteLanguage([]string{"alert", "account_locked.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username})
							t.Channel.Close()
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "lock", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "unlock",
				Description:        "unlock a users account",
				CommandPermissions: []string{"admin"}, CommandSplit: "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "unlock", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "unlock", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.Unlock(user.Username); err != nil { //error handles properly without issus
							language.ExecuteLanguage([]string{"users", "unlock", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "unlock", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			//{
			//	SubcommandName: "kick",
			//	Description: "kick user for x amount of time",
			//	CommandPermissions: []string{"admin"}, CommandSplit: " ",
			//	SubCommandFunction: func(s *sessions.Session, cmd []string) error {
			//
			//		return nil
			//	},
			//},
			{
				SubcommandName:     "sessions=",
				Description:        "edit max amount of sessions",
				CommandPermissions: []string{"admin"}, CommandSplit: "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//detects if the object happens
					//this will start the routes without issues
					if strings.Count(cmd[1], "=") <= 0 || len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "sessions", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets the password properly and safely
					//this will be properly used inside this function without issues
					NewSessions, err := strconv.Atoi(strings.Split(cmd[1], "=")[1])
					if err != nil { //renders the syntax error properly
						return language.ExecuteLanguage([]string{"users", "sessions", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}
					//ranges through all the arguments
					//this will ensure its done properly without issues happening
					for arg := 2; arg < len(cmd); arg++ { //ranges through the arguments given properly
						//tries to find the user inside the database
						//this will make sure the user they are trying to update a password is valid
						usr, err := database.Conn.FindUser(cmd[arg]) //tries to find the user properly
						if err != nil {                              //error handles the statement correctly
							language.ExecuteLanguage([]string{"users", "sessions", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, usr.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", usr.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to correctly update the users password
						//this will ensure its done without errors happening
						if err := database.Conn.Sessions(usr.Username, NewSessions); err != nil { //renders the database error if needed properly
							language.ExecuteLanguage([]string{"users", "sessions", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username})
							continue
						}
						//this will properly try to adjust the information without issues
						//broadcasts the message to the remote host without issues happening
						s.FunctionRemote(usr.Username, func(t *sessions.Session) {
							t.User.MaxSessions = NewSessions //updates the users maxsessions properly
							sessions.Sessions[t.ID].User.MaxSessions = NewSessions
							if err := language.ExecuteLanguage([]string{"alerts", "sessions-changed.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promotor": s.User.Username, "oldsessions": strconv.Itoa(usr.MaxSessions), "newsessions": strconv.Itoa(NewSessions)}); err != nil {
								t.Write("\x1b[0m\x1b7\x1b[1A\r\x1b[2K Sessions changed from " + strconv.Itoa(usr.MaxSessions) + " -> " + strconv.Itoa(NewSessions) + "\x1b[0m\x1b8")
								return //kills the function properly
							}
						})
						//renders the information properly
						//this will make sure its done properly without errors
						language.ExecuteLanguage([]string{"users", "sessions", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": usr.Username, "oldsessions": strconv.Itoa(usr.MaxSessions), "newsessions": strconv.Itoa(NewSessions)})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "remove_mfa",
				Description:        "removes the mfa from the user",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//this will ensure nothing which is invalid gets around
					//makes sure its secure without any issues happening on purpose
					if len(cmd) <= 2 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "remove_mfa", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the args properly
					//this will ensure its done without any errors
					for arg := 2; arg < len(cmd); arg++ { //loops through properly
						//tries to get the user properly without issues
						//this will allow for better handling without errors
						user, err := database.Conn.FindUser(cmd[arg]) //tries to find
						if err != nil || user == nil {                //error handles properly without issues
							language.ExecuteLanguage([]string{"users", "remove_mfa", "invalidUser.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//checks to see if they can modify
						//if this statement returns true they cant modify
						if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
							if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
								s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
							}
							continue //continues looping properly without issues
						}

						//tries to properly insert into the database
						//this will ensure its done without any errors
						if err := database.Conn.RmMFA(user.Username); err != nil { //error handles properly without issue
							language.ExecuteLanguage([]string{"users", "remove_mfa", "database-fault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[arg]})
							continue
						}

						//tries to connect to all sessions
						//this will ensure its done without any errors
						s.FunctionRemote(user.Username, func(t *sessions.Session) { //properly executes the function on the session
							language.ExecuteLanguage([]string{"alert", "mfa_reset.itl"}, t.Channel, deployment.Engine, t, map[string]string{"promoter": s.User.Username})
							t.Channel.Close()
						})

						//renders the success message properly
						//this will ensure its done without any errors
						language.ExecuteLanguage([]string{"users", "remove_mfa", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username})
						continue
					}
					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "recent",
				Description:        "recent login locations",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					if len(cmd) < 3 { //syntax issue presented at this point properly
						return language.ExecuteLanguage([]string{"users", "recent", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//grabs all login locations from the database
					//allows us to filter into the system properly
					locations, err := database.Conn.GetLogins(cmd[2])
					if err != nil { //renders the error sheet properly and safely without issues
						return language.ExecuteLanguage([]string{"users", "recent", "database.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": cmd[2]})
					}

					var recorded map[string]*Location = make(map[string]*Location)

					//ranges through the locations properly
					//this will ensure its done without issues
					for _, system := range locations {

						//gets the recorded system inside map properly
						//this will ensure its done without issues happening
						if _, ok := recorded[strings.Split(system.Address, ":")[0]]; ok {
							recorded[strings.Split(system.Address, ":")[0]].Last = time.Unix(system.TimeStore, 0) //sets last time
							recorded[strings.Split(system.Address, ":")[0]].Times++                               //updates the times founded
						} else {
							recorded[strings.Split(system.Address, ":")[0]] = &Location{ //creates new
								IP:    strings.Split(system.Address, ":")[0], //set ipv4
								Times: 1,                                     //sets amount of times
								First: time.Unix(system.TimeStore, 0),        //time store
							}
						}
					}

					table := simpletable.New()

					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "recent", "headers", "ip.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "recent", "headers", "times.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "recent", "headers", "first.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "recent", "headers", "last.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through the recent properly
					//this will ensure its done without issues
					for _, recent := range recorded {

						rk := []*simpletable.Cell{ //saves into the system properly and safely. ensures its done without errors
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "recent", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", recent.IP)},                                        //id
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "recent", "values", "times.txt").Containing, lexer.Escapes), "<<$times>>", strconv.Itoa(recent.Times))},                 //username
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "recent", "values", "first.txt").Containing, lexer.Escapes), "<<$first>>", tools.ResolveTimeStamp(recent.First, true))}, //maxtime
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "recent", "values", "last.txt").Containing, lexer.Escapes), "<<$last>>", tools.ResolveTimeStamp(recent.Last, true))},    //conns
						}

						table.Body.Cells = append(table.Body.Cells, rk)
					}

					return pager.MakeTable("users", table, s).TextureTable()
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "tokens",
				Description:        "view all redeemable tokens",
				CommandPermissions: []string{"admin", "moderator"}, CommandSplit: "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//grabs the tokens properly
					//ensures we have them all properly
					toks, err := database.Conn.GrabTokens()
					if err != nil { //err handles
						return err
					}

					//displays our headers properly and safely
					//this will ensure its done within safely and properly
					s.Write(Centre("Token", 33, "") + "|" + Centre("Information", s.Length-34, "") + "\r\n")

					//ranges through all tokens
					//displays the information properly
					for _, token := range toks {

						switch token.Bundle.Type {

						case 1: //user
							s.Write(fmt.Sprintf("%s| maxtime: %d cooldown: %d concurrents: %d", Centre(token.Token, 33, ""), token.Bundle.User.MaxTime, token.Bundle.User.Cooldown, token.Bundle.User.Concurrents) + "\r\n")
						}
						fmt.Println(token.Token, token.Bundle.Type, token.Bundle.User)
					}

					return nil
				},
				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
			{
				SubcommandName:     "regen=",
				Description:        "regenerate a users apikey",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       "=", SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					if strings.Count(cmd[1], "=") <= 0 || len(strings.Split(cmd[1], "=")[1]) <= 0 { //checks for one not being added
						return language.ExecuteLanguage([]string{"users", "regenerate", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//checks to validate the user exists properly and safely
					user, err := database.Conn.FindUser(strings.Split(cmd[1], "=")[1])
					if err != nil || user == nil { //invalid user might be detected
						return language.ExecuteLanguage([]string{"users", "regenerate", "invalid-user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": strings.Split(cmd[1], "=")[1]})
					}

					//checks to see if they can modify
					//if this statement returns true they cant modify
					if util.SearchTracer(s.TracerMatrix, user.Parent) && !s.CanAccess("admin") { //executes the higher statement render properly
						if err := language.ExecuteLanguage([]string{"users", strings.ToLower(strings.Split(cmd[1], "=")[0]), "tracer-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": user.Username}); err != nil {
							s.Write(fmt.Sprintf("the user (%s) obtains higher permissions than yourself (%s)\r\n", user.Username, s.User.Username))
						}
						return nil //continues looping properly without issues
					}

					//will generate the new strong apikey for usage
					apikey := tools.CreateStrongPassword(toml.ApiToml.API.KeyLen)

					//will try to update the apikey inside the database
					if err := database.Conn.APIKey(user.Username, apikey); err != nil {
						return language.ExecuteLanguage([]string{"users", "regenerate", "database-error.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": strings.Split(cmd[1], "=")[1]})
					}

					//returns the success information within the network correctly
					return language.ExecuteLanguage([]string{"users", "regenerate", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username": strings.Split(cmd[1], "=")[1], "apikey": apikey})
				},

				//AutoComplete allows for callback completing
				AutoComplete: func(s *sessions.Session) []string {
					//gets all users properly from the database
					//this will ensure its done without issues happening
					systemAccounts, err := database.Conn.GetUsers()
					if err != nil { //err handles properly
						return []string{"Error"}
					}

					//stores our future array full properly
					var system []string = make([]string, 0)

					//ranges through all accounts properly
					//this will ensure its done properly and safely
					for _, account := range systemAccounts { //ranges
						system = append(system, account.Username)
					}

					return system
				},
			},
		},
	})
}

type Location struct {
	IP    string
	Times int
	First time.Time
	Last  time.Time
}

// resolves the boolean into the string format
// allows for better control without issues happening
func MakeBoolean(b bool) string { //returns the string type
	if b { //gets the true banner properly
		return lexer.AnsiUtil(views.GetView("true.txt").Containing, lexer.Escapes)
	} else { //gets the false banner properly
		return lexer.AnsiUtil(views.GetView("false.txt").Containing, lexer.Escapes)
	}
}

// builds the line properly without issues happening
func BuildLine(h1, v1 string, h2, v2 string, s *sessions.Session) string { //returns a string ofset properly
	return PaddingRight(h1, 20) + PaddingRight(v1, 20) + PaddingRight(h2, 20) + PaddingRight(v2, 20) + "\r\n"
}

// centres the text properly
// this will return the string properly
func Centre(s string, c int, dst string) string { //string
	//for loops through properly
	//this will ensure its done without any issues
	for p := 0; p < c; p++ { //loops through properly
		if p == c/2-len(s)/2 { //compares to middle
			dst += s        //saves in properly
			p += len(s) - 1 //skips chars properly
		} else {
			dst += " "
		}
	}
	return dst
}

func FilterRank(session *sessions.Session, tableID string, access ...string) error {

	//tries to correctly get all the users
	//this will ensure its done without issues happening on reqeust
	users, err := database.Conn.GetUsers() //gets the users properly without issues
	if err != nil {                        //basic error handling without issues happening on request and returns the error statement
		return language.ExecuteLanguage([]string{"users", "databaseErr.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	}

	//creates the simpletable
	//this will store our information
	table := simpletable.New() //makes the structure

	//sets table header
	//this will be used to define information
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{ //sets the values properly without issues happening
			{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-id.txt").Containing, lexer.Escapes)},
			{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-username.txt").Containing, lexer.Escapes)},
			{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-maxtime.txt").Containing, lexer.Escapes)},
			{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-concurrents.txt").Containing, lexer.Escapes)},
			{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-cooldown.txt").Containing, lexer.Escapes)},
			{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("users", "headers", "header-ranks.txt").Containing, lexer.Escapes)},
		},
	}

	//ranges through all the users properly
	//this will insert each user into the database without issues happening on request
	for _, usr := range users {
		//creates the rank information
		//this will ensure its done without issues
		r := ranks.MakeRank(usr.Username) //makes the rank
		//tries to correctly sync the render without issues happening
		//this will make sure its done without issues happening on request
		if err := r.SyncWithString(usr.Ranks); err != nil {
			return err //returns the error correctly and properly
		}

		if len(access) > 0 {
			if !r.CanAccessArray(access) {
				continue
			}
		}

		//properly deploys the rank into string
		//this will ensure its done without errors happening
		ranks, err := r.DeployRanks(true) //deploys into string format
		if err != nil {                   //error handles the syntax properly without issues
			return err //returns the error properly
		}

		var user string = strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-username.txt").Containing, lexer.Escapes), "<<$username>>", usr.Username)
		if len(session.GetSessions(usr.Username)) > 0 {
			user = strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-user-online.txt").Containing, lexer.Escapes), "<<$username>>", usr.Username)
		}

		//creates the store properly without issues
		rk := []*simpletable.Cell{ //fills with the information properly without issues
			{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.Identity))}, //id
			{Align: simpletable.AlignLeft, Text: user}, //username
			{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-maxtime.txt").Containing, lexer.Escapes), "<<$maxtime>>", HandleTime(usr.MaxTime))},             //maxtime
			{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-concurrents.txt").Containing, lexer.Escapes), "<<$concurrents>>", HandleTime(usr.Concurrents))}, //conns
			{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-cooldown.txt").Containing, lexer.Escapes), "<<$cooldown>>", strconv.Itoa(usr.Cooldown))},        //cooldown
			{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("users", "values", "value-ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))},                //ranks
		}
		//saves into the array correctly
		//this will be properly founded later onwards
		table.Body.Cells = append(table.Body.Cells, rk)
	}

	//renders the table properly
	//this will ensure its done without issues
	return pager.MakeTable(tableID, table, session).TextureTable()
}

func HandleTime(period int) string {
	if period == 0 {
		return "∞"
	}
	return strconv.Itoa(period)
}
