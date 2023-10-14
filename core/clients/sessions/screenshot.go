package sessions

import "strings"

//Capture will get current session outlook
func (s *Session) Capture() string {
	//stores output
	var que string = ""

	//ranges through the frames
	for _, system := range s.Written { 
		
		//clear ansi escape detected
		if strings.Contains(system, "\033c") {
			que = ""; continue
		}

		//title input detected
		if strings.Contains(system, "\033[0;") {
			continue
		}

		que += system //saves in array
	}

	return que
}