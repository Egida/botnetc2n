package parser

import (
	"Nosviak2/core/sources/language/lexer"
)


//peeks the next token properly
//this will allow the system to view the next token
func (p *Parser) peek(pos int) *lexer.Token {
	//checks the system position
	//makes sure it doesn't pass the position proc
	if p.position + pos > len(p.lex.Tokens()) {
		return nil
	}
	//gets the next token inside the array
	//this will return to the main function
	return &p.lex.Tokens()[p.position + pos]
}