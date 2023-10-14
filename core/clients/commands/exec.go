package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/ranks"
	"os/exec"

	"os"

	"strconv"
	"strings"
)

//properly tries to run the command
//this will ensure its done without any errors
func RunBinCommand(session *sessions.Session, b *models.BinCommand, cmds []string) error {

	//checks for invalid permissions
	//allows for proper hanlding without issues
	if len(b.Access) > 0 && !session.CanAccessArray(b.Access) { //checks access to the bin command properly
		return language.ExecuteLanguage([]string{"errors", "command403.itl"}, session.Channel, deployment.Engine, session, map[string]string{"command":b.Name})
	}

	//ranges through the env variables properly
	//this will ensure its done and registered without issues
	for _, s := range b.Env { //ranges through properly without issues
		s = EnvSettings(s, session) //runs the settings sortor properly
		if strings.Count(s,"=") != 1 { //checks for invalid ratio properly
			continue //continues looping properly without issues happening
		}
		
		//sets the env object properly without issues
		if err := os.Setenv(strings.Split(s,"=")[0], strings.Split(s,"=")[1]); err != nil {
			return err //returns the error properly without issues
		}
	}

	//enforces the min args rule
	//this will ensure its followed without issues
	if b.MinArgs > len(cmds) - 1{ //len checking properly and renders the error properly without issues happening
		return language.ExecuteLanguage([]string{"errors", "bin_missing_args.itl"}, session.Channel, deployment.Engine, session, map[string]string{"cmd":cmds[0], "minArgs":strconv.Itoa(b.MinArgs), "givenArgs":strconv.Itoa(len(cmds))})
	}

	//ranges through all the args inside the command
	//this will replace all of commands in their properly
	for proc := range b.Args { //ranges through properly safely
		//ranges through all the args given properly
		//this will ensure its done without any errors
		for p := range cmds[1:] { //ranges through properly
			//replaces the arg without issues happening
			b.Args[proc] = strings.ReplaceAll(b.Args[proc], "<<arg("+strconv.Itoa(p)+")>>", cmds[1:][p])
		}
	}


	//creates the command information
	//this will make sure its done without any errors
	cmd := exec.Command(b.Runtime, b.Args...) //command route
	cmd.Stdout = session.Channel //in
	cmd.Stderr = session.Channel //err
	//cmd.Stdin = session.Channel  //out
	//starts the command properly
	//this will ensure its done without any errors
	if err := cmd.Start(); err != nil { //err handles
		return err //error handles properly without issues
	}
	//waits for the command to finish
	//this will ensure we dont keep passing
	if err := cmd.Wait(); err != nil { //err handles
		return err //returns the error which happened
	}

	//employs the prompt writer properly
	//this will ensure its done without any errors
	return language.ExecuteLanguage([]string{"prompt.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
}



//stores the input string properly
//this will ensure its done without any errors happening
func EnvSettings(s string, session *sessions.Session) string {
	//ranges through all the ranks
	//this will set a value to each inside the string
	for name := range ranks.PresetRanks { //stores the information properly
		s = strings.ReplaceAll(s, "<<$"+name+">>", strconv.FormatBool(session.CanAccess(name)))
	}
	//sets all the user data properly
	//this will ensure its done without any errors
	s = strings.ReplaceAll(s, "<<$length>>", strconv.Itoa(session.Length))
	s = strings.ReplaceAll(s, "<<$height>>", strconv.Itoa(session.Height))
	s = strings.ReplaceAll(s, "<<$ip>>", session.Target)
	s = strings.ReplaceAll(s, "<<$username>>", session.User.Username)
	s = strings.ReplaceAll(s, "<<$id>>", strconv.Itoa(session.User.Identity))
	s = strings.ReplaceAll(s, "<<$maxtime>>", strconv.Itoa(session.User.MaxTime))
	s = strings.ReplaceAll(s, "<<$cooldown>>", strconv.Itoa(session.User.Cooldown))
	s = strings.ReplaceAll(s, "<<$concurrents>>", strconv.Itoa(session.User.Concurrents))
	s = strings.ReplaceAll(s, "<<$maxsessions>>", strconv.Itoa(session.User.MaxSessions))
	s = strings.ReplaceAll(s, "<<$theme>>", session.User.Theme)
	s = strings.ReplaceAll(s, "<<$length>>", strconv.Itoa(session.Length))
	return s //returns the edited string properly
}