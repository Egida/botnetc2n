package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"
	"strconv"
	"strings"
)

//compiles the selection of the tokens into one properly
//this will allow for proper controlling without issues happening
func (e *Evaluator) compileString(tokens []lexer.Token) (*lexer.Token, error) {
	//stores all the different tokens properly without issues
	//allows for better controlling without issues happening...
	var operators []lexer.Token = make([]lexer.Token, 0) //creates the array
	var objects []lexer.Token = make([]lexer.Token, 0) //creates the array

	//properly ranges through all the tokens
	//this will execute and compile the string literal format
	for position := 0; position < len(tokens); position++ {


		//allows easier control on detecting operators
		//this will allow for properly controlling without issues happenign
		if e.validateOperator(&tokens[position]) {
			//saves the object into the array
			//this will properly try to save into the array
			operators = append(operators, tokens[position])	
		} else {
			//stores the token properly
			//this makes sure its done properly
			var token = tokens[position]

			//this will make sure its done properly
			//this ensures its properly done without errors
			if token.TokenType() == lexer.Indent { //supports for lexer indents
				//searchs for the variable without issues happening
				//this will ensure its done properly without errors happening
				s, err := e.findScope(token.Literal())
				if err != nil {
					//tries to parser for function
					//allows for proper handle without issues
					path, err := parser.MakeTokens(tokens, 0).ExecuteFunction()
					if err != nil {
						return nil, err
					}
					//tries to correctly execute the function
					//allows for proper subject handle without issues
					mem, err := e.executeFunction(path)
					if err != nil {
						//returns the error
						return nil, err
					}
					//validates the length
					//makes sure its correctly done without issues
					if len(mem) > 1 {
						//returns the error properly without issues
						return nil, errors.New("mismatch type, one wanted, "+strconv.Itoa(len(mem))+" given")
					}

					//inserts into the array correctly and properly
					//this will ensure its done properly and we can access it
					token = *lexer.NewToken(mem[0].Literal, mem[0].Type, token.Position())
					position += len(path.Tokens) - 1
				} else {
					//updates the value correctly
					//allows for proper supportive information
					token = *s.TokenValue
				}
			}

			//makes sure all the types are the same without issues
			//this will make sure that they all match without issues happening
			if token.TokenType() != lexer.String { //tries to match the information without issues
				return nil, ErrTypeMatchSeq //returns the error correctly
			}

			//saves into the array correctly
			//this will ensure its properly done without issues happening
			objects = append(objects, token) //saves into the system
		}
	}

	//temp object to store the outcome
	//this will allow for better handle within the system
	var value string = objects[0].Literal() //creates the new memory object

	//ranges through the different objects properly
	//this will make sure its properly done without issues
	for stringAgent, operator := 1, 0; stringAgent < len(objects) && operator < len(operators); stringAgent, operator = stringAgent + 1, operator + 1 {
		//switchs the systems without issues happening
		//this will make sure its properly done without issues
		switch operators[operator].TokenType() {

		case lexer.Addition: //addition in strings properly
			//properly sorts without issues happening
			//makes sure its properly done without issues
			value = value + objects[stringAgent].Literal(); continue
		}
	}

	//returns the new token structure properly
	//this will allow us to properly handle without errors happening
	return lexer.NewToken(strings.ReplaceAll(value, "\"", ""), lexer.String, objects[0].Position()), nil
}