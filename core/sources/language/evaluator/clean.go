package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
)

//properly sorts the return command without issues happening
//this will make sure its properly done without errors happening on request
func (e *Evaluator) sortReturns(r *parser.ReturnReply) ([]Object, error) {

	//properly tries to sort the information without issues happening
	//this will make sure its properly done without errors happening on request
	var tok []Object = make([]Object, 0)

	//ranges throughout all the different objects correctly
	//this will ensure its done properly without errors happening making it safer
	for _, array := range r.Values {

		//properly compiles the tokens without errors
		//this will ensure its done properly without errors
		token, err := e.compileTokens(array, lexer.EOF)
		if err != nil {
			//returns the error correctly
			return nil, err
		}

		//saves into the array correctly and properly
		//this will make sure its done without issues happening
		tok = append(tok, Object{ //saves as object
			Literal: token.Literal(), //sets literal properly
			Type: token.TokenType(), //sets tokentype properly
		})
	}
	
	//returns the values properly without issues
	//this will ensure its done properly without issues happening
	return tok, nil
}