package shorts

import (
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	//"Nosviak2/core/sources/layouts/functions"
	"Nosviak2/core/sources/layouts/packages"
)

func Register(c map[string]string, e *evaluator.Evaluator) *evaluator.Evaluator {
	//registers all the packages
	//this will make sure they have been completely registered
	for _, packages := range packages.Packages {
		e.AddPackage(&packages) //registers the package properly
	}
	//registers all the functions
	//this will make sure they have been completely registered
	//for _, function := range functions.Functions {
	//	e.AddFunction(function.FunctionName, function.Function) //registers the function properly
	//}
	//registers the custom values properly
	//this will ensure they have all been registered without issues
	for key, value := range c {
		e.VariableFunction(key, lexer.String, value) //registers the object properly
	}

	//returns the output correctly
	//this will ensure its done without issues happening
	return e
}