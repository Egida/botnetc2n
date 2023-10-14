package webhooks

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/views"
	"log"

	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	//stores all the valid webhooks which we loaded
	//this will ensure they are all properly handled
	Webhooking map[string]string = make(map[string]string)
	mutex      sync.Mutex //basic termlox control
)

//completes clears the branding map
func Reset() { //resets everything inside
	Webhooking = make(map[string]string)
}

//loads all the webhooks properly
//this will ensure its done without any errors
func RenderWebhooks() error { //return error
	//tries to read all the assets properly
	//this will ensure its done without any errors happening
	assets, err := views.ReadForever(filepath.Join(deployment.Assets, "webhooks"))
	if err != nil { //basic error handling for the req
		return err //returns the error
	}

	//loads all the webhooks and ranges
	//this will render each object into the map
	for item := range assets { //ranges through
		//detects if verbose is enabled or not
		//this will read everything if enabled properly
		if toml.ConfigurationToml.Itl.Verbose { //detects properly
			log.Printf("%s\r\n", assets[item]) //renders the object properly
		}
		//reads the file properly without issues
		//this will ensure without any errors issues
		Value, err := ioutil.ReadFile(assets[item]) //reads the file
		if err != nil { //error handles properly without issues
			return err //returns the error
		}

		mutex.Lock() //saves into the array correctly and properly
		Webhooking[strings.Replace(assets[item], deployment.Assets+deployment.Runtime()+"webhooks"+deployment.Runtime()+"", "", -1)] = string(Value)
		mutex.Unlock() //savesinto the map properly
	}
	return nil
}