package tools

import (
	"strings"
)

//stores the sanatize the system properly
//this will ensure its done without issues happening
func SanatizeTool(contains string) string {

	var Charaters []string = []string{ //all valid charaters which can be used properly
		"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m", //letters
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", //numbers
		"_", ".", "#", "!", "/", "?", "@", "£", "$", "%",
	}

	//ranges through contains properly
	for _, private := range contains {

		//checks for valid input properly
		//this will ensure we only replace invalid charaters
		if NeedleHaystackOne(Charaters, strings.ToLower(string(private))) {
			continue
		}

		//contains and replaces properly
		contains = strings.ReplaceAll(contains, strings.ToLower(string(private)), "")
	}

	return contains //returns the string properly
}