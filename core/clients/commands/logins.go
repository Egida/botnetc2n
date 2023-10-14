package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/views"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "logins",
		Aliases:            []string{"login"},
		CommandPermissions: []string{"admin", "moderator"},
		CommandDescription: "display all login attempts",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//tries to get the logins properly
			//this will ensure its done without issues happening
			logins, err := database.Conn.GetLogins(s.User.Username) //tries to get the logins
			if err != nil {                                         //error handles the statement properly without issues
				return language.ExecuteLanguage([]string{"commands", "logins", "databaseFault.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//creates the new simpletable information
			//this will ensure its done properly without errors
			table := simpletable.New()

			//assigns the headers properly
			//these will be used inside the table without issues
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "headers", "username.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "headers", "ip.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "headers", "banner.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "headers", "date.txt").Containing, lexer.Escapes)},
				},
			}

			//ranges through all the logins
			//these will help within the systems controller
			for login := range logins { //ranges through the logins
				row := []*simpletable.Cell{ //stores the row information
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", logins[login].Username)},                               //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", strings.Split(logins[login].Address, ":")[0])},                     //username
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "values", "banner.txt").Containing, lexer.Escapes), "<<$banner>>", logins[login].Banner)},                                     //maxtime
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "values", "date.txt").Containing, lexer.Escapes), "<<$date>>", time.Unix(logins[login].TimeStore, 0).Format("02 Jan 15:04"))}, //conns
				}

				//saves into the information properly
				//this will just store the information without errors
				table.Body.Cells = append(table.Body.Cells, row)
			}

			//renders the pager correctly
			//this will make sure its done without errors
			return pager.MakeTable("logins", table, s).TextureTable()
		},

		SubCommands: []SubCommand{
			{
				SubcommandName:     "all",
				Description:        "view all logins which have occurred",
				CommandPermissions: []string{"admin"}, CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//tries to get the logins properly
					//this will ensure its done without issues happening
					logins, err := database.Conn.AllLogins() //tries to get the logins
					if err != nil {                          //error handles the statement properly without issues
						return language.ExecuteLanguage([]string{"commands", "logins", "all", "databaseFault.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//creates the new simpletable information
					//this will ensure its done properly without errors
					table := simpletable.New()

					//assigns the headers properly
					//these will be used inside the table without issues
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "all", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "all", "headers", "ip.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "all", "headers", "banner.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "all", "headers", "date.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through all the logins
					//these will help within the systems controller
					for login := range logins { //ranges through the logins
						row := []*simpletable.Cell{ //stores the row information
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "all", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", logins[login].Username)},                               //id
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "all", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", strings.Split(logins[login].Address, ":")[0])},                     //username
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "all", "values", "banner.txt").Containing, lexer.Escapes), "<<$banner>>", logins[login].Banner)},                                     //maxtime
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "all", "values", "date.txt").Containing, lexer.Escapes), "<<$date>>", time.Unix(logins[login].TimeStore, 0).Format("02 Jan 15:04"))}, //conns
						}

						//saves into the information properly
						//this will just store the information without errors
						table.Body.Cells = append(table.Body.Cells, row)
					}

					//renders the pager correctly
					//this will make sure its done without errors
					return pager.MakeTable("logins", table, s).TextureTable()
				},
			},
			{
				SubcommandName:     "user=",
				Description:        "view logins for a certain user",
				CommandPermissions: []string{"admin", "moderator"},
				CommandSplit:       "=", SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					if strings.Count(cmd[1], "=") <= 0 { //invalid syntax detection here properly
						return language.ExecuteLanguage([]string{"commands", "logins", "user", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//tries to get the logins properly
					//this will ensure its done without issues happening
					logins, err := database.Conn.GetLogins(strings.Split(cmd[1], "=")[1]) //tries to get the logins
					if err != nil {                                                       //error handles the statement properly without issues
						return language.ExecuteLanguage([]string{"commands", "logins", "user", "databaseFault.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//creates the new simpletable information
					//this will ensure its done properly without errors
					table := simpletable.New()

					//assigns the headers properly
					//these will be used inside the table without issues
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "user", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "user", "headers", "ip.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "user", "headers", "banner.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "logins", "user", "headers", "date.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through all the logins
					//these will help within the systems controller
					for login := range logins { //ranges through the logins
						row := []*simpletable.Cell{ //stores the row information
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "user", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", logins[login].Username)},                               //id
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "user", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", strings.Split(logins[login].Address, ":")[0])},                     //username
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "user", "values", "banner.txt").Containing, lexer.Escapes), "<<$banner>>", logins[login].Banner)},                                     //maxtime
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "logins", "user", "values", "date.txt").Containing, lexer.Escapes), "<<$date>>", time.Unix(logins[login].TimeStore, 0).Format("02 Jan 15:04"))}, //conns
						}

						//saves into the information properly
						//this will just store the information without errors
						table.Body.Cells = append(table.Body.Cells, row)
					}

					//renders the pager correctly
					//this will make sure its done without errors
					return pager.MakeTable("logins", table, s).TextureTable()
				},
				//completes the callback for the system properly
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
