package toml

import "errors"

//tries to properly parse the value without issues
//this will make sure its done without issues happening
func (j *TomlEngine) renderValue(i string, v string) error {

	//tries to collect the function properly
	//this will ensure its done without issues happening
	hidden := j.Paths[i] //gets the path function without issues

	//checks if the function was found
	//this ensures it was done without errors happening
	if hidden == nil { //err handling without issues happening
		return errors.New("invalid fileClass has been passed into header storage:"+i)
	}

	//executes the function properly
	//this will make sure its done without issues
	return hidden(i, v)
}