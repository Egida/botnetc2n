package lexer

import "errors"

var (
	//when the type given isn't closed correctly like string mainly
	ErrTypeOpenedNotClosed error = errors.New("type which was opened was never closed with mustFinish enabled")
)