package routes

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/database"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/views"
	"errors"

	"golang.org/x/term"
)

var (
	ErrInvalid error = errors.New("newLogin function returned invalid properly")
)

//this will enforce the user to change there password without issues
//makes sure they change there password and if they dont close there session properly
func NewLogin(s *sessions.Session) error { //returns error properly
	//properly executes the language without issues
	//this will ensure its done without any errors happening
	if err := language.ExecuteLanguage([]string{"views", "newuser", "banner.itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
		return err //returns the error properly without issues
	}

	//tries to get the prompts properly
	//this will grab the username properly
	password := views.GetView("views", "newuser", "password.tfx")
	if password == nil { //error handles the statement properly
		return errors.New("missing branding peice correctly: views/newuser/password.tfx")
	}

	//gets the header information properly
	//this will allow for better handling without issues
	confirmpassword := views.GetView("views", "newuser", "confirm_password.tfx")
	if confirmpassword == nil { //error handles the statement properly
		return errors.New("missing branding peice correctly: views/newuser/confirm_password.tfx")
	}

	//tries to read the line properly
	//this will ensure its done without any errors
	passwordPrompt, err := term.NewTerminal(s.Channel, lexer.AnsiUtil(password.Containing, lexer.Escapes)).ReadLine()
	if err != nil { //error handles the system properly without issues
		return err //returns the error properly without issues
	}

	//tries to read the line properly
	//this will ensure its done without any errors
	confirmPasswordPrompt, err := term.NewTerminal(s.Channel, lexer.AnsiUtil(confirmpassword.Containing, lexer.Escapes)).ReadLine()
	if err != nil { //error handles the system properly without issues
		return err //returns the error properly without issues
	}

	//checks if the password is the same as there old one
	//this will ensure its not allowing duping for password schemas
	if database.HashProduct(passwordPrompt) == s.User.Password {
		//tries to render the same as password util
		//this will ensure its done without any errors happening
		if err := language.ExecuteLanguage([]string{"views", "newuser", "same-oldpassword.itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
			return err //returns the error properly
		}; return ErrInvalid //returns the unauth error properly
	}

	//tries to update the information
	//this will ensure its done without any errors happening on request
	if err := database.Conn.Password(confirmPasswordPrompt, s.User.Username); err != nil { //error handles
		//returns the invalid password prompt properly
		//this will ensure its done without any errors happening
		if err := language.ExecuteLanguage([]string{"views", "newuser", "password-error.itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
			return err //returns the error properly
		}; return ErrInvalid //returns the unauth error properly
	}

	//updates the sessions password aswell
	//this will make sure its done without any errors
	sessions.Sessions[s.ID].User.Password = confirmPasswordPrompt
	sessions.Sessions[s.ID].User.NewUser = false //removes new user

	//returns on that note properly
	//this will ensure its done without errors
	return database.Conn.DisableNewUser(s.User.Username)
}