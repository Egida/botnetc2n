package evaluator

import (
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/language/parser"
	"errors"
	"strings"

	"github.com/oleksandr/conditions"
)

func (e *Evaluator) ConditionsWorker(p *parser.Conditional) error {

	//this will properly make sure its safe
	//makes sure its done safely without issues happening
	for section := 0; section < len(p.Sections); section++ {

		//checks for nil sections properly
		//this will ensure its done without issues happening
		if p.Sections[section] == nil || len(p.Sections) <= 0 {
			return errors.New("invalid system has happened on request without issues")
		}

		//tries to validate the token section without issues
		//this will make sure its not invalid without errors happening
		if p.Sections[section][0].TokenType() != lexer.Indent {
			continue //continues the loop without issues
		}

		//this will trigger the variable section
		//makes sure its not ignored without issues happening
		if len(p.Sections[section]) == 1 {
			//tries to find the scope without issues happening
			//this will make sure its safe without issues happening
			Value, err := e.findScope(p.Sections[section][0].Literal())
			if err != nil || Value == nil { //error handles the request without issues
				return errors.New("can't find the variable `"+p.Sections[section][0].Literal()+"`")
			}

			//creates the array correctly and properly
			p.Sections[section] = make([]lexer.Token, 0) //clears array
			p.Sections[section] = append(p.Sections[section], *Value.TokenValue) //inserts into array
			continue //continues to loop again properly
		}

		//tries to correctly parse the function without issues happening
		//this will ensure its done safely without errors happening on request
		Path, err := parser.MakeTokens(p.Sections[section], 0).ExecuteFunction()
		if err != nil || Path == nil { //error handles correctly and properly
			return errors.New("failed to completely parse function within condition statement")
		}

		//tries to correctly execute the function without issues
		//this will make sure its done without issues happening on request
		Objects, err := e.executeFunction(Path) //executes the function
		if err != nil || len(Objects) == 0 {
			return errors.New("either function returns no arguments or error happened")
		}

		if len(Objects) < 1 {
			return errors.New("function must only return one value within system")
		}

		p.Sections[section] = make([]lexer.Token, 0) //clears array
		p.Sections[section] = append(p.Sections[section], *lexer.NewToken(Objects[0].Literal, Objects[0].Type, Path.Labels[0].Position())) //inserts into array
		continue
	}

	var rawUTL string = ""
	//ranges througout the system sections again
	//this time we will compress into a string without issues
	for section := 0; section < len(p.Sections); section++ {
		rawUTL += e.joinLabels(p.Sections[section]) + " "
	}
	
	//parses the system links without issues happening
	//this will ensure its done without issues happening
	P := conditions.NewParser(strings.NewReader(rawUTL))
	expr, err := P.Parse() //parses the expression properly
	if err != nil {
		return err
	}

	//tries to correctly execute the interpreter
	//this will make sure its done without issues happening
	Status, err := conditions.Evaluate(expr, nil)
	if err != nil { //error handles and returns error
		return err
	}

	//this will choose what body to execute
	//makes sure the wrong body isn't executes as an issue
	if Status {
		//tries to correctly execute the system
		//this will make sure its not ignored without issues happening
		if _, err := MakeEvalWithMore(p.PositiveBody, e.GuideScope, e.wr, e.templateRcog, e.packages, e.functions, e.internalFunctions, e.session).FollowGuide(); err != nil {
			return err //returns the error correctly
		}
	} else if !Status && len(p.NegativeBody) > 0 {
		//tries to correctly execute the system
		//this will make sure its not ignored without issues happening
		if _, err := MakeEvalWithMore(p.NegativeBody, e.GuideScope, e.wr, e.templateRcog, e.packages, e.functions, e.internalFunctions, e.session).FollowGuide(); err != nil {
			return err //returns the error correctly
		}
	}
	return nil
}