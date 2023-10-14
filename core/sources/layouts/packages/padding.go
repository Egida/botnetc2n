package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/functions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"io"
	"strings"
	"unicode/utf8"

	"errors"
	"strconv"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "padding",
		Functions: map[string]evaluator.Builtin{
			//sets the padding from the right to left
			//this will ensure its done without issues happening
			"padRight" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//ensures the length is valid
				//makes sure its properly done without issues
				if len(args) != 2 { //checks length properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside the statement")
				}
				//gets the spacer properly
				//this will make sure its done properly without issues
				spacer, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}
				//stores the text properly
				var Text string = args[0].Literal()
				//for loops throughout the spacer without issues
				//this will ensure its done without errors happening
				for pos := len(Text); pos < spacer; pos++ {
					Text += " "
				}
				//returns the object properly
				//this will make sure its done without issues
				return evaluator.ArrayObject(evaluator.Object{Literal: Text, Type: lexer.String}), nil
			},

			//sets the padding from the right to left with custom padding
			//this will ensure its done without issues happening
			"padcustomRight" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//ensures the length is valid
				//makes sure its properly done without issues
				if len(args) != 3 { //checks length properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside the statement")
				}
				//gets the spacer properly
				//this will make sure its done properly without issues
				spacer, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}
				//stores the text properly
				var Text string = args[0].Literal()
				//for loops throughout the spacer without issues
				//this will ensure its done without errors happening
				for pos := len(Text); pos < spacer; pos++ {
					Text += args[2].Literal()
				}
				//returns the object properly
				//this will make sure its done without issues
				return evaluator.ArrayObject(evaluator.Object{Literal: Text, Type: lexer.String}), nil
			},

			//sets the padding from the left to the right
			//this will ensure its done without issues happening
			"padLeft" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//ensures the length is valid
				//makes sure its properly done without issues
				if len(args) != 2 { //checks length properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside the statement")
				}
				//gets the spacer properly
				//this will make sure its done properly without issues
				spacer, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}
				//stores the text properly
				var Text string = args[0].Literal()
				//for loops throughout the spacer without issues
				//this will ensure its done without errors happening
				for pos := len(Text); pos < spacer; pos++ {
					Text = " " + Text
				}
				
				//returns the object properly
				//this will make sure its done without issues
				return evaluator.ArrayObject(evaluator.Object{Literal: Text, Type: lexer.String}), nil
			},

			//sets the padding from the left to the right with custom padding
			//this will ensure its done without issues happening
			"padcustomLeft" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//ensures the length is valid
				//makes sure its properly done without issues
				if len(args) != 3 { //checks length properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside the statement")
				}
				//gets the spacer properly
				//this will make sure its done properly without issues
				spacer, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}
				//stores the text properly
				var Text string = args[0].Literal()
				//for loops throughout the spacer without issues
				//this will ensure its done without errors happening
				for pos := len(Text); pos < spacer; pos++ {
					Text = args[2].Literal() + Text
				}
				//returns the object properly
				//this will make sure its done without issues
				return evaluator.ArrayObject(evaluator.Object{Literal: Text, Type: lexer.String}), nil
			},

			//centres the first argument given inside this proeprly
			"centre": func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the length properly
				//tries to stop nil pointers properly
				if len(args) != 2 { //checks length properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside the statement")
				}

				//converts the value properly
				//this will ensure its usable properly
				spacer, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}

				var out string
				//loops through properly
				//this will try to properly centre without issues
				for target := 0; target < spacer; target++ { //loops
					//this will try to properly check for half pos
					if spacer / 2 - utf8.RuneCountInString(args[0].Literal()) / 2 == target { //checks for half
						out += args[0].Literal()
						target += utf8.RuneCountInString(args[0].Literal())
						continue
					}

					out += " "
				}
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: out}), nil
			},

			"centrecustom": func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the length properly
				//tries to stop nil pointers properly
				if len(args) != 3 { //checks length properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside the statement")
				}

				//converts the value properly
				//this will ensure its usable properly
				spacer, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}

				//converts the value properly
				//this will ensure its usable properly
				lo, err := strconv.Atoi(args[2].Literal())
				if err != nil { //error handles properly without issues
					return make([]evaluator.Object, 0), err
				}

				var out string
				//loops through properly
				//this will try to properly centre without issues
				for target := 0; target < spacer; target++ { //loops
					//this will try to properly check for half pos
					if spacer / 2 - lo / 2 == target { //checks for half
						out += args[0].Literal()
						target += lo
						continue
					}

					out += " "
				}
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: out}), nil
			},

			"lower" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				var lowered string = ""
				for _, system := range args {

					if system.TokenType() != lexer.String {
						continue
					}

					lowered += system.Literal()
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: strings.ToLower(lowered)}), nil
			},

			"upper" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				var lowered string = ""
				for _, system := range args {

					if system.TokenType() != lexer.String {
						continue
					}

					lowered += system.Literal()
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: strings.ToLower(lowered)}), nil
			},

			//stores the ansi object properly
			//this will allow us to build of it safely
			"ansi" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				var build string = ""

				//ranges through all tokens given
				//this will ensure its done without issues
				for _, systemToken := range args { //ranges
					build += systemToken.Literal()
				}

				return evaluator.ArrayObject(evaluator.Object{Literal: functions.String(build), Type: lexer.String}), nil
			},
		},
	})
}