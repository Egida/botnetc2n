package views

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"io/ioutil"
	"log"
	"path/filepath"
)

var (
	//stored in array type formula
	//this will make sure its done correctly
	Subject []EngineView = make([]EngineView, 0)
)

//clears all objects from the map without issues happening
func Reset() { //removes all the branding from the map
	Subject = make([]EngineView, 0) //clears the map properly
}

//stores the engineView properly
//this will make sure its done properly and safely
type EngineView struct { //stored in type structure
	//stores the location of the file
	//this will allow us to properly find the file
	PathWalk string //stored in type string correctly
	//stores what the file contains properly
	//this will make sure its done properly without issues
	Containing string //stored in type string correctly
}

//properly tries to gather the peices
//this will make sure its done correctly and safely
func GatherPeices(dir string) error {
	//gets all the files inside properly
	//this will help for better control without issues happening
	Objects, err := ReadForever(dir) //reads the item correctly and properly
	if err != nil { //error handles properly and safely without issues
		return err //returns the error correctly
	}

	//we will also try to read the theme vendor
	//this will ensure its done without any errors
	ThemeObjects, err := ReadForever(filepath.Join(deployment.Assets, "themes"))
	if err != nil { //error handles properly and safely
		return err //returns the error correctly and properly
	}

	//saves the theme objects properly
	Objects = append(Objects, ThemeObjects...)

	//ranges througout the objects
	//this will ensure its done properly without errors
	for peice := range Objects {
		//detects if verbose is enabled or not
		//this will read everything if enabled properly
		if toml.ConfigurationToml.Itl.Verbose { //detects properly
			log.Printf("%s\r\n", Objects[peice]) //renders the object properly
		}
		//tries to read the file correctly
		//this will make sure its done safely without errors
		contains, err := ioutil.ReadFile(Objects[peice])
		if err != nil { //error handles properly and safely
			return err //returns the error correctly
		}

		//fills the subject information correctly
		//this will make sure its done correctly and safely
		Subject = append(Subject, EngineView{PathWalk: Objects[peice], Containing: string(contains)})
		continue //forces the system to keep looping properly
	}

	return nil
}