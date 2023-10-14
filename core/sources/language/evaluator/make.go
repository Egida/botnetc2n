package evaluator

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"

	"io"
	"strings"

	"github.com/valyala/fasttemplate"
)

//stores the execution information
//this will guide us inside the function bodies inside the eval
type Evaluator struct {
	//takes the array of systems
	//this will properly handle inside the eval
	NodeBodies []parser.Node
	//stores the array of tokens properly
	//this will allow for better handle without issues
	GuideScope []Scope
	//stores all the different packages correctly
	//this will ensure its done properly without issues happening
	packages []Package
	//stores all the builtin functions properly
	//this will be used when trying to execute the functions
	functions []Function
	//stores all the internal functions properly
	//this will ensure its properly and completely done
	internalFunctions []InternalFunction
	//this will store what we write to properly
	//ensures its properly done without issues happening
	wr io.Writer
	//stores the template engine recognition plants
	//this will help us to set custom recog plants without issues
	templateRcog [2]string
	//stores the session information correctly
	//this will make sure its foundc correctly without issues happening
	session *sessions.Session
}

//stores the eval guie scope properly
//this will ensure its done correctly without issues
func (e *Evaluator) Guide() []Scope { //allows for access
	return e.GuideScope //returns the array
}


func (e *Evaluator) Add(f []InternalFunction) {
	e.internalFunctions = append(e.internalFunctions, f...)
}


func (e *Evaluator) GrabFunc() []InternalFunction {
	return e.internalFunctions
}

//creates the eval with more information
//this will allow for better handling without issues
func MakeEvalWithMore(bodies []parser.Node, scope []Scope, wr io.Writer, prefixs [2]string, packages []Package, functions []Function, internalFunctions []InternalFunction, session *sessions.Session) *Evaluator {
	//returns the structure properly
	return &Evaluator{
		//takes the node bodies properly
		//this will help us to properly handle
		NodeBodies: bodies,
		//takes the scope and places into var
		//this will store our variables which have been found
		GuideScope: scope,
		//sets the prefixs properly without issues
		//this will make sure its done without issues happening
		templateRcog: prefixs,
		//stores the writer correctly without issues
		//this will ensure its done properly without issues happening
		wr: wr,
		//creates the default information correctly
		//this will ensure its done properly without issues happening
		packages: packages,
		functions: functions,
		internalFunctions: internalFunctions,
		session: session,
	}
}

//makes the evaluator properly and safely
//this will ensure its done correctly without issues
func MakeEval(bodies []parser.Node, scope []Scope, wr io.Writer, prefixs [2]string, session *sessions.Session) *Evaluator {
	//returns the structure properly
	return &Evaluator{
		//takes the node bodies properly
		//this will help us to properly handle
		NodeBodies: bodies,
		//takes the scope and places into var
		//this will store our variables which have been found
		GuideScope: scope,
		//sets the prefixs properly without issues
		//this will make sure its done without issues happening
		templateRcog: prefixs,
		//stores the writer correctly without issues
		//this will ensure its done properly without issues happening
		wr: wr,
		//creates the default information correctly
		//this will ensure its done properly without issues happening
		packages: make([]Package, 0),
		functions: make([]Function, 0),
		internalFunctions: make([]InternalFunction, 0),
		//stores the extra information
		//this will make sure its done correctly
		session: session,
	}
}

//adds the function into the array
//allows the instance to control/execute the function
func (e *Evaluator) AddFunction(name string, b Builtin) {
	//adds into the array correctly
	//this will allow the header function to access
	e.functions = append(e.functions, Function{FunctionName: name, Function: b})
}

//adds the package without issues happening
//this will ensure its done properly without errors
func (e *Evaluator) AddPackage(p *Package) {
	//saves into the array correctly
	//this will ensure its done properly without issues
	e.packages = append(e.packages, *p)
}

func (e *Evaluator) MakeNewRoute(bodies []parser.Node, we io.Writer, s *sessions.Session) *Evaluator {
	return &Evaluator{
		NodeBodies: bodies,
		GuideScope: e.GuideScope,
		wr: we,
		templateRcog: e.templateRcog,
		packages: e.packages,
		functions: e.functions,
		internalFunctions: e.internalFunctions,
		session: s,
	}
}

