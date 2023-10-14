package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/tools"
	"io"
	"strconv"
	"time"

	"golang.org/x/term"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "channel",
		Functions: map[string]evaluator.Builtin{


			//closes the connection with the channel properly
			"close" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				if err := s.Channel.Close(); err != nil { //returns the objective output correctly
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}
				
				//returns true properly and safely
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "true"}), nil
			},

			//closes the connection with the channel properly
			"connected" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(s.Connected.Unix())), Type: lexer.Int}), nil
			},
			//shakes the connection properly
			"shake" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return make([]evaluator.Object, 0), tools.ShakeTerminal(1, 11 * time.Millisecond, s.Channel)
			},

			"input" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//reads the terminal properly
				//this will ensure its done without any errors
				val, err := term.NewTerminal(s.Channel, "").ReadLine()
				if err != nil { //error handles properly and safely
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: ""}), nil
				}

				//returns the raw value properly
				//this will ensure its done without any errors
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: val}), nil
			},

			//x = horizontal
			"x" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				position, err := tools.GetWindow(s.Channel) //grabs the windows channel properly
				if err != nil { //err handles properly
					return evaluator.ArrayObject(evaluator.Object{Literal: "0", Type: lexer.Int}), nil
				}
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(position.Horizontal), Type: lexer.Int}), nil
			},

			//y = vertical
			"y" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				position, err := tools.GetWindow(s.Channel) //grabs the windows channel properly
				if err != nil { //err handles properly
					return evaluator.ArrayObject(evaluator.Object{Literal: "0", Type: lexer.Int}), nil
				}
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(position.Vertical), Type: lexer.Int}), nil
			},
		},
	})
}