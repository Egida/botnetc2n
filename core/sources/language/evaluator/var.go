package evaluator

import (
	"errors"
)

//correctly properly tries to find the scope
//the scope will store information about a variables
func (e *Evaluator) findScope(m string) (*Scope, error) {
	//ranges through the different objects
	//this will allow properly handle without issues
	for s := range e.GuideScope {
		//tries to match the name properly
		//this will allow us to properly try to match without issues
		if e.GuideScope[s].Name == m {
			//returns the scope correctly
			//this will allow us to properly handle
			return &e.GuideScope[s], nil
		}
	}
	//returns the error correctly
	//this will properly without issues happening
	return nil, errors.New("unknown variables couldn't be found correctly")
}

func (e *Evaluator) updateObject(scope *Scope) {

	//ranges througout the map correctly
	//this will ensure its done properly without errors
	for value := range e.GuideScope {

		if e.GuideScope[value].Name == scope.Name {
			e.GuideScope[value] = *scope //updates the scope
			return //stops the function correctly
		}
	}
}