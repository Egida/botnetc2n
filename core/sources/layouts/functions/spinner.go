package functions

import (
	"Nosviak2/core/clients/animations"
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"
	"io"
	"strconv"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "spinner",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			if len(args) != 2 || args[0].TokenType() != lexer.String || args[1].TokenType() != lexer.Boolean { //checks the length and the type properly
				return make([]evaluator.Object, 0), errors.New("missing args to perform this tag correctly")
			}

			//parses the secondary item
			//this will ensure its done without issues
			secondary, err := strconv.ParseBool(args[1].Literal())
			if err != nil { //error handles the second one properly
				return make([]evaluator.Object, 0), err //returns error
			}

			//gets the spinners current frame
			//this will ensure its done without issues happening
			frame, err := animations.AccessCurrentFrame(args[0].Literal())
			if err != nil || len(frame) == 0 { //checks the frame properly without issues
				return nil, err //returns the error correctly and properly
			}

			//decides the route
			//this completes 0 parsing
			if !secondary {
				//returns the object properly
				//this will ensure its done without any errors
				return evaluator.ArrayObject(evaluator.Object{Literal: frame, Type: lexer.String}), nil
			}

			//runs the parser properly
			//this will ensure its done without issues
			nodes, err := parser.MakeParserRun(lexer.Make(frame, true).RunTarget())
			if err != nil { //returns the error properly
				return make([]evaluator.Object, 0), err
			}
			//returns the object properly
			//this will ensure its done without any errors
			return evaluator.MakeEval(nodes, e.GuideScope, wr, deployment.Engine, s).FollowGuide()
		},
	})
}