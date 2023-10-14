package functions

//formats all the frames into a string properly
//this will hopefully make sure its done without errors
func TerminalQR(frames [][]bool, source string) (string, error) {

	//gets the array of bools for line
	//this will ensure its done without errors
	for line, row := range frames {

		//validates properly
		//ensures its done properly
		if Passed(line, len(row)) { //check if
			continue //continues looping 
		}

		//ranges through the row properly and safely
		//this will ensure its done without errors happening
		for charaterPos, charater := range row { //ranges through
			
			//validates properly
			//ensures its done properly
			if Passed(charaterPos, len(frames)) { //check if
				continue //continues looping 
			}

			if charater {
				source += "\x1b[48;5;7m  \x1b[0m"
			} else {
				source += "\x1b[48;5;0m  \x1b[0m"
			}
		}

		//new line input here properly
		//this will ensure its done properly
		source += "\r\n" //new line inset here
	}

	//returns the source properly
	//this will ensure its done properly
	return source, nil
}

//checks the system properly and safely
//this will hopefully make sure its done safely
func Passed(pointer int, backwards int) bool { // checks if the system passed properly
	return pointer <= 2 || pointer == backwards - 1 || backwards - 2 == pointer || backwards - 3 == pointer
}