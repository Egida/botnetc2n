package lexer

import (
	"strings"
)

//stores information about the token type
//this will allow us to set different token types without issues
type TokenType int 


const (
	EOF				TokenType = -1
	Int				TokenType = 1 // 1, 2, 3, 4, 5, 6, 7, 8, 9, 0
	String 			TokenType = 2 // "qwertyuiopasdfghjklzxcvbnm"
	Indent			TokenType = 3 // qwertyuiopasdfghjklzxcvbnm
	Boolean			TokenType = 4 // true, false
	Text 			TokenType = 5 // >?Hello boss how are you<?
	

	NEQ				TokenType = 10 // !=
	Bang 			TokenType = 11 // !
	Equal 			TokenType = 12 // ==
	Assign 			TokenType = 13 // =
	Addition		TokenType = 14 // + //op
	Subtraction 	TokenType = 15 // - //op
	Divide 			TokenType = 16 // / //op
	Multiply    	TokenType = 17 // * //op
	BracketOpen 	TokenType = 18 // (
	BracketClose	TokenType = 19 // )
	BraceOpen 		TokenType = 20 // [
	BraceClose 		TokenType = 21 // ]
	ParentheseOpen  TokenType = 22 // {
	ParentheseClose	TokenType = 23 // }
	GreaterThan 	TokenType = 24 // >
	GreaterEqual	TokenType = 25 // >=
	LessThan		TokenType = 26 // <
	LessEqual		TokenType = 27 // <=
	SemiColon		TokenType = 28 // ;
	Colon 			TokenType = 29 // :
	Modulus 		TokenType = 30 // %
	Dot 			TokenType = 31 // .
	Comma 			TokenType = 32 // ,
	Comment 		TokenType = 33 // //
	Question		TokenType = 34 // ?

	NewBody 		TokenType = 40 // <?
	EndBody			TokenType = 41 // ?>
)


func (t *TokenType) GetString() string {
	switch *t {

	case String:
		return "string"
	case Int:
		return "int"
	case Boolean:
		return "bool"
	}
	return "EOF"
}

var (
	//this will check if the object passed inwards is a string
	//makes sure we don't ignore the end point without issues making it safe
	stringChecker func(t string, p int) bool = func(t string, p int) bool {
		if t == "\"" {
			//returns true
			//this will make sure they know its a string
			return true
		} else {
			//returns false as it isn't a string
			//returns false correctly and properly
			return false
		}
	}

	//this will check if the object passed inwards is a indent
	//makes sure we don't ignore the end point without issues making it safe
	indentChecker func(t string, p int) bool = func(t string, p int) bool {
		//switchs the current input without issues
		//this will make sure its properly done without issues
		switch strings.ToLower(t) {
		//stores all the different charaters on a qwerty keyboard
		case "q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m", "_":
			//returns false as its valid
			return false
		default:
			if p >= 1 {
				if !intChecker(t, p) {
					return false
				} else {
					return true
				}
			}
			//returns true as its invalid
			return true
		}
	}


	//checks if the object is an integer
	//this will correctly work the integer out without issues
	intChecker func(t string, p int) bool = func(t string, p int) bool {
		//switchs the value without issues
		//this will make sure its properly done
		switch t {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			//returns false as its valid
			return false
		default:
			return true
		}
	}

)