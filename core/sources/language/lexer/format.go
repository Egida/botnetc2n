package lexer

import "strconv"

//this will properly fill and create a model which can be used
//this will ensure its properly done without issues making it safer without issues
func (l *Lexer) FormattedTokens() []string {
	var formatted []string = make([]string, len(l.Tokens()))
	//for loops throughout all the tokens
	//this will ensure its properly done without issues
	for pos, token := range l.Tokens() {
		//saves into the array properly
		formatted[pos] = token.formatLine()
	}
	//returns the formatted system without issues
	//this will make sure its done safely without issues
	return formatted
}

//properly formats a line with the information
//this will ensure its done safely without issues happening
func (t *Token) formatLine() string {
	//formats the line properly without issues happening
	return Padding(strconv.Itoa(t.position.Row()) + ":" + strconv.Itoa(t.position.Column()), 10) + Padding(strconv.Itoa(int(t.TokenType())), 10) + t.Literal()
}


//pads the system without issues
//this will ensure its properly done without issues
func Padding(text string, wanted int) string {
	//works the amount of system information
	//this will make sure its properly done without issues
	for p := len(text); p < wanted; p++ {
		text += " "
	}
	return text
}