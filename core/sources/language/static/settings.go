package static

import (
	deployment "Nosviak2/core/configs"
	"io/ioutil"
	"path/filepath"
)

var ( //stores all properly and safely without issues
	Controls map[string]string = make(map[string]string)
)

//gets all possible objects within
//this will ensure its done without errors
func GetStatic(dir string) error { //err handles

	//tries to read the dir properly
	//this will make sure its done without errors
	dirs, err := ioutil.ReadDir(filepath.Join(deployment.Assets, dir))
	if err != nil { //err handles properly
		return err
	}

	//ranges through all the files
	//this will make sure its done without errors
	for _, settingConfig := range dirs { //ranges through

		//reads the file properly and safely
		//this will ensure its done without errors
		system, err := ioutil.ReadFile(filepath.Join(deployment.Assets, dir, settingConfig.Name()))
		if err != nil { //err handles properly
			return err
		}
		
		//saves into the system properly and safely
		Controls[settingConfig.Name()] = string(system)
	}

	return nil
}