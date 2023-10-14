package branding

var ( //users/...

	//databaseErr.itl
	USERS_DatabaseErr string = " While trying to access the users dataset an unforeseen error occurred!\r\n"

	//syntax-ranks.itl
	USERS_SyntaxRanks string = " Syntax: users <<rank()>>=<boolean> <usernames>\r\n"
)

var ( //add_days

	//add_days\database-fault.itl
	ADD_DAYS_database_fault string = " Failed to update `<<username()>>` without issues happening\r\n"+
		"\r\n"


	//add_days\invalidUser.itl
	ADD_DAYS_invalidUser string = " The user <<username()>> is an unclassified user!\r\n"+
		"\r\n"


	//add_days\success.itl
	ADD_DAYS_success string = " Successfully added <<days()>> hours to <<username()>> properly\r\n"+
		"\r\n"


	//add_days\syntax.itl
	ADD_DAYS_syntax string = " syntax: users add-days=<int> <usernames...>\r\n"+
		"\r\n"


	//add_days\tracer-error.itl
	ADD_DAYS_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //add_hours

	//add_hours\database-fault.itl
	ADD_HOURS_database_fault string = " Failed to update `<<username()>>` without issues happening\r\n"+
		"\r\n"


	//add_hours\invalidUser.itl
	ADD_HOURS_invalidUser string = " The user <<username()>> is an unclassified user!\r\n"+
		"\r\n"


	//add_hours\success.itl
	ADD_HOURS_success string = " Successfully added <<days()>> days to <<username()>> properly\r\n"+
		"\r\n"


	//add_hours\syntax.itl
	ADD_HOURS_syntax string = " syntax: users add-days=<int> <usernames...>\r\n"+
		"\r\n"


	//add_hours\tracer-error.itl
	ADD_HOURS_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //admin

	//admin\database-error.itl
	ADMIN_database_error string = " unfortunately we have failed to update <<username()>>'s database structure\r\n"+
		"\r\n"


	//admin\giveFault.itl
	ADMIN_giveFault string = " Failed to give `<<username()>>` the admin level properly\r\n"+
		"\r\n"


	//admin\has.itl
	ADMIN_has string = " the user `<<username()>>` already has admin status!\r\n"+
		"\r\n"


	//admin\invalidUser.itl
	ADMIN_invalidUser string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"


	//admin\not.itl
	ADMIN_not string = " the user `<<username()>>` currently has not got admin status!\r\n"+
		"\r\n"


	//admin\success-false.itl
	ADMIN_success_false string = " We have correctly demoted `\\x1b[38;5;9m<<username()>>\\x1b[0m` from admin!\r\n"+
		"\r\n"


	//admin\success-true.itl
	ADMIN_success_true string = " We have correctly promoted `\\x1b[38;5;2m<<username()>>\\x1b[0m` to admin!\r\n"+
		"\r\n"


	//admin\syntax.itl
	ADMIN_syntax string = " Syntax: users admin=<boolean> <username>...\r\n"+
		"\r\n"


	//admin\takeFault.itl
	ADMIN_takeFault string = "Failed to take `<<username()>>'s` the admin level properly\r\n"+
		"\r\n"


	//admin\tracer-error.itl
	ADMIN_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //banned

	//banned\controlFault.itl
	BANNED_controlFault string = " Failed to control <<username()>>'s banned rank correctly\r\n"+
		"\r\n"


	//banned\false-has.itl
	BANNED_false_has string = " the user \"<<username()>>\" is already banned\r\n"+
		"\r\n"


	//banned\giveFault.itl
	BANNED_giveFault string = " failed to ban <<username()>> properly!\r\n"+
		"\r\n"


	//banned\success-false.itl
	BANNED_success_false string = " Success! correctly unbanned <<username()>> from <<cnc()>>\r\n"+
		"\r\n"


	//banned\success-true.itl
	BANNED_success_true string = " Success! correctly banned <<username()>> from <<cnc()>>\r\n"+
		"\r\n"


	//banned\takeFault.itl
	BANNED_takeFault string = " failed to unban \"<<username()>>\" properly!\r\n"+
		"\r\n"


	//banned\tracer-error.itl
	BANNED_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"


	//banned\true-has.itl
	BANNED_true_has string = " the user \"<<username()>>\" is already banned!\r\n"+
		"\r\n"


	//banned\user-EOF.itl
	BANNED_user_EOF string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"
)

