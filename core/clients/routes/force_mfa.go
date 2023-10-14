package routes

import (
	"Nosviak2/core/clients/sessions"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/functions"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/views"
	"errors"
	"net/url"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/xlzd/gotp"
	"golang.org/x/term"
)

//wants mfa enforced properly and safely
//this will make sure its done without errors happening
func WantMFA(session *sessions.Session) error { //err handles

	//resizes the screen properly and safely
	//this will ensure its done without errors happening
	err := language.ExecuteLanguage([]string{"views", "force_mfa", "resize_screen.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	if err != nil { //err handles properly
		return err
	}

	//gets the resize system properly and safely
	//this will make sure its done without errors happening on purpose
	source := views.GetView("views", "force_mfa", "mfa_displayed.itl")

	//awaits until the screen has been resized
	//this will ensure we can completely fit the qr code
	for { //loops through properly

		//checks the length properly and safely without errors
		//ensures its done without errors happening on purpose
		if session.WindowSize.Height < 55 + len(strings.Split(source.Containing, "\r\n")) && session.WindowSize.Length < 90 {
			continue
		}

		break //breaks looping
	}

	//generates our secret properly and safely
	//this will ensure its done without errors happening
	incomingSec, err := functions.GenerateSecret() //generates
	if err != nil { //err handles properly
		return err
	}

	//creates our otp method properly and safely
	//this will ensure its done without errors happening
	data, err := qrcode.New("otpauth://totp/" + url.QueryEscape(toml.CatpchaToml.Mfa.App) + ":" + url.QueryEscape(session.User.Username) + "?secret=" + incomingSec + "&issuer=" + url.QueryEscape(toml.CatpchaToml.Mfa.App) + "&digits=6&period=30", qrcode.Medium)
	if err != nil { //err handles properly
		return err
	}

	//creates our terminal system properly and safely
	//this will ensure its done without any errors happening
	qrcode, err := functions.TerminalQR(data.Bitmap(), "")
	if err != nil { //err handles properly
		return err
	}

	//resizes the screen properly and safely
	//this will ensure its done without errors happening
	err = language.ExecuteLanguage([]string{"views", "force_mfa", "mfa_displayed.itl"}, session.Channel, deployment.Engine, session, map[string]string{"qrcode":qrcode, "secret":incomingSec})
	if err != nil { //err handles properly
		return err
	}

	//enters the system prompt properly
	//this will ensure its done without errors happening
	err = language.ExecuteLanguage([]string{"views", "force_mfa", "prompt.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	if err != nil { //err handles properly
		return err
	}

	//reads the line input properly and safely
	//this will make sure its done without errors happening
	systemIN, err := term.NewTerminal(session.Channel, "").ReadLine()
	if err != nil { //err handles properly
		return err
	}

	//err handles properly and safely without issues
	//allows the system without errors happening on purpose
	if gotp.NewDefaultTOTP(incomingSec).Now() == systemIN {
		return nil //returns nil and safely
	}

	//enters the system prompt properly
	//this will ensure its done without errors happening
	err = language.ExecuteLanguage([]string{"views", "force_mfa", "invalid_otp.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	if err != nil { //err handles properly
		return err
	}

	return errors.New("invalid mfa input properly")
}