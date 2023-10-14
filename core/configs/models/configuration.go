package models 



//used inside the config.json dataset
//this will allow the options ]o have an affect inside the program
type ConfigurationJson struct {
	//holds information about the database
	//helps with authentication and generally configuring the attempts
	Database struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
	} `json:"database"`
	Masters struct {
		Server struct {
			Protocol    string `json:"protocol"`
			Listener    string `json:"listener"`
			DynamicAuth bool   `json:"dynamicAuth"`
		} `json:"server"`
		MaxAuthAttempts    int    `json:"maxAuthAttempts"`
		ServerKey          string `json:"serverKey"`
		BeforePasswdPrompt struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		} `json:"beforePasswdPrompt"`
		Accounts struct {
			PasswordLength int `json:"passwordLength"`
			MaxSessions    int `json:"maxSessions"`
			DaysExpiry     int `json:"daysExpiry"`
		} `json:"accounts"`
	} `json:"masters"`
}