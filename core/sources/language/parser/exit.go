package parser

import (
	"Nosviak2/core/sources/language/lexer"
)

//allows for functions to properly return objects
//allows for better control without issues happening on request
type ReturnReply struct { //stores the values
	//stores the different values being returned
	//this will allow for better system handling without issues
	Values [][]lexer.Token
	Tokens []lexer.Token
}

//properly parses the return values without issues
//ensures that its properly done without errors happening
func (p *Parser) parseReturn() (*ReturnReply, error) {
	//properly parses the statement without issues
	//this ensures its properly done without errors
	var returnValue *ReturnReply = &ReturnReply{
		Values: make([][]lexer.Token, 1), //stores the different values used properly
	}
	//ranges throughout the system properly
	//this ensures it done properly without errors happening
	for tok := p.position + 1; tok < len(p.lex.Tokens()); tok++ {
		returnValue.Tokens = append(returnValue.Tokens, p.lex.Tokens()[tok])
		//makes sure the token stays within the rules
		//this will make sure its valid without issues
		if p.lex.Tokens()[tok].TokenType() == lexer.Comma {
			//creates the new array correctly and properly without issues
			returnValue.Values = append(returnValue.Values, make([]lexer.Token, 0)); continue
		}
		//saves correctly into the array properly without issues
		//this will ensure its properly done without errors happening
		returnValue.Values[len(returnValue.Values)-1] = append(returnValue.Values[len(returnValue.Values)-1], p.lex.Tokens()[tok])
	}

	//returns the structure properly
	//ensures its done properly without issues happening
	return returnValue, nil
}