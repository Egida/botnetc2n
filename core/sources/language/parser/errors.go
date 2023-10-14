package parser

import "errors"

var (
	//when a parser default setting has a fault starting type
	ErrParseRootFault error = errors.New("invalid parse root was founded")
	//when the parser found an routeName which doesn't start with a indent
	ErrRouteNameInvalid error = errors.New("the routeName given wasn't type indent")
	//when the parsers tokenType option isn't valid and isn't safe
	ErrInvalidAssign error = errors.New("parser found an invalid option inside the system")
	//when the parser found a suppose type but it isn't valid and safe
	ErrInvalidTypeKeyword error = errors.New("keyword wanted but type wasn't valid")
	//when the parser finds a nil pointer inside a statement
	ErrNilPointer error = errors.New("nil pointer found during parser section")
)