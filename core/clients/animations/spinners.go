package animations

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"Nosviak2/core/sources/language/shorts"
	"Nosviak2/core/sources/layouts/toml"
	"errors"
	"strings"
	"time"
	"sync"
)


var (
	//stores all the running spinners for each session
	//allows for better control without any errors happening
	running map[int64]bool = make(map[int64]bool)
	mutexSpec sync.Mutex
)
//properly renders all the frames
//this will ensure its done without issues happening on reqeust
func RunSpinner(spinner string, s *sessions.Session, custom map[string]string) error {
	//sets the finish ratio to false
	//this will make sure its done without issues
	mutexSpec.Lock() //saves into running
	running[s.ID] = false //ensure its done without issues
	mutexSpec.Unlock() //unlocks again properly


	//tries to cache the spinner frames
	//this will ensure its done without issues happening
	frames := toml.Spinners.Spins[spinner] //tries to get without issues
	if frames == nil { //error handles properly without issues happening on requests
		return errors.New("failed to grab the frames from the toml file")
	}

	//starts the for loop properly
	//this will ensure its done without errors
	for { //starts the for looping properly
	
		//ranges through all the frames properly
		//this will make sure its done without errors happening
		for _, frame := range frames.Frames { //ranges through the frames without issues happening
			//checks if the spinner has finished properly
			//this will make sure its done without issues happening
			if running[s.ID] { //checks for the finish correctly and properly
				return nil //clears and end the functions
			} 
			frame = "\\033[2K\\r" + frame + "\\x1b[0m"//adds the new line system properly

			//checks if they want itl on the spinner
			//this will allow for better editing without issues happening
			if frames.Launch {
				//tries to start the parser properly and safely
				//this will ensure its done without errors happening
				nodes, err := parser.MakeParserRun(lexer.Make(frame, true).RunTarget())
				if err != nil { //error handles the parser statement
					return err //returns the error correctly
				}

				//registers the packages and the custom variables
				//normal functions will not be able to be registered properly
				e := shorts.Register(custom, evaluator.MakeEval(nodes, make([]evaluator.Scope, 0), s.Channel, deployment.Engine, s))
				if _, err := e.FollowGuide(); err != nil { //error handles the statement
					continue //continues the loop properly
				}
			} else {
				//ranges through the frames properly
				//this will allow for some basic tfx codes
				for ref, value := range custom { //replaces the object properly
					frame = strings.ReplaceAll(frame, "<<"+ref+"()>>", value) //replaces
				}
				
				//writes the frame properly
				//this will ensure its done without issues
				if err := s.Write(lexer.AnsiUtil(frame, lexer.Escapes)); err != nil { //error handles
					continue //continues looping properly
				}
			}

			//sleeps for the amount of ticks given properly
			//this will ensure its done without any errors happening
			time.Sleep(time.Duration(frames.Ticks) * time.Millisecond)
		}
	}
}


//ends the spinner
//this will ensure its done without issues
func End(s *sessions.Session) { //changes to true properly
	s.Write("\033[2K\r") //clears the line
	mutexSpec.Lock() //locks mutex again
	running[s.ID] = true //stops spinner
	mutexSpec.Unlock() //unlocks again
}