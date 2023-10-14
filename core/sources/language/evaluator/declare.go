package evaluator

import (
	"Nosviak2/core/sources/language/parser"
)

//comparses the variable into a token subject
//this will allow us to use the variable in anyplace
func (e *Evaluator) comparseVariable(r *parser.DeclareRoute) (*Scope, error) {
	//creates the scope properly
	var object *Scope = &Scope{
		//this will help with better handle without issues happening
		Name: r.RouteName(), //sets the route name properly
	}
	//properly tries to compiles the tokens into one
	//this will safely and properly try to work out the information
	seq, err := e.compileTokens(r.Values(), r.GivenType())
	if err != nil {
		//returns the error correctly and safely
		//this will allow the header function to be alerted
		return nil, err //returns the error
	}
	//updates the token value without issues happening
	//this will make sure its done properly without issues happening
	object.TokenValue = seq
	return object, nil
}