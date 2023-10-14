package parser

import (
	"Nosviak2/core/sources/language/lexer"
)

//properly formats all the tokens which are one line
//this allows for proper handling without issues happening on request
type LineFormat struct {
	//stores the text inside here
	//this will allow for properly handling without issues happening
	Lines []string //stored in type arrays properly

	//stores all the tokens used inside the array properly
	//this will allow for properly handle without issues without errors
	Tokens []lexer.Token
}

//compresses the text into lines properly
//allows for proper execution without issues happening allow for safely textures
func (p *Parser) textCompression() (*LineFormat, error) {

	//stores the future structure without issues happening
	//this will store all the future information without errors
	var line *LineFormat = &LineFormat{}

	//ranges througout the systems properly
	//this will try to format everysingle texture without errors happening
	for text := p.position; text < len(p.lex.Tokens()); text++ {
		//stores the current token
		//allows for proper handling without errors
		current := p.lex.Tokens()[text]


		//validate source without issues happening
		//this will properly be formatted without errors happening
		if current.TokenType() == lexer.Text && current.Position().Row() == p.lex.Tokens()[p.position].Position().Row() {
			//saves into the system properly without issues happening
			//this will also properly save into 
			line.Lines = append(line.Lines, current.Literal())
			line.Tokens = append(line.Tokens, current) //saves current into tokens aswell
		} else {
			return line, nil
		}
	}

	//returns the system information without issues
	//this will ensure its properly and safely done without issues
	return line, nil
}