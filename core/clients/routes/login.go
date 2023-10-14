package routes

import (
	"Nosviak2/core/clients/views/util"
	"Nosviak2/core/configs"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/tfx"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"io"
	"strconv"

	"errors"
	"time"

	"golang.org/x/crypto/ssh"
)

//properly tries to work the information without issues happen
//this will make sure its done without any errors happening on reqeust
func LoginRequest(channel ssh.Channel, conn *ssh.ServerConn) (*database.User, error) {

	//gets the header information properly
	//this will allow for better handling without issues
	header := views.GetView("views", "login", "header.tfx")
	if header == nil { //error handles the statement properly
		return nil, errors.New("missing branding peice correctly: login/header.tfx")
	}
	//gets the header information properly
	//this will allow for better handling without issues
	username := views.GetView("views", "login", "username.tfx")
	if header == nil { //error handles the statement properly
		return nil, errors.New("missing branding peice correctly: login/username.tfx")
	}
	//gets the header information properly
	//this will allow for better handling without issues
	password := views.GetView("views", "login", "password.tfx")
	if header == nil { //error handles the statement properly
		return nil, errors.New("missing branding peice correctly: login/username.tfx")
	}

	termp := termfx.New()
	//registers the date function without issues
	//this will ensure its done without any errors happening
	termp.RegisterFunction("date", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(time.Now().Format("Mon 2 Jan 15:04:05"))) //formats the time
	})

	//registers the api name properly
	//this will ensure its done without any errors
	termp.RegisterVariable("cnc", toml.ConfigurationToml.AppSettings.AppName)
	termp.RegisterVariable("version", deployment.Version) //stores the version properly

	//tries to correctly execute the header
	//this will ensure its done without any errors happening
	Raw, err := termp.ExecuteString(header.Containing) //renders the string
	if err != nil { //error handles the statement without issues and errors happening
		return nil, err //returns the error
	}

	//writes the header properly without issues
	//this will ensure its done without any errors happening
	if _, err := channel.Write([]byte(lexer.AnsiUtil(Raw, lexer.Escapes))); err != nil {
		return nil, err //returns the error correctly and properly
	}
	//renders the username and acts with the prompt
	//this will ensure its done without any errors happening
	if _, err := channel.Write([]byte(lexer.AnsiUtil(username.Containing, lexer.Escapes))); err != nil {
		return nil, err //returns the error correctly
	}

	//properly reads the line without issues
	//this will act as the username without errors
	usernameReq, err := util.TermReader(channel, toml.ConfigurationToml.Login.MaxUsernameInput, false, "")
	if err != nil { //error handles the reader properly
		return nil, err
	}

	//tries to find a valid user with that name
	//this ensures its done without any errors happening
	user, err := database.Conn.FindUser(usernameReq) //tries to find the user properly
	if err != nil || user == nil { //error handles the req properly without any issues happening
		//tries to render the invalid username render item
		//this will ensure its done correctly and properly without errors
		if _, err := channel.Write([]byte(lexer.AnsiUtil(views.GetView("views", "login", "invalid-username.tfx").Containing, lexer.Escapes))); err != nil {
			return nil, err //returns the error correctly and properly
		}
		//sleeps for the duration needed properly
		//this will ensure its done without any errors
		time.Sleep(10 * time.Second)
		//returns the invalid username error
		//this will close the session without issues
		return nil, errors.New(conn.RemoteAddr().String()+" has provided any invalid username "+usernameReq)
	}
	//ranges through all the auth attempts properly
	//this will ensure its done without any errors happening
	for attempt := 0; attempt < json.ConfigSettings.Masters.MaxAuthAttempts; attempt++ {

		//prints the login screen with the user filled in properly
		//allows for better control without issues happening on purpose
		if err := PrintLoginWithUser(Raw, username.Containing, channel, usernameReq); err != nil {
			return nil, err //error handles properly without issues happening
		}

		//renders the username and acts with the prompt
		//this will ensure its done without any errors happening
		if _, err := channel.Write([]byte(lexer.AnsiUtil(password.Containing, lexer.Escapes))); err != nil {
			return nil, err //returns the error correctly
		}
		//properly reads the line without issues
		//this will act as the password without errors
		passwordReq, err := util.TermReader(channel, toml.ConfigurationToml.Login.MaxPasswordInput, true, toml.ConfigurationToml.Login.MaskingCharater)
		if err != nil { //error handles the reader properly
			continue
		}


		//compares the passwords when hashed properly
		//this will ensure its done without any errors happening
		if user.Password != database.HashProduct(passwordReq) {
			//renders the invalid password properly
			//this will ensure its done without any errors happening
			if _, err := channel.Write([]byte(lexer.AnsiUtil(views.GetView("views", "login", "invalid-password.tfx").Containing, lexer.Escapes))); err != nil {
				return nil, err //returns the error correctly and properly
			}
			//executes the terminal shaker properly
			//this will shake the terminal window properly
			tools.ShakeTerminal(5, time.Duration(11 * time.Millisecond), channel)

			time.Sleep(2 * time.Second)
			//this will ensure its not ignored properly
			//this will make sure its safe without issues
			continue //continues looping properly without issues
		}

		//returns nil properly and safely
		//this will return the username without issues happening
		return user, nil //this will ensure its done without any errors
	}
	

	//too many attempts banner here
	//this will try to load without issues happening
	if _, err := channel.Write([]byte(lexer.AnsiUtil(views.GetView("views", "login", "too-many-attempts.tfx").Containing, lexer.Escapes))); err != nil {
		return nil, err //returns the error properly without issues happening
	}

	//renders the banner properly
	//this will ensure its done without issues
	time.Sleep(5 * time.Second) //sleeps for the duration

	//returns the error properly without any issues
	//this will make sure its safe without any errors happening
	return nil, errors.New(conn.RemoteAddr().String()+" has provided an invalid password for "+usernameReq+" "+strconv.Itoa(json.ConfigSettings.Masters.MaxAuthAttempts)+" times")
}




//renders the login screen with the username filed
//this will ensure its done without any errors happening
func PrintLoginWithUser(header string, username string, channel ssh.Channel, givenUser string) error {
	//renders the header properly without issues
	//this will ensure its done without any errors happening
	if _, err := channel.Write([]byte("\033c"+lexer.AnsiUtil(header, lexer.Escapes))); err != nil {
		return err //returns the error properly
	}

	//renders the username properly without issues
	//this will ensure its done without any errors happening
	if _, err := channel.Write([]byte(lexer.AnsiUtil(username + givenUser, lexer.Escapes))); err != nil {
		return err //returns the error properly
	}; 
	
	//returns nil as no errors happening
	//allows for better control without issues
	return nil
}