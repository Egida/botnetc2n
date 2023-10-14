package sessions

import (

	"strings"
)

//writes to the remote host properly
//this will ensure its done without errors happening
func (s *Session) Write(i ...string) error {
	//saves properly into the array without issues happening
	//if len(s.Written) > 0 { //checks the length properly and safely
		s.Written = append(s.Written, i...) //saves into the array
	//}


	//tries to correctly write to the remote host without issues
	//this will ensure its done without any issues happening on reqeust
	if _, err := s.Channel.Write([]byte(strings.Join(i, " "))); err != nil {
		return err //returns the error correctly and properly
	}; return nil //returns nil as it worked!
}


//execute a function graph on a remote session
//this will allow access to the information with all the session details
func (s *Session) FunctionRemote(user string, f func(t *Session)) {
	//ranges through all the sessions which are the users
	//this will ensure its done without issues happening on request
	for _, session := range s.GetSessions(user) { //gets all the users sessions
		f(&session) //properly executes the session information without issues happening
	}
}