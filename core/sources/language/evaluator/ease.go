package evaluator

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/lexer"
	"io"

	"fmt"
	"reflect"
)

//this will properly register everysingle function without issues
//this will ensure its properly done without errors happening making it safe
//variable functions are function which are intended to compeltely return that value only and thats it
func (e *Evaluator) VariableFunction(header string, TokenType lexer.TokenType, literal string) {
	//returns the function array correctly
	//this will ensure its properly done without errors happening
	var Func *Function = &Function{ //sets the params properly
		//sets the header properly
		//this will ensure its done properly without issues
		FunctionName: header,
		//sets the function body correctly
		//this will make sure its completely and properly done
		Function: func(args []lexer.Token, session *sessions.Session, e *Evaluator, wr io.Writer) ([]Object, error) {
			//uses the quick create function to create the objects
			//this will make sure its properly done without issues happening
			return ArrayObject(Object{Type: TokenType, Literal: literal}), nil
		},
	}

	//saves into the function array correctly and properly
	//this will ensure its properly done without errors happening on requests
	e.functions = append(e.functions, *Func) //inserts into the array correctly and properly
}

//correctly tries to insert the functions under a function
//this will ensure its done properly without errors happening on request
func (e *Evaluator) DropPackageStructure(pkg string, s interface{}) {

	//gets the type of information
	//this will make sure its properly done without errors
	inspect := reflect.TypeOf(s) //sets the reflect properly

	//this will store properly without issues
	//this will ensure its done safely and correctly
	var functions map[string]Builtin = make(map[string]Builtin)

	//ranges throughout the fields correctly
	//this will allow us to properly and safely without errors
	for i := 0; i < inspect.NumField(); i++ {
		//properly sorts the information without issues happening
		//this will make sure its done without issues happening so it performs safely
		field, value := inspect.Field(i), reflect.ValueOf(s).Field(i)

		//sets the default type
		//this will ensure its done properly without issues
		var Value lexer.TokenType = lexer.EOF
		var ObjectValue interface{} = nil

		//properly added switch dynamic tool
		switch field.Type.String() { //switch dynamic
		case "string": //string detection
			Value = lexer.String
			ObjectValue = value.String()
		case "int": //int detection
			Value = lexer.Int
			ObjectValue = value.Int()
		case "bool": //bool detection
			Value = lexer.Boolean
			ObjectValue = value.Bool()
		}

		//registers into the map correctly without issues
		//this will make sure its done safely without errors happening
		functions[field.Name] = func(args []lexer.Token, session *sessions.Session, e *Evaluator, wr io.Writer) ([]Object, error) {
			return ArrayObject(Object{Type: Value, Literal: fmt.Sprint(ObjectValue)}), nil
		}; continue
	}

	//inserts the package correctly
	//stops issues happening one request
	e.AddPackage(MakePackage(pkg, functions))
}

//correctly places everysingle object into the functions array
//we will properly use the variableFunction method to insert without issues
func (e *Evaluator) DropStructure(s interface{}) {

	//gets the type of information
	//this will make sure its properly done without errors
	inspect := reflect.TypeOf(s) //sets the reflect type properly

	//ranges through everysingle field
	//this will ensure its properly done without errors
	for i := 0; i < inspect.NumField(); i++ {
		//properly sorts the information without issues happening
		//this will make sure its done without issues happening so it performs safely
		field, value := inspect.Field(i), reflect.ValueOf(s).Field(i)

		//sets the default type
		//this will ensure its done properly without issues
		var Value lexer.TokenType = lexer.EOF

		//properly added switch dynamic tool
		switch field.Type.String() { //switch dynamic
		case "string": //string detection
			Value = lexer.String
		case "int": //int detection
			Value = lexer.Int
		case "bool": //bool detection
			Value = lexer.Boolean
		}

		//tries to correctly create the object
		//this will make sure its properly done without issues
		e.VariableFunction(field.Name, Value, value.String())
	}
}