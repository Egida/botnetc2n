package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"strings"
)

//properly supports the system without issues
//this will ensure its done properly without issues happening
func (e *Evaluator) support(b ...bool) bool {
	//stores the current value
	//allows for better handle without issues 
	var current int = 0

	//ranges through the boolean
	//properly switchs the information without issues
	for bool := range b {
		//switchs the system information
		//allows for proper control without issues
		switch b[bool] {
		case true: 		//detection for true
			current++ //one postive
		case false: //detection for false
			current = current - 2 //two negative
		}
	}
	//inputs the rule properly
	//allows for proper suppoer
	if current <= 0 {
		//return false
		return false
	} else {
		//return true
		return true
	}
}

//properly joins the labels without issues
//this will allow for proper token joining without issues
//ensures its done without issues happening on request
func (e *Evaluator) joinLabels(tokens []lexer.Token) string {
	//stores the complete string outcome properly
	//this will ensure its done properly without errors
	var complete string = ""
	//ranges through all the tokens without issues happening
	//this will ensure its done properly without an error showing
	for token := range tokens {
		
		if len(tokens[token].Literal()) == 0 {
			complete += "\"" + tokens[token].Literal() + "\""; continue
		}

		if tokens[token].TokenType() == lexer.String && strings.Split(tokens[token].Literal(), "")[0] != "\"" && strings.Split(tokens[token].Literal(), "")[len(strings.Split(tokens[token].Literal(), ""))-1] != "\"" {
			complete += "\"" + tokens[token].Literal() + "\""; continue
		}
		complete += tokens[token].Literal()
	}
	//returns the literal format correctly
	//this will ensure its done properly without issues
	return complete
}


//checks if the object already exists
//this will allow for the object to be updated
func (e *Evaluator) varExists(header string) bool {
	//ranges througout all the different objects
	//this will ensure its done properly without issues
	for object := range e.GuideScope {
		//compares the two different components
		//this will make sure its properly done without issues
		if e.GuideScope[object].Name == header {
			//returns true as it exists
			return true
		}
	}
	//returns false as its invalid
	//this will ensure its not done properly
	return false
}