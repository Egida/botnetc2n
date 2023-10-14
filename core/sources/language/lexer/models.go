package lexer

//stores the main lexer information
//this will allow the lexer at anypoint to control items without issues
type Lexer struct { //lexer items storage structure
	//stores the source being passed inwards
	//this will allow for better controlling without issues
	target [][]string //stores the target source without issues
	//stores the position properly without issues
	//this will make sure its properly setup without issues
	position *Position
	//stores all the token we have created
	//this will allow us to properly track without issues
	tokens []Token
	//this will store if capture text is enabled
	//allows for proper handling without issues happening on request
	TextCapture bool
}

func CreateLexer(tokens []Token) *Lexer {
	return &Lexer{
		target: make([][]string, 0),
		position: &Position{
			row: tokens[len(tokens)-1].Position().Row(),
			column: tokens[len(tokens)-1].Position().Column(),
		},
		tokens: tokens,
	}
}

//stores the position of an object
//this will allow for better control without issues
type Position struct {
	//this will store the col & row of the item
	//makes sure we can correctly find the object without issues
	column, row int
}

//allows us to access the position information
//this certain function allows us to access the col without issues
func (p *Position) Column() int {
	//returns the column properly
	return p.column
}

//allows us to access the position information
//this certain function allows us to access thw row without issues
func (p *Position) Row() int {
	//returns the row properly
	return p.row
}

//allows us to access the target source properly
//this will allow for other function to access the lexer source
func (l *Lexer) Target() [][]string {
	//returns the lexer source properly
	return l.target
}

//allows us to access the position of the scanner
//this will make sure its properly done without issues
func (l *Lexer) Position() *Position {
	//returns the position correctly
	return l.position
}

//this will allow us to access the tokens 
//allows for custom parsers with this lexer
func (l *Lexer) Tokens() []Token {
	//returns the tokens properly
	return l.tokens
}