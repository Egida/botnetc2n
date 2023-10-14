package routes

import (
	attacks "Nosviak2/core/attack"
	"Nosviak2/core/clients/commands"
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/layouts/toml"
	"sort"
	"strings"
	"unicode/utf8"
)

type Terminal struct {
	s sessions.Session //session information stored

	Current []string //stores all current system
	Range   int //stores the current range inside
	AttackCleanUp bool //stores the attack clean up guide
}


func NewTerm(s *sessions.Session) *Terminal {
	return &Terminal{s: *s, Current: make([]string, 0), Range: 0}
}

//AutoCompleteCallBack predicts the callback from a command properly
func (t *Terminal) AutoCompleteCallBack(line string, pos int, key rune) (newLine string, newPos int, ok bool) {

	//target feedback and more
	if toml.AttacksToml.Attacks.Callback {
		//possible attack method detection here currently
		if strings.Count(line, " ") >= 1 && attacks.ValidateAttackPrefix(strings.Split(line, " ")[0], toml.AttacksToml.Attacks.Prefix) {
			group, _ := attacks.BufferCallback(&t.s, line, t.AttackCleanUp) //sorts the group properly
			t.AttackCleanUp = group //sets the attack clean up mode properly
			return //runs the buffer for attacks callback properly
		}
	}

	//only allows tab charaters through
	//this will also ensure they are at the end
	if key != '\t' || pos != len(line) {
		return 
	}

	//position of the command properly
	//this will ensure its done safely
	switch strings.Count(line, " ") {

	case 0: //main command header

		//checks for the attack prefix properly
		//this will ensure its done without issues happening
		if strings.Contains(strings.Split(line, " ")[0], toml.AttacksToml.Attacks.Prefix) {

			//range through the system here properly
			//this has detected scrolling properly and safely
			if t.Current != nil && len(t.Current) > t.Range && t.Current[t.Range] == strings.Split(strings.Replace(line, toml.AttacksToml.Attacks.Prefix, "", -1), " ")[0] {

				//checks within guidelines properly
				if t.Range + 1 <= len(t.Current) - 1 { //checks guideline properly
					t.Range++; return "."+t.Current[t.Range], len("."+t.Current[t.Range]), true
				}

				t.Range = 0 //stores the current system range properly
				return "."+t.Current[t.Range], len("."+t.Current[t.Range]), true
			}	

			//tries to fetch the method properly and safely
			//this will ensure its done without issues happening
			controls := LookupMethods(strings.Split(strings.Replace(line, toml.AttacksToml.Attacks.Prefix, "", -1), " ")[0])
			if len(controls) <= 0 { //length checks properly
				return //return the system
			}


			t.Current = controls//holds commands
			t.Range   = 0 //stores the range properly. holds the position safely
	
			//returns the current position properly and safely
			//this will ensure its done properly and safely properly
			return "."+t.Current[t.Range], utf8.RuneCountInString("."+t.Current[t.Range]), true
		}	

		//range through the system here properly
		//this has detected scrolling properly and safely
		if t.Current != nil && len(t.Current) > t.Range && t.Current[t.Range] == strings.Split(line, " ")[0] {

			//checks within guidelines properly
			if t.Range + 1 <= len(t.Current) - 1 { //checks guideline properly
				t.Range++; return t.Current[t.Range], len(t.Current[t.Range]), true
			}

			t.Range = 0 //stores the current system range properly
			return t.Current[t.Range], len(t.Current[t.Range]), true
		}

		//LookupCommand predicts the command properly
		//produces an array which will help within the state
		possible := LookupCommand(line)
		if len(possible) <= 0 { //length
			return //returns the loops
		}


		t.Current = ToArray(possible, make([]string, 0)) //holds commands
		t.Range   = 0 //stores the range properly. holds the position safely

		//returns the current position properly and safely
		//this will ensure its done properly and safely properly
		return t.Current[t.Range], utf8.RuneCountInString(t.Current[t.Range]), true

	case 1: //subcommand position here properly
		
		//range through the system here properly
		//this has detected scrolling properly and safely
		if t.Current != nil && len(t.Current) > t.Range && t.Current[t.Range] == strings.Split(line, " ")[1] {
		
			//checks within guidelines properly
			if t.Range + 1 <= len(t.Current) - 1 { //checks guideline properly
				t.Range++; return strings.Split(line, " ")[0] + " " + t.Current[t.Range], len(strings.Split(line, " ")[0] + " " + t.Current[t.Range]), true
			}
		
			t.Range = 0 //stores the current system range properly
			return strings.Split(line, " ")[0] + " " + t.Current[t.Range], len(strings.Split(line, " ")[0] + " " + t.Current[t.Range]), true
		}

		//gets the command from the array properly
		//this will ensure it only works on valid systems
		prefixCommand := commands.TryCommand(strings.Split(line, " ")[0])
		if prefixCommand == nil { //invalid command
			return 
		}

		//tries to fetch all predicts subcommands properly
		//this will ensure its done without issues happening
		subcommand := LookupSubCommand(strings.Split(line, " ")[1], prefixCommand)
		if len(subcommand) <= 0 {
			return
		}

		t.Current = ToArraySub(subcommand, make([]string, 0)) //range
		t.Range   = 0 //sets the current range properly

	
		//subcommand formula has been set at this point properly
		return strings.Split(line, " ")[0] + " " + t.Current[t.Range], len(strings.Split(line, " ")[0] + " " + t.Current[t.Range]), true

	default: //anything higher or equal to 2 will be display
		
		//range through the system here properly
		//this has detected scrolling properly and safely
		if t.Current != nil && len(t.Current) > t.Range && t.Current[t.Range] == strings.Split(line, " ")[2] {
		
			//checks within guidelines properly
			if t.Range + 1 <= len(t.Current) - 1 { //checks guideline properly
				t.Range++; return strings.Join(strings.Split(line, " ")[:strings.Count(line, " ")], " ") + " " + t.Current[t.Range], len(strings.Join(strings.Split(line, " ")[:strings.Count(line, " ")], " ") + " " + t.Current[t.Range]), true
			}
		
			t.Range = 0 //stores the current system range properly
			return strings.Join(strings.Split(line, " ")[:strings.Count(line, " ")], " ") + " " + t.Current[t.Range], len(strings.Join(strings.Split(line, " ")[:strings.Count(line, " ")], " ") + " " + t.Current[t.Range]), true
		}

		//grabs the main command for the system properly
		prefixCommand := commands.TryCommand(strings.Split(line, " ")[0])
		if prefixCommand == nil { //checks for the nil pointer properly
			return
		}

		prefixSubCommand := prefixCommand.FindSubs(strings.Split(line, " ")[1])
		if prefixSubCommand == nil || prefixSubCommand.AutoComplete == nil { //checks for the nil pointer properly
			return
		}

		//gets the canidates subcommands properly
		candidates := prefixSubCommand.AutoComplete(&t.s)

		//predicts all possible subcommands properly
		//this will ensure its done without issues happening
		Cands := ArrayMatch(strings.Split(line, " ")[2], candidates)
		if len(Cands) <= 0 { //checks the length properly
			return 
		}

		t.Current = Cands //stores all possible elements
		t.Range   = 0 //sets default pos to 0

		//returns the string properly and 
		return strings.Join(strings.Split(line, " ")[:strings.Count(line, " ")], " ") + " " + t.Current[t.Range], len(strings.Join(strings.Split(line, " ")[:strings.Count(line, " ")], " ") + " " + t.Current[t.Range]), true
	}
}

