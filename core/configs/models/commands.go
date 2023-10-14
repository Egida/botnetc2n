package models




//stores the command information
//this will allow the instance to edit it without issues
type CustomCommand struct { //stored in type structure properly
	Description string        `json:"description"`
	Permissions []string      `json:"permissions"`
	Aliases     []string      `json:"aliases"`
	Subcommands []struct{ //stores the structure properly
			Address     string        `json:"name"`
			Description string        `json:"description"`
			Permissions []string      `json:"permissions"`
	} `json:"subcommands"`
}