package parser 


const (
	//sets all the different keywords for the system
	//this will help within the parser & interpreter detecting structure
	FUNCTION string = "func"
	VAR string = "var"
	CONST string = "const"
	STRING, INT, BOOLEANTRUE, BOOLEANFALSE, BOOLEAN string = "string", "int", "true", "false", "bool"
	RETURN = "return" //reply user
	IF, ELSE = "if", "else"
)