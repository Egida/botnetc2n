package parser

import (
	"Nosviak2/core/sources/language/lexer"
	"errors"
)

//correctly reads a body until a close is detected
//this will ensure its done properly without issues happening
func (p *Parser) ReadBodyUntil(position int, tokens []lexer.Token, open lexer.TokenType, close lexer.TokenType) ([]lexer.Token, error) {

	//stores all the different information inside
	//this will store every single bracket opening without issues
	var Open int = 0 //set to 0 as default

	//stores the future tokens correctly and properly
	//this will allow for better and safer handling without issues
	var toks []lexer.Token = make([]lexer.Token, 0)

	//ranges thorugh the array of tokens
	//this will ensure its done properly without issues happening
	for position := position; position < len(tokens); position++ {
		//saves into the array correctly
		//this will ensure its done properly without issues happening
		toks = append(toks, tokens[position])

		//detects the open seq properly
		//this will properly handle without issues happening
		if tokens[position].TokenType() == open {
			//adds another position onto the open counter
			//this adds another group to close with inside the function
			Open++
		}

		if tokens[position].TokenType() == close {
			//checks if there is any waiting openings
			//this will ensure we close at the right stage
			if Open > 0 {
				//removes a group and continues
				//this will make sure its not ignored properly
				Open--; continue
			} else {
				//returns the token correctly and properly
				return toks, nil
			}
		}


	}

	return nil, errors.New("body opened but never closed")
}