//LookupCommand grabs all command with the same prefix
//this function is used within the tab completion system
func LookupCommand(prefix string) ([]commands.Command) {

	//takes the prefix in properly
	//lowers down into lower case string
	prefix = strings.ToLower(prefix)

	//stores all commands within the system
	var commandStore []commands.Command = make([]commands.Command, 0)

	//ranges through the command properly
	//this will allow us to check the command
	for command, index := range commands.Commands  {

		//checks for the command properly
		if strings.HasPrefix(command, prefix) {
			commandStore = append(commandStore, *index); continue
		}
	}

	return commandStore //returns the commands
}

//LookupSubCommand looks up the subcommand properly
func LookupSubCommand(prefix string, command *commands.Command) ([]commands.SubCommand) {
	//takes the prefix in properly
	//lowers down into lower case string
	prefix = strings.ToLower(prefix)

	//stores all commands within the system
	var commandStore []commands.SubCommand = make([]commands.SubCommand, 0)

	//ranges through the command properly
	//this will allow us to check the command
	for _, index := range command.SubCommands  {

		//checks for the command properly
		if strings.HasPrefix(index.SubcommandName, prefix) {
			commandStore = append(commandStore, index); continue
		}
	}

	return commandStore //returns the commands
}


//GetSubcommands gets all the attacks method
func LookupMethods(prefix string) ([]string) {

	//takes the prefix in properly
	//lowers down into lower case string
	prefix = strings.ToLower(prefix)

	//stores all commands within the system
	var commands []string = make([]string, 0)

	//ranges through the command properly
	//this will allow us to check the command
	for _, index := range attacks.AllMethods(make([]*attacks.Method, 0)) {

		//checks for the command properly
		if strings.HasPrefix(index.Name, prefix) {
			commands = append(commands, index.Name); continue
		}
	}

	return commands
}

//ArrayMatch gets all possible matchs from the array
func ArrayMatch(prefix string, array []string) ([]string) {

	//takes the prefix in properly
	//lowers down into lower case string
	prefix = strings.ToLower(prefix)

	//stores all commands within the system
	var commands []string = make([]string, 0)

	//ranges through the command properly
	//this will allow us to check the command
	for _, index := range array  {

		//checks for the command properly
		if strings.HasPrefix(index, prefix) {
			commands = append(commands, index); continue
		}
	}

	return commands
}

func ToArray(s []commands.Command, src []string) []string {
	for _, system := range s {
		src = append(src, system.CommandName)
	}
	sort.Strings(src)
	return src
}

func ToArraySub(s []commands.SubCommand, src []string) []string {
	for _, system := range s {
		src = append(src, system.SubcommandName)
	}
	sort.Strings(src)
	return src
}
