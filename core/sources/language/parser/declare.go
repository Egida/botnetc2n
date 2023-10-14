package parser

import (
	"Nosviak2/core/sources/language/lexer"
)

//stores the information about the declare root
//this will ensure its properly done without issues happening
type DeclareRoute struct {
	//stores the declare route name
	//this will ensure its properly done without issues happening
	routeName string //stores the routeName without issues happening

	//this option is properly optional for better configs
	//this will allow the system to set variable types without issues
	givenType lexer.TokenType

	//stores what has been assigned without issues
	//this will store all the different arguments without issues
	values []lexer.Token //stored in type array without issues happening

	//stores all the tokens which have been collected inside
	//this will ensure its properly done without errors happening
	tokens []lexer.Token
}


func (d *DeclareRoute) Tokens() []lexer.Token {
	return d.tokens
}

//allows us to access the routeName
//this will properly allow us to access it
func (d *DeclareRoute) RouteName() string {
	//returns the string properly
	return d.routeName
}

//allows us to access the given lexer type
//this will properly handle without issues happening
func (d *DeclareRoute) GivenType() lexer.TokenType {
	//returns the tokenType properly
	return d.givenType
}

//allows us to access the values lexer type
//this will properly handle without issues happening
func (d *DeclareRoute) Values() []lexer.Token {
	//returns the values properly
	return d.values
}

//properly parsers the declare route without issues happening
//this will make sure its properly done without issues happening 
func (p *Parser) HandleDeclare(position int) (*DeclareRoute, error) {
	//this will be slowly filled with information
	//allows for better information without issues happening
	var fill *DeclareRoute = &DeclareRoute{}//fills overtime

	//saves the first keyword into array
	//this will properly be saved into the array
	fill.tokens = append(fill.tokens, p.lex.Tokens()[p.position])

	//gets the next position inside the array
	//this will hold the declareroutes routeName properly
	routeName := p.peek(1)
	
	//makes sure the next charater is an indent
	//this will ensure its properly done without issues
	if routeName.TokenType() != lexer.Indent || routeName == nil {
		//returns the error correctly
		//only routeName's can be indents...
		return nil, ErrRouteNameInvalid
	} else {
		//this will properly allow the position to be skipped forward
		//this allows us to properly handle without issues happening on ref...
		p.position++ //adds one position ontop without issues happening
		fill.routeName = routeName.Literal() //sets the filled name properly
		fill.tokens = append(fill.tokens, *routeName) //saves into the array correctly
	}

	maybeType := p.peek(1)
	//this will ensure the next token isn't a nil pointer
	//makes sure we don't accept nil pointers without issues happening
	if maybeType == nil {
		//returns the error correctly
		return nil, ErrNilPointer
	}
	//we will try to check if they have given it a type name
	//this will ensure its properly done without issues being given
	if maybeType.TokenType() != lexer.Assign {
	
		//this will ensure that the route is safe
		//this makes sure that the guid isn't invalid
		if maybeType.TokenType() != lexer.Indent {
			//this will return the error properly
			//makes sure they know about the possible issues happening
			return nil, ErrInvalidAssign
		}

		//this will show that their is a possible type assign happening
		//this will be logged & formatted into the structure safely without issues
		if maybeType.Literal() != STRING && maybeType.Literal() != INT && maybeType.Literal() != BOOLEAN {
			//returns the invalid keyword type error without issues
			return nil, ErrInvalidTypeKeyword
		}

		//this will now be accepted as a possible type without issues happening
		//this will be enforced into the state without issues happening
		if maybeType.Literal() == STRING {
			//string support properly
			fill.givenType = lexer.String
		} else if maybeType.Literal() == INT {
			//int support properly
			fill.givenType = lexer.Int
		} else if maybeType.Literal() == BOOLEAN {
			//boolean support properly
			fill.givenType = lexer.Boolean
		}

		//this will store into the tokens without issues happening
		//this will ensure we have access to the routes tokens without errors happening
		fill.tokens = append(fill.tokens, *maybeType)

		//skips another position without issues
		//this will safely and properly skip the position
		next := p.peek(1)
		//checks for nil pointers
		//this will ensure its properly done
		if next == nil {
			//returns the error correctly
			return nil, ErrNilPointer
		}
		p.position++; fill.tokens = append(fill.tokens, *next)

		//skips a position ahead properly
		//this will skip 2 spots ahead without issues happening
		p.position += 2 //based cause of the value selection support
	} else {
		//skips one position ahead properly
		//this will only skip 1 to stop value ignoring
		p.position += 2 //skips the positions
		fill.tokens = append(fill.tokens, *maybeType) //saves into the array
		fill.givenType = lexer.EOF
	}


	//loops through the rules without issues happening
	//this will make sure its properly done without issues happening
	for proc := p.position; proc < len(p.lex.Tokens()); proc++ {
		//saves into the system without issues
		//this will ensure we don't miss the system finish
		if p.lex.Tokens()[proc].TokenType() == lexer.SemiColon || fill.tokens[len(fill.tokens)-1].Position().Row() < p.lex.Tokens()[proc].Position().Row() {
			break
		} else {
			//saves into the array correctly and safely
			//this will ensure we have access to it without errors happening
			fill.values = append(fill.values, p.lex.Tokens()[proc])
		}
	}

	//saves into the tokens system properly
	//this will allow us to properly messure the amount
	fill.tokens = append(fill.tokens, fill.values...)

	return fill, nil
}