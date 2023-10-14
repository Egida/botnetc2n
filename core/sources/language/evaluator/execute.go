package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"
)

//this will guide the information to the exit properly
//this will ensure its done properly without errors happening
type Object struct {
	//stores the object type
	//this will ensure its done properly
	Type lexer.TokenType

	//stores the literal format properly
	//this will ensure its done properly without issues happening
	Literal string //stores the token literal without issues
}

//arrays the object correctly
func ArrayObject(a ...Object) []Object {
	return a
}


//properly tries to correctly execute the function
//this will ensure its done properly without errors happening on request
//this will safely and properly execute the function path which has been requests
func (e *Evaluator) executeFunction(f *parser.FunctionPath) ([]Object, error) {

	//makes sure the amount of labels if valid
	//ensures its properly done without issues
	if len(f.Labels) % 2 == 0 && len(f.Labels) > 3 {
		//returns error for invalid function
		return nil, errors.New("invalid function execution path")
	}

	//this will store the future args without issues happening
	//ensures that its properly done without errors happening
	var args []lexer.Token = make([]lexer.Token, 0)

	//ranges throughout the args without issues
	for _, arg := range f.Arguments {

		if len(arg) <= 0 {
			continue
		}

		//properly compiles the arguments into one
		//this will execute each reach inside the function
		value, err := e.compileTokens(arg, lexer.EOF)
		if err != nil {
			//returns the error correctly
			return make([]Object, 0), err
		}
		//saves into the array correctly
		//this will ensure its saved and we can use it without issues happening
		args = append(args, *value)
	}

	//properly tries to resolve the function
	//this ensures its found the correct function without issues happening
	if len(f.Labels) >= 3 {
		//this will properly sort without isses
		//package and header will be used to find the function without issues
		packages, header := f.Labels[0], f.Labels[len(f.Labels)-1]

		//tries to correctly validate without issues
		//this ensures that the different types have been validates
		if packages.TokenType() != lexer.Indent || header.TokenType() != lexer.Indent {
			//returns the error as its invalid
			//this ensures its done properly without issues happening
			return make([]Object, 0), errors.New("invalid function path location has been passed")
		}

		//tries to correctly find the package
		//this ensures its done properly without issues happeing
		pkg := e.locatePackage(packages.Literal())
		if pkg == nil {
			//returns the package error correctly
			//this will ensure its done properly without errors
			return make([]Object, 0), errors.New("invalid function path location has been passed")
		}

		//tries to correctly find this packageFunction
		//this ensures its correct and valid without issues
		function := pkg.Functions[header.Literal()]
		if function == nil {
			//returns nil as the function doesn't exist
			//this will help with validating functions/information
			return make([]Object, 0), errors.New("invalid function has been passed when it doesn't exist")
		}

		//properly executes the function without issues happening when asked
		//this ensures its done properly without errors happening on calling
		return function(args, e.session, e, e.wr)
	} else{ 
		//gets the function header
		//this will be used to locate the function without issues
		header := f.Labels[0] //gets the header token

		//tries to locate the function name
		//this will ensure its found properly without issues
		function := e.locateFunction(header.Literal())
		if function == nil || header.TokenType() != lexer.Indent {
			//this will properly try to locate the system
			//allows for proper handling without issues happening
			if function == nil { //checks if its a function error
				//trys to correctly locate the system without issues happening
				//this will find the structure for the internal function without issues happening
				internal := e.getInternal(header.Literal())
				//makes sure we correctly found the internal function
				//ensures its properly done without issues
				if internal == nil {
					return nil, errors.New("invalid function request has been made")
				}

				//checks for missing args properly
				//this will make sure its not ignored properly
				if len(internal.Args) != len(args) {
					return nil, errors.New("missing args for calling internal function")
				}

				var guide []Scope = e.GuideScope

				//ranges througout the args given properly
				//this will ensure its properly done without issues happening
				for p := range internal.Args {
					//makes sure the types are correct
					//this will ensure that matching is going safely ahead
					if internal.Args[p].Type != args[p].TokenType() {
						//returns the error correct and properly
						return nil, errors.New("unmatched arguments has been given wanted: "+internal.Args[p].Literal+" "+internal.Args[p].Type.GetString())
					} else {
						guide = append(guide, Scope{Name: internal.Args[p].Literal, TokenValue: &args[p]})
					}
				}

				//transfer all the information without issues happening
				//this will make sure its done properly without issues happening on request
				var ELock = &Evaluator{
					NodeBodies: internal.Nodes,
					GuideScope: guide,
					packages: e.packages,
					functions: e.functions,
					internalFunctions: e.internalFunctions,
					wr: e.wr,
					templateRcog: e.templateRcog,
					session: e.session,
				}

				//executes the node bodies correctly
				//this will ensure its properly done without issues happening
				return ELock.FollowGuide()
			}
			//returns the error properly without issues
			return make([]Object, 0), errors.New("invalid function request has been made")
		}
		//no package has been used here
		//this will be were we have to safely handle the information
		return function.Function(args, e.session, e, e.wr)
	}
	
}