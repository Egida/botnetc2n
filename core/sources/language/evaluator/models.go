package evaluator

import "Nosviak2/core/sources/language/lexer"

//aka variable structure properly
//this stores our different variables properly
type Scope struct {
	//stores the routeName properly
	//allows for the variable to be called with the name given
	Name string //set as type string properly, the literal form
	//stores the value properly
	//we will properly format/comparse the tokens without issues
	TokenValue *lexer.Token
}