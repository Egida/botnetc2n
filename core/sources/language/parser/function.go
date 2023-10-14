package parser

import (
	"Nosviak2/core/sources/language/lexer"
	"errors"
)

//stores the main function information
//this ensure its properly done without issues
type Function struct {
	//stores the main function name without issues
	//this ensures its done properly without issues happening
	Header lexer.Token

	//stores the wanted arguments properly
	//this will ensure when executed its properly done without issues
	Args []lexer.Token //example would be like (user string)

	//stores what the system returns properly
	//this ensures its done properly without issues happening
	Returns []lexer.Token //example would be like (string, int)

	//stores the body information properly
	//this will ensure its done properly without issues happening
	Bodies []Node //stores all the parse able nodes inside the value

	//stores everysingle token found without issues
	//this will ensure its properly done without issues happening
	Tokens []lexer.Token //stores all the tokens properly
}


//parses the complete function body without issues
//this ensures its properly done without errors happening on request
func (p *Parser) parseFunction() (*Function, error) {

	//stores the future information without issues 
	//this will ensure its done properly without issues happening
	var functionBody *Function = &Function{}

	//skipping on position forward should display our function name
	//this will be used inside the function structure properly without issues
	header := p.peek(1) //tries to correctly peek the charater

	//updates the token array correctly and properly without issues happeing
	//this will ensure its done properly without issues happening
	functionBody.Tokens =append(functionBody.Tokens, p.lex.Tokens()[p.position]) //main keyword
	functionBody.Tokens =append(functionBody.Tokens, p.lex.Tokens()[p.position+1]) //header information properly

	//properly tries to correctly validate the token
	//this ensures its properly done without issues happening
	if header.TokenType() != lexer.Indent { //returns the error if prompted
		return nil, errors.New("invalid header used inside function request")
	}

	//properly saves the header
	//this ensures its safely stored if it passed the checks
	functionBody.Header = *header

	//properly parses the args safey
	//this will ensure its done properly without issues happening
	for position := p.position + 2; position < len(p.lex.Tokens()); position++ {

		//saves into the args statement if valid
		//this will ensure its done properly without issues happening
		functionBody.Args = append(functionBody.Args, p.lex.Tokens()[position])

		//checks for the closing information
		//this will try to detect the closing bracket without issues
		if p.lex.Tokens()[position].TokenType() == lexer.BracketClose {
			break //breaks from the loop properly
		}

		//tries to correctly handle without issues happening
		//this will ensure its done properly without issues happening
		if p.lex.Tokens()[position].Position().Row() > p.lex.Tokens()[p.position].Position().Row() {
			return nil, errors.New("function expected `(` gotton `newLine`") //returns the error correctly
		}
	}

	//pushes the args properly
	//this will ensure its done properly without issues happening
	functionBody.Tokens = append(functionBody.Tokens, functionBody.Args...)

	//updates the position render properly
	//this will store the next position without issues happening
	current := p.position+len(functionBody.Args)+2 //sets the new position

	//detects the return seq properly
	//this will allow us to properly handle without issues happening
	if p.lex.Tokens()[current].TokenType() == lexer.BracketOpen {
		//ranges throughout all the return passages without issues
		//this will ensure its done properly without issues happening
		for renderReturn := current; renderReturn < len(p.lex.Tokens()); renderReturn++ {
			//saves into the return arguments properly
			//this will ensure its properly saved without errors happening
			functionBody.Returns = append(functionBody.Returns, p.lex.Tokens()[renderReturn])

			//checks if the statement has been complete properly
			//ensures that its properly done without issues happening
			if p.lex.Tokens()[renderReturn].TokenType() == lexer.BracketClose {
				break //breaks from the continue loop properly
			}

			//checks for a newline inside the engine
			//this will count as an error properly without issues
			if p.lex.Tokens()[renderReturn].Position().Row() > p.lex.Tokens()[p.position].Position().Row() {
				return nil, errors.New("function expected `(` gotton `newLine`") //returns the error correctly
			}
		}

		//pushes the returns properly
		//this will ensure its done properly without issues happening
		functionBody.Tokens = append(functionBody.Tokens, functionBody.Returns...)
	}


	//finds the correct body start without issues
	//this ensures its done properly without errors happening
    bodyStart := p.lex.Tokens()[p.position+len(functionBody.Tokens)]
	functionBody.Tokens = append(functionBody.Tokens, bodyStart)

	tokens, err := p.ReadBodyUntil(p.position+len(functionBody.Tokens)+1, p.lex.Tokens(), lexer.ParentheseOpen, lexer.ParentheseClose)
	if err != nil {
		//returns the error correctly and safely
		//this ensures its properly done without issues happening
		return nil, err
	}

	//saves into the array correctly
	//ensures its done properly without issues
	functionBody.Tokens = append(functionBody.Tokens, tokens...)

	//updates the token array correctly
	//this will ensure its done properly without issues
	tokens = tokens[:len(tokens)-1]
	

	//correctly executes the parser without issues happening
	//this will return the array of nodes without issues happening
	nodes, err := MakeTokens(tokens, 0).RunPath()
	if err != nil {
		//returns the error correctly and properly
		return nil, err
	}

	//updates the bodies information
	//this will ensure its done properly without issues
	functionBody.Bodies = nodes //saves into the array

	return functionBody, nil
}