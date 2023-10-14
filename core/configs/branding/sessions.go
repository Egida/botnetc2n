package branding


var ( //sessions/...

	//max_sessions_reached.itl
	SESSION_MaxSessions string = "<<clear()>>You have reached your max sessions!\r\n"+
	"open Sessions: <<user.opensessions()>> max Sessions: <<user.maxsessions()>>\r\n"+
	"<<sleep(10000)>>"
)

var ( //headers

	//headers\connected.txt
	HEADERS_connected string = 		"*\\x1b[38;5;15mConnected*\r\n"


	//headers\id.txt
	HEADERS_id string = 		"*\\x1b[38;5;15m#*\r\n"


	//headers\idle.txt
	HEADERS_idle string = 		"*\\x1b[38;5;15mIdle*\r\n"


	//headers\ip.txt
	HEADERS_ip string = 		"*\\x1b[38;5;15mIP*\r\n"


	//headers\ranks.txt
	HEADERS_ranks string = 		"*\\x1b[38;5;15mRanks*\r\n"


	//headers\username.txt
	HEADERS_username string = 		"*\\x1b[38;5;15mUsername*\r\n"
)

var ( //message

	//message\invalid_id.itl
	MESSAGE_invalid_id string = " Invalid ID has been given within the unit properly!\r\n"+
		"\r\n"


	//message\mismatch_id_username.itl
	MESSAGE_mismatch_id_username string = " No user exists within the params given!\r\n"+
		"\r\n"


	//message\syntax.itl
	MESSAGE_syntax string = " Syntax: sessions message <message> [user_tag]\r\n"+
		" [user_tag]: @root root@12345\r\n"+
		"\r\n"
)

var ( //values

	//values\connected.txt
	VALUES_connected string = 		"*\\x1b[38;5;15m<<$connected>>*\r\n"


	//values\id.txt
	VALUES_id string = 		"*\\x1b[38;5;15m<<$id>>*\r\n"


	//values\idle.txt
	VALUES_idle string = 		"*\\x1b[38;5;15m<<$idle>>*\r\n"


	//values\ip.txt
	VALUES_ip string = 		"*\\x1b[38;5;15m<<$ip>>*\r\n"


	//values\ranks.txt
	VALUES_ranks string = 		"*<<$ranks>>*\r\n"


	//values\username.txt
	VALUES_username string = 		"*\\x1b[38;5;15m<<$username>>*\r\n"
)