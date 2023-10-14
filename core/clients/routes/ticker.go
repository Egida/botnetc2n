package routes

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"log"
	"time"
)

//stores the rotation properly
//this will ensure its done without any errors
var rotations bool = true 

//this will store the title worker properly
//this executes on each rotation and is broadcasted to each session
func TitleWorkerFunction() {
	//ranges through the sessions properly
	//this will allow us to broadcast to each session without issues
	for {
		//ranges through all the sessions
		//this will broadcast the message to each session without issues
		for _, session := range sessions.Sessions {

			//detects the current rotation
			if rotations && toml.ConfigurationToml.AppSettings.CursorBlink { //blink enabled
				session.Write("\033[?25l")
			} else { //blink disabled
				session.Write("\033[?25h\033[?0c")
			}

			//resets the custom title properly
			//this will remove the custom title information
			if session.CustomTitleReset <= time.Now().Unix() { //checks
				session.Title = "" //checks properly without issues
			}

			//checks the length properly
			//this will ensure the custom title is rendered
			if len(session.Title) > 0 { //checks properly without issues
				session.Write("\033]0;"+session.Title+"\007"); continue //continues looping
			}


			//checks the different route types properly
			//this will ensure its done without any errors
			if toml.ConfigurationToml.TitleWorker.Route == "itl" {
				//renders the title to the main profile using itl
				//this will make sure its correctly done without issues happening
				if err := language.ExecuteLanguage([]string{"title.itl"}, session.Channel, deployment.Engine, session, make(map[string]string)); err != nil {
					delete(sessions.Sessions, session.ID) //deletes the session properly without issues happening
					session.Channel.Close() //closes the session properly without issues happening
					continue //continues the loop properly without issues happening
				}
			} else { //executes using termfx
				//this will use the termfx route without issues happening
				//allows for better control without issues happening on reqeust
				frame, err := language.MakeTermFX([]string{"termfx-title.tfx"}, session)
				if err != nil { //error handles properly without issues happening
					continue //continues looping without issues happening
				}

				//tries to write the title without issues
				//this will ensure its done without any errors happening
				if err := session.Write("\033]0;"+frame+"\007"); err != nil {
					log.Printf("[TERMFX TITLE] reason: %s\r\n", err.Error()); continue
				}
			}
		}
		//sleeps for the duration given properly
		//this will make sure its done without issues happening
		time.Sleep(time.Duration(toml.ConfigurationToml.TitleWorker.Duration) * tools.ResolveString(toml.ConfigurationToml.TitleWorker.TimeUnit))

		//flicks the rotation properly
		//this will ensure its done without any errors
		rotations = !rotations //flips properly without issues
	}
}