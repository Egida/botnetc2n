package language

import "strconv"

//returns the default map correctly and properly
func CreateDefaut(user string, session int64) map[string]string {
	//returns the map properly without issues happening and makes sure its secure without errors
	return map[string]string{"username": user, "session": strconv.Itoa(int(session))}
}