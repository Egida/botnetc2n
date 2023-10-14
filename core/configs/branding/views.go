package branding



var ( //redeem

	//redeem\inputs.ini
	REDEEM_inputs string = "# controls the main input menu\r\n"+
		"# this allows for proper control\r\n"+
		"max_token_input=20\r\n"+
		"maskingCharater=\"●\"\r\n"+
		"\r\n"+
		"# username input\r\n"+
		"username_max_input=20\r\n"+
		"username_maskCharater=\"\"\r\n"+
		"\r\n"+
		"# password input\r\n"+
		"password_max_input=20\r\n"+
		"password_maskCharater=\"●\"\r\n"+
		"\r\n"+
		"\r\n"


	//redeem\invalidToken.tfx
	REDEEM_invalidToken string = 		"\\033[12;0f\\x1b[0m                               > invalid Token has been given\r\n"


	//redeem\password.tfx
	REDEEM_password string = 		"\\033[14;0f\\x1b[0m\\x1b[38;5;15m                               Password: \\x1b[30m\\x1b[47m\r\n"


	//redeem\redeem.tfx
	REDEEM_redeem string = "\\033c\r\n"+
		"                  Enter the token you wish to redeem below!\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"                               \\x1b[38;5;15mToken: \\x1b[30m\\x1b[47m                     \\x1b[0m\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\\x1b[0m <<$cnc>> 2022                                                             <<$version>>\\x1b[0m \r\n"


	//redeem\token.tfx
	REDEEM_token string = 		"\\033[12;0f\\x1b[0m\\x1b[38;5;15m                               Token: \\x1b[30m\\x1b[47m\r\n"


	//redeem\user-dup.tfx
	REDEEM_user_dup string = 		"\\033[16;0f\\x1b[0m\\x1b[38;5;15m                               > that username already exists on an account\r\n"


	//redeem\user-error.tfx
	REDEEM_user_error string = 		"\\033[16;0f\\x1b[0m\\x1b[38;5;15m                               > Fault within internal workers happened\r\n"


	//redeem\user-success.tfx
	REDEEM_user_success string = 		"\\033[16;0f\\x1b[0m\\x1b[38;5;15m                               > Correctly inserted the user without issues\r\n"


	//redeem\user.tfx
	REDEEM_user string = "\\033c\\x1b[38;5;205m\r\n"+
		"      __/\\\\\\\\\\_____/\\\\\\__/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\__/\\\\\\______________/\\\\\\_        \r\n"+
		"       _\\/\\\\\\\\\\\\___\\/\\\\\\_\\/\\\\\\///////////__\\/\\\\\\_____________\\/\\\\\\_       \r\n"+
		"        _\\/\\\\\\/\\\\\\__\\/\\\\\\_\\/\\\\\\_____________\\/\\\\\\_____________\\/\\\\\\_      \r\n"+
		"         _\\/\\\\\\//\\\\\\_\\/\\\\\\_\\/\\\\\\\\\\\\\\\\\\\\\\_____\\//\\\\\\____/\\\\\\____/\\\\\\__     \r\n"+
		"          _\\/\\\\\\\\//\\\\\\\\/\\\\\\_\\/\\\\\\///////_______\\//\\\\\\__/\\\\\\\\\\__/\\\\\\___    \r\n"+
		"           _\\/\\\\\\_\\//\\\\\\/\\\\\\_\\/\\\\\\_______________\\//\\\\\\/\\\\\\/\\\\\\/\\\\\\____   \r\n"+
		"            _\\/\\\\\\__\\//\\\\\\\\\\\\_\\/\\\\\\________________\\//\\\\\\\\\\\\//\\\\\\\\\\_____  \r\n"+
		"             _\\/\\\\\\___\\//\\\\\\\\\\_\\/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\_____\\//\\\\\\__\\//\\\\\\______ \r\n"+
		"              _\\///_____\\/////__\\///////////////_______\\///____\\///_______\r\n"+
		"              \\x1b[0m\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"                               \\x1b[38;5;15mUsername: \\x1b[30m\\x1b[47m                     \\x1b[0m\r\n"+
		"\r\n"+
		"                               \\x1b[38;5;15mPassword: \\x1b[30m\\x1b[47m                     \\x1b[0m\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\\x1b[0m <<$cnc>> 2022                                                             <<$version>>\\x1b[0m \r\n"


	//redeem\username.tfx
	REDEEM_username string = 		"\\033[12;0f\\x1b[0m\\x1b[38;5;15m                               Username: \\x1b[30m\\x1b[47m\r\n"
)

var ( //newuser

	//newuser\banner.itl
	NEWUSER_banner string = "<<clear()>>\\033[38;5;105m\r\n"+
		"      __/\\\\\\\\\\_____/\\\\\\__/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\__/\\\\\\______________/\\\\\\_        \r\n"+
		"       _\\/\\\\\\\\\\\\___\\/\\\\\\_\\/\\\\\\///////////__\\/\\\\\\_____________\\/\\\\\\_       \r\n"+
		"        _\\/\\\\\\/\\\\\\__\\/\\\\\\_\\/\\\\\\_____________\\/\\\\\\_____________\\/\\\\\\_      \r\n"+
		"         _\\/\\\\\\//\\\\\\_\\/\\\\\\_\\/\\\\\\\\\\\\\\\\\\\\\\_____\\//\\\\\\____/\\\\\\____/\\\\\\__     \r\n"+
		"          _\\/\\\\\\\\//\\\\\\\\/\\\\\\_\\/\\\\\\///////_______\\//\\\\\\__/\\\\\\\\\\__/\\\\\\___    \r\n"+
		"           _\\/\\\\\\_\\//\\\\\\/\\\\\\_\\/\\\\\\_______________\\//\\\\\\/\\\\\\/\\\\\\/\\\\\\____   \r\n"+
		"            _\\/\\\\\\__\\//\\\\\\\\\\\\_\\/\\\\\\________________\\//\\\\\\\\\\\\//\\\\\\\\\\_____  \r\n"+
		"             _\\/\\\\\\___\\//\\\\\\\\\\_\\/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\_____\\//\\\\\\__\\//\\\\\\______ \r\n"+
		"              _\\///_____\\/////__\\///////////////_______\\///____\\///_______\r\n"+
		"	      \\033[38;5;15m<<time.unix(time.now(), \"Mon 2 Jan 15:04:05\")>>\\033[0m\r\n"+
		"\r\n"+
		"\r\n"+
		" \r\n"+
		" \r\n"+
		"                   \\x1b[38;5;15mPassword: \\x1b[30m\\x1b[47m                             \\x1b[0m\r\n"+
		" \r\n"+
		"                   \\x1b[38;5;15mConfirm Password: \\x1b[30m\\x1b[47m                     \\x1b[0m\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"


	//newuser\confirm_password.tfx
	NEWUSER_confirm_password string = 		"\\033[16;0f\\x1b[0m\\x1b[38;5;15m                   Confirm Password: \\x1b[30m\\x1b[47m\r\n"


	//newuser\password-error.itl
	NEWUSER_password_error string = "\\x1b[0m \r\n"+
		"                   > Password has failed to be updated properly\r\n"+
		"<<sleep(10000)>>\r\n"


	//newuser\password.tfx
	NEWUSER_password string = 		"\\033[14;0f\\x1b[0m\\x1b[38;5;15m                           Password: \\x1b[30m\\x1b[47m\r\n"


	//newuser\same-oldpassword.itl
	NEWUSER_same_oldpassword string = "\\x1b[0m \r\n"+
		"                   > Password given matchs your old password!\r\n"+
		"<<sleep(10000)>>\r\n"
)

var ( //mfa

	//mfa\invalid_code.itl
	MFA_invalid_code string = "You have given an invalid MFA Code!\r\n"+
		"<<sleep(10000)>>\r\n"


	//mfa\mfa-splash.itl
	MFA_mfa_splash string = "<<clear()>>\r\n"+
		"<<padding.centre(\"Enter your MFA Code!\", user.length())>>\r\n"+
		" \r\n"+
		" \r\n"+
		"\r\n"


	//mfa\prompt.itl
	MFA_prompt string = 		"\\x1b[0mMFA> \r\n"
)

var ( //login

	//login\header.tfx
	LOGIN_header string = "\\033c\r\n"+
		"\\033[38;5;15m   __/\\\\\\____________________________________________________________        \r\n"+
		"    _\\/\\\\\\____________________________________________________________       \r\n"+
		"     _\\/\\\\\\_____________________________/\\\\\\\\\\\\\\\\___/\\\\\\_______________      \r\n"+
		"      _\\/\\\\\\_________________/\\\\\\\\\\_____/\\\\\\////\\\\\\_\\///___/\\\\/\\\\\\\\\\\\___     \r\n"+
		"       _\\/\\\\\\_______________/\\\\\\///\\\\\\__\\//\\\\\\\\\\\\\\\\\\__/\\\\\\_\\/\\\\\\////\\\\\\__    \r\n"+
		"        _\\/\\\\\\______________/\\\\\\__\\//\\\\\\__\\///////\\\\\\_\\/\\\\\\_\\/\\\\\\__\\//\\\\\\_   \r\n"+
		"         _\\/\\\\\\_____________\\//\\\\\\__/\\\\\\___/\\\\_____\\\\\\_\\/\\\\\\_\\/\\\\\\___\\/\\\\\\_  \r\n"+
		"          _\\/\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\__\\///\\\\\\\\\\/___\\//\\\\\\\\\\\\\\\\__\\/\\\\\\_\\/\\\\\\___\\/\\\\\\_ \r\n"+
		"           _\\///////////////_____\\/////______\\////////___\\///__\\///____\\///__\r\n"+
		"           \\033[0m<<date()>>\r\n"+
		"\r\n"+
		"                             \\x1b[4m                             \\x1b[0m\r\n"+
		"                   \\x1b[38;5;15mUsername> \\x1b[38;5;15m\\x1b[48;5;235m\\x1b[4m                             \\x1b[0m\r\n"+
		"                             \\x1b[4m                             \\x1b[0m\r\n"+
		"                   \\x1b[38;5;15mPassword> \\x1b[38;5;15m\\x1b[48;5;235m\\x1b[4m                             \\x1b[0m\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\r\n"+
		"\\x1b[48;5;15m\\x1b[38;5;16m <<$cnc>> 2022                                                           <<$version>> \\x1b[0m\r\n"


	//login\invalid-password.tfx
	LOGIN_invalid_password string = "\\033[18;0f\\x1b[0m                               > invalid password has been given\r\n"+
		"                               > please try again in 3 seconds\r\n"


	//login\invalid-username.tfx
	LOGIN_invalid_username string = "\\033[13;0f\\x1b[0m                   \\x1b[48;5;235m\\x1b[38;5;9m Invalid username [\\x1b[38;5;11mALERT\\x1b[38;5;9m]               \\x1b[0m\r\n"+
		"\\033[14;0f                   \\x1b[48;5;235m\\x1b[38;5;15m the username which you have entered has\\x1b[0m\r\n"+
		"\\033[15;0f                   \\x1b[48;5;235m\\x1b[38;5;15m been classified as invalid therefor you\\x1b[0m\r\n"+
		"\\033[16;0f                   \\x1b[48;5;235m\\x1b[38;5;15m can't login to this account right now! \\x1b[0m\r\n"


	//login\password.tfx
	LOGIN_password string = 		"\\033[14;60f\\x1b[0m\\x1b[38;5;2m●\\x1b[0m\\033[16;0f\\x1b[0m                   \\x1b[38;5;15mPassword> \\x1b[38;5;15m\\x1b[48;5;235m\\x1b[4m\r\n"


	//login\too-many-attempts.tfx
	LOGIN_too_many_attempts string = "<<clear()>>\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[s\r\n"+
		"\\x1b[13A\\x1b[48;5;9m\\x1b[38;5;11m                                Access Denied\r\n"+
		"\\x1b[48;5;9m\\x1b[38;5;16m        You have reached the max amount of login attempts per session\\x1b[u\\x1b[0m\r\n"


	//login\username.tfx
	LOGIN_username string = 		"\\033[14;0f                   \\x1b[38;5;15mUsername> \\x1b[38;5;15m\\x1b[48;5;235m\\x1b[4m\r\n"
)

var ( //force_mfa

	//force_mfa\invalid_otp.itl
	FORCE_MFA_invalid_otp string = "invalid OTP system has been given properly!\r\n"+
		"<<sleep(10000)>>\r\n"


	//force_mfa\mfa_displayed.itl
	FORCE_MFA_mfa_displayed string = "<<clear()>>\r\n"+
		"<<qrcode()>>\r\n"+
		"Please enter this secret: <<secret()>>\r\n"+
		"On either Authy or Google\r\n"


	//force_mfa\prompt.itl
	FORCE_MFA_prompt string = 		"\\x1b[0mMFA>\r\n"


	//force_mfa\resize_screen.itl
	FORCE_MFA_resize_screen string = "<<clear()>>\r\n"+
		"	Hey <<user.username()>>, Resize your screen to fix the MFA Code!\r\n"+
		" \r\n"+
		" \r\n"
)

var ( //expired

	//expired\plan_expired.itl
	EXPIRED_plan_expired string = "<<clear()>>\r\n"+
		"<?\r\n"+
		"	clear()\r\n"+
		"	var app string = cnc()\r\n"+
		"?>\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[0m\r\n"+
		"\\x1b[48;5;9m                                                                                \\x1b[s\r\n"+
		"\\x1b[13A\\x1b[48;5;9m\\x1b[38;5;11m                                Access Denied\r\n"+
		"\\x1b[48;5;9m\\x1b[38;5;16m<<padding.centre(\"You're account on \"+app+\" has expired!\", user.length())>>\r\n"+
		"\\x1b[u\\x1b[0m<<sleep(10000)>>\r\n"


	//expired\title.tfx
	EXPIRED_title string = 		"Incorrect Expiry!\r\n"
)

var ( //captcha

	//captcha\banner.itl
	CAPTCHA_banner string = "\\x1b[0m<?\r\n"+
		"	clear()\r\n"+
		"	const cnc string = cnc()\r\n"+
		"	const question1 = QueryOne()\r\n"+
		"	const question2 = QueryTwo()\r\n"+
		"?>\r\n"+
		" \r\n"+
		" \r\n"+
		"<<padding.centre(\"Hello there!\", 80)>>\r\n"+
		"<<padding.centre(\"complete the mathmatical question below\", 80)>>\r\n"+
		"<<padding.centre(\"to gain access to \"+cnc+\"!\", 80)>>\r\n"+
		" \r\n"+
		"<<padding.centre(\"Question: \"+question1+\" + \"+question2, 80)>>\r\n"+
		" \r\n"+
		" \r\n"+
		"\r\n"


	//captcha\incorrect.itl
	CAPTCHA_incorrect string = "Incorrect answer given! you have <<attempts()>> left...\r\n"+
		"<<sleep(1000)>>\r\n"


	//captcha\incorrect_max.itl
	CAPTCHA_incorrect_max string = "Incorrect answer given! you have 0 left...\r\n"+
		"<<sleep(10000)>>\r\n"+
		"Goodbye!\r\n"+
		"<<sleep(10000)>>\r\n"


	//captcha\prompt.itl
	CAPTCHA_prompt string = 		"\\x1b[0m \\x1b[48;5;234m\\x1b[38;5;9m <<user.username()>> ● Captcha@<<cnc()>> \\x1b[0m ►► \\x1b[97m\r\n"


	//captcha\title.itl
	CAPTCHA_title string = 		"\\033]0; Hey <<user.username()>>! complete the mathmatical question presented below to gain access\\007\r\n"
)

var ( //banned

	//banned\user-banned.itl
	BANNED_user_banned string = "Your account has been banned from <<cnc()>>!\r\n"+
		"<<sleep(10000)>>\r\n"
)