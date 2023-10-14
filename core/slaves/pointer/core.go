package pointer

import (
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/slaves/qbot"
	"Nosviak2/core/sources/layouts/toml"
	"strings"
)

//builds the string properly
//this will ensure its done without issues
func BuildString() int { //returns an int properly

	var pointer int = 0

	//ranges through all the systems inside the file
	//this will make sure its done without any errors happening
	for _, declare := range strings.Split(toml.ConfigurationToml.Pointer.Write, ",") {
		switch declare { //switchs the declare properly
		case "mirai": //mirai slaves
			pointer += len(mirai.MiraiSlaves.All)
		case "qbot": //qbot slaves
			pointer += len(qbot.QbotClients)
		}
	}
	
	return pointer
}