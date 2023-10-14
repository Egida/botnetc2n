package evaluator

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"fmt"
	"io"
	"strings"
)

//stores the template driver information
//this will properly and safely store the template driver execution route
func (e *Evaluator) driver(tag string, s *sessions.Session, wr io.Writer) (int, error) {

	//validates if the request is type function
	//if the request is type function we will route accordingly without issues
	if strings.Contains(tag, "(") && strings.Contains(tag, ")") && strings.Split(tag, "")[0] != "$" {
		//properly pushes the information through the lexer
		//this will help with parsing without issues happening
		l := lexer.Make(tag, false) //creates the new lexer properly

		//correctly executes the lexer without issues
		//this will ensure its safely executed without issues happening
		if _,err := l.RunTarget(); err != nil {
			//prints directly to the main object without issues happening on request
			//this will make sure the users knows about the issue when it happens on request
			return wr.Write([]byte(fmt.Sprintf("[#Unknown function %s#]", tag)))
		}

		//creates and properly runs the parser without issues happening
		//this will ensure its done properly without issues happening on request
		nodes, err := parser.MakeParser(l).RunPath()
		if err != nil || nodes[0].Route() != parser.FuncExe {
			//prints directly to the main object without issues happening on request
			//this will make sure the users knows about the issue when it happens on request
			return wr.Write([]byte(fmt.Sprintf("[#Unknown function %s#]", tag)))
		}

		//properly tries to execute the route without issues happening
		//this will ensure its safely executed without issues happening on request
		objects, err := MakeEvalWithMore(nodes, e.GuideScope, e.wr, e.templateRcog, e.packages, e.functions, e.internalFunctions, e.session).executeFunction(nodes[0].Path())
		if err != nil {
			//prints directly to the main object without issues happening on request
			//this will make sure the users knows about the issue when it happens on request
			return wr.Write([]byte(fmt.Sprintf("[#Unknown function %s#]", tag)))
		}

		//correctly tries to diag the error
		//this will make sure its not ignored without issues
		if len(objects) <= 0 { //length isnt 0 properly
			return 0, nil
		}
	
		return wr.Write([]byte(lexer.AnsiUtil(objects[0].Literal, lexer.Escapes)))
	}

	if strings.Split(tag, "")[0] != "$" {
		return wr.Write([]byte(fmt.Sprintf("[#Unknown factor found within template %s^#]", tag)))
	}

	//replaces the doller sign properly
	//allows for better system handle without issues
	tag = strings.Replace(tag, "$", "", -1)
	

	//tries to correctly find the object within
	//this will ensure its done properly without errors happening
	object, err := e.findScope(tag) //tries to correctly find the object
	if err != nil {
		//prints directly to the main object without issues happening on request
		//this will make sure the users knows about the issue when it happens on request
		return wr.Write([]byte(fmt.Sprintf("[#Unknown tag %s#]", tag)))
	}

	//writes the literal direct to the writer
	//this will ensure the user has got access without issues happening
	return wr.Write([]byte(lexer.AnsiUtil(object.TokenValue.Literal(), lexer.Escapes)))
}