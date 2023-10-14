package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/clients/views/pager"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/views"
	"strconv"
	"strings"

	"github.com/alexeyco/simpletable"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "who",
		Aliases:            []string{"whoami", "me"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "sessions opened in your username",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			
			var target string = s.User.Username

			//checks for special targeted usernames
			//this will allow us to properly view the users
			if len(cmd) >= 2 && s.CanAccessArray([]string{"admin", "moderator"}){ //checks length properly without issues
				target = cmd[len(cmd)-1] //gets the targeted user properly without any errors happening
			}

			//creates the simpletable
			//this will store our information
			table := simpletable.New() //makes the structure
			//sets table header
			//this will be used to define information
			table.Header = &simpletable.Header{
				Cells: []*simpletable.Cell{ //sets the values properly without issues happening
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "who", "headers", "username.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "who", "headers", "connected.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "who", "headers", "ip.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "who", "headers", "timestamp.txt").Containing, lexer.Escapes)},
					{Align: simpletable.AlignCenter, Text: lexer.AnsiUtil(views.GetView("commands", "who", "headers", "client.txt").Containing, lexer.Escapes)},
				},
			}
			//ranges through every opens ession
			//this will make sure its done without issues
			for _, usr := range s.GetSessions(target) { //ranges through all users sessions
				//creates the store properly without issues
				rk := []*simpletable.Cell{ //fills with the information properly without issues
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "who", "values", "username.txt").Containing, lexer.Escapes), "<<$username>>", usr.User.Username)}, //id
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "who", "values", "connected.txt").Containing, lexer.Escapes), "<<$connected>>", usr.Connected.Format("15:04:05"))}, //username
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "who", "values", "ip.txt").Containing, lexer.Escapes), "<<$ip>>", usr.Target)}, //maxtime
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "who", "values", "timestamp.txt").Containing, lexer.Escapes), "<<$timestamp>>", strconv.Itoa(int(usr.ID)))}, //maxtime
					{Align: simpletable.AlignCenter, Text: strings.ReplaceAll(lexer.AnsiUtil(views.GetView("commands", "who", "values", "client.txt").Containing, lexer.Escapes), "<<$client>>", GuessClient(string(usr.Conn.ClientVersion())))}, //maxtime
				}

				//saves into the array correctly
				//this will be properly founded later onwards
				table.Body.Cells = append(table.Body.Cells, rk)
			}

			//renders the table properly
			//this will ensure its done without issues
			return pager.MakeTable("who", table, s).TextureTable()
		},
	})
}

//tries to guess the current ssh client
//this will ensure we have got it properly
func GuessClient(Version string) string { //returns a string properly
	return strings.ToLower(strings.Split(strings.ReplaceAll(Version, "SSH-2.0-", ""), "_")[0])
}