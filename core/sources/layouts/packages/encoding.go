package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "encoding",
		Functions: map[string]evaluator.Builtin{
			"sha256" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 1 { //tries to validate the length
					return make([]evaluator.Object, 0), errors.New("missing object inside argument properly")
				}
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: hex.EncodeToString(sha256.New().Sum([]byte(args[0].Literal())))}), nil
			},
			"sha1" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 1 { //tries to validate the length
					return make([]evaluator.Object, 0), errors.New("missing object inside argument properly")
				}
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: hex.EncodeToString(sha1.New().Sum([]byte(args[0].Literal())))}), nil
			},
			"sha512" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 1 { //tries to validate the length
					return make([]evaluator.Object, 0), errors.New("missing object inside argument properly")
				}
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: hex.EncodeToString(sha512.New().Sum([]byte(args[0].Literal())))}), nil
			},
			"base64" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 1 { //tries to validate the length
					return make([]evaluator.Object, 0), errors.New("missing object inside argument properly")
				}
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: base64.RawStdEncoding.EncodeToString([]byte(args[0].Literal()))}), nil
			},
			"base32" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 1 { //tries to validate the length
					return make([]evaluator.Object, 0), errors.New("missing object inside argument properly")
				}
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: base32.StdEncoding.EncodeToString([]byte(args[0].Literal()))}), nil
			},
			"md5" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 1 { //tries to validate the length
					return make([]evaluator.Object, 0), errors.New("missing object inside argument properly")
				}
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: hex.EncodeToString(md5.New().Sum([]byte(args[0].Literal())))}), nil
			},
		},
	})
}