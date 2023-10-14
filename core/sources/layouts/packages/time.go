package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/tools"
	"io"

	"errors"
	"strconv"
	"time"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "time",
		Functions: map[string]evaluator.Builtin{

			"now" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(time.Now().Unix())), Type: lexer.Int}), nil
			},

			"unix" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//this will ensure its done without any errors
				if len(args) < 2 { //checks the args given properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside function call")
				}

				//tries to convert the first digit properly
				//this will ensure its done without any errors
				unix, err := strconv.Atoi(args[0].Literal()) //converts into int properly
				if err != nil || args[0].TokenType() != lexer.Int { //error handles properly
					return make([]evaluator.Object, 0), err //returns err
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: time.Unix(int64(unix), 0).Format(args[1].Literal())}), nil
			},

			//since the unix time happened
			//this will work out the duration properly
			"since" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//this will ensure its done without any errors
				if len(args) < 1 { //checks the args given properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside function call")
				}

				//tries to convert the first digit properly
				//this will ensure its done without any errors
				unix, err := strconv.Atoi(args[0].Literal()) //converts into int properly
				if err != nil || args[0].TokenType() != lexer.Int { //error handles properly
					return make([]evaluator.Object, 0), err //returns err
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: tools.ResolveTimeStamp(time.Unix(int64(unix), 0), true)}), nil
			},

			//until the unix time is current
			//this will work out the duration until the unix stamp is now
			"until" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//this will ensure its done without any errors
				if len(args) < 1 { //checks the args given properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside function call")
				}

				//tries to convert the first digit properly
				//this will ensure its done without any errors
				unix, err := strconv.Atoi(args[0].Literal()) //converts into int properly
				if err != nil || args[0].TokenType() != lexer.Int { //error handles properly
					return make([]evaluator.Object, 0), err //returns err
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: tools.ResolveTimeStampUnix(time.Unix(int64(unix), 0), true)}), nil
			},

		},
	})
}