package routes

import (
	attacks "Nosviak2/core/attack"
	"Nosviak2/core/clients/commands"
	"Nosviak2/core/clients/sessions"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/logs"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/tools"
	"log"
	"path/filepath"

	"time"

	"strings"

	"golang.org/x/term"
)

//creates the new prompt properly
//this will safely handle without issues happening
func NewPrompt(session *sessions.Session) error {

	//tries to print the home screen
	//this will render everytime this function is called onwards
	if err := language.ExecuteLanguage([]string{"welcome.itl"}, session.Channel, deployment.Engine, session, make(map[string]string)); err != nil {
		return err //returns the error correctly and properly without issues happening
	}

	//creates the new terminal prompt properly
	//this will be what accepts the inputs without issues
	application := term.NewTerminal(session.Channel, "")

	//sets the auto complete callback system properly
	//this will ensure its done without any errors happening
	application.AutoCompleteCallback = NewTerm(session).AutoCompleteCallBack

	//ranges through inputs properly
	//this will ensure its done correctly without issues
	for { //loops through inputs correctly

		//prints the prompt correctly
		//this will allow for better handling control for the users views which are selected properly
		if err := language.ExecuteLanguage([]string{"prompt.itl"}, session.Channel, deployment.Engine, session, make(map[string]string)); err != nil {
			return err //returns the error correctly and properly without issues happening
		}

		//reads the users input correctly
		//this will ensure its done correctly without issues happening
		message, err := application.ReadLine()
		//saves into the written book
		//this will allow for recall without issues
		session.Written = append(session.Written, message+"\r\n")

		//this will ensure its done without issues
		if err != nil { //error handles the read
			//creates a new terminal app properly
			//this will make sure its done correctly without issues happening
			application = term.NewTerminal(session.Channel, "")
			continue
		}

		sessions.Mutex.Lock()
		sessions.Sessions[session.ID].Idle = time.Now()
		sessions.Mutex.Unlock()

		//properly pushes the command to be executed
		//this will execute the command without issues happening
		if err := ExecuteCommand(message, session); err != nil { //error handles the statement properly
			if err := language.ExecuteLanguage([]string{"errors", "commandError.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": message, "error": err.Error()}); err != nil {
				return err //returns the error correctly and properly without issues happening
			}
		}
	}

}

//executes the incoming command properly
//this will make sure it executes without issues happening
func ExecuteCommand(cmd string, session *sessions.Session) error {
	//splits the command properly
	//this will be used properly without issues happening
	var command []string = strings.Split(cmd, " ") //splits
	if command[0] == "" {                          //validates the 0 length properly
		return nil //returns and ends the function properly
	}

	//executes the language properly
	//this will ensure its done without errors
	if err := language.ExecuteLanguage([]string{"commands", "before-command.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": command[0]}); err != nil {
		return err
	}

	//tries to correctly write the log into the file
	//this will ensure its done without any errors happening
	if err := logs.WriteLog(filepath.Join(deployment.Assets, "logs", "commands.json"), logs.CommandLog{Command: cmd, Args: command, Username: session.User.Username, Time: time.Now()}); err != nil {
		log.Printf("logging fault: %s\r\n", err.Error()) //alerts the main terminal properly
	}

	//this validates that its an attack request
	//this will ensure its done without issues happening
	if attacks.ValidateAttackPrefix(strings.ToLower(command[0]), toml.AttacksToml.Attacks.Prefix) {
		//removes the prefix from the command reqeust properly
		//this will ensure its done without issues happening on reqeust
		command[0] = strings.Replace(command[0], toml.AttacksToml.Attacks.Prefix, "", 1)
		return attacks.MakeAttack(command, session).RunTarget() //runs the attack package properly
	}

	//tries to search for the command
	//this will ensure its done correctly without issues
	cmds := commands.TryCommand(strings.ToLower(command[0])) //tries to get the command

	//checks if the command is equal to nil
	//this will allow us to properly handle without issues
	if cmds == nil { //checks if the command is nil properly and renders the invalid command properly
		return language.ExecuteLanguage([]string{"errors", "command404.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": command[0]})
	}

	//tries to check if the command is blacklisted
	//this will allow for better control without issues
	blocked, _ := tools.NeedleHaystack(toml.ConfigurationToml.AppSettings.NoCmds, strings.ToLower(command[0]))
	if blocked && len(cmds.CustomCommand) <= 0 { //checks for the nil pointer properly
		return language.ExecuteLanguage([]string{"errors", "command404.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": command[0]})
	}

	//creates the rank instance
	//this will start the string handle properly
	rank := ranks.MakeRank(session.User.Username) //makes instance
	rank.SyncWithString(session.User.Ranks)       //syncs with the user

	//checks to see if they can access the object
	//this will allow for better handling without issues
	if len(cmds.CommandPermissions) > 0 && !rank.CanAccessArray(cmds.CommandPermissions) { //ensures the user can access without issues
		return language.ExecuteLanguage([]string{"errors", "command403.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": command[0]})
	}

	//checks for possible subcommands
	//this will ensure they are executed without issues
	if len(command) >= 2 && len(cmds.SubCommands) >= 1 {
		//tries to check for the subcommand properly
		//this will make sure its done without errors happening
		subcommand := cmds.FindSubs(command[1]) //tries to find the subcommand properly

		//checks if the command is nil
		//this will ensure its done without issues
		if subcommand == nil || subcommand.RenderRef { //checks for nil pointers
			if cmds.InvalidSubCommand != nil { //returns invalid subcommand
				return cmds.InvalidSubCommand(session, command) //this will make sure if its set we execute
			} //executes the default branding properly without issues happening on request
			return language.ExecuteLanguage([]string{"errors", "subcommand404.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": command[0], "subcommand": command[1]})
		}

		//properly tries to see if the user can access this object without issues
		//this will ensure either its valid or not without issues and makes sure its valid
		if len(subcommand.CommandPermissions) > 0 && !rank.CanAccessArray(subcommand.CommandPermissions) {
			return language.ExecuteLanguage([]string{"errors", "subcommand403.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command": command[0], "subcommand": command[1]})
		}

		//tries to execute the function properly
		//this will make sure its done without issues happening
		return subcommand.SubCommandFunction(session, command)
	}

	//executes the normal command properly
	//this will ensure its done without any errors happening
	if cmds.CommandFunction == nil && cmds.CustomCommand != "" { //renders the command properly without issues happening
		return language.ExecuteLanguageText(cmds.CustomCommand, session.Channel, deployment.Engine, session, make(map[string]string))
	}

	//checks for bin commands properly
	//this will ensure its not ignored without issues happening
	if cmds.CommandFunction == nil && cmds.BinCommand != nil { //follows bin route
		return commands.RunBinCommand(session, cmds.BinCommand, command) //returns the runbincommand properly
	}

	//tries to execute the function correctly
	//this will safely and properly execute without issues
	//return cmds.CommandFunction(session, command)
	return cmds.CommandFunction(session, command)
}
