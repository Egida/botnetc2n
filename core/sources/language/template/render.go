package template

import (
	"Nosviak2/core/clients/sessions"
	"errors"
	"unicode/utf8"
)

//stores information about the current instance
//this will allow for proper and safe control of the instances
type TemplateEngine struct {
	//stores the current line we will render
	//this will allow for better control without issues
	lines string //stored in type string properly
	//stores the current position
	//this will make sure its done without errors happening
	position int //stored in type int safely
	//this will store what we are looking for
	//makes sure its properly done without issues happening
	enginePrefix [2]string //stored in type array properly
}

//creates the new template without issues
//this will ensure its done without errors happening
func MakeTemplate(line string, prefix [2]string) *TemplateEngine {
	return &TemplateEngine{ //returns the template properly
		lines: line, //sets the line properly without issues happening
		position: 0, //sets the position properly without issues
		enginePrefix: prefix, //sets the engines prefix properly
	}
}

//correctly runs the engine without issues
//this will ensure its done properly without errors
func (t *TemplateEngine) RunEngine(f func(tag string, s *sessions.Session) (int, error), wr *sessions.Session) error {

	//ranges througout the source properly
	//this will allow for proper control without errors happening
	for pos := 0; pos < utf8.RuneCountInString(t.lines); pos++ { //for loops througout with issues
		//updates the position inside the engine
		//this will ensure its properly done without issues
		t.position = pos //updates position

		//checks if the current line is a inlet
		//this will start with all the system checks without issues
		if t.takeEngineIN() {

			//stores the collect tag without issues happeing
			//this will make sure its properly done without errors happening on request
			tag, finish := "", false //stored inside a string item properly and safely without issues

			//ranges througout the system
			//allows for proper controlling without issues
			for texture := pos + utf8.RuneCountInString(t.enginePrefix[0]); texture < utf8.RuneCountInString(t.lines); texture++ {
				//this will also update the position
				//allows for proper handling without issues
				t.position = texture

				//detects the engine out properly
				//this will make sure its done without issues happening
				if t.takeEngineOUT() {
					//sets the finished to true, removes the error chance correctly
					//this will make sure the header function knows about the header function error
					finish = true; break //breaks and finishs correctly
				} else {
					//saves into the tag correctly
					//this will allow the header function to access without errors happening
					tag += string(t.lines[texture]); continue
				}
			}
			
			//checks if the engine was closed
			//this will make sure its properly been closed without error
			if !finish { //this detects that is was closed without issues
				return errors.New("tag opened but never closed without valid reason for why")
			}

			//this will also update the main looper without issues
			//allows for proper handling without issues happening on request
			pos = t.position + utf8.RuneCountInString(t.enginePrefix[1]) -1 //skips the correct amount without issues

			//tries to correctly execute the function without issues
			//this will make sure its been correctly executed without valid issues
			_, errors := f(tag, wr) //properly executes under the header function without issues
			if errors != nil { //tries to correctly error handle without issues happening without a reason
				//returns the error correcty and safely
				return errors //error returned properly
			}; continue //continues if so
		} else { //else we will render the information without a reason
			//writes to the text to the writer without issues happening
			//this will ensure that the writer has collected the information
			if err := wr.Write(string(t.lines[pos])); err != nil {
				//returns the error correctly and properly without issues
				return err //returns the error correctly
			}
		}
	}


	return nil
}