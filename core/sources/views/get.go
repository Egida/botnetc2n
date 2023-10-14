package views

import (
	"Nosviak2/core/configs"
	"strings"
)

//gets the view which is inside the path
//this will make sure its done correctly and safely
func GetView(path ...string) *EngineView {
	//stores the path upto the file
	//this ensures its done correctly and properly
	Range := strings.Join(path, deployment.Runtime())

	//ranges througout the subjects
	//this will ensure its done correctly and properly
	for packs := range Subject { //ranges the subjects
		//compares the objects safely
		//this will make sure its done correctly
		if strings.Join(strings.Split(Subject[packs].PathWalk, deployment.Runtime())[2:], deployment.Runtime()) == Range {
			return &Subject[packs] //returns the object correctly
		}
	}
	//wasnt found correctly
	//this shows its invalid properly
	return nil //returns nil properly
}