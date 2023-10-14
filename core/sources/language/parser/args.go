package parser

import "Nosviak2/core/sources/language/lexer"

func (p *Parser) parseArgs(tokenArgs []lexer.Token) [][]lexer.Token {

	//stores the output correctly
	//this will ensure its done properly without issues happening
	var pushed [][]lexer.Token = make([][]lexer.Token, 1)

	//ranges througout the system correctly
	//this will ensure its done properly without issues happening
	for token := range tokenArgs {

		//properly seperates the new col without issues happening
		//this will ensure its done properly without issues happening
		if tokenArgs[token].TokenType() == lexer.Comma {
			pushed = append(pushed, make([]lexer.Token, 0)); continue //appends properly
		} else {
			//appends into the current one without issues happening
			pushed[len(pushed)-1] = append(pushed[len(pushed)-1], tokenArgs[token])
		}
	}
	//returns the parsed system without issues
	//this will ensure its done properly without errors happening
	return pushed //returns the object properly
}