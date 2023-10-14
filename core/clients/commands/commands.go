package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/views"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "command",
		Aliases:            []string{"cmds", "help"},
		CommandDescription: "display information about commands",
		CommandPermissions: []string{"admin", "moderator"},
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-id.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-name.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-description.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-ranks.txt").Containing, lexer.Escapes)},
				},
			}

			//this will just render the back information without issues
			var spins int = 0
			//ranges through all the commands
			//this will register into the structure properly
			for _, cmd := range Commands { //ranges through
				spins++ //spins through the rotations without errors happening
				//this will ensure its done without error happening on request etc
				//creates the store properly without issues were issues shouldnt happen
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(spins))},                                               //id
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-name.txt").Containing, lexer.Escapes), "<<$name>>", cmd.CommandName)},                                                 //username
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-description.txt").Containing, lexer.Escapes), "<<$description>>", cmd.CommandDescription)},                            //maxtime
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks.CreateSystem(cmd.CommandPermissions), " "))}, //conns
				}
				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//creates the pager correctly
			//this will ensure its done without errors happening
			return pager.MakeTable("commands", table, s).TextureTable()
		},

		//invalid subcommand detection properly
		//this will ensure its done without errors happening
		InvalidSubCommand: func(s *sessions.Session, cmd []string) error {

			target := TryCommand(strings.ToLower(cmd[1])) //tries to get subcommand
			if target == nil {                            //invalid subcommand detection properly here
				return language.ExecuteLanguage([]string{"commands", "commands", "unclassified.itl"}, s.Channel, deployment.Engine, s, map[string]string{"command": cmd[1]})
			}

			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-id.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-name.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-description.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "commands", "headers", "header-ranks.txt").Containing, lexer.Escapes)},
				},
			}

			//this will just render the back information without issues
			var spins int = 0
			//ranges through all the commands
			//this will register into the structure properly
			for _, cmd := range target.SubCommands { //ranges through
				spins++ //spins through the rotations without errors happening
				//this will ensure its done without error happening on request etc
				//creates the store properly without issues were issues shouldnt happen
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(spins))},                                               //id
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-name.txt").Containing, lexer.Escapes), "<<$name>>", cmd.SubcommandName)},                                              //username
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-description.txt").Containing, lexer.Escapes), "<<$description>>", cmd.Description)},                                   //maxtime
					{Align: simpletable.AlignLeft, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "commands", "values", "value-ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks.CreateSystem(cmd.CommandPermissions), " "))}, //conns
				}
				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//creates the pager correctly
			//this will ensure its done without errors happening
			return pager.MakeTable("commands", table, s).TextureTable()
		},

		SubCommands: []SubCommand{

			{
				SubcommandName:     "describe",
				Description:        "displays information about the command path given",
				CommandPermissions: make([]string, 0), CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//invalid syntax pointer
					//this will ensure its done
					if len(cmd) < 3 { //length check the object
						return language.ExecuteLanguage([]string{"commands", "commands", "describe", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to fetch the command properly
					var targetObject *Command = TryCommand(cmd[2])
					if targetObject == nil { //invalid pointer detection
						return language.ExecuteLanguage([]string{"commands", "commands", "describe", "unknown_command.itl"}, s.Channel, deployment.Engine, s, map[string]string{"command": cmd[2]})
					}

					//objective checker properly
					if len(cmd) > 3 { //properly enforces it
						objective := targetObject.FindSubs(cmd[3]) //checks properly
						if objective == nil {                      //properly checks without issues
							return language.ExecuteLanguage([]string{"commands", "commands", "describe", "unknown_subcommand.itl"}, s.Channel, deployment.Engine, s, map[string]string{"subcommand": cmd[3]})
						}

						//stores the object information
						targetObject = &Command{CommandName: objective.SubcommandName, CommandDescription: objective.Description, CommandPermissions: objective.CommandPermissions}
					}

					//writes the information properly
					//this will be used within the system properly
					return s.Write(fmt.Sprintf("Name: %s\r\nDescription: %s\r\nPermissions: %s\r\n", targetObject.CommandName, targetObject.CommandDescription, strings.Join(ranks.CreateSystem(targetObject.CommandPermissions), " ")))
				},
				AutoComplete: func(s *sessions.Session) []string {
					return ToSystemArray(make([]string, 0))
				},
			},
		},
	})
}

// converts all command into array
// allows us to properly handle without issues
func ToSystemArray(src []string) []string {
	for Command := range Commands {
		src = append(src, Command)
	}
	return src
}
