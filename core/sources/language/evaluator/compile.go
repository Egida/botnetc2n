package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"
	"strconv"
)

//properly controls/compiles the tokens without issues
//this will allow for better controlling without issues happening
func (e *Evaluator) compileTokens(tokens []lexer.Token, mustType lexer.TokenType) (*lexer.Token, error) {
	//properly tries to correctly find the token type
	//this will work what the token type should be allows support for multisystems
	if mustType == lexer.EOF {

		//makes sure the values tokens is longer than 0
		//this will use the the first token
		if len(tokens) <= 0 {
			//returns the error correctly and properly
			return nil, nil
		}

		//adds support for the system without issues
		//this will properly store the information without errors
		var token *lexer.Token = &tokens[0]

		//detects possible variables for support
		//this will allow variables to be used without issues happening
		if token.TokenType() == lexer.Indent {
			//tries to correctly find the variable
			//if the variable system fails we will look for the function
			object, err := e.findScope(token.Literal())
			if err != nil || object == nil { //error handles properly
				//tries to parser for function
				//allows for proper handle without issues
				path, err := parser.MakeTokens(tokens, 0).ExecuteFunction()
				if err != nil {
					return nil, err
				}
				//tries to correctly execute the function
				//allows for proper subject handle without issues
				mem, err := e.executeFunction(path)
				if err != nil {
					//returns the error
					return nil, err
				}
				//validates the length
				//makes sure its correctly done without issues
				if len(mem) > 1 {
					//returns the error properly without issues
					return nil, errors.New("mismatch type, one wanted, "+strconv.Itoa(len(mem))+" given")
				}
				//inserts into the array correctly and properly
				//this will ensure its done properly and we can access it
				token = lexer.NewToken(mem[0].Literal, mem[0].Type, token.Position())
			} else {
				//sets the object properly
				//allows for better support without issues happening
				token = object.TokenValue //sets the value properly without issues 
			}
		}
		//updates the token type properly
		//this will ensure its done properly without issues happening
		mustType = token.TokenType() //forces the update without issues happening
	}


	switch mustType {

	case lexer.String:
		//returns the subject compiled
		//this will properly try to compile the information
		return e.compileString(tokens)
	case lexer.Int:
		//returns the subject compiled
		//this will properly try to compile the information
		return e.compileInt(tokens)
	case lexer.Boolean:
		//returns the subject compiled
		//this will properly try to compile the information
		return e.compileBoolean(tokens)
	}

	return nil, nil
}
