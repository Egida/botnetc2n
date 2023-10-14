package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs/models"
	"strings"
	"sync"
)

var (
	Commands map[string]*Command = make(map[string]*Command)
	mutex    sync.Mutex
)

type Command struct {
	CommandName        string
	Aliases            []string
	CommandDescription string
	CommandPermissions []string
	CommandFunction    func(s *sessions.Session, cmd []string) error
	SubCommands        []SubCommand
	InvalidSubCommand  func(s *sessions.Session, cmd []string) error
	CustomCommand      string
	BinCommand         *models.BinCommand
}
type SubCommand struct {
	SubcommandName     string
	Description        string
	CommandPermissions []string
	CommandSplit       string
	SubCommandFunction func(s *sessions.Session, cmd []string) error
	RenderRef          bool
	AutoComplete       func(s *sessions.Session) []string
}

func MakeCommand(c *Command) {
	mutex.Lock()
	defer mutex.Unlock()
	Commands[c.CommandName] = c
}

func TryCommand(command string) *Command {
	if cmd := Commands[command]; cmd != nil {
		return cmd
	}
	for c := range Commands {
		for aliases := range Commands[c].Aliases {
			if strings.ToLower(Commands[c].Aliases[aliases]) == command {
				return Commands[c]
			}
		}
	}
	return nil
}

func (c *Command) FindSubs(inlet string) *SubCommand {
	for pos := range c.SubCommands {
		if c.SubCommands[pos].SubcommandName == strings.SplitAfter(inlet, c.SubCommands[pos].CommandSplit)[0] {
			return &c.SubCommands[pos]
		}
	}
	return nil
}
