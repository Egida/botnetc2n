package tools

import "strings"

//basic needle haystack function
//this will rapidly search through the array
func NeedleHaystack(haystack []string, needle string) (bool, int) {
	//ranges through the array
	//this will ensure its done without errors
	for h := range haystack {
		//compares the options properly
		if haystack[h] == needle { //makes sure they match
			return true, h //returns true properly without issues
		}
	}; return false, 0
}

//basic needle haystack function
//this will rapidly search through the array
func NeedleHaystackOne(haystack []string, needle string) (bool) {
	//ranges through the array
	//this will ensure its done without errors
	for h := range haystack {
		//compares the options properly
		if haystack[h] == needle { //makes sure they match
			return true //returns true properly without issues
		}
	}; return false
}

//basic needle haystack function
//this will rapidly search through the array
func NeedleHaystackContains(haystack []string, needle string) (bool) {
	//ranges through the array
	//this will ensure its done without errors
	for h := range haystack {
		//compares the options properly
		if strings.Contains(haystack[h], needle) { //makes sure they match
			return true //returns true properly without issues
		}
	}; return false
}

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