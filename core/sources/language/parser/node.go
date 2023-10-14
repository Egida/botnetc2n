package parser



//stores in type int properly
//this will allow the interpreter to execute safely and correctly
type RouteNodeExe int

const (
	Declare RouteNodeExe = 1 //sets the declare properly without issues
	FuncExe RouteNodeExe = 2 //sets the funcexecution properly without issues
	FuncReg RouteNodeExe = 3 //when registering a function correctly properly
	ReturnP RouteNodeExe = 4 //when the function returns the values properly
	Texture RouteNodeExe = 5 //when text is properly detected without a proper cause
	Conditi RouteNodeExe = 6 //conditions statement properly without issues happening
)


//nodes store each parser feature
//this allows for better execution routes inside the system
type Node struct {
	//stores the routenodeexe route properly
	//this will safely and properly store the route
	exe RouteNodeExe //stores the routenodeexe

	//stores the declare structure correctly
	//this will ensure its properly done without issues happening
	declare *DeclareRoute //stores the parsed structure

	//stores the execution path properly
	//this will ensure its done properly without issues
	path *FunctionPath //stores the parsed structure
	
	//stores the function register path correctly
	//this will make sure its done properly without issues happening
	function *Function

	//stores the return information properly
	//makes sure its properly done without issues happening
	returns *ReturnReply

	//stores the text part
	//this will properly store the text structure
	text *LineFormat

	//stores the condition statements correctly
	//this will make sure its done without issues happening
	conditions *Conditional
}

func (n *Node) Condition() *Conditional {
	return n.conditions
}

func (n *Node) Text() *LineFormat {
	return n.text
}

func (n *Node) Return() *ReturnReply {
	return n.returns
}

func (n *Node) Function() *Function {
	return n.function
}

func (n *Node) Route() RouteNodeExe {
	return n.exe
}

func (n *Node) Declare() *DeclareRoute {
	return n.declare
}

func (n *Node) Path() *FunctionPath {
	return n.path
}