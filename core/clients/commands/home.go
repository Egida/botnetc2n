package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "home",
		Aliases:            []string{"welcome"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "redirect your self back home",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//renders the clear splash information properly and safely without issues happening on request
			return language.ExecuteLanguage([]string{"welcome.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
		},
	})
}