//follows the interpreter guid without issues happening when prompted
//this will allow for better control without issues happening when being called
func (e *Evaluator) FollowGuide() ([]Object, error) {
	//ranges through the node bodies properly
	for bodies := 0; bodies < len(e.NodeBodies); bodies++ {
		//sets as easy subject for current rotation
		//this will help to properly without issues happening
		node := e.NodeBodies[bodies]

		//detects the different routes
		//allows for proper guides to be inplace
		switch node.Route() {
			
		case parser.Conditi:
			//executes the condition worker properly without issues
			if err := e.ConditionsWorker(node.Condition()); err != nil {
				return nil, err
			}; continue

		//uses the butilin template engine to render text
		//this will allow you to execute with the objects inside a text based engine
		case parser.Texture: //proper text based engine

			//stores the target source properly
			//this will allow for better configuration without issues
			var target string = strings.Join(node.Text().Lines, "")
			//creates the new safe template engine without issues
			//this will ensure its done properly without issues happening on request
			temp, err := fasttemplate.NewTemplate(target, e.templateRcog[0], e.templateRcog[1])
			if err != nil { //error handles the statement properly
				return make([]Object, 0), err //returns the error
			}

			//properly tries to execute without issues happening
			//this will allow for better system handling without issues
			tag, err := temp.ExecuteFuncStringWithErr(func(w io.Writer, tag string) (int, error) {return e.driver(tag, e.session, w)})
			if err != nil { //error handles properly without issues happening
				return make([]Object, 0), err //returns the error properly
			}


				//regexpCompression := regexp.MustCompile(`<main>(.*)</main>`)
				//if err == nil {
				//	for _, tag := range regexpCompression.FindAllStringSubmatch(string(tag), -1) {
				//		for _, second := range tag {
				//			fmt.Println(second)
				//		}
				//	}
				//}
			
			tag = WithinLine(tag) //runs the body checks within the system
			
			//tries to write it properly
			//this will ensure its done without any errors
			if err := e.session.Write(lexer.AnsiUtil(tag, lexer.Escapes)); err != nil { //error
				return make([]Object, 0), err //returns error
			}; continue
		case parser.ReturnP:
			//properly sorts the return without issues
			//this ensures its done properly without issues
			return e.sortReturns(node.Return()) //returns the values properly

		case parser.FuncReg:
			//tries to correctly regiser the function
			//this ensures its properly done without errors
			err := e.registerInternal(node.Function())
			if err != nil {
				return nil, err
			}

			//forces the continue properly without issues
			continue
		case parser.Declare: //declare route properly
			//properly comparses the object
			//this will store information without issues
			sc, err := e.comparseVariable(node.Declare())
			if err != nil {
				//returns the error properly
				return nil, err
			}

			//checks if the object is a const
			//this won't edit it if it doesn't exist
			if e.varExists(sc.Name) && node.Declare().Tokens()[0].Literal() == parser.CONST {
				continue //continues looping without it changing
			} else if e.varExists(sc.Name) { //checks if the object exists
				e.updateObject(sc); continue //updates and continues
			}
			//adds into the scope properly
			//this allows for better handle without issues
			e.GuideScope = append(e.GuideScope, *sc); continue
		case parser.FuncExe:
			//properly tries to execute the function
			//this will allow for better handle without issues
			values, err := e.executeFunction(node.Path())
			if err != nil {
				//returns the error properly
				return nil, err
			}

			_ = values
			continue
		}
	}

	return nil, nil
}

//tries to get the object from the scope
//this will ensure its done without errors happening
func (e *Evaluator) GetObject(name string) (*Scope, error) {
	//ranges through all scopes
	//this will allow us to access an object
	for _, pointer := range e.GuideScope{
		//compares the given objects
		//this will ensure its done properly without errors
		if pointer.Name == name {
			return &pointer, nil
		}
	}

	//returns nil properly without issues happening
	//allows for propery controlling without issues happening
	return nil, errors.New("failed to find the object properly")
} 