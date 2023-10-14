package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/views"
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {
	MakeCommand(&Command{
		CommandName:        "ranks",
		Aliases:            []string{"rank"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "view all active ranks",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ranks", "headers", "name.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ranks", "headers", "description.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ranks", "headers", "symbol.txt").Containing, lexer.Escapes)},
				},
			}
			
			// ranges through all ranks inside the toml
			for name, rank := range ranks.PresetRanks { 

				rk := []*simpletable.Cell{ 
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ranks", "values", "name.txt").Containing, lexer.Escapes), "<<$name>>", name)}, //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ranks", "values", "description.txt").Containing, lexer.Escapes), "<<$description>>", rank.RankDescription)}, //username
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ranks", "values", "symbol.txt").Containing, lexer.Escapes), "<<$symbol>>", fmt.Sprintf("\x1b[0m\x1b[%sm\x1b[%sm %s \x1b[0m", strings.Join(ranks.Convert(rank.MainColour), ";"), strings.Join(ranks.Convert(rank.SecondColour), ";"), rank.SignatureCharater))}, //maxtime
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("ranks", table, s).TextureTable()
		},
	})
}