package parser

import (
	"Nosviak2/core/sources/language/lexer"
)

type Parser struct {
	//stores the lexer information properly
	//this will make sure its done properly without issues happening
	lex *lexer.Lexer //stores the lexer past structures without issues
	//stores the current token position
	//this will allow for better handle without issues happening
	position int
}

//creates the new parser path
//this will ensure its properly done
func MakeParser(l *lexer.Lexer) *Parser {
	//creates the new parser structure
	//this will ensure its properly done without issues
	return &Parser{
		lex: l, //sets the lexer information
	}
}

//properly tries to run the parser without issues
//this will ensure its done correctly without issues allowing easier execution
func MakeParserRun(l *lexer.Lexer, e error) ([]Node, error) {
	//checks if there was an error properly
	//this will ensure its done correctly and safely
	if e != nil {
		//returns the error
		return make([]Node, 0), e
	}

	//runs the parser the correctly
	//this will return the nodes without issues happening
	return MakeParser(l).RunPath()
}

//creates the parser structure without needing the lexer
//this allows for properly handling without issues happening
func MakeTokens(t []lexer.Token, pos int) *Parser {
	if len(t) == 0 {
		return nil
	}
	return &Parser{ //returns the parser structure
		lex: lexer.CreateLexer(t), //creates the new lexer structure
		position: pos, //sets the position without issues
	}
}

//correctly executes the parser using the keywords mode
//this will ensure its properly done without issues happening
func (p *Parser) RunPath() ([]Node, error) {

	//stores the many parsed nodes
	//this will ensure its properly done without issues happening
	var routes []Node = make([]Node, 0) //stored in array

	//ranges through the tokens with for loops
	//this will allow breaking within the system without issues
	for token := 0; token < len(p.lex.Tokens()); token++ {
		//sets the current token properly
		//this will make accessing the token easier without issues
		Token := p.lex.Tokens()[token]
		//this will set the current position
		//makes sure its properly done without issues
		p.position = token //renders the token properly

		//we will only accept idents as keywords
		//this will ensure its properly found a path
		if Token.TokenType() != lexer.Indent {
			if Token.TokenType() == lexer.Comment {
				token += p.Comments() - 1; continue
			}
			//detects text types properly
			//this will be used inside the system
			if Token.TokenType() == lexer.Text {
				//properly works the system without issues
				//this will ensure its done properly without errors
				Line, err := p.textCompression()
				if err != nil {
					//returns the error
					//allows for proper control
					return nil, err
				}

				//skips the safe amount without errors happening
				token += len(Line.Tokens) - 1

				//saves into the array correctly
				//allows for proper control without issues happening
				routes = append(routes, Node{exe: Texture, text: Line}); continue
			} else if Token.TokenType() == lexer.NewBody || Token.TokenType() == lexer.EndBody {
				continue
			}
			//forces it to return an error
			//this will alert that an invalid parse root has happened
			continue
		}

		switch Token.Literal() {

		case IF:
			//parses the conditional statements
			//this will make sure its done without issues
			Con, err := p.parseConditions()
			if err != nil {
				return nil, err
			}

			token += len(Con.Tokens) - 1

			routes = append(routes, Node{exe: Conditi, conditions: Con})
			continue
		case RETURN:
			//properly parses the return statements without issues
			//this will make sure its properly done without errors happening
			ret, err := p.parseReturn()
			if err != nil {
				//returns the error correctly
				return nil, err
			}

			token += len(ret.Tokens)

			//tries to correctly store into the array
			//this will ensure its done properly without issues
			routes = append(routes, Node{exe: ReturnP, declare: nil, path: nil, function: nil, returns: ret}) //saves properly
			continue

		//detects the keyword for function statement
		//makes sure its properly done without issues happening
		case FUNCTION:
			//properly executes the function route
			//this will ensure its done properly without errors
			function, err := p.parseFunction()
			if err != nil {
				//returns the error correctly
				return nil, err
			}

			//adds support for the skipping
			//this will correctly skip the value amount
			token += len(function.Tokens) - 1//skips the amount

			//tries to correctly store into the array
			//this will ensure its done properly without issues
			routes = append(routes, Node{exe: FuncReg, declare: nil, path: nil, function: function, returns: nil}) //saves properly
			continue

		//detects the var/const statements without issues
		//this will make sure we don't have issues with the system
		case VAR, CONST:
			//this will properly handle the declare route
			//this will make sure its properly done without issues happening
			route, err := p.HandleDeclare(p.position)
			if err != nil {
				//returns the error
				return nil, err
			}

			//adds support for the skipping
			//this will correctly skip the value amount
			token += len(route.tokens) - 1 //skips the amount

			//tries to correctly store into the array
			//this will ensure its done properly without issues
			routes = append(routes, Node{exe: Declare, declare: route, returns: nil, path: nil, function: nil}) //saves properly
			continue

		default:
			//parses the function route without issues
			//allows for better & safer handle without issues
			route, err := p.ExecuteFunction()
			if err != nil {
				//returns the error
				return nil, err
			}

			//support for charater skipping properly
			//this will allow for better handle without issues
			token += len(route.Tokens) - 1 //skips the amount

			//tries to correctly store into the array
			//this will ensure its done properly without issues
			routes = append(routes, Node{exe: FuncExe, declare: nil, path: route})
			continue
		}
	}




	return routes, nil
}