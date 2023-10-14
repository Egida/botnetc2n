package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/views"
	"strconv"

	"strings"
)

func init() {

	MakeCommand(&Command{
		CommandName:        "tokens",
		Aliases:            []string{"itltokens"},
		CommandPermissions: []string{"admin"},
		CommandDescription: "Cast all lexed tokens from a file",
		CommandFunction: func(s *sessions.Session, cmd []string) error {
			//tries to check the arguments check
			//this will ensure its done without issues happening
			if len(cmd) <= 1 { //checks the arguments length properly
				return language.ExecuteLanguage([]string{"commands", "tokens", "syntax.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
			}

			//tries to cache the branding peice
			//this will ensure its done without errors
			peice := views.GetView(strings.Join(cmd[1:], "\\"))
			if peice == nil { //makes sure its not equal to nil properly
				return language.ExecuteLanguage([]string{"commands", "tokens", "branding-eof.itl"}, s.Channel, deployment.Engine, s, map[string]string{"target":strings.Join(cmd[1:], "\\")})
			}

			//properly ranges through the system
			//this will ensure its done without errors happening
			l, err := lexer.Make(peice.Containing, true).RunTarget()
			if err != nil { //makes sure its not equal to nil properly and this will handle the system without issues
				return language.ExecuteLanguage([]string{"commands", "tokens", "branding-eof.itl"}, s.Channel, deployment.Engine, s, map[string]string{"target":strings.Join(cmd[1:], "\\")})
			}

			//ranges through all the tokens
			//this will allow for better handling without issues
			for _, token := range l.Tokens() { //ranges through all the tokens properly
				err := s.Write(PaddingRight(strconv.Itoa(token.Position().Row())+":"+strconv.Itoa(token.Position().Column()), 12)+PaddingRight(strconv.Itoa(int(token.TokenType())), 12)+token.Literal()+"\r\n")
				if err != nil { //error handles the system properly without issues
					continue //this will continue the for looping without issues happening 
				}
			}
			return nil
		},
	})
}

func PaddingRight(text string, size int) string {
	for p := len(text); p < size; p++ {
		text += " "
	}
	return text
}