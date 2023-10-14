package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"strconv"
	"strings"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "broadcast",
		Aliases:            []string{"announce"},
		CommandPermissions: []string{"admin"},
		CommandDescription: "message all online sessions",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//checks length properly
			//this will act as a syntax checker properly
			if len(cmd) <= 1 { //only accepting messages with at least 2 bodies
				return language.ExecuteLanguage([]string{"commands", "broadcast", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//this will be launched to the sessions
			msg := cmd[1:] //stores the message
			
			//ranges through the sessions
			//this will ensure its done without errors
			for _, session := range sessions.Sessions { //ranges through
				//checks that we dont broadcast to myself
				//makes sure its done without issues properly
				if session.ID == s.ID { //checks the id properly
					continue //continues looping properly without issues
				}

				//writes the alert properly without issues happening 
				//this will make sure its done without any errors happening
				language.ExecuteLanguage([]string{"alerts", "broadcast.itl"}, session.Channel, deployment.Engine, session, map[string]string{"sender":s.User.Username, "message":lexer.AnsiUtil(strings.Join(msg, " "), lexer.Escapes)})
			}

			//renders the success message properly
			//this will alert me that its done without any issues
			return language.ExecuteLanguage([]string{"commands", "broadcast", "success.itl"}, s.Channel, deployment.Engine, s, map[string]string{"message":strings.Join(msg, " "), "sent":strconv.Itoa(len(sessions.Sessions) - 1)})
		},
	})
}