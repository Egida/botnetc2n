package commands

import (
	attacks "Nosviak2/core/attack"
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "allm",
		Aliases:            []string{"attkss", "bkg9j"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "view all the possibly attack methods",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "methods", "headers", "name.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "methods", "headers", "description.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "methods", "headers", "ranks.txt").Containing, lexer.Escapes)},
				},
			}

			//stores all captured methods properly
			var capture []string = make([]string, 0)

			//ranges through all the methods properly
			for _, method := range attacks.AllMethods(make([]*attacks.Method, 0)) {
				capture = append(capture, method.Name)
			}

			//orders in alpha order
			sort.Strings(capture)

			//ranges through all the methods properly
			//this will register all the methods properly without issues
			for _, captured := range capture { //ranges through all api methods
				method := attacks.QueryMethod(captured)
				if method == nil {
					continue
				}

				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "methods", "values", "name.txt").Containing, lexer.Escapes), "<<$name>>", method.Name)},                                                 //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "methods", "values", "description.txt").Containing, lexer.Escapes), "<<$description>>", method.Description)},                            //username
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "methods", "values", "ranks.txt").Containing, lexer.Escapes), "<<$ranks>>", strings.Join(ranks.CreateSystem(method.Permissions), " "))}, //maxtime
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("methods", table, s).TextureTable()
		},

		SubCommands: []SubCommand{
			{
				SubcommandName:     "launched",
				Description:        "view all floods launched from all users",
				CommandPermissions: []string{"admin", "moderator"}, CommandSplit: " ",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					//ranges through all floods properly
					//this will try to grab them from the database
					floods, err := database.Conn.GlobalSent() //all sent
					if err != nil {                           //err handles properly and safely
						return language.ExecuteLanguage([]string{"attacks", "launched", "database-error.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//makes the simpletable properly
					//used within the system properly
					table := simpletable.New()

					//sets table header
					//this will be used to define information
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "date.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "method.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "target.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "duration.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "finished.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through all attacks launched
					//this will allow it to be inserted within it
					for _, systemLaunched := range floods { //ranges

						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "date.txt").Containing, lexer.Escapes), "<<$date>>", time.Unix(systemLaunched.Created, 0).Format("2/01/2006 15:04:05"))},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", systemLaunched.Username)},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "method.txt").Containing, lexer.Escapes), "<<$method>>", systemLaunched.Method)},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "target.txt").Containing, lexer.Escapes), "<<$target>>", Format(systemLaunched.Target))},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "duration.txt").Containing, lexer.Escapes), "<<$duration>>", strconv.Itoa(systemLaunched.Duration))},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "finished.txt").Containing, lexer.Escapes), "<<$finished>>", tools.ResolveTimeStamp(time.Unix(systemLaunched.Finish, 0), false))},
						}

						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}

					//renders the table properly
					//this will ensure its done without issues
					return pager.MakeTable("methods", table, s).TextureTable()
				},
			},

			{
				SubcommandName:     "sent=",
				Description:        "amount of attacks sent from usr",
				CommandPermissions: []string{"admin", "moderator"}, CommandSplit: "=",
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {

					if strings.Count(cmd[1], "=") <= 0 { //invalid syntax detection here properly
						return language.ExecuteLanguage([]string{"attacks", "sent", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//ranges through all floods properly
					//this will try to grab them from the database
					floods, err := database.Conn.UserSent(strings.Split(cmd[1], "=")[1]) //all sent
					if err != nil {                                                      //err handles properly and safely
						return language.ExecuteLanguage([]string{"attacks", "launched", "database-error.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					//makes the simpletable properly
					//used within the system properly
					table := simpletable.New()

					//sets table header
					//this will be used to define information
					table.Header = &simpletable.Header{
						Cells: []*simpletable.Cell{ //sets the values properly without issues happening
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "date.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "username.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "method.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "target.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "duration.txt").Containing, lexer.Escapes)},
							{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("attacks", "launched", "headers", "finished.txt").Containing, lexer.Escapes)},
						},
					}

					//ranges through all attacks launched
					//this will allow it to be inserted within it
					for _, systemLaunched := range floods { //ranges

						//creates the store properly without issues
						rk := []*simpletable.Cell{ //fills with the information properly without issues
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "date.txt").Containing, lexer.Escapes), "<<$date>>", time.Unix(systemLaunched.Created, 0).Format("2/01/2006 15:04:05"))},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", systemLaunched.Username)},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "method.txt").Containing, lexer.Escapes), "<<$method>>", systemLaunched.Method)},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "target.txt").Containing, lexer.Escapes), "<<$target>>", Format(systemLaunched.Target))},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "duration.txt").Containing, lexer.Escapes), "<<$duration>>", strconv.Itoa(systemLaunched.Duration))},
							{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("attacks", "launched", "values", "finished.txt").Containing, lexer.Escapes), "<<$finished>>", tools.ResolveTimeStamp(time.Unix(systemLaunched.Finish, 0), false))},
						}

						//saves into the array correctly
						//this will be properly founded later onwards
						table.Body.Cells = append(table.Body.Cells, rk)
					}

					//renders the table properly
					//this will ensure its done without issues
					return pager.MakeTable("methods", table, s).TextureTable()
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

			{
				SubcommandName:     "enable",
				Description:        "enables floods across the global platform",
				CommandPermissions: []string{"admin"}, CommandSplit: " ", //admin only control
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//checks if they are already enabled properly
					//this will enable the attacks properly without issues
					if attacks.AttacksEnabled { //message if they are already enabled
						return language.ExecuteLanguage([]string{"attacks", "enable", "already-enabled.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					attacks.AttacksEnabled = true //enables
					//lets the system know they have enabled attacks
					//this will ensure its done without issues happening
					return language.ExecuteLanguage([]string{"attacks", "enable", "success.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
				},
			},

			{
				SubcommandName:     "enable-api",
				Description:        "enables floods across the api platform",
				CommandPermissions: []string{"admin"}, CommandSplit: " ", //admin only control
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//checks if they are already enabled properly
					//this will enable the attacks properly without issues
					if attacks.AttacksEnabled { //message if they are already enabled
						return language.ExecuteLanguage([]string{"attacks", "enable-api", "already-enabled.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					attacks.APIAttacksEnabled = true //enables
					//lets the system know they have enabled attacks
					//this will ensure its done without issues happening
					return language.ExecuteLanguage([]string{"attacks", "enable-api", "success.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
				},
			},

			{
				SubcommandName:     "disable",
				Description:        "disable floods across the global platform",
				CommandPermissions: []string{"admin"}, CommandSplit: " ", //admin only control
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//checks if they are already disbled properly
					//this will ensure its done without issues happening
					if !attacks.AttacksEnabled { //message if they are already enabled
						return language.ExecuteLanguage([]string{"attacks", "disable", "already-disable.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					attacks.AttacksEnabled = false //enables
					//lets the system know they have enabled attacks
					//this will ensure its done without issues happening
					return language.ExecuteLanguage([]string{"attacks", "disable", "success.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
				},
			},

			{
				SubcommandName:     "disable-api",
				Description:        "disable floods across the api platform",
				CommandPermissions: []string{"admin"}, CommandSplit: " ", //admin only control
				SubCommandFunction: func(s *sessions.Session, cmd []string) error {
					//checks if they are already disbled properly
					//this will ensure its done without issues happening
					if !attacks.AttacksEnabled { //message if they are already enabled
						return language.ExecuteLanguage([]string{"attacks", "disable-api", "already-disable.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
					}

					attacks.APIAttacksEnabled = false //enables
					//lets the system know they have enabled attacks
					//this will ensure its done without issues happening
					return language.ExecuteLanguage([]string{"attacks", "disable-api", "success.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
				},
			},
		},
	})
}

// formats the target properly
// this will ensure its done without errors
func Format(targ string) string { //string returned
	if utf8.RuneCountInString(targ) <= 10 {
		return targ
	}

	return strings.Join(strings.Split(targ, "")[:9], "")
}
