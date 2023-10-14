package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/views"
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "themes",
		Aliases:            []string{"theme"},
		CommandPermissions: []string{"admin", "moderator"},
		CommandDescription: "control your users theme properly",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "themes", "headers", "name.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "themes", "headers", "description.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "themes", "headers", "colours.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "themes", "headers", "ranks.txt").Containing, lexer.Escapes)},
				},
			}

			//ranges through all the themes
			//this will make sure its done without issues
			for name, theme := range toml.ThemeConfig.Theme { //ranges through all the themes
				//checks if hidden or disabled theme
				//this will ensure its done without any errors
				if theme.Hidden || !theme.Enabled {
					continue
				} //continues looping properly

				var colours []string = make([]string, 0)
				for _, system := range theme.Decor.Colours {
					R, G, B := system[0], system[1], system[2]
					colours = append(colours, fmt.Sprintf("\x1b[0m\x1b[48;2;%d;%d;%dm \x1b[0m", R, G, B))
				}

				//creates the store properly without issues and makes sure they can view it without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues and errors happening
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "themes", "values", "name.txt").Containing, lexer.Escapes), "<<$name>>", name)},                                                     //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "themes", "values", "description.txt").Containing, lexer.Escapes), "<<$description>>", theme.Description)},                          //username
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "themes", "values", "colours.txt").Containing, lexer.Escapes), "<<$colours>>", strings.Join(colours, ""))},                          //username
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "themes", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks.CreateSystem(theme.Permissions), " "))}, //maxtime
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("themes", table, s).TextureTable()
		},

		SubCommands: []SubCommand{
			{
				SubcommandName:     "change",
				Description:        "change your current theme properly",
				CommandPermissions: make([]string, 0), CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//checks the length
					//ensures its done without any errors
					if len(cmd) <= 2 { //length checks properly and displays the syntax issue if needed
						return language.ExecuteLanguage([]string{"commands", "themes", "change", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//gets the string for the wanted theme
					//this will ensure its done without any errors
					future := cmd[2] //gets the cmd properly without issues

					//gets the theme properly
					//this will be what we change to properly
					them := toml.ThemeConfig.Theme[strings.ToLower(future)] //lowers string
					if them == nil {                                        //checks if the theme was found properly
						return language.ExecuteLanguage([]string{"commands", "themes", "change", "invalidtheme.itl"}, s.Channel, deployment.Engine, s, map[string]string{"theme": cmd[2]})
					}

					//adds support for the rank system inside the settings
					//this will get all the ranks and sync them without issues and secures
					rsystem := ranks.MakeRank(s.User.Username)
					rsystem.SyncWithString(s.User.Ranks)

					//checks if they can access the array options
					//this will add support for the permissions based system
					if rsystem.CanAccessArray(them.Permissions) { //checks permissions and ensures its safe without issues
						return language.ExecuteLanguage([]string{"commands", "themes", "change", "permissionsFault.itl"}, s.Channel, deployment.Engine, s, map[string]string{"theme": cmd[2]})
					}

					//tries to update the theme value properly
					//this will allow the system to be updated without errors
					if err := database.Conn.Theme(s.User.Username, future); err != nil { //error handles
						return language.ExecuteLanguage([]string{"commands", "themes", "change", "fault-theme.itl"}, s.Channel, deployment.Engine, s, map[string]string{"theme": cmd[2]})
					}

					//checks if we are updating the sessions
					//this will allow for all sessions to be updated
					if them.UpdateSesss { //checks if its valid without issues
						//updates inside all the open sessions
						//this will ensure its done without any errors
						s.FunctionRemote(s.User.Username, func(t *sessions.Session) {
							sessions.Sessions[t.ID].User.Theme = future                              //updates the central database aswell properly
							sessions.Sessions[t.ID].BrandingPath = strings.Split(them.Branding, "/") //updates the theme properly
							//checks for the default theme properly
							//this will ensure we can reset the colours
							if strings.ToLower(future) == "default" { //sets back to default colours
								sessions.Sessions[t.ID].Colours = toml.DecorationToml.Gradient.Colours
							} else { //sets the decor colours properly
								//this will ensure its done without any errors
								sessions.Sessions[t.ID].Colours = them.Decor.Colours
							}
						}) ///updates and properly sets the system data without issues
					}

					//returns the success message properly
					//this will ensure its done without any errors
					return language.ExecuteLanguage([]string{"commands", "themes", "change", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"theme": cmd[2]})
				},

				//AutoComplete will allow for extra args within the command to be completed
				AutoComplete: func(s *sessions.Session) []string {

					var capture []string = make([]string, 0)

					//ranges through all themes
					//we will only ignore the current theme
					for name, _ := range toml.ThemeConfig.Theme {
						if name == s.User.Theme { //ignores default
							continue //continues looping
						}

						//saves into array
						capture = append(capture, name)
					}

					return capture
				},
			},
		},
	})
}
