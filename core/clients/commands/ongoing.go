package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	"Nosviak2/core/database"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/views"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "ongoing",
		Aliases:            []string{"running"},
		CommandPermissions: []string{"admin", "moderator"},
		CommandDescription: "display all global ongoing attacks",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "id.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "username.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "method.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "target.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "port.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "length.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "ongoing", "headers", "finish.txt").Containing, lexer.Escapes)},
				},
			}

			//gets all of the users attacks currently
			//this will ensure its done without errors happening
			Attacks, err := database.Conn.GlobalRunning()
			if err != nil || len(Attacks) <= 0 { //error handles properly without issues
				return language.ExecuteLanguage([]string{"commands", "ongoing", "no-running.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//ranges through every opens ession
			//this will make sure its done without issues
			for _, attack := range Attacks { //ranges through all users sessions
				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "id.txt").Containing, lexer.Escapes), "<<$id>>", strconv.Itoa(attack.ID))}, 
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", attack.Username)},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "method.txt").Containing, lexer.Escapes), "<<$method>>", attack.Method)},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "target.txt").Containing, lexer.Escapes), "<<$target>>", attack.Target)},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "port.txt").Containing, lexer.Escapes), "<<$port>>", strconv.Itoa(attack.Port))},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "length.txt").Containing, lexer.Escapes), "<<$length>>", strconv.Itoa(attack.Duration))},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "ongoing", "values", "finish.txt").Containing, lexer.Escapes), "<<$finish>>", fmt.Sprintf("%.2f", time.Until(time.Unix(attack.Finish, 64)).Seconds()))},
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("myrunning", table, s).TextureTable()
		},
	})
}