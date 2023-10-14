package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/database"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"

	"golang.org/x/term"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "password",
		Aliases:            []string{"passwd", "mypassword"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "update your accounts password properly",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//prints the language banner properly
			//this will ensure its done without issues happening
			if err := language.ExecuteLanguage([]string{"commands", "password", "banner.itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
				return err //returns the error correctly
			}
			//renders the prompt information properly
			//this will complete the prompt without errors happening
			Pass, err := Prompt([]string{"commands", "password", "password.itl"}, s)
			if err != nil { //error handles the prompt statement properly without issues
				return err //returns the error correctly and properly
			}
			//renders the prompt information properly
			//this will complete the prompt without errors happening
			ConfirmPass, err := Prompt([]string{"commands", "password", "confirm-password.itl"}, s)
			if err != nil { //error handles the prompt statement properly without issues
				return err //returns the error correctly and properly
			}
			//checks if the password confirm
			//this will ensure its done without issues
			if Pass != ConfirmPass { //checks they match and returns the error if they dont properly
				return language.ExecuteLanguage([]string{"commands", "password", "not-matching.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}
			//tries to correctly update the password
			//this will ensure the users password has been 
			if err := database.Conn.Password(Pass, s.User.Username); err != nil { //returns the error correctly and properly
				return language.ExecuteLanguage([]string{"commands", "password", "error.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//updates the session password properly
			//this will make sure its done without any errors
			sessions.Sessions[s.ID].User.Password = Pass
			
			//returns the success message properly
			//this will allow them to view it was correct
			return language.ExecuteLanguage([]string{"commands", "password", "success.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
		},
	})
}

//properly builds a prompt without issues
//this will ensure its complete without errors happening
func Prompt(file []string, s *sessions.Session) (string, error) { //returns string and error
	if err := language.ExecuteLanguage(file, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
		return "", err //returns the error correctly and properly without issues
	}
	//properly reads the input properly
	//this will ensure its done without issues
	return term.NewTerminal(s.Channel, "").ReadLine()
}