package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"

)

func init() {
	MakeCommand(&Command{
		CommandName:        "clear",
		Aliases:            []string{"cls"},
		CommandPermissions: make([]string, 0),
		CommandDescription: "clear all past rendered modules",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			return language.ExecuteLanguage([]string{"clear-splash.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
		},
	})
}