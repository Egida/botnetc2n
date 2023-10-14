package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "sessions",
		Aliases:            []string{"session", "online"},
		CommandPermissions: []string{"admin", "moderator"},
		CommandDescription: "moderate & manage open sessions",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure

			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "id.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "username.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "connected.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ip.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "idle.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ranks.txt").Containing, lexer.Escapes)},
				},
			}


			//ranges through every opens ession
			//this will make sure its done without issues
			for _, usr := range sessions.Sessions {
				//creates the rank information
				//this will ensure its done without issues
				r := ranks.MakeRank(usr.User.Username) //makes the rank
				//tries to correctly sync the render without issues happening
				//this will make sure its done without issues happening on request
				if err := r.SyncWithString(usr.User.Ranks); err != nil {
					return err //returns the error correctly and properly
				}
				//properly deploys the rank into string
				//this will ensure its done without errors happening
				ranks, err := r.DeployRanks(true) //deploys into string format
				if err != nil { //error handles the syntax properly without issues
					return err //returns the error properly
				}

				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.User.Identity))}, //id
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", usr.User.Username)}, //username
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "connected.txt").Containing, lexer.Escapes), "<<$connected>>", usr.Connected.Format("15:04:05"))}, //maxtime
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", usr.Target)}, //conns
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "idle.txt").Containing, lexer.Escapes), "<<$idle>>", tools.ResolveTimeStamp(usr.Idle, false))}, //cooldown
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))}, //ranks
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("sessions", table, s).TextureTable()
		},
		InvalidSubCommand: func(s *sessions.Session, cmd []string) error {

			//checks if the rank is valid properly
			//this will ensure only valid options are ok
			if _, ok := ranks.PresetRanks[cmd[1]]; !ok {
				return language.ExecuteLanguage([]string{"errors", "command403.itl"}, s.Channel, deployment.Engine, s, map[string]string{"command":cmd[0]}) //executes properly
			}
			
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure

			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "id.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "username.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "connected.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ip.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "idle.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ranks.txt").Containing, lexer.Escapes)},
				},
			}


			//ranges through every opens ession
			//this will make sure its done without issues
			for _, usr := range sessions.Sessions {
				//creates the rank information
				//this will ensure its done without issues
				r := ranks.MakeRank(usr.User.Username) //makes the rank
				//tries to correctly sync the render without issues happening
				//this will make sure its done without issues happening on request
				if err := r.SyncWithString(usr.User.Ranks); err != nil {
					return err //returns the error correctly and properly
				}

				if !r.CanAccess(cmd[1]) {
					continue
				}


				//properly deploys the rank into string
				//this will ensure its done without errors happening
				ranks, err := r.DeployRanks(true) //deploys into string format
				if err != nil { //error handles the syntax properly without issues
					return err //returns the error properly
				}

				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.User.Identity))}, //id
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", usr.User.Username)}, //username
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "connected.txt").Containing, lexer.Escapes), "<<$connected>>", usr.Connected.Format("15:04:05"))}, //maxtime
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", usr.Target)}, //conns
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "idle.txt").Containing, lexer.Escapes), "<<$idle>>", tools.ResolveTimeStamp(usr.Idle, false))}, //cooldown
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))}, //ranks
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("sessions", table, s).TextureTable()
		},
		SubCommands: []SubCommand{
			{
				SubcommandName: "list",
				Description: "list all open sessions",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit: " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//creates the simpletable
					//this will store our information
					table := simpletable.New() //makes the structure

					//sets table header
					//this will be used to define information
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "list", "headers", "id.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "list", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "list", "headers", "connected.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "list", "headers", "ip.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "list", "headers", "idle.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "list", "headers", "ranks.txt").Containing, lexer.Escapes)},
						},
					}
				
				
					//ranges through every opens ession
					//this will make sure its done without issues
					for _, usr := range sessions.Sessions {
						//creates the rank information
						//this will ensure its done without issues
						r := ranks.MakeRank(usr.User.Username) //makes the rank
						//tries to correctly sync the render without issues happening
						//this will make sure its done without issues happening on request
						if err := r.SyncWithString(usr.User.Ranks); err != nil {
							return err //returns the error correctly and properly
						}
						//properly deploys the rank into string
						//this will ensure its done without errors happening
						ranks, err := r.DeployRanks(true) //deploys into string format
						if err != nil { //error handles the syntax properly without issues
							return err //returns the error properly
						}
					
						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "list", "values", "id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.User.Identity))}, //id
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "list", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", usr.User.Username)}, //username
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "list", "values", "connected.txt").Containing, lexer.Escapes), "<<$connected>>", usr.Connected.Format("15:04:05"))}, //maxtime
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "list", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", usr.Target)}, //conns
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "list", "values", "idle.txt").Containing, lexer.Escapes), "<<$idle>>", tools.ResolveTimeStamp(usr.Idle, false))}, //cooldown
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "list", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))}, //ranks
						}
					
						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}
				
					//renders the table properly
					//this will ensure its done without issues
					return pager.MakeTable("sessions", table, s).TextureTable()
				},
			},
			{
				SubcommandName: "admin",
				Description: "view sessions with admin",
				CommandPermissions: []string{"admin"},
				CommandSplit: " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//creates the simpletable
					//this will store our information
					table := simpletable.New() //makes the structure

					//sets table header
					//this will be used to define information
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "id.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "connected.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ip.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "idle.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ranks.txt").Containing, lexer.Escapes)},
						},
					}
				
				
					//ranges through every opens ession
					//this will make sure its done without issues
					for _, usr := range sessions.Sessions {
						//creates the rank information
						//this will ensure its done without issues
						r := ranks.MakeRank(usr.User.Username) //makes the rank
						//tries to correctly sync the render without issues happening
						//this will make sure its done without issues happening on request
						if err := r.SyncWithString(usr.User.Ranks); err != nil {
							return err //returns the error correctly and properly
						}
					
						if !r.CanAccess(cmd[1]) {
							continue
						}
					
					
						//properly deploys the rank into string
						//this will ensure its done without errors happening
						ranks, err := r.DeployRanks(true) //deploys into string format
						if err != nil { //error handles the syntax properly without issues
							return err //returns the error properly
						}
					
						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.User.Identity))}, //id
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", usr.User.Username)}, //username
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "connected.txt").Containing, lexer.Escapes), "<<$connected>>", usr.Connected.Format("15:04:05"))}, //maxtime
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", usr.Target)}, //conns
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "idle.txt").Containing, lexer.Escapes), "<<$idle>>", tools.ResolveTimeStamp(usr.Idle, false))}, //cooldown
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))}, //ranks
						}
					
						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}
				
					//renders the table properly
					//this will ensure its done without issues
					return pager.MakeTable("sessions", table, s).TextureTable()
				},
			},
			{
				SubcommandName: "moderator",
				Description: "view sessions with moderator",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit: " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//creates the simpletable
					//this will store our information
					table := simpletable.New() //makes the structure
						
					//sets table header
					//this will be used to define information
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "id.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "connected.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ip.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "idle.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("sessions", "headers", "ranks.txt").Containing, lexer.Escapes)},
						},
					}
				
				
					//ranges through every opens ession
					//this will make sure its done without issues
					for _, usr := range sessions.Sessions {
						//creates the rank information
						//this will ensure its done without issues
						r := ranks.MakeRank(usr.User.Username) //makes the rank
						//tries to correctly sync the render without issues happening
						//this will make sure its done without issues happening on request
						if err := r.SyncWithString(usr.User.Ranks); err != nil {
							return err //returns the error correctly and properly
						}
					
						if !r.CanAccess(cmd[1]) {
							continue
						}
					
					
						//properly deploys the rank into string
						//this will ensure its done without errors happening
						ranks, err := r.DeployRanks(true) //deploys into string format
						if err != nil { //error handles the syntax properly without issues
							return err //returns the error properly
						}
					
						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(usr.User.Identity))}, //id
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", usr.User.Username)}, //username
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "connected.txt").Containing, lexer.Escapes), "<<$connected>>", usr.Connected.Format("15:04:05"))}, //maxtime
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", usr.Target)}, //conns
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "idle.txt").Containing, lexer.Escapes), "<<$idle>>", tools.ResolveTimeStamp(usr.Idle, false))}, //cooldown
							{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("sessions", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks, " "))}, //ranks
						}
					
						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}
				
					//renders the table properly
					//this will ensure its done without issues
					return pager.MakeTable("sessions", table, s).TextureTable()
				},
			},
			{
				SubcommandName: "message", //message an active session
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				Description: "message active sessions",
				CommandSplit: " ", //no split to be active within commands
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//tries to validate the length
					//this will ensure its done without any errors
					if len(cmd) < 4 { //checks the length properly
						return language.ExecuteLanguage([]string{"sessions", "message", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//stores all the targeted sessions
					//this will be what we broadcast to correctly
					var operants []sessions.Session = make([]sessions.Session, 0)
					var messages []string           = make([]string, 0) //message

					//tries to build the message information
					//this will be what we broadcast correctly
					for pos := range cmd[2:] { //ranges through all args
						//checks for a syntax highlight properly
						//this will ensure its done without any errors
						if strings.Contains(cmd[2:][pos], ":") { //checks for user syntax
							//checks for a individual session detection
							//this will collect only that session properly
							if strings.Split(cmd[2:][pos], ":")[0] == "" { //checks
								//inserts all of the users sessions properly without any issues happening
								operants = append(operants, s.GetSessions(strings.ReplaceAll(cmd[2:][pos], ":", ""))...); continue
							}

							//converts the id properly without issues
							//this will try to convert without any errors
							id, err := strconv.Atoi(strings.Split(cmd[2:][pos], ":")[1])
							if err != nil { //error handles properly and makes sure its done without any errors
								return language.ExecuteLanguage([]string{"sessions", "message", "invalid_id.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
							}

							//tries to find with the id given
							//this will ensure its done without any errors
							session := s.GetWithID(int64(id)) //tries to find the session
							if session == nil || session.User.Username != strings.Split(cmd[2:][pos], ":")[0] { //renders the branding peice properly
								return language.ExecuteLanguage([]string{"sessions", "message", "mismatch_id_username.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
							}

							//saves into the array properly
							//this will ensure its done without any errors happening
							operants = append(operants, *session);continue
						}
						//saves an as argument properly
						//this will ensure its done without any errors happening
						messages = append(messages, cmd[2:][pos]); continue
					}

					//ranges through all the sessions and broadcasts the message
					//this will ensure its done without any errors happening on purpose
					for active := range operants { //ranges through without any errors
						//tries to launch the message without any errors happening properly
						//this will ensure its done without any errors happening on purpose making it safer
						err := language.ExecuteLanguage([]string{"alerts", "session-message.itl"}, operants[active].Channel, deployment.Engine, &operants[active], map[string]string{"message":strings.Join(messages, " ")})
						if err != nil { //error handles the alert properly
							//this will try to launch the broadcast safely without errors
							operants[active].Write("Message>", strings.Join(messages, " ")) //writes
						}; continue //continues looping without any errors
					}
					return nil
				},
				//AutoComplete allows for statement auto callback
				AutoComplete: func(s *sessions.Session) []string {
	
					//stores all the storage properly
					//this will ensure its done properly
					var storage []string = make([]string, 0)

					for _, open := range sessions.Sessions { //ranges through and saves properly

						//makes sure it is not our session properly
						//this will allow it to not be shown within the tab
						if open.User.Username == s.User.Username && s.ID == open.ID {
							continue
						}

						storage = append(storage, open.User.Username + ":" + strconv.Itoa(int(open.ID)))
						
						//checks for the main system properly and safely
						if !tools.NeedleHaystackOne(storage, ":" + open.User.Username) {
							storage = append(storage, ":" + open.User.Username)
						}
					}

					//returns the modules properly
					//this will allow it to be used within it
					return storage
				},
			},
			{
				SubcommandName: "maximise",
				Description: "Expand a terminal window",
				CommandPermissions: []string{"admin", "moderator", "sessions"},
				CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//tries to validate the length
					//this will ensure its done without any errors
					if len(cmd) < 3 { //checks the length properly
						return language.ExecuteLanguage([]string{"sessions", "maximise", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//stores all the targeted sessions
					//this will be what we broadcast to correctly
					var operants []sessions.Session = make([]sessions.Session, 0)

					//tries to build the message information
					//this will be what we broadcast correctly
					for pos := range cmd[2:] { //ranges through all args
						//checks for a syntax highlight properly
						//this will ensure its done without any errors
						if strings.Contains(cmd[2:][pos], ":") { //checks for user syntax
							//checks for a individual session detection
							//this will collect only that session properly
							if strings.Split(cmd[2:][pos], ":")[0] == "" { //checks
								//inserts all of the users sessions properly without any issues happening
								operants = append(operants, s.GetSessions(strings.ReplaceAll(cmd[2:][pos], ":", ""))...); continue
							}

							//converts the id properly without issues
							//this will try to convert without any errors
							id, err := strconv.Atoi(strings.Split(cmd[2:][pos], ":")[1])
							if err != nil { //error handles properly and makes sure its done without any errors
								return language.ExecuteLanguage([]string{"sessions", "maximise", "invalid_id.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
							}

							//tries to find with the id given
							//this will ensure its done without any errors
							session := s.GetWithID(int64(id)) //tries to find the session
							if session == nil || session.User.Username != strings.Split(cmd[2:][pos], ":")[0] { //renders the branding peice properly
								return language.ExecuteLanguage([]string{"sessions", "maximise", "mismatch_id_username.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
							}

							//saves into the array properly
							//this will ensure its done without any errors happening
							operants = append(operants, *session);continue
						}
					
						return language.ExecuteLanguage([]string{"sessions", "maximise", "invalid_id.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the sessions and broadcasts the message
					//this will ensure its done without any errors happening on purpose
					for active := range operants { //ranges through without any errors

						//this will properly maximse the session
						//ensures its done properly and safely within it
						operants[active].Write("\x1b[9;2;0t") //maxises


						continue //continues looping without any errors
					}
					return nil
				},
				//AutoComplete allows for statement auto callback
				AutoComplete: func(s *sessions.Session) []string {
	
					//stores all the storage properly
					//this will ensure its done properly
					var storage []string = make([]string, 0)

					for _, open := range sessions.Sessions { //ranges through and saves properly

						//makes sure it is not our session properly
						//this will allow it to not be shown within the tab
						if open.User.Username == s.User.Username && s.ID == open.ID {
							continue
						}

						storage = append(storage, open.User.Username + ":" + strconv.Itoa(int(open.ID)))
						
						//checks for the main system properly and safely
						if !tools.NeedleHaystackOne(storage, ":" + open.User.Username) {
							storage = append(storage, ":" + open.User.Username)
						}
					}

					//returns the modules properly
					//this will allow it to be used within it
					return storage
				},
			},
			{
				SubcommandName: "kick",
				Description: "kick an active session",
				CommandPermissions: []string{"admin", "moderator", "reseller"},
				CommandSplit: " ", SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//tries to validate the length
					//this will ensure its done without any errors
					if len(cmd) < 3 { //checks the length properly
						return language.ExecuteLanguage([]string{"sessions", "kick", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//stores all the targeted sessions
					//this will be what we broadcast to correctly
					var operants []sessions.Session = make([]sessions.Session, 0)

					//tries to build the message information
					//this will be what we broadcast correctly
					for pos := range cmd[2:] { //ranges through all args
						//checks for a syntax highlight properly
						//this will ensure its done without any errors
						if strings.Contains(cmd[2:][pos], ":") { //checks for user syntax
							//checks for a individual session detection
							//this will collect only that session properly
							if strings.Split(cmd[2:][pos], ":")[0] == "" { //checks
								//inserts all of the users sessions properly without any issues happening
								operants = append(operants, s.GetSessions(strings.ReplaceAll(cmd[2:][pos], ":", ""))...); continue
							}

							//converts the id properly without issues
							//this will try to convert without any errors
							id, err := strconv.Atoi(strings.Split(cmd[2:][pos], ":")[1])
							if err != nil { //error handles properly and makes sure its done without any errors
								return language.ExecuteLanguage([]string{"sessions", "kick", "invalid_id.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
							}

							//tries to find with the id given
							//this will ensure its done without any errors
							session := s.GetWithID(int64(id)) //tries to find the session
							if session == nil || session.User.Username != strings.Split(cmd[2:][pos], ":")[0] { //renders the branding peice properly
								return language.ExecuteLanguage([]string{"sessions", "kick", "mismatch_id_username.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
							}

							//saves into the array properly
							//this will ensure its done without any errors happening
							operants = append(operants, *session);continue
						}
					
						return language.ExecuteLanguage([]string{"sessions", "kick", "invalid_id.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all the sessions and broadcasts the message
					//this will ensure its done without any errors happening on purpose
					for active := range operants { //ranges through without any errors
						operants[active].Channel.Close() //tries to close the channel
						language.ExecuteLanguage([]string{"sessions", "kick", "kicked.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":operants[active].User.Username})
						continue //continues looping without any errors
					}
					return nil
				},
				//AutoComplete allows for statement auto callback
				AutoComplete: func(s *sessions.Session) []string {
					
					//stores all the storage properly
					//this will ensure its done properly
					var storage []string = make([]string, 0)

					for _, open := range sessions.Sessions { //ranges through and saves properly

						//makes sure it is not our session properly
						//this will allow it to not be shown within the tab
						if open.User.Username == s.User.Username && s.ID == open.ID {
							continue
						}

						storage = append(storage, open.User.Username + ":" + strconv.Itoa(int(open.ID)))
						
						//checks for the main system properly and safely
						if !tools.NeedleHaystackOne(storage, ":" + open.User.Username) {
							storage = append(storage, ":" + open.User.Username)
						}
					}

					//returns the modules properly
					//this will allow it to be used within it
					return storage
				},
			},
			{
				SubcommandName: "observe",
				Description: "observe a sessions activity",
				CommandPermissions: []string{"admin", "moderator"}, CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//detects invalid syntax within the string
					if len(cmd) != 3 || !strings.Contains(cmd[2], ":"){
						return language.ExecuteLanguage([]string{"sessions", "observe", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//converts the id into int type
					convert, err := strconv.Atoi(strings.Join(strings.Split(cmd[2], ":")[1:], ":"))
					if err != nil {
						return language.ExecuteLanguage([]string{"sessions", "observe", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets session
					target := sessions.Sessions[int64(convert)]
					if target == nil || target.User.Username != strings.Split(cmd[2], ":")[0] { //prints the unknown user message
						return language.ExecuteLanguage([]string{"sessions", "observe", "unknown_username.itl"}, s.Channel, deployment.Engine, s, map[string]string{"username":strings.Split(cmd[2], ":")[0], "id":strings.Split(cmd[2], ":")[1]})
					}

					language.ExecuteLanguage([]string{"alerts", "being_observed.itl"}, target.Channel, deployment.Engine, target, map[string]string{"observer":s.User.Username})

					before := s.Capture() //captures current window size properly

					var close bool = false
					go func() { //routine for detecting exit
						buf := make([]byte, 1) //stores buffer
						s.Channel.Read(buf) //reads from term
						close=true
					}()

					var clicks int = 0 //current length of system

					target.Viewers++
					for {
						if close { //detect close
							target.Viewers--
							break //ends
						}

						//checks for new frames inside the system properly
						if target == nil || clicks != len(target.Written) && strings.Contains(target.Written[len(target.Written)-1], "\033]0;") {
							continue
						}

						clicks = len(target.Written) //updates length stops frame skipping

						//writes to the terminal
						s.Write("\033c"+target.Capture(), "\x1b[0;0f\x1b[101;30m"+Centre("Currently observing "+target.User.Username+", as "+s.User.Username, s.Length, "")+"\x1b[0m") //writes source
						//time.Sleep(500 * time.Millisecond) //frame buffer on the system
					}	

					return s.Write("\033c"+before)
				},

				//AutoComplete allows for statement auto callback
				AutoComplete: func(s *sessions.Session) []string {
					//stores all the storage properly
					//this will ensure its done properly
					var storage []string = make([]string, 0)

					//ranges through the sessions properly
					for _, open := range sessions.Sessions {
						storage = append(storage, open.User.Username + ":" + strconv.Itoa(int(open.ID)))
					}

					//returns the modules properly
					//this will allow it to be used within it
					return storage
				},
			},
		},

	})
}