var ( //bypass-blacklist

	//bypass-blacklist\controlFault.itl
	BYPASSBLACKLIST_controlFault string = " Failed to control <<username()>>'s bypass-blacklist rank correctly\r\n"+
		"\r\n"


	//bypass-blacklist\false-has.itl
	BYPASSBLACKLIST_false_has string = " the user \"<<username()>>\" already can bypass the blacklist!\r\n"+
		"\r\n"


	//bypass-blacklist\giveFault.itl
	BYPASSBLACKLIST_giveFault string = " failed to give the user bypass blacklist permissions, <<username()>>\r\n"+
		"\r\n"


	//bypass-blacklist\success-false.itl
	BYPASSBLACKLIST_success_false string = " Success! correctly taken <<username()>> bypass blacklist permissions\r\n"+
		"\r\n"


	//bypass-blacklist\success-true.itl
	BYPASSBLACKLIST_success_true string = "Success! correctly given <<username()>> bypass blacklist permissions\r\n"+
		"\r\n"


	//bypass-blacklist\takeFault.itl
	BYPASSBLACKLIST_takeFault string = " failed to take \"<<username()>>\" bypass blacklist permissions!\r\n"+
		"\r\n"


	//bypass-blacklist\tracer-error.itl
	BYPASSBLACKLIST_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"


	//bypass-blacklist\true-has.itl
	BYPASSBLACKLIST_true_has string = " the user \"<<username()>>\" can already bypass the blacklist!\r\n"+
		"\r\n"


	//bypass-blacklist\user-EOF.itl
	BYPASSBLACKLIST_user_EOF string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"
)

var ( //cooldown

	//cooldown\database-error.itl
	COOLDOWN_database_error string = " unfortunately we have failed to update <<username()>>'s database structure\r\n"+
		"\r\n"


	//cooldown\invalidUser.itl
	COOLDOWN_invalidUser string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"


	//cooldown\success.itl
	COOLDOWN_success string = " Success! cooldown changed from <<oldcooldown()>> -> <<newcooldown()>> for <<username()>>\r\n"+
		"\r\n"


	//cooldown\syntax.itl
	COOLDOWN_syntax string = " Syntax: users cooldown=<int> <usernames>\r\n"+
		"\r\n"


	//cooldown\tracer-error.itl
	COOLDOWN_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //concurrents

	//concurrents\database-error.itl
	CONCURRENTS_database_error string = " unfortunately we have failed to update <<username()>>'s database structure\r\n"+
		"\r\n"


	//concurrents\invalidUser.itl
	CONCURRENTS_invalidUser string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"


	//concurrents\success.itl
	CONCURRENTS_success string = " Success! concurrents changed from <<oldconcurrents()>> -> <<newconcurrents()>> for <<username()>>\r\n"+
		"\r\n"


	//concurrents\syntax.itl
	CONCURRENTS_syntax string = " Syntax: users concurrents=<int> <usernames>\r\n"+
		"\r\n"


	//concurrents\tracer-error.itl
	CONCURRENTS_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //create

	//create\concurrents-atoi.itl
	CREATE_concurrents_atoi string = "\\x1b[0man error occurred when trying to parse (<<concurrents()>>) the concurrents for <<username()>>\r\n"+
		" \r\n"


	//create\concurrents.itl
	CREATE_concurrents string = 		"\\x1b[38;5;15mconcurrents (<<user.concurrents()>>) :: \r\n"


	//create\cooldown-atoi.itl
	CREATE_cooldown_atoi string = "\\x1b[0man error occurred when trying to parse (<<cooldown()>>) the cooldown for <<username()>>\r\n"+
		" \r\n"


	//create\cooldown.itl
	CREATE_cooldown string = 		"\\x1b[38;5;15mcooldown (<<user.cooldown()>>) :: \r\n"


	//create\creation-error.itl
	CREATE_creation_error string = "\\x1b[0mFailed to correctly create the user <<username()>> inside the database\r\n"+
		" \r\n"


	//create\maxtime-atoi.itl
	CREATE_maxtime_atoi string = "\\x1b[0man error occurred when trying to parse (<<maxtime()>>) the maxtime for <<username()>>\r\n"+
		" \r\n"


	//create\maxtime.itl
	CREATE_maxtime string = 		"\\x1b[38;5;15mmaxtime (<<user.maxtime()>>) :: \r\n"


	//create\password.itl
	CREATE_password string = "\\x1b[38;5;15mRandomly generated password -> <<colour.marshal(password(), \"230,255,0\", \"255,0,100\", \"138,43,226\")>>\r\n"+
		" \r\n"


	//create\success.itl
	CREATE_success string = "\\033[38;5;2mSuccess\\x1b[0m! <<username()>> has been created inside the database!\r\n"+
		" \r\n"


	//create\username.itl
	CREATE_username string = 		"\\x1b[38;5;15musername :: \r\n"


	//create\usr-already.itl
	CREATE_usr_already string = "the user <<username()>> already exists!\r\n"+
		" \r\n"
)

