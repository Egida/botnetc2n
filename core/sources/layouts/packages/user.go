package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/ranks"
	"errors"
	"io"
	"strings"
	"time"

	"strconv"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "user",
		Functions: map[string]evaluator.Builtin{
			"logins" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				logins, err := database.Conn.GetLogins(s.User.Username)
				if err != nil {
					return make([]evaluator.Object, 0), err
				}
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(logins))}), nil
			},
			//allows us to access the window size properly
			//this will be inserted without issues happening on request
			"length" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.Length)}), nil
			},

			//allows us to access the window size properly
			//this will be inserted without issues happening on request
			"height" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.Height)}), nil
			},

			//allows us to access the ip properly
			//this will be inserted without issues happening on request
			"ip" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: s.Target}), nil
			},
			
			//allows us to access the username properly
			//this will be inserted without issues happening on request
			"username" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: s.User.Username}), nil
			},

			//allows us to access the maxslaves for the account properly
			//this will be inserted without issues happening on request
			"maxslaves" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the user model correctly and safely
				//this will make sure its done correctly without issues happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.MaxSlaves)}), nil
			},

			//allows us to access the id properly
			//this will make sure its done correctly without issues happening
			"id" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.Identity)}), nil
			},

			//allows us to access the maxtime properly
			//this will make sure its done correctly without issues happening
			"maxtime" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.MaxTime)}), nil
			},

			//allows us to access the cooldown properly
			//this will make sure its done correctly without issues happening
			"cooldown" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.Cooldown)}), nil
			},

			//allows us to access the concurrents properly
			//this will make sure its done correctly without issues happening
			"concurrents" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.Concurrents)}), nil
			},

			//allows us to access the maxsessions properly
			//this will make sure its done correctly without issues happening
			"maxsessions" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.MaxSessions)}), nil
			},

			//allows us to access the maxsessions properly
			//this will make sure its done correctly without issues happening
			"opensessions" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(s.CountOpenSessions()))}), nil
			},

			//allows us to access the maxsessions properly
			//this will make sure its done correctly without issues happening
			"theme" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: s.User.Theme}), nil
			},

			//allows us to access the maxsessions properly
			//this will make sure its done correctly without issues happening
			"ranks" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the users ranks properly
				//this will ensure its done without any issues
				r, err := s.Ranks.DeployRanks(false) //gets the ranks properly
				if err != nil { //error handles properly
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: ""}), nil
				}
				//returns the value properly
				//this will make sure its done correctly
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: strings.Join(r, " ")}), nil
			},

			"CanAccess" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the arguments properly
				//this will ensure its done without issues happening
				if len(args) < 1 { //checks the length correctly and properly and returns error
					return make([]evaluator.Object, 0), errors.New("missing arguments properly without reason")
				}

				rank := ranks.MakeRank(s.User.Username) //creates the new rank information
				rank.SyncWithString(s.User.Ranks) //syncs the rank string properly

				//returns the output properly without issues
				//this will make sure its properly done without errors
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: strconv.FormatBool(rank.CanAccess(args[0].Literal()))}), nil
			},

			"HasPermissions" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the arguments properly
				//this will ensure its done without issues happening
				if len(args) < 1 { //checks the length correctly and properly and returns error
					return make([]evaluator.Object, 0), errors.New("missing arguments properly without reason")
				}

				rank := ranks.MakeRank(s.User.Username) //creates the new rank information
				rank.SyncWithString(s.User.Ranks) //syncs the rank string properly

				//returns the output properly without issues
				//this will make sure its properly done without errors
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: strconv.FormatBool(rank.CanAccessArray(Combine(args, make([]string, 0))))}), nil
			},

			//gets the users parent properly
			"parent" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(s.User.Parent)}), nil
			},

			//gets the users total attacks sent
			"totalAttacks" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the users attacks properly
				//this will ensure we can gather them without issues
				sent, err := database.Conn.UserSent(s.User.Username) //gets all the users attacks
				if err != nil { //error handles properly without issues
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(0)}), nil
				}

				//returns the amount properly without issues
				//this will ensure its done without any errors happening
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(len(sent))}), nil
			},

			"created" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(s.User.Created))}), nil
			},

			"updated" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(s.User.Updated))}), nil
			},

			"expiry" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(s.User.Expiry))}), nil
			},

			"days" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(time.Until(time.Unix(s.User.Expiry, 0)).Hours()/24))}), nil
			},

			"hours" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(time.Until(time.Unix(s.User.Expiry, 0)).Hours()))}), nil
			},

			"minutes" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(time.Until(time.Unix(s.User.Expiry, 0)).Minutes()))}), nil
			},

			"seconds" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(time.Until(time.Unix(s.User.Expiry, 0)).Seconds()))}), nil
			},

			//gets the last target properly
			//this will get the last target the user attacked
			"lastTarget" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the attacks properly
				//this will access all the attacks properly
				attks, err := database.Conn.UserSent(s.User.Username) //gets the user
				if err != nil || len(attks) <= 0 { //error handles properly without issues happening on reqeust
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: "0.0.0.0"}), nil
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: attks[len(attks)-1].Target}), nil
			},

			//gets the last method properly
			//this will get the last method the user attacked
			"lastMethod" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the attacks properly
				//this will access all the attacks properly
				attks, err := database.Conn.UserSent(s.User.Username) //gets the user
				if err != nil || len(attks) <= 0 { //error handles properly without issues happening on reqeust
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: "[EOF]"}), nil
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: attks[len(attks)-1].Method}), nil
			},

			//gets the last duration properly
			//this will get the last duration the user attacked
			"lastDuration" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the attacks properly
				//this will access all the attacks properly
				attks, err := database.Conn.UserSent(s.User.Username) //gets the user
				if err != nil || len(attks) <= 0 { //error handles properly without issues happening on reqeust
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: "0"}), nil
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(attks[len(attks)-1].Duration)}), nil
			},

			//gets the last port properly
			//this will get the last port the user attacked
			"lastPort" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the attacks properly
				//this will access all the attacks properly
				attks, err := database.Conn.UserSent(s.User.Username) //gets the user
				if err != nil || len(attks) <= 0 { //error handles properly without issues happening on reqeust
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: "0"}), nil
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(attks[len(attks)-1].Port)}), nil
			},

			//gets the last port properly
			//this will get the last port the user attacked
			"lastFinishedUnix" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the attacks properly
				//this will access all the attacks properly
				attks, err := database.Conn.UserSent(s.User.Username) //gets the user
				if err != nil || len(attks) <= 0 { //error handles properly without issues happening on reqeust
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(time.Now().Unix()))}), nil
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(attks[len(attks)-1].Finish))}), nil
			},

			//gets the last port properly
			//this will get the last port the user attacked
			"lastCreatedUnix" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//gets all the attacks properly
				//this will access all the attacks properly
				attks, err := database.Conn.UserSent(s.User.Username) //gets the user
				if err != nil || len(attks) <= 0 { //error handles properly without issues happening on reqeust
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(time.Now().Unix()))}), nil
				}

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Int, Literal: strconv.Itoa(int(attks[len(attks)-1].Created))}), nil
			},

			//how many viewers which are watching the sessions screen
			"viewers" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(s.Viewers), Type: lexer.Int}), nil
			},
			"title" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the arg length without issues happening
				if len(args) < 2 { //checks the arg length properly without issues
					return make([]evaluator.Object, 0), errors.New("missing arguments properly without reason")
				}
				//tries to convert properly
				//this will ensure its done without any errors
				durtimeout, err := strconv.Atoi(args[0].Literal())
				if err != nil { //err handles properly without issues
					return make([]evaluator.Object, 0), err //returns err
				}
				//sets the different information properly
				//this will reset back to default settings
				s.Title = strings.Join(Combine(args[1:], make([]string, 0)), "") //sets the title properly
				s.CustomTitleReset = time.Now().Add(time.Duration(durtimeout) * time.Millisecond).Unix() //sets reset properly
				return make([]evaluator.Object, 0), nil //stores the values properly without issues happening
			},

			//awards the user the rank given inside the arguments
			//this will ensure its done without any errors happening
			"award" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				//checks the length properly
				//this will ensure its done without errors
				if len(args) != 1 { //makes sure only 1 argument is given
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}

				//gets the first rank properly and safely
				//this will ensure its done without errors
				if err := s.Ranks.GiveRank(args[0].Literal()); err != nil {
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}

				//compress into the string properly
				//this will ensure its done without issues
				compress, err := s.MakeString() //made within a string
				if err != nil { //error handles properly without issues happening
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}
				

				//tries to push the ranks inside the database
				//this will check that they have been pushed properly
				if err := database.Conn.EditRanks(compress, s.User.Username); err != nil {
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}

				//pushes the update to the current session
				//this will ensure its done without errors happening
				sessions.Sessions[s.ID].Ranks = s.Ranks
				sessions.Sessions[s.ID].User.Ranks = compress

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "true"}), nil
			},

			//revoke the user the rank given inside the arguments
			//this will ensure its done without any errors happening
			"revoke" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				//checks the length properly
				//this will ensure its done without errors
				if len(args) != 1 { //makes sure only 1 argument is given
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}

				//gets the first rank properly and safely
				//this will ensure its done without errors
				if err := s.Ranks.RemoveRank(args[0].Literal()); err != nil {
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}

				//compress into the string properly
				//this will ensure its done without issues
				compress, err := s.MakeString() //made within a string
				if err != nil { //error handles properly without issues happening
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}
				

				//tries to push the ranks inside the database
				//this will check that they have been pushed properly
				if err := database.Conn.EditRanks(compress, s.User.Username); err != nil {
					return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "false"}), nil
				}

				//pushes the update to the current session
				//this will ensure its done without errors happening
				sessions.Sessions[s.ID].Ranks = s.Ranks
				sessions.Sessions[s.ID].User.Ranks = compress

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.Boolean, Literal: "true"}), nil
			},
		},
	})
}

//combines the tokens into an array of strings
//this will combine the array properly and safely
func Combine(l []lexer.Token, src []string) []string {
	//ranges through all the tokens
	//this will ensure its done without any errors
	for _, tok := range l { //ranges through l properly
		src = append(src, tok.Literal()) //saves into properly
	}; return src
}