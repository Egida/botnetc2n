package lexer

import (
	"strings"
)

//stores all the different value ansi escapes
//this will allow the header function to add additional ones
var Escapes map[string]string = map[string]string{
	"\\x1b":"\x1b", "\\u001b":"\u001b", "\\033":"\033", 
	"\\r":"\r", "\\n":"\n", "\\a":"\a", 
	"\\b":"\b", "\\t":"\t", "\\v":"\v",
	"\\f":"\f", "\\007":"\007",
}

//creates the new lexer instance
//this will allow for better controlling without issues
func Make(target string, capture bool) *Lexer {

	//this will properly store the source without issues
	//this will allow us to parse the input into the allowed type
	var TargetParsed [][]string = make([][]string, len(strings.Split(target, "\r\n")))

	//ranges through the source properly
	//this will split the source line by line without issues
	for position := range strings.Split(target, "\r\n") {
		//splits the current line by charater by charater for the main target
		TargetParsed[position] = strings.Split(strings.Split(target, "\r\n")[position], "")
	}

	//returns the lexer structure properly
	//this will allow for more control without issues
	return &Lexer{
		//sets the new parsed source without issues
		//this will make sure its secure and safe without issues
		target: TargetParsed, //sets the target properly
		//sets the position marker up properly
		//this will make sure its properly done without issues
		position: &Position{
			row: 0, column: 0,
		},
		TextCapture: capture,
		//sets the tokens properly
		//this will make sure its properly done
		tokens: make([]Token, 0),
	}
} 


//this will properly run the lexer without issues
//this makes sure its properly executed without issues
func (lex *Lexer) RunTarget() (*Lexer, error) {
	//ranges through each line inside the source
	//this will ensure its safely done without issues making it safer
	for line := range lex.Target() {
		//this will update the line
		//makes sure its properly done
		lex.position.row = line

		//ranges through each col properly
		//this will allow skipping so we can skip big types
		for column := 0; column < len(lex.Target()[line]); column++ {
			//this will update the column properly
			//allows for better control without issues
			lex.position.column = column

			//checks if the text capture is enabled
			//this will make sure its not ignored without issues
			if lex.TextCapture && lex.Target()[lex.position.row][lex.position.column] != "<" {
				if lex.position.column == len(lex.Target()[lex.position.row]) - 1 && len(lex.Target()) - 1 > lex.position.row {
					lex.tokens = append(lex.tokens, *lex.newToken(lex.Target()[lex.position.row][lex.position.column]+"\r\n", Text)); continue
				}
				//inserts into the token array without issues happening
				//this will ensure its done properly without issues happening
				lex.tokens = append(lex.tokens, *lex.newToken(lex.Target()[lex.position.row][lex.position.column], Text))
				continue //continues the loop properly without issues
			}
			//runs the charater anaylise properly
			//this will ensure its properly done without issues
			token, err := lex.charater()
			//makes sure there wasn't an error
			//this will also ensure the token doesnt equal nil
			if err != nil || token == nil {
				//returns the error if that happened
				//this will ensure they know about the reason
				if err != nil {
					//returns the error correctly
					//this will make sure we know about the error
					return nil, err
				} else {
					//returns the token error properly
					//this will make sure they know about the reason
					continue
				}
			}
			//skips the certain positions
			//this will make sure we don't ignore it without reasons
			column += token.size //skips the positions without issues

			//inserts the tokens correctly
			//this will allow us to properly handle without issues
			lex.tokens = append(lex.tokens, *token); continue
		}
	}

	return lex, nil
}

