package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"io"
	"runtime"
	"strconv"

	"github.com/pbnjay/memory"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "sys",
		Functions: map[string]evaluator.Builtin{
			//access the servers os format properly
			"os" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: runtime.GOOS, Type: lexer.String}), nil //string format
			},

			//access the servers arch format properly
			"arch" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: runtime.GOARCH, Type: lexer.String}), nil //string format
			},

			//access the servers go version format properly
			"goversion" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: runtime.Version(), Type: lexer.String}), nil //string format
			},

			//access the servers os format properly
			"cpu" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(runtime.NumCPU()), Type: lexer.Int}), nil //string format
			},

			//access the amount of allocated memmory properly
			"routines" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(runtime.NumGoroutine()), Type: lexer.Int}), nil //string format
			},
			

			//access the amount of allocated memmory properly
			"allocated" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				var run runtime.MemStats; runtime.ReadMemStats(&run)
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(bToMb(run.TotalAlloc))), Type: lexer.Int}), nil //string format
			},

			//access the amount of allocated memmory properly
			"system" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				var run runtime.MemStats; runtime.ReadMemStats(&run)
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(bToMb(run.Sys))), Type: lexer.Int}), nil //string format
			},

			//access the amount of allocated memmory properly
			"memory" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				var run runtime.MemStats; runtime.ReadMemStats(&run)
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(bToMb(memory.TotalMemory()))), Type: lexer.Int}), nil //string format
			},

			//access the amount of allocated memmory properly
			"freememory" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				var run runtime.MemStats; runtime.ReadMemStats(&run)
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(bToMb(memory.FreeMemory()))), Type: lexer.Int}), nil //string format
			},
		},
	})
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024 / 1024
}