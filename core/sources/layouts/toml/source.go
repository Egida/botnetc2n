package toml

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

//stores information about the current
//this allows for proper handling without issues happening
type TomlEngine struct {
	//stores the directory which holds the information
	//this ensures its safely and properly done without issues happening
	Directory string //we will scan this folder for all Toml entities
	//these hierarchy objects must be loaded to perform it safely
	//this will ensure its done without issues happening on requests
	Hierarchy []string //stores all the hierarchy issues
	//stores all the different routes without issues
	//this ensures its done without errors happening on request
	Paths map[string]func(file string, value string) error
}

//creates the engine structure
//this ensures its done without issues
func MakeEngine(dir string, hierarchy []string) *TomlEngine {
	//returns the structure properly
	return &TomlEngine{
		//sets the dir for scanning properly
		//this will make sure its done properly without issues
		Directory: dir,
		//sets the hierarchy properly
		//this will be forced to scan without issues
		Hierarchy: hierarchy,
		//sets the path functions properly
		//this will allow for better handling without issues
		Paths: Objects,
	}
}

//this will properly execute the render engine without issues happening
//this allows for proper handling without issues happening on requests
func (j *TomlEngine) RunEngine() error {
	//tries to correctly read the file
	//this will ensure its done without issues happening
	render, err := ioutil.ReadDir(j.Directory)
	if err != nil { //errors handle properly
		return err //returns the error
	}

	//creates an array to the size of the files
	//this will allow a slot for all information to be created
	var rendered []string = make([]string, len(render))

	//ranges througout the system locations
	//this will ensure its done without issues happening
	for app := range render {

		//makes sure only Toml files are accepted
		//this will make sure nothing else is passed as a Toml file
		if strings.Split(render[app].Name(), ".")[len(strings.Split(render[app].Name(), "."))-1] != "toml" {
			continue //forces the continue properly
		}
		
		//tries to read the file properly
		//this will make sure its done without issues
		object, err := ioutil.ReadFile(filepath.Join(j.Directory, render[app].Name())) 
		if err != nil { //error handling properly without issues
			return err //returns the error
		}

		//tries to correctly parse the value without issues
		//this will make sure its done without errors happening
		err = j.renderValue(render[app].Name(), string(object))
		if err != nil { //error handles properly without issues
			return err //returns the error without issues
		}

		//stores into the array properly
		//this will ensure its done without issues
		rendered[app] = render[app].Name()
	}

	//this is the last function inside
	//this validates the rest of paths without issues
	return j.workForced(rendered) //returns the function outcome
}
