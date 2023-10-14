package lexer

import "strings"

//tries to correctly peek the next charater
//this will make sure its properly done without issues
func (l *Lexer) peek() rune {
	//checks for end of line information
	//this will make sure its properly done without issues
	if len(l.Target()[l.position.row]) <= l.position.column + 1 {
		//returns nil as its invalid
		//this will make sure nothing else happens within the system
		return 0
	}
	//returns the next charater properly without issues
	return rune(l.Target()[l.position.row][l.position.column + 1][0])
}

//completely adds support for the ansi escapes
//this will ensure they have been sorted without issues happening
func AnsiUtil(src string, escapes map[string]string) string {
	//ranges throughout all the escapes properly
	//this will ensure its done properly without issues happening
	for escape := range escapes {
		//properly sorts the escapes without issues
		src = strings.ReplaceAll(src, escape, escapes[escape])
	} //returns the source once done
	return src
}