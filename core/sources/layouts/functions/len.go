package functions

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"errors"
	"io"
	"strconv"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "len",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

			if len(args) != 1 {
				return make([]evaluator.Object, 0), errors.New("missing args to perform this tag correctly")
			}

			return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(len(args[0].Literal())), Type: lexer.Int}), nil
		},
	})
}