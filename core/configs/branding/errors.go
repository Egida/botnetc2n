package branding

var ( //errors

	//errors\bin_missing_args.itl
	ERRORS_bin_missing_args string = " Missing arguments for that statement\r\n"+
		" Wanted: <<minArgs()>>\r\n"+
		" Given: <<givenArgs()>>\r\n"+
		" Cmd: <<cmd()>>\r\n"+
		"\r\n"


	//errors\command403.itl
	ERRORS_command403 string = " missing permissions for `<<command()>>`\r\n"+
		"\r\n"


	//errors\command404.itl
	ERRORS_command404 string = "\\x1b[K    `<<command()>>` is an unclassified command\r\n"+
		"\r\n"


	//errors\commandError.itl
	ERRORS_commandError string = "<?\r\n"+
		"	if user.CanAccess(\"admin\") {\r\n"+
		"		echo(\"Command error: \", error(), \"\\r\\n\")\r\n"+
		"	} else {\r\n"+
		"		echo(\"Error happened while trying to execute \"+command(), \"\\r\\n\")\r\n"+
		"	}\r\n"+
		"?>\r\n"


	//errors\subcommand403.itl
	ERRORS_subcommand403 string = " missing permissions for `<<command()>> <<subcommand()>>`\r\n"+
		"\r\n"


	//errors\subcommand404.itl
	ERRORS_subcommand404 string = " `<<command()>> <<subcommand()>>` is an unclassified subcommand\r\n"+
		"\r\n"
)