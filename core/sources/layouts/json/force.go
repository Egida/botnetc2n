package json

import (
	"errors"
)

//this will properly work the forced input
//this ensures its done without issues happening
func (j *JsonEngine) workForced(given []string) error {
	//ranges througout the given objects
	//makes sure its done without issues happening
	for manage := range j.Hierarchy {
		//tries to find the object
		//this will make sure its gotten without issues
		err := j.workArray(j.Hierarchy[manage], given) //works the array properly
		if err == nil { //basic error handling
			continue //returns the error
		}

		//returns the Hierarchy error correctly
		//this will make sure its not ignored properly
		return errors.New("hierarchy file is missing from the files listed")
	}

	//this will ensure all objects were located properly
	//this is a plus sign from the function meaning its safe
	return nil
}

//this will check inside the array properly
//makes sure its without issues happening on request
func (j *JsonEngine) workArray(object string, given []string) error {
	//ranges througout the HObjects properly
	//this will ensure its done without issues happening
	for tok := range given {

		//completes the check as its valid
		//this will return nil as it was found properly
		if given[tok] == object {
			return nil //returns nil as its invalid
		}
	}

	//returns the error correctly
	//this will allow the header functions to know why
	return errors.New("object wasn't located properly within the array")
}