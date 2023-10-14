package functions

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"io"
	"strconv"
	"time"

	"errors"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "sleep",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			//checks the length
			//this will make sure its done properly without issues
			if len(args) != 1 { //error handling the system properly
				return make([]evaluator.Object, 0), errors.New("missing args to perform this tag correctly")
			}

			//tries to correctly convert
			//this will make sure its done without issues
			convert, err := strconv.Atoi(args[0].Literal())
			if err != nil { //error handles the system properly
				return make([]evaluator.Object, 0), err //returns error
			}

			//sleeps for the duration given properly
			//this will ensure its done properly without issues
			time.Sleep(time.Duration(convert) * time.Millisecond)
			return make([]evaluator.Object, 0), nil
		},
	})
}