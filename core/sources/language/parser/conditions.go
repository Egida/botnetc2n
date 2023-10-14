package parser

import (
	"Nosviak2/core/sources/language/lexer"
	"errors"
)

//stores the information properly without issues happening
//this will make sure its done properly without issues producing
type Conditional struct {
	//stores all the different token sections
	//this will allow for proper control without issues happening
	Sections [][]lexer.Token //stored in this array for multi token usage
	PositiveBody []Node //this will execute when the actions are true
	NegativeBody []Node //this will execute when the actions are false
	Tokens []lexer.Token
}


//properly tries to parse the conditions
//this will make sure its done without errors happening
func (p *Parser) parseConditions() (*Conditional, error) {
	//this will store all possible information without issues happening
	//makes sure its done without issues happening on request
	var Con *Conditional = &Conditional{}
	var Object [][]lexer.Token = make([][]lexer.Token, 1)
	var exitProperly bool = false; var Rotations int = 0
	Con.Tokens = append(Con.Tokens, p.lex.Tokens()[p.position])

	//ranges throughout all the objects without issues
	//this will make sure its done without errors happening on request
	for Rotations = p.position + 1; Rotations < len(p.lex.Tokens()); Rotations++ {
		Con.Tokens = append(Con.Tokens, p.lex.Tokens()[Rotations])

		//tries to detect the closure
		//this will make sure its done properly without issues
		if p.lex.Tokens()[Rotations].TokenType() == lexer.ParentheseOpen {
			exitProperly = true; break //breaks from the loop properly
		}

		//switchs the token type properly for value handling
		switch p.lex.Tokens()[Rotations].TokenType() { //operator detection properly
		case lexer.NEQ, lexer.Equal, lexer.Addition, lexer.Subtraction, lexer.Divide, lexer.Multiply, lexer.GreaterThan, lexer.GreaterEqual, lexer.LessThan, lexer.LessEqual, lexer.Modulus:
			Object = append(Object, []lexer.Token{p.lex.Tokens()[Rotations]}) //creates the new array properly without issues
			Object = append(Object, make([]lexer.Token, 0)) //makes the new array correctly without issues
		default: //object detection properly
			//saves into the array correctly and properly without issues happening
			Object[len(Object)-1] = append(Object[len(Object)-1], p.lex.Tokens()[Rotations]); continue
		}
	}

	//this will detect if an error
	//makes sure we report the error properly
	if !exitProperly || len(p.lex.Tokens()) <= Rotations { //returns thr error
		return nil, errors.New("syntax error involding a conditions statement")
	} else {
		Con.Sections = Object //saves into array slot
	}


	//properly tries to render the tokens
	//this will make sure its done without issues happening
	Tokens, err := p.ReadBodyUntil(Rotations + 1, p.lex.Tokens(), lexer.ParentheseOpen, lexer.ParentheseClose)
	if err != nil {
		//returns the error
		return nil, err
	}
	Con.Tokens = append(Con.Tokens, Tokens...)

	//tries to parse the new body without issues
	//this will make sure its done properly without issues happening
	PosNodes, err := MakeParser(lexer.CreateLexer(Tokens)).RunPath()
	if err != nil { //error handling
		//returns the error correctly
		return nil, err
	} else {
		Con.PositiveBody = PosNodes
	}

	if len(p.lex.Tokens()) < Rotations+len(Tokens)+1 || len(p.lex.Tokens()) < Rotations+len(Tokens)+2 {
		return Con, nil
	}

	//this will make the else body parsing start
	//makes sure its not done without issues happening on request
	if p.lex.Tokens()[Rotations+len(Tokens)+1].Literal() == ELSE {
		Con.Tokens = append(Con.Tokens, p.lex.Tokens()[Rotations+len(Tokens)+1])
		Con.Tokens = append(Con.Tokens, p.lex.Tokens()[Rotations+len(Tokens)+2])
		
		//properly tries to render the negTokens
		//this makes sure its done without issues happening
		NegTokens, err := p.ReadBodyUntil(Rotations+len(Tokens)+3, p.lex.Tokens(), lexer.ParentheseOpen, lexer.ParentheseClose)
		if err != nil {//returns the error
			return nil, err
		}

		Con.Tokens = append(Con.Tokens, NegTokens...)
		//tries to parse the new body without issues
		//this will make sure its done properly without issues happening
		NegNodes, err := MakeParser(lexer.CreateLexer(NegTokens)).RunPath()
		if err != nil { //error handling
			//returns the error correctly
			return nil, err
		} else { //saves into neg nodes properly
			Con.NegativeBody = NegNodes
		}
	}

	return Con, nil
}