package lexer

import "unicode/utf8"

//this will store the information properly
//allows for better control without issues and error
type Token struct {
	//stores the literal format of the token
	//this will mainly help inside the interpreter
	literal string

	//stores the token type properly
	//this will allow for better control without issues
	tokenType TokenType

	//stores the token position without issues
	//this will allow for better handle without issues
	position *Position

	//stores how large the token is properly
	//this is mostly used in handling the token inside the scanner
	size int //simple item to store how many places it takes up
}


//creates the new token properly
//this will allow us to properly insert into the structure
func (l *Lexer) newToken(literal string, tokenType TokenType) *Token {
	var correct *Position = &Position{column: l.position.column, row: l.position.row}
	//returns the token structure
	//this will allow for better handle without issues
	return &Token{
		//sets the literal format
		//this will allow for proper handle
		literal: literal, //sets the literal format
		//sets the tokentype properly
		//this will ensure its properly done without issues
		tokenType: tokenType,
		//sets the position correctly
		//this will ensure its properly done without issues
		position: correct,
		//sets the size properly
		//this will ensure we skip the correct length
		size: utf8.RuneCountInString(literal) - 1,
	}
}

//creates the new token properly
//this makes sure its properly done without issues happening
func NewToken(literal string, tokenType TokenType, position *Position) *Token {
	//returns the token structure
	//this will allow for better handle without issues
	return &Token{
		//sets the literal format
		//this will allow for proper handle
		literal: literal, //sets the literal format
		//sets the tokentype properly
		//this will ensure its properly done without issues
		tokenType: tokenType,
		//sets the position correctly
		//this will ensure its properly done without issues
		position: position,
		//sets the size properly
		//this will ensure we skip the correct length
		size: utf8.RuneCountInString(literal) - 1,
	}
}

//allows us to access the different information
//this will ensure its properly done without issues
func (t *Token) TokenType() TokenType {
	//returns the tokentype correctly 
	return t.tokenType
}

//allows us to access the literal format
//this will ensure its properly done without issues
func (t *Token) Literal() string {
	//returns the literal format properly
	return t.literal
}

//allows us to access the position
//this will ensure its properly done without issues
func (t *Token) Position() *Position {
	//returns the position structure
	return t.position
}