var ( //createtoken

	//createtoken\concurrents-atoi.itl
	CREATETOKEN_concurrents_atoi string = "\\x1b[0man error occurred when trying to parse (<<concurrents()>>) the concurrents\r\n"+
		" \r\n"


	//createtoken\concurrents.itl
	CREATETOKEN_concurrents string = 		"\\x1b[38;5;15mconcurrents (<<user.concurrents()>>) :: \r\n"


	//createtoken\cooldown-atoi.itl
	CREATETOKEN_cooldown_atoi string = "\\x1b[0man error occurred when trying to parse (<<cooldown()>>) the cooldown\r\n"+
		" \r\n"


	//createtoken\cooldown.itl
	CREATETOKEN_cooldown string = 		"\\x1b[38;5;15mcooldown (<<user.cooldown()>>) :: \r\n"


	//createtoken\creation-error.itl
	CREATETOKEN_creation_error string = "\\x1b[0mFailed to correctly create the token inside the database\r\n"+
		" \r\n"


	//createtoken\maxtime-atoi.itl
	CREATETOKEN_maxtime_atoi string = "\\x1b[0man error occurred when trying to parse (<<maxtime()>>) the maxtime\r\n"+
		" \r\n"


	//createtoken\maxtime.itl
	CREATETOKEN_maxtime string = 		"\\x1b[38;5;15mmaxtime (<<user.maxtime()>>) :: \r\n"


	//createtoken\success.itl
	CREATETOKEN_success string = "\\033[38;5;2mSuccess\\x1b[0m! created inside the database!\r\n"+
		" Token: <<token()>>\r\n"+
		" \r\n"
)

var ( //filter

	//filter\syntax.itl
	FILTER_syntax string = " syntax: users filter=<rank>\r\n"+
		"\r\n"
)

var ( //headers

	//headers\header-concurrents.txt
	HEADERS_header_concurrents string = 		"*\\x1b[38;5;15mConns*\r\n"


	//headers\header-cooldown.txt
	HEADERS_header_cooldown string = 		"*\\x1b[38;5;15mCooldown*\r\n"


	//headers\header-id.txt
	HEADERS_header_id string = 		"*\\x1b[38;5;15m#*\r\n"


	//headers\header-maxtime.txt
	HEADERS_header_maxtime string = 		"*\\x1b[38;5;15mTime*\r\n"


	//headers\header-ranks.txt
	HEADERS_header_ranks string = 		"*\\x1b[38;5;15mRanks*\r\n"


	//headers\header-username.txt
	HEADERS_header_username string = 		"*\\x1b[38;5;15mUser*\r\n"
)

