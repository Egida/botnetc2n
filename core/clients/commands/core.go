package commands

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//stores the command information
//this will ensure its done without any errors
type CustomCommand struct { //stores in type structure
	//stores the filepath properly
	//this will ensure its done without any errors
	Pathway []string

	//this will ensure its done without any errors
	Body string //stores the command body
}

//stores the path properly
//this will ensure its done without errors
func EngineLoader(dir ...string) (int, error) {
	//tries to read the file properly
	//this will ensure its done without any errors
	paths, err := ioutil.ReadDir(filepath.Join(dir...))
	if err != nil { //error handles the reqeust properly
		return 0, err //returns the error given properly
	}
	//ranges through the array properly
	//this will ensure its done without any errors
	for _, cmd := range paths { //ranges through properly
		//tries to read the file properly
		//this will ensure its done without any errors
		file, err := ioutil.ReadFile(filepath.Join(filepath.Join(dir...), cmd.Name()))
		if err != nil { //error handles the statement properly
			return 0, err //returns the error properly
		}
		//tries to correctly execute the body
		//this will parse without issues happening
		body := strings.Join(strings.Split(string(file), "?>")[1:], "?>")
		//feeds the system into the parse
		//this will ensure it can be properly parsed
		nodes, err := parser.MakeParserRun(lexer.Make(strings.SplitAfter(string(file), "?>")[0], true).RunTarget())
		if err != nil { //error handles properly without issues
			return 0, err //returns the error
		}
		//feeds the first value into the interpreter without issues
		//this will ensure its properly executed and we can access without errors
		eval := evaluator.MakeEval(nodes, make([]evaluator.Scope, 0), nil, deployment.Engine, nil)
		if _, err := eval.FollowGuide(); err != nil { //follows the guide path without errors
			return 0, err //returns the error properly
		}
		//guides the name properly
		//this will ensure its done without any errors
		NameObject, err := eval.GetObject("name") //gets the object
		if err != nil { //error handles the get statement properly
			return 0, err //returns the error properly
		}
		//guides the description properly
		//this will ensure its done without any errors
		DescriptionObject, err := eval.GetObject("description") //gets the object
		if err != nil { //error handles the get statement properly
			return 0, err //returns the error properly
		}
		//guides the permissions properly
		//this will ensure its done without any errors
		PermissionsObject, err := eval.GetObject("permissions") //gets the object
		if err != nil { //error handles the get statement properly
			return 0, err //returns the error properly
		}
		//guides the aliases properly
		//this will ensure its done without any errors
		aliasesObject, err := eval.GetObject("aliases") //gets the object
		if err != nil { //error handles the get statement properly
			return 0, err //returns the error properly
		}

		//properly stores the commandperms
		//this will ensure its done without errors
		var cmds []string = make([]string, 0)
		if len(PermissionsObject.TokenValue.Literal()) > 0 {
			cmds = strings.Split(PermissionsObject.TokenValue.Literal(), ",")
		}
		//tries to register into the map
		//this will ensure its done without any errors
		Commands[NameObject.TokenValue.Literal()] = &Command{
			Aliases: strings.Split(aliasesObject.TokenValue.Literal(), ","), CommandPermissions: cmds,
			CommandDescription: DescriptionObject.TokenValue.Literal(), //registers the description
			CommandName: NameObject.TokenValue.Literal(), //registers command name
			CustomCommand: body, //registers the custom command body
			CommandFunction: nil, //renders the function
			SubCommands: make([]SubCommand, 0),
			InvalidSubCommand: nil, //renders invalid
		}
	}


	//returns the length properly
	//this will ensure its done without errors
	return len(paths), nil
}