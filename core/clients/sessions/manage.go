package sessions



//counts how many sessions are open properly
//this will return the int amount of how many are open
func (s *Session) CountOpenSessions() []Session {

	var self []Session = make([]Session, 0) //stores the amount
	//ranges through all the different sessions
	//we will check the username of each session we check
	for sp := range Sessions {

		//checks the username status
		//this will ensure its done without issues
		if Sessions[sp].User.Username == s.User.Username {
			self = append(self, *Sessions[sp]) //adds another session in properly
		}
	}

	//returns how many sessions
	//this will ensure we can properly access
	return self //returns the self amount open
}

//counts how many sessions are open properly
//this will return the int amount of how many are open
func (s *Session) GetSessions(usr string) []Session {

	var self []Session = make([]Session, 0) //stores the amount
	//ranges through all the different sessions
	//we will check the username of each session we check
	for sp := range Sessions {

		//checks the username status
		//this will ensure its done without issues
		if Sessions[sp].User.Username == usr {
			self = append(self, *Sessions[sp]) //adds another session in properly
		}
	}

	//returns how many sessions
	//this will ensure we can properly access
	return self //returns the self amount open
}

//tries to find the session with the id 
//this will ensure its done without any errors
func (s *Session) GetWithID(id int64) *Session {
	//ranges through all sessions properly
	//this will ensure its done without any errors
	for i := range Sessions { //ranges through properly
		if i == id {return Sessions[i]} //returns properly
	}; return nil
}