var ( //maxtime

	//maxtime\database-error.itl
	MAXTIME_database_error string = " unfortunately we have failed to update <<username()>>'s database structure\r\n"+
		"\r\n"


	//maxtime\invalidUser.itl
	MAXTIME_invalidUser string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"


	//maxtime\success.itl
	MAXTIME_success string = " Success! MaxTime changed from <<oldmaxtime()>> -> <<newmaxtime()>> for <<username()>>\r\n"+
		"\r\n"


	//maxtime\syntax.itl
	MAXTIME_syntax string = " Syntax: users maxtime=<int> <usernames>\r\n"+
		"\r\n"


	//maxtime\tracer-error.itl
	MAXTIME_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //moderator

	//moderator\database-error.itl
	MODERATOR_database_error string = 		"unfortunately we have failed to update <<username()>>'s database structure\r\n"


	//moderator\giveFault.itl
	MODERATOR_giveFault string = " Failed to give `<<username()>>` the moderator level properly\r\n"+
		"\r\n"


	//moderator\has.itl
	MODERATOR_has string = " the user `<<username()>>` already has moderator status!\r\n"+
		"\r\n"


	//moderator\invalidUser.itl
	MODERATOR_invalidUser string = " the user \"<<username()>>\" is an unclassified user!\r\n"+
		"\r\n"


	//moderator\not.itl
	MODERATOR_not string = " the user `<<username()>>` currently has not got moderator status!\r\n"+
		"\r\n"


	//moderator\success-false.itl
	MODERATOR_success_false string = "We have correctly demoted `\\x1b[38;5;9m<<username()>>\\x1b[0m` from moderator!\r\n"+
		"\r\n"


	//moderator\success-true.itl
	MODERATOR_success_true string = " We have correctly promoted `\\x1b[38;5;2m<<username()>>\\x1b[0m` to moderator!\r\n"+
		"\r\n"


	//moderator\syntax.itl
	MODERATOR_syntax string = " Syntax: users moderator=<boolean> <username>...\r\n"+
		"\r\n"


	//moderator\takeFault.itl
	MODERATOR_takeFault string = "Failed to take `<<username()>>'s` the moderator level properly\r\n"+
		"\r\n"


	//moderator\tracer-error.itl
	MODERATOR_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //newuser

	//newuser\already_false.itl
	NEWUSER_already_false string = " `<<username()>>` already isn't a newuser!\r\n"+
		"\r\n"


	//newuser\already_true.itl
	NEWUSER_already_true string = " `<<username()>>` is already classified as a newuser\r\n"+
		"\r\n"


	//newuser\database-fault.itl
	NEWUSER_database_fault string = " We were currently unable to edit `<<username()>>'s` newuser col\r\n"+
		"\r\n"


	//newuser\invalidUser.itl
	NEWUSER_invalidUser string = " `<<username()>>` is an unclassified username! \r\n"+
		"\r\n"


	//newuser\success_false.itl
	NEWUSER_success_false string = " Correctly removed `<<username()>>` newuser status!\r\n"+
		"\r\n"


	//newuser\success_true.itl
	NEWUSER_success_true string = " Correctly gave `<<username()>>` newuser status!\r\n"+
		"\r\n"


	//newuser\syntax.itl
	NEWUSER_syntax string = " users newuser=<option> <users>\r\n"+
		"\r\n"


	//newuser\tracer-error.itl
	NEWUSER_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //password

	//password\database-error.itl
	PASSWORD_database_error string = " Sorry, we have failed to correctly update <<username()>>'s password!\r\n"+
		"\r\n"


	//password\invalidUser.itl
	PASSWORD_invalidUser string = " the user (<<username()>>) is classed as invalid!\r\n"+
		"\r\n"


	//password\success.itl
	PASSWORD_success string = " Success! we have updated <<username()>>'s password\r\n"+
		"\r\n"


	//password\syntax.itl
	PASSWORD_syntax string = " Syntax: users password=<string> <usernames>\r\n"+
		"\r\n"


	//password\tracer-error.itl
	PASSWORD_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //remove

	//remove\database-EOF.itl
	REMOVE_database_EOF string = " Sorry, while trying to remove \"\\x1b[38;5;9m<<username()>>\\x1b[0m\" an error happened!\r\n"+
		"\r\n"


	//remove\success.itl
	REMOVE_success string = " Success! we have correctly removed <<username()>> from the database\r\n"+
		"\r\n"


	//remove\syntax.itl
	REMOVE_syntax string = " Syntax: users remove <users>\r\n"+
		"\r\n"


	//remove\tracer-error.itl
	REMOVE_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"


	//remove\user-EOF.itl
	REMOVE_user_EOF string = " the user (<<username()>>) is classed as invalid properly!\r\n"+
		"\r\n"
)

var ( //slaves

	//slaves\database-fault.itl
	SLAVES_database_fault string = "We were currently unable to edit `<<username()>>'s` max_slaves col\r\n"+
		"\r\n"


	//slaves\invalidUser.itl
	SLAVES_invalidUser string = " `<<username()>>` is an unclassified username!\r\n"+
		"\r\n"


	//slaves\success.itl
	SLAVES_success string = " Correctly updated `<<username()>>` max_slave interface!\r\n"+
		"\r\n"


	//slaves\syntax.itl
	SLAVES_syntax string = " Syntax: users slaves=<amount> <users>\r\n"+
		"\r\n"


	//slaves\tracer-error.itl
	SLAVES_tracer_error string = " the <<username()>> obtains higher permissions than yourself!\r\n"+
		"\r\n"
)

var ( //values

	//values\value-concurrents.txt
	VALUES_value_concurrents string = 		"*\\x1b[38;5;15m<<$concurrents>>*\r\n"


	//values\value-cooldown.txt
	VALUES_value_cooldown string = 		"*\\x1b[38;5;15m<<$cooldown>>*\r\n"


	//values\value-id.txt
	VALUES_value_id string = 		"*\\x1b[38;5;15m<<$id>>*\r\n"


	//values\value-maxtime.txt
	VALUES_value_maxtime string = 		"*\\x1b[38;5;15m<<$maxtime>>*\r\n"


	//values\value-ranks.txt
	VALUES_value_ranks string = 		"*<<$ranks>>*\r\n"


	//values\value-username.txt
	VALUES_value_username string = 		"*\\x1b[38;5;15m<<$username>>*\r\n"
)

var ( //view

	//view\invalid-user.itl
	VIEW_invalid_user string = " `<<user()>>` is an unclassified username!\r\n"+
		"\r\n"


	//view\syntax.itl
	VIEW_syntax string = " syntax: users view <username>\r\n"+
		"\r\n"
)