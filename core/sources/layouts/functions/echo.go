package functions

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"bytes"
	"errors"
	"io"
)

func init() {
	//registers the function properly
	//this will ensure its done without issues happening
	RegisterFunction(&evaluator.Function{
		FunctionName: "echo",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			//validates the args
			//this will make sure there are no args missing
			if len(args) < 1 {
				//returns the error correctly if there is an arg missing
				return make([]evaluator.Object, 0), errors.New("missing args to perform this tag correctly")
			}

			//stores the byte buffer properly
			//this holds the future text properly
			var buffer bytes.Buffer = *bytes.NewBuffer(nil)

			//saves each arguments into the array
			//this will ensure its save to use without errors
			for tok := range args {
				//applys into the array correctly
				buffer.Write([]byte(args[tok].Literal()))
			}

			err := s.Write(lexer.AnsiUtil(buffer.String(), lexer.Escapes))
		
			//writes to the buffer
			//this will write the buffer to the remote host properly			
			return make([]evaluator.Object, 0), err
		},
	})
}