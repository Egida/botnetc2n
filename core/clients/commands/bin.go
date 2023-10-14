package commands

import (
	"Nosviak2/core/configs/models"
	"io/ioutil"
	"path/filepath"
	//"sync"

	"github.com/naoina/toml"
)


//gets all the bin command objects
//this will ensure its done without any issues
func GetBinSettings(dir ...string) error { //returns error
	//this will try to read the dir properly
	//allows for better control wtihout issues happening
	dirPath, err := ioutil.ReadDir(filepath.Join(dir...))
	if err != nil { //error handles properly without issues
		return err //returns the error properly without issues
	}

	//ranges through all the points given
	//this will ensure its done without any errors
	for rf := range dirPath { //reads through the dir properly
		//checks for toml files only
		//this will ensure its done without issues
		if filepath.Ext(dirPath[rf].Name()) != ".toml" {
			continue //only allows toml files properly
		}

		//tries to read the current rotation file
		//this will ensure its done without any errors
		byteval, err := ioutil.ReadFile(filepath.Join(filepath.Join(dir...), dirPath[rf].Name()))
		if err != nil { //tries to read the value proeprly
			return err //returns the error which was found properly
		}

		//input command properly
		//future to be parsed into execs
		var in models.BinCommand
		//tries to parse the object
		//this will ensure its done without any errors
		if err := toml.Unmarshal(byteval, &in); err != nil {
			return err //returns the error properly
		}

		mutex.Lock()
		//tries to fill the information
		//this will fill with all the stores
		Commands[in.Name] = &Command{
			CommandName: in.Name,
			Aliases: make([]string, 0),
			CommandDescription: in.Description,
			CommandPermissions: in.Access,
			CommandFunction: nil, SubCommands: make([]SubCommand, 0),
			CustomCommand: "", BinCommand: &in, //sets the command properly
		}; mutex.Unlock() //unlocks the settings properly		
	}
	return nil
}