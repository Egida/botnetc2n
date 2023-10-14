package evaluator

import "Nosviak2/core/sources/language/lexer"

//tries to correctly validate the token given as operator
//this will allow us to completely validate without issues happening
func (e *Evaluator) validateOperator(t *lexer.Token) bool {
	//switchs the information properly and safely
	switch t.TokenType() {
	//tries to correctly validate the operator given
	//this will compare the tokenType with the different options
	case lexer.Addition, lexer.Subtraction, lexer.Divide, lexer.Multiply:
		return true
	default:
		return false
	}
}