//tries to correctly lex the current charater
//this will ensure its safe without issues making it properly handled
func (lex *Lexer) charater() (*Token, error) {
	//switchs and allows for proper control
	//this will ensure its properly handled without issues
	switch rune(lex.Target()[lex.position.row][lex.position.column][0]) {

	case '?':
		if lex.peek() == '>' {
			//disables the text capture
			//this will ensure its done properly without issues happening
			lex.TextCapture = true
			//creates the token correctly without issues happening
			//this will ensure its done without issues happening
			return lex.newToken("?>", EndBody), nil
		} else if lex.TextCapture {
			return lex.newToken(lex.target[lex.position.row][lex.position.row], Text), nil
		}
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("?", Question), nil
	case ',':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken(",", Comma), nil
	case '.':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken(".", Dot), nil
	case ';':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken(";", SemiColon), nil
	case ':':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken(":", Colon), nil
	//adds support for greater
	//properly supports greaterThan of greaterEqual
	case '>':
		//checks for greaterEqual tokens
		//this will make sure they are properly fitted
		if lex.peek() == '=' {
			//returns the token correctly
			//this will allow for better handle without issues
			return lex.newToken(">=", GreaterEqual), nil
		}
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken(">", GreaterThan), nil
	case '<':
		//checks for lessEqual tokens
		//this will make sure they are properly fitted
		if lex.peek() == '=' {
			//returns the token correctly
			//this will allow for better handle without issues
			return lex.newToken("<=", LessEqual), nil
		} else if lex.peek() == '?' {
			//enables the text capture
			//this will ensure its done properly without issues happening
			lex.TextCapture = false
			//creates the token correctly without issues happening
			//this will ensure its done without issues happening
			return lex.newToken("<?", NewBody), nil
		} else if lex.TextCapture {
			return lex.newToken(lex.Target()[lex.position.row][lex.position.column], Text), nil
		}
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("<", LessThan), nil
	//adds support for exclamation
	//this will add suport for Bang & NEQ
	case '!':
		//gets the next charater inside the array
		if lex.peek() == '=' {
			//returns the token correctly
			//this will allow for better handle without issues
			return lex.newToken("!=", NEQ), nil
		} 

		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("!", Bang), nil

		//adds support for equals
		//this will add support for Equals & Assign
	case '=':
		//gets the next charater inside the array
		if lex.peek() == '=' {
			//returns the token correctly
			//this will allow for better handle without issues
			return lex.newToken("==", Equal), nil
		}

		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("=", Assign), nil

	case '+':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("+", Addition), nil
	case '-':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("-", Subtraction), nil
	case '/':
		if lex.peek() == '/' {
			//checks for comments properly
			//allows for comments to be created
			return lex.newToken("//", Comment), nil
		}
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("/", Divide), nil
	case '*':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("*", Multiply), nil
	case '(':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("(", BracketOpen), nil
	case ')':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken(")", BracketClose), nil
	case '[':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("[", BraceOpen), nil
	case ']':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("]", BraceClose), nil
	case '{':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("{", ParentheseOpen), nil
	case '}':
		//returns the token correctly
		//this will allow for better handle without issues
		return lex.newToken("}", ParentheseClose), nil


	//string support properly
	//this will allow for string tokens without issues
	case '"':
		//this will work the token type with the must finish
		//this will make sure its properly handled without issues happening
		return lex.workType(stringChecker, String, true, 1) //returns the token properly
	default:
		//adds support for more context types without issues
		if !indentChecker(lex.Target()[lex.position.row][lex.position.column], 0) {
			tok, err := lex.workType(indentChecker, Indent, false, 0)
			//checks for the boolean state without issues
			//this will properly check without issues happening
			if tok.Literal() == "true" || tok.Literal() == "false" {
				//sets type to boolean
				tok.tokenType = Boolean
			}
			//returns the token and error
			//this will allow for better handle without issues
			return tok, err
		} else if !intChecker(lex.Target()[lex.position.row][lex.position.column], 0) {
			return lex.workType(intChecker, Int, false, 0)
		}
	}

	return nil, nil
}



//this will properly work out the system without issues
//this allows for better control without issues happening
func (l *Lexer) workType(f func(t string, p int) bool, tokenType TokenType, mustFinish bool, ahead int) (*Token, error) {
	//stores all the supported charaters without issues
	//this allows for better control without issues happening
	var serial string = ""
	//ranges through from the col support without issues
	//this allows for properly control without issues happening
	for columnProfile := l.position.column + ahead; columnProfile < len(l.Target()[l.position.row]); columnProfile++ {
		//checks if the input meets in the catorgys properly
		//this will either save into the string or finish the function
		if f(l.Target()[l.position.row][columnProfile], columnProfile - l.position.column) {
			//returns the information without issues happening
			//allows for better control without issues happening
			if tokenType == 2 {
				serial = "\"" + serial + "\""
			}
			return l.newToken(serial, tokenType), nil
		} else {
			//saves into the string propetly without issues
			serial += l.Target()[l.position.row][columnProfile]
		}
	}

	//this will ensure an error is returned
	//makes sure we don't ignore the errors
	if mustFinish {
		//returns the error correctly and properly
		return nil, ErrTypeOpenedNotClosed
	}
	//returns the token properly
	//this will make sure its safely done without issues
	return l.newToken(serial, tokenType), nil
}