package views

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

//loads all the different branding peices properly
//this will make sure its done without issues happening and it safely
func ReadForever(dir string) ([]string, error) {
	//tries to read the file correctly
	//this will make sure its done correctly without issues
	dirRender, err := ioutil.ReadDir(dir)
	if err != nil { //error handles the statement
		return make([]string, 0), err //returns error
	}
	//stores all the different appFiles
	//this will make sure its properly done without issues
	var Spacer []string = make([]string, 0)
	//ranges througout the dir gathering all the files
	//this will make sure its not ignored and makes it safer
	for PeiceObject := range dirRender {
		//this will check if its a folder properly
		//allows for proper control without issues happening
		if strings.Count(dirRender[PeiceObject].Name(), ".") <= 0 {
			//sets the location information properly
			//this will store the item peice without issues
			Path := filepath.Join(dir, dirRender[PeiceObject].Name())
			//reads the new object correctly
			//this will read that dir without issues
			Vad, err := ReadForever(Path) //reads the dir without issues
			if err != nil { //error handling properly which returns the error
				return make([]string, 0), err
			}
			//saves into the spacer array properly
			//this will make sure its done correct and properly
			Spacer = append(Spacer, Vad...); continue
		}
		//saves into the array properly
		//this will make sure its done correctly without issues
		Spacer = append(Spacer, filepath.Join(dir, dirRender[PeiceObject].Name()))
	}
	//returns the output correctly
	//this will make sure its done correctly and properly
	return Spacer, nil
}