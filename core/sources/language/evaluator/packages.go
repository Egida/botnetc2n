package evaluator

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/lexer"
	"io"
)


type Builtin func(args []lexer.Token, s *sessions.Session, e *Evaluator, wr io.Writer) ([]Object, error)

//stores the package header properly
//this will ensure its done properly without issues happening
type Package struct {
	//stores the package name correctly
	//this will ensure its done properly without issues happening
	Package string //stores the package name correctly and safely
	//stores all the different subfunctions
	//this will ensure its done properly without issues happening
	Functions map[string]Builtin
}

//stores the function information
//stored inside an array allowing for better control without errors
type Function struct {
	//stores the function name correctly
	//allows for better handle without issues
	FunctionName string //stores the functionName 
	//stores the function body correctly
	//this will be executed on request without issues
	Function Builtin
}


//locates the function correctly
//this will return the functions structure without issues
func (e *Evaluator) locateFunction(name string) *Function {
	//ranges throughout the array of functions
	//this will compare each name without issues
	for _, function := range e.functions {

		//compares the different function paths
		//this will ensure its done properly without errors
		if function.FunctionName == name {
			//returns the function body correctly
			return &function
		}
	}
	//returns nil as it wasnt found properly
	//ensures its done properly without errors
	return nil
}

//creates the new package properly
//this will ensure its done properly without issues happening
func MakePackage(name string, functions map[string]Builtin) *Package {
	//creates the new package correctly
	//this will ensure its done properly without issues happening
	return &Package{
		//sets the package name correctly
		//this will ensure its done properly
		Package: name, //sets the package name
		//sets the different function map properly
		//this will ensure its done properly without issues happening
		Functions: functions, //inserts into the map correctly
	}
}

//tries to properly find the package
//this ensures the correct package has been found without issues
func (e *Evaluator) locatePackage(packages string) *Package {
	//ranges throughout the package array
	//this will allow us to properly handle without issues happening
	for pkg := range e.packages {

		//compares the different labels correctly
		//this ensures that the package located is correctly
		if e.packages[pkg].Package == packages {
			//returns the package properly
			return &e.packages[pkg]
		}
	}
	//returns nil as its invalid
	//this ensures its done properly
	return nil
}