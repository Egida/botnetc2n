package deployment

import "runtime"

var (
	//stores the template information properly
	//this will help with better storage without issues happening
	Engine [2]string = [2]string{"<<", ">>"} //sets the open engine parts

	//stores if nosviak2 is running inside debug mode
	//this will allow for proper handling without issues
	DebugMode bool = false

	//stores all the information properly
	//this will allow for better control without issues
	JsonHierarchy []string = []string{"assets"+Runtime()+"config.json", "assets"+Runtime()+"attacks"+Runtime()+"apis.json", "assets"+Runtime()+"commands.json"}
	TomlHierarchy []string = []string{"server.toml", "attacks.toml", "decoration.toml", "ranks.toml", "spinners.toml", "webhooks.toml", "api.toml", "themes.toml", "ip_rewrite.toml", "entry.toml", "fake_slaves.toml"}

	//sets the creators information properly
	Creators []string = []string{"FB"} //only dev = FB
)

//gets the runtime split needed
//this will ensure its done without any errors
func Runtime() string { //returns string
	if runtime.GOOS == "windows" {
		return "\\" //windows
	}; return "/" //detects runtime
}