package commands


//removes all the commands properly
//this will ensure its done without any errors happening
//this will be used mainly for the reload seq without errors
func RemoveObjects() { //removes all the objects properly and safely
	//ranges through all the commands
	//this will ensure its done without any errors
	for c, s := range Commands { //ranges through
		//tries to validate the system without issues happening
		//this will ensure its done without any errors happening
		if len(s.CustomCommand) <= 0 || s.CustomCommand == "" {
			continue //continues the loop properly
		}
		//removes from the map properly
		//this will ensure its done without any errors
		delete(Commands, c) //removes from the map properly
	}
}