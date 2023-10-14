package parser 


//properly handles/skips the comment section
//this will make sure it isn't passed into the interpreter with issues possibly happening
func (p *Parser) Comments() int {
	var i int = 0
	//ranges throughout the tokens until newLine
	//this will make sure its done without issues happening
	for i = p.position; i < len(p.lex.Tokens()); i++ {

		//checks for the new line properly
		//makes sure it isn't ignored without issues happening
		if p.lex.Tokens()[i].Position().Row() > p.lex.Tokens()[p.position].Position().Row() {
			return i - p.position //works out the size without issues happening
		}
	}

	return i - p.position //works out the size without issues happening
}