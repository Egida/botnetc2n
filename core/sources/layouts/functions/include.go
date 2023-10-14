package functions

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"Nosviak2/core/sources/views"
	"errors"
	"io"
	"strings"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "include",
		//clears the screen properly and safer
		//this will make sure its done correctly without issues
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

			//checks the length given properly
			//this will ensure its done without any issues
			if len(args) != 1 { //length checks
				return make([]evaluator.Object, 0), errors.New("missing argument properly")
			}

			//tries to find the file properly
			//this will make sure its done without issues
			brand := views.GetView(strings.Split(args[0].Literal(), "/")...)
			if brand == nil {
				return make([]evaluator.Object, 0), nil
			}


			//runs the parser properly
			//this will ensure its done without issues
			nodes, err := parser.MakeParserRun(lexer.Make(brand.Containing, true).RunTarget())
			if err != nil { //error handles the ref properly
				return make([]evaluator.Object, 0), err
			}

			//properly tries to make the eval
			//this will be used to guide the execution
			eval := e.MakeNewRoute(nodes, s.Channel, s)

			//tries to run the evaluator
			//this will ensure its done without issues
			if _, err := eval.FollowGuide(); err != nil {
				return make([]evaluator.Object, 0), err //returns error
			}


			//saves into the guide scope properly
			//this will try to safely handle without issues
			e.GuideScope = append(e.GuideScope, eval.Guide()...)
			e.Add(eval.GrabFunc())

			//correctly tries to write without issues
			//this will make sure its done without issues making it safer
			return make([]evaluator.Object, 0), nil
		},
	})
}