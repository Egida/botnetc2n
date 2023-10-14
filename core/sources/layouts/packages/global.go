package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/database"
	"Nosviak2/core/slaves/fakes"
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/slaves/propagation"
	"Nosviak2/core/slaves/qbot"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"errors"
	"io"
	"strconv"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "global",
		Functions: map[string]evaluator.Builtin{

			//gets the amount of online users
			//this will ensure its done without issues
			"online" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(sessions.Sessions))}), nil
			},

			
			//creates the cnc created time properly
			//this will show the created instance time properly
			"started" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(database.Conn.Time.Unix()))}), nil
			},


			//gets all of the ongoing attacks properly
			//this will ensure its done without issues happening
			"ongoing" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the running attacks properly
				//this will ensure its done without errors
				running, err := database.Conn.GlobalRunning()
				if err != nil { //returns 0 as none could be found properly
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: "0"}), nil
				}
				//returns the literal running properly
				//this will ensure its done without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(running))}), nil
			},

			//gets all of the attacks sent properly
			//this will ensure its done without issues happening
			"totalattacks" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the running attacks properly
				//this will ensure its done without errors
				launched, err := database.Conn.GlobalSent() //gets all sent
				if err != nil { //returns 0 as none could be found properly
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: "0"}), nil
				}
				//returns the literal running properly
				//this will ensure its done without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(launched))}), nil
			},

			"myrunning" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//properly tries to push the ongoing attacks
				//this will allow for better control without issues
				linear, err := database.Conn.Attacking(s.User.Username) //gets the running
				if err != nil { //error handles the statement properly without issues
					return make([]evaluator.Object, 0), err //returns the error
				}
				//returns the object properly without issues happening
				//this will allow for better control within internal functions
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(linear))}), nil
			},

			//access the slave count for mirai properly
			//the different features will allow for more then one mirai etc to be linked
			"mirai" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(mirai.MiraiSlaves.All))}), nil
			},

			//access the slave count for qbot properly
			//the different features will allow for more then one qbot etc to be linked
			"qbot" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(qbot.QbotClients))}), nil
			},

			//access the slave count for qbot properly
			//the different features will allow for more then one propagated etc to be linked
			"propagated" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				var current int = 0

				//ranges through props properly
				//allows us to get the correct amount
				for _, amount := range propagation.Reps {
					current += amount //adds amount
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(current)}), nil
			},

			//access the slave count for mirai properly
			//the different features will allow for more then one mirai etc to be linked
			"fake" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) <= 0 { //length checks properly within the system
					return make([]evaluator.Object, 0), errors.New("missing args properly")
				}
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(fakes.FakeSlaves[args[0].Literal()])}), nil
			},
		},
	})
}