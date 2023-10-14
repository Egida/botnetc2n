package evaluator

import "errors"

var (
	//Error when the token are invalid properly and invalid
	ErrTypeMatchSeq error = errors.New("tokens appears different inside the array")
)