package parser

import (
	"Nosviak2/core/sources/language/lexer"
	"errors"
)

//stores information about the function being executed properly
//this will allow for better system handling without issues happening when asked for
type FunctionPath struct {
	//stores the Labels correctly
	//these will act the the Tokens calling the function name
	Labels []lexer.Token //stores the actions leading upto the function

	//stores all the arguments
	//allows for arguments to be used inside the function calls
	Arguments [][]lexer.Token //uncompiled arguments array properly

	//stores all the registered Tokens
	//this will allow for product skipping inside the keyword powered parser
	Tokens []lexer.Token
}

//properly parses the function statement
//allows for more advanced issues without error happenings
func (p *Parser) ExecuteFunction() (*FunctionPath, error) {

	//stores all the upcoming information properly
	//allows us to return this information without issues happening
	var functionPath *FunctionPath = &FunctionPath{}

	//for loops throughout the system
	//this will constantly for loop until the end has been reached
	for proc := p.position; proc < len(p.lex.Tokens()); proc++ {
		//stores the current rotations token
		//allows for better handling without issues happening
		Tokens := p.lex.Tokens()[proc]

		//checks the Tokens type properly without issues
		//this will ensure its done properly without errors happening
		if Tokens.TokenType() == lexer.SemiColon || p.lex.Tokens()[p.position].Position().Row() < Tokens.Position().Row() {
			//returns the error properly without issues happening
			return nil, errors.New("failed to detect functionPath closure")
		}
		//saves into the label path correctly
		//allows for proper and safe handle without issues happening
		functionPath.Labels = append(functionPath.Labels, Tokens)

		//detects the finish correctly
		//this will happen when the function is complete
		if Tokens.TokenType() == lexer.BracketOpen {
			break //breaks the function closure
		}
	}

	//saves into the Tokens properly
	//this will allows us to properly skip the correct position
	functionPath.Tokens = append(functionPath.Tokens, functionPath.Labels...)

	
	//removes the bracket information properly
	//this will ensure that we have still skipped the correct positions
	functionPath.Labels = functionPath.Labels[:len(functionPath.Labels)-1]


	//properly renders the body without issues
	//this will properly and safely read every token until closure
	args, err := p.ReadBodyUntil(p.position+len(functionPath.Labels)+1, p.lex.Tokens(), lexer.BracketOpen, lexer.BracketClose)
	if err != nil {
		//returns the error correctly
		return nil, err
	}
	

	//saves into the Tokens properly
	//allows for better system handling without issues
	functionPath.Tokens = append(functionPath.Tokens, args...)

	//properly parses the arguments without issues happening
	//this will make sure its properly done without issues happening
	args = args[:len(args)-1]

	//parses the arguments properly
	//this will ensure its done properly without errors happening
	argsParsed := p.parseArgs(args); functionPath.Arguments = argsParsed
	//returns the information correctly and properly
	//this will ensure its done properly without errors happening
	return functionPath, nil
}