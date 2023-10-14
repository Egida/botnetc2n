package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"
)

//stores the internal functions properly
//this will allow for better handling without issues happeing
type InternalFunction struct {
	//stores the function name correctly
	//allows for properly handling without issues happening
	Header string //stores the function header properly
	//stores the internal nodes correctly and properly
	//this will ensure if its correct and properly without errors
	Nodes []parser.Node //stores all the nodes properly
	//stores all the args properly without issues
	//this will ensure its done properly without issues
	Args []Object //stores all the args properly
	//stores everything which it returns properly
	//allows for proper system management without errors happening
	Returns []lexer.Token
}

//tries to correctly find the function with that name
//this will make sure its properly done without issues happening
func (e *Evaluator) getInternal(name string) *InternalFunction {
	//ranges throughout the functions array
	//makes sure its properly done without issues happening
	for _, f := range e.internalFunctions {

		//compares the name properly
		//makes sure we only return the correct function
		if f.Header == name {
			//returns the function correctly
			return &f
		}
		
	}

	return nil
}

//tries to correctly execute the internal function
//this will ensure its done properly without errors happening
//pretty much converts the internal function into the structure properly
func (e *Evaluator) registerInternal(reg *parser.Function) (error) {
	//correctly fill the options with the object
	//this ensures its properly done without errors happening
	var New *InternalFunction = &InternalFunction{
		Header: reg.Header.Literal(),
		Nodes: reg.Bodies,
		Returns: reg.Returns,
	}
	//stores all the future arguments without issues happening
	//this will ensure its done properly without errors happening
	var args []Object = make([]Object, 0)
	//ranges throughout every single argument inside the wanted
	//this will ensure its done properly without issues happening unwantedly
	for tokens := range reg.Args {
		//adds support for more rules without issues 
		if len(args) <= 0 {
			args = append(args, Object{Type: 0, Literal: ""})
		}
		//forces another rotation is one of these are detected properly
		//makes sure its properly done without errors happening on request
		if reg.Args[tokens].TokenType() == lexer.BracketClose || reg.Args[tokens].TokenType() == lexer.BracketOpen {
			continue //forces the loop again properly
		}

		if reg.Args[tokens].Literal() == "string" {
			args[len(args)-1].Type = lexer.String //creates the string
		} else if reg.Args[tokens].Literal() == "int" {
			args[len(args)-1].Type = lexer.Int //creates the int
		} else if reg.Args[tokens].Literal() == "bool" {
			args[len(args)-1].Type = lexer.Boolean //creates the boolean
		} else if reg.Args[tokens].TokenType() == lexer.Comma {
			args = append(args, Object{}) //creates the new object
		} else if reg.Args[tokens].TokenType() == lexer.Indent {
			args[len(args)-1].Literal = reg.Args[tokens].Literal()
		} else {
			//returns the error correctly
			return errors.New("invalid argument given as function arguments")
		}
	}
	//forces the update on the args
	//this will ensure its done properly without issues
	New.Args = args

	//saves the new internal function properly
	//ensures it properly done without issues and safely
	e.internalFunctions = append(e.internalFunctions, *New)
	return nil
}