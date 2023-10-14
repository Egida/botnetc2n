package branding

var ( //alerts

	//alerts\admin-false.itl
	ALERTS_admin_false string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're admin status has been revoked by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\admin-true.itl
	ALERTS_admin_true string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You have been promoted to admin by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\broadcast.itl
	ALERTS_broadcast string = "<?\r\n"+
		"	var user string = sender()\r\n"+
		"	var msg string = message()\r\n"+
		"?>\r\n"+
		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"broadcast from \"+user+\">\", \"255,255,0\", \"255,0,255\", \"0,255,255\")>> \\x1b[38;5;11m(\\x1b[4m<<$msg>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\concurrents-changed.itl
	ALERTS_concurrents_changed string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're concurrents field has been updated by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m) <<oldconcurrents()>> -> <<newconcurrents()>>\\x1b[0m\\x1b8\r\n"


	//alerts\cooldown-changed.itl
	ALERTS_cooldown_changed string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're cooldown field has been updated by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m) <<oldcooldown()>> -> <<newcooldown()>>\\x1b[0m\\x1b8\r\n"


	//alerts\maxtime-changed.itl
	ALERTS_maxtime_changed string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're maxtime field has been updated by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m) <<oldmaxtime()>> -> <<newmaxtime()>>\\x1b[0m\\x1b8\r\n"


	//alerts\moderator-false.itl
	ALERTS_moderator_false string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're moderator status has been revoked by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\moderator-true.itl
	ALERTS_moderator_true string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You have been promoted to moderator by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\password-changed.itl
	ALERTS_password_changed string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're password field has been updated by \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\session-message.itl
	ALERTS_session_message string = 		"\\x1b[0m\\x1b7\\x1b[1A\\r\\x1b[2K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"Message> \", \"255,255,0\", \"255,0,255\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<message()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\\x1b[0m\\x1b8\r\n"


	//alerts\user-removed.itl
	ALERTS_user_removed string = "\\r\\x1b[K\\x1b[38;5;15m[<<colour.marshal(\"System\", \"255,215,0\", \"255,127,127\")>>\\x1b[38;5;15m] <<colour.marshal(\"You're user has been removed by \", \"255,255,0\", \"0,255,255\")>>\\x1b[38;5;11m(\\x1b[4m<<promotor()>>\\x1b[0m\\x1b[38;5;11m)\\x1b[0m\r\n"+
		"<<sleep(8000)>>\r\n"
)