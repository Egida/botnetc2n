package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"

	"golang.org/x/term"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "chat",
		Aliases:            []string{"chats", "talk"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "enrole you're self inside the chatroom",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//renders the banner properly
			//this will ensure its done without any errors
			if err := language.ExecuteLanguage([]string{"commands", "chat", "banner.itl"}, s.Channel, deployment.Engine, s, make(map[string]string)); err != nil {
				return err //returns the error properly without issues
			}

			s.Chat = true //enables chat incoming properly

			//for loops through messages
			//this will ensure its done without any errors
			for { //loops until break is found properly and safely
				//executes the prompt properly and safely
				//this will ensure its done without any errors happening
				language.ExecuteLanguage([]string{"commands", "chat", "prompt.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
				msg, err := term.NewTerminal(s.Channel, "").ReadLine()
				if err != nil { //error handles the read statement
					return err //returns the error properly
				}

				//checks for the exit messages
				//this will ensure its done without any errors
				if msg == "exit" { //detects the exit command
					s.Chat = false //removes the chat props
					return nil //closes the chat sessions
				}

				//ranges through all the sessions
				//this will ensure its done without any errors
				for se := range sessions.Sessions { //ranges through the sessions
					if sessions.Sessions[se].ID == s.ID || !sessions.Sessions[se].Chat {continue} //stops broadcasts from
					//this will render the message properly
					//allows for better control without issues happening
					language.ExecuteLanguage([]string{"commands", "chat", "incoming-message.itl"}, sessions.Sessions[se].Channel, deployment.Engine, sessions.Sessions[se], map[string]string{"sender":s.User.Username, "message":msg})
					language.ExecuteLanguage([]string{"commands", "chat", "prompt.itl"}, sessions.Sessions[se].Channel, deployment.Engine, sessions.Sessions[se], make(map[string]string))
				}


			}
		},
	})
}