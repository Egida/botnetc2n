package slaves

import "strings"

//Sanatizes the string properly
//this will ensure its done without errors
func Sanatize(s string) string { //returns string
	Name := strings.ReplaceAll(s, "\n", "") //replaces some symbols
	Name = strings.ReplaceAll(Name, "\r", "") //replaces some symbols
	Name = strings.ReplaceAll(Name, "\t", "") //replaces some symbols
	Name = strings.ReplaceAll(Name, "\x1b", "") //replaces some symbols
	Name = strings.ReplaceAll(Name, "\r\n", "") //replaces some symbols
	Name = strings.ReplaceAll(Name, "\x00", "") //replaces some symbols
	return Name
}