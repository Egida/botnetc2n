package routes

import (
	"Nosviak2/core/clients/views/util"
	"Nosviak2/core/database"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language/lexer"
	termfx "Nosviak2/core/sources/language/tfx"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/views"
	"errors"
	"io"
	"log"
	"time"

	"golang.org/x/crypto/ssh"

	tml "github.com/naoina/toml"
)

type RedeemConfigure struct {
	MaxTokenInput int    `toml:"max_token_input"`
	MaskCharater  string `toml:"maskingCharater"`

	UsernameMaxLen int `toml:"username_max_input"`
	UsernameMaskChar string `toml:"username_maskCharater"`

	PasswordMaxLen int `toml:"password_max_input"`
	PasswordMaskChar string `toml:"password_maskCharater"`
}

//stores the redeem route routes properly
//this will execute when someone tries to redeem a token
func RedeemRoute(ch ssh.Channel, conn *ssh.ServerConn) error {
	//tries to render the system information
	//allows for better control without issues happening
	mainPart := views.GetView("views", "redeem", "redeem.tfx")
	if mainPart == nil { //checks if it was found properly
		return errors.New("missing views/redeem/redeem.tfx") //returns error
	}

	tfx := termfx.New()
	//executes the termfx solutions
	//this will allow for proper termfx support
	tfx.RegisterVariable("cnc", toml.ConfigurationToml.AppSettings.AppName)
	tfx.RegisterVariable("version", deployment.Version) //app version properly
	tfx.RegisterFunction("date", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(time.Now().Format("Mon 2 Jan 15:04:05")))
	})

	//executes the banner string properly
	//this will ensure its done without any errors
	raw, err := tfx.ExecuteString(mainPart.Containing)
	if err != nil { //error handles properly without issues
		return nil //returns the error properly
	}

	//writes the banner properly
	//this will ensure its done without any errors
	if _, err := ch.Write([]byte(lexer.AnsiUtil(raw, lexer.Escapes))); err != nil {
		return err //returns the error properly without issues
	}

	//gets the prompt properly
	//this will ensure its done without issues
	token := views.GetView("views", "redeem", "token.tfx")
	if token == nil { //checks for nil pointers properly
		return errors.New("missing views/redeem/redeem.tfx")
	}

	var redeem RedeemConfigure
	//umarshals and parses the input properly
	//this will ensure its done without errors happening
	if err := tml.Unmarshal([]byte(views.GetView("views", "redeem", "inputs.ini").Containing), &redeem); err != nil {
		return err //returns the error properly and safely
	}

	var masking bool = false
	if len(redeem.MaskCharater) > 0 {
		masking = true
	}
	

	//writes the prompt seq properly and safely
	//this will ensure its done without errors happening
	if _, err := ch.Write([]byte(lexer.AnsiUtil(token.Containing, lexer.Escapes))); err != nil {
		return err //returns the error properly
	}

	//this will ensure it has been properly redeemed
	//allows for proper control without errors happening
	in, err := util.TermReader(ch, redeem.MaxTokenInput, masking, redeem.MaskCharater)
	if err != nil { //error handles properly and safely
		return err
	}

	//tries to find the redeem token
	//this will ensure its done without any errors
	Tok, err := database.Conn.Redeem(in) //redeems the token properly
	if err != nil || Tok == nil { //error handles properly without issues happening
		//writes the invalid token information
		//this will alert the main header about it properly
		if deployment.DebugMode {
			log.Printf("[USER CREATION].[REDEEM] %s\r\n", err.Error())
		}
		if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "invalidToken.tfx").Containing, lexer.Escapes))); err != nil {
			return err //returns the error properly without issues
		}
		//sleeps for 10 seconds properly
		//stops rendering issues happening
		time.Sleep(10 * time.Second); return nil
	}

	//makes sure its follows the correct route
	//ensures its not matched as a different route
	switch Tok.Type {
	//makes the user a new account
	//this will insert into the database
	case database.RedeemUser: //redeem user
		//tries to render the system information
		//allows for better control without issues happening
		mainPart := views.GetView("views", "redeem", "user.tfx")
		if mainPart == nil { //checks if it was found properly
			return errors.New("missing views/redeem/redeem.tfx") //returns error
		}

		tfx := termfx.New()
		//executes the termfx solutions
		//this will allow for proper termfx support
		tfx.RegisterVariable("cnc", toml.ConfigurationToml.AppSettings.AppName)
		tfx.RegisterVariable("version", deployment.Version) //app version properly
		tfx.RegisterFunction("date", func(session io.Writer, args string) (int, error) {
			return session.Write([]byte(time.Now().Format("Mon 2 Jan 15:04:05")))
		})

		//executes the banner string properly
		//this will ensure its done without any errors
		raw, err := tfx.ExecuteString(mainPart.Containing)
		if err != nil { //error handles properly without issues
			return nil //returns the error properly
		}

		//writes the banner properly
		//this will ensure its done without any errors
		if _, err := ch.Write([]byte(lexer.AnsiUtil(raw, lexer.Escapes))); err != nil {
			return err //returns the error properly without issues
		}

		//writes our central username prompt
		//this will ensure its done without errors happening
		if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "username.tfx").Containing, lexer.Escapes))); err != nil {
			return err //returns the error properly
		}

		var userMask bool = false
		if len(redeem.UsernameMaskChar) > 0 {
			userMask = true
		}

		//tries to write into the system properly
		//this will ensure its done without errors happening
		userNew, err := util.TermReader(ch, redeem.UsernameMaxLen, userMask, redeem.UsernameMaskChar)
		if err != nil { //error handles properly and safely
			return err
		}
		
		//writes our central username prompt
		//this will ensure its done without errors happening
		if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "username.tfx").Containing, lexer.Escapes))); err != nil {
			return err //returns the error properly
		}

		var passMask bool = false
		if len(redeem.PasswordMaskChar) > 0 {
			userMask = true
		}

		//tries to write into the system properly
		//this will ensure its done without errors happening
		passNew, err := util.TermReader(ch, redeem.PasswordMaxLen, passMask, redeem.PasswordMaskChar)
		if err != nil { //error handles properly and safely
			return err
		}
		

		//tries to find the user inside the database
		//this will try to stop any username dups properly
		if user, err := database.Conn.FindUser(userNew); err == nil || user != nil {
			//writes the user error properly and safely
			//this will ensure its done without any errors 
			if deployment.DebugMode {
				log.Printf("[USER CREATION].[REDEEM] %s\r\n", err.Error())
			}
			if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "user-dup.tfx").Containing, lexer.Escapes))); err != nil {
				return err //returns the error properly without issues
			}
			time.Sleep(10 * time.Second) //sleeps for 10 seconds properly
			return nil
		}

		//tries to properly remove the token
		//this will ensure its done without any errors
		if err := database.Conn.RemoveRedeem(in); err != nil {
			//writes the user error properly and safely
			//this will ensure its done without any errors 
			if deployment.DebugMode {
				log.Printf("[USER CREATION].[REDEEM] %s\r\n", err.Error())
			}
			if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "user-error.tfx").Containing, lexer.Escapes))); err != nil {
				return err //returns the error properly without issues
			}
			time.Sleep(10 * time.Second) //sleeps for 10 seconds properly
			return nil
		}

		Tok.User.Username = userNew  //username
		Tok.User.Password = passNew //password

		//tries to create the user without issues
		//tries to properly insert it into the col properly
		if err := database.Conn.MakeUser(Tok.User); err != nil {
			//writes the user error properly and safely
			//this will ensure its done without any errors 
			if deployment.DebugMode {
				log.Printf("[USER CREATION].[REDEEM] %s\r\n", err.Error())
			}
			if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "user-error.tfx").Containing, lexer.Escapes))); err != nil {
				return err //returns the error properly without issues
			}
			time.Sleep(10 * time.Second) //sleeps for 10 seconds properly
			return nil
		}

		//renders the success message properly
		//this will ensure its done without errors happening
		if _, err := ch.Write([]byte(lexer.AnsiUtil(views.GetView("views", "redeem", "user-success.tfx").Containing, lexer.Escapes))); err != nil {
			return err //returns the error properly without issues
		}
		time.Sleep(10 * time.Second) //sleeps for 10 seconds properly
		return nil
	}


	return nil
}