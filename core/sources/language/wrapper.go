package language

import (
	"Nosviak2/core/clients/sessions"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"Nosviak2/core/sources/layouts/functions"
	"Nosviak2/core/sources/layouts/packages"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/views"
	"Nosviak2/core/sources/webhooks"
	"errors"
	"strings"

	"io"
)

//tries to correctly execute the language
//this will ensure its executed properly without issues happening
func ExecuteLanguage(render []string, wr io.Writer, template [2]string, session *sessions.Session, register map[string]string) error {
	//adds support for the theme render
	//this will allow for proper management without issues
	path := RenderParser(render, session) //properly tries without issues


	//tries to get properly without issues
	//this will get without issues happening
	value := views.GetView(path...) //gets within the theme properly
	if value == nil { //checks if the theme was found properly
		//tries to get the default properly theme without issues happening
		def := views.GetView(render...) //gets the default properly without issues
		if def == nil { //error handles properly without issues happening on reqeust
			return errors.New(strings.Join(render, "/")+" is classed as an invalid branding object")
		}

		//updates the default properly
		//this will select the default branding without issues
		value = def //updates to the default properly without issues
	}

	//checks if webhooks are enabled in this area without issues happening on purpose
	if tools.NeedleHaystackOne(toml.WebhookingToml.Webhooks.Trigger, strings.Join(render, deployment.Runtime())) {
		//performs the webhook event properly
		//this will ensure its done without any errors
		go webhooks.PerformEventWebhook(render, session, register)
	}
	
	//creates the new lexer structure with capture
	//capture will enable usage of the template engine properly without issues
	l := lexer.Make(value.Containing, true) //creates the structure properly
	//tries to correctly execute and run the lexer
	//this will ensure its done properly without errors
	if _, err := l.RunTarget(); err != nil { //error handling for issues properly
		return err //returns the error correctly
	}
	//executes the complete lexer set properly
	//this will make sure its done properly without issues happening
	nodes, err := parser.MakeParser(l).RunPath()
	if err != nil { //error handling for issues properly and safely
		return err //returns the error correctly
	}
	//follows the guide properly without issues
	//this will make sure its correctly followed without issues happening
	eval := evaluator.MakeEval(nodes, make([]evaluator.Scope, 0), wr, template, session)
	//ranges and registers the packages
	//this will make sure its done correctly
	for _, packageP := range packages.Packages {
		eval.AddPackage(&packageP) //adds the package correctly
	}
	//ranges and registers the functions
	//this will make sure its done correctly
	for _, functionP := range functions.Functions { //saves into the array correctly
		eval.AddFunction(functionP.FunctionName, functionP.Function)
	}

	//registers all the dynamic functions
	//this will be set by the header function without issues
	for key, value := range register { //ranges through the register
		eval.VariableFunction(key, lexer.String, value)
	}
	//executes the eval handling properly
	//this will make sure its done correctly
	if _, err := eval.FollowGuide(); err != nil {
		return err //returns the error correctly
	}

	return nil
}

//tries to correctly execute the language
//this will ensure its executed properly without issues happening
func ExecuteLanguageText(nodes string, wr io.Writer, template [2]string, session *sessions.Session, register map[string]string) error {
	//creates the new lexer structure with capture
	//capture will enable usage of the template engine properly without issues
	l := lexer.Make(nodes, true) //creates the structure properly
	//tries to correctly execute and run the lexer
	//this will ensure its done properly without errors
	if _, err := l.RunTarget(); err != nil { //error handling for issues properly
		return err //returns the error correctly
	}
	//executes the complete lexer set properly
	//this will make sure its done properly without issues happening
	nodesP, err := parser.MakeParser(l).RunPath()
	if err != nil { //error handling for issues properly and safely
		return err //returns the error correctly
	}
	//follows the guide properly without issues
	//this will make sure its correctly followed without issues happening
	eval := evaluator.MakeEval(nodesP, make([]evaluator.Scope, 0), wr, template, session)
	//ranges and registers the packages
	//this will make sure its done correctly
	for _, packageP := range packages.Packages {
		eval.AddPackage(&packageP) //adds the package correctly
	}
	//ranges and registers the functions
	//this will make sure its done correctly
	for _, functionP := range functions.Functions { //saves into the array correctly
		eval.AddFunction(functionP.FunctionName, functionP.Function)
	}

	//registers all the dynamic functions
	//this will be set by the header function without issues
	for key, value := range register { //ranges through the register
		eval.VariableFunction(key, lexer.String, value)
	}
	//executes the eval handling properly
	//this will make sure its done correctly
	if _, err := eval.FollowGuide(); err != nil {
		return err //returns the error correctly
	}

	return nil
}

//this will correctly parse the path route
//allows for proper management without issues
func RenderParser(render []string, s *sessions.Session) []string {
	if s.User.Theme == "default" { //checks for the default properly
		return render //returns the default properly without parsing statements
	}

	//saves into the array correctly and properly
	path := append(s.BrandingPath, render...)
	return path //saves into properly
}