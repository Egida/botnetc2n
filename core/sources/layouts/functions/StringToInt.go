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
		FunctionName: "atoi",
		//clears the screen properly and safer
		//this will make sure its done correctly without issues
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			if len(args) < 1 { //checks the arg length properly and safely
				return make([]evaluator.Object, 0), errors.New("missing args properly")
			}
			//tries to convert properly
			//this will ensure its done without any errors
			conv, err := strconv.Atoi(args[0].Literal()) //atoi
			if err != nil || args[0].TokenType() != lexer.String { //checks the type
				return make([]evaluator.Object, 0), errors.New("error atoi forced")
			}

			//correctly tries to write without issues
			//this will make sure its done without issues making it safer
			return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(conv), Type: lexer.Int}), nil
		},
	})

	RegisterFunction(&evaluator.Function{
		FunctionName: "itoa",
		//clears the screen properly and safer
		//this will make sure its done correctly without issues
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			if len(args) < 1 { //checks the arg length properly and safely
				return make([]evaluator.Object, 0), errors.New("missing args properly")
			}

			//correctly tries to write without issues
			//this will make sure its done without issues making it safer
			return evaluator.ArrayObject(evaluator.Object{Literal: args[0].Literal(), Type: lexer.String}), nil
		},
	})
}