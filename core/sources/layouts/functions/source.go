package functions

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/layouts/toml"
	"io"
)

func init() {

	RegisterFunction(&evaluator.Function{
		FunctionName: "cnc",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			return evaluator.ArrayObject(evaluator.Object{Literal: toml.ConfigurationToml.AppSettings.AppName, Type: lexer.String}), nil
		},
	})

	RegisterFunction(&evaluator.Function{
		FunctionName: "version",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			return evaluator.ArrayObject(evaluator.Object{Literal: deployment.Version, Type: lexer.String}), nil
		},
	})
	RegisterFunction(&evaluator.Function{
		FunctionName: "attackprefix",
		Function: func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
			return evaluator.ArrayObject(evaluator.Object{Literal: toml.AttacksToml.Attacks.Prefix, Type: lexer.String}), nil
		},
	})
}