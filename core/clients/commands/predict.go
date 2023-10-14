package commands

import (
	"fmt"
	"strings"
)

//creates our new cache
//this is used to store information
func New() *TermCache {
	return &TermCache{}
}

//stores the structure cache
//used to hold information properly
type TermCache struct {
	Possibly []string //all possible commands
	System int
}

//AutoCompleteCallBack completes the callback for tab
//allows for tab completion with statements within the system
func (t *TermCache) AutoCompleteCallBack(line string, pos int, key rune) (newLine string, newPos int, ok bool) {

	//checks for tab completion properly
	//this will only accept tab through the system
	if key != '\t' || pos != len(line) {
		return //ends the system
	}

	switch strings.Count(line, " ") {


	case 0: //main command header



		if len(t.Possibly) > t.System && strings.Split(line, " ")[0] == t.Possibly[t.System] {
			fmt.Println(t.Possibly[t.System], strings.Split(line, " ")[0], t.Possibly)
			if t.System + 1 < len(t.Possibly) {
				t.System++
				return t.Possibly[t.System], len(t.Possibly[t.System]), true
			}

			t.System = 0
			return t.Possibly[t.System], len(t.Possibly[t.System]), true
		}

		t.Possibly = make([]string, 0) //resets the array properly

		resolved := LookupCommand(line)
		fmt.Println(line, len(resolved))
		if len(resolved) <= 0 {
			return
		}

		if len(resolved) == 1 {
			return resolved[0].CommandName, len(resolved[0].CommandName), true
		}

		t.Possibly = ToArray(resolved, make([]string, 0)) //changes from []Command to []string
		return resolved[0].CommandName, len(resolved[0].CommandName), true //returns the Command
	}


	return
}

func ToArray(s []Command, src []string) []string {
	for _, system := range s {
		src = append(src, system.CommandName)
	}
	return src
}

//LookupCommand grabs all command with the same prefix
//this function is used within the tab completion system
func LookupCommand(prefix string) ([]Command) {

	//takes the prefix in properly
	//lowers down into lower case string
	prefix = strings.ToLower(prefix)

	//stores all commands within the system
	var commands []Command = make([]Command, 0)

	//ranges through the command properly
	//this will allow us to check the command
	for command, index := range Commands {

		//checks for the command properly
		if strings.HasPrefix(command, prefix) {
			commands = append(commands, *index); continue
		}
	}

	return commands //returns the commands
}