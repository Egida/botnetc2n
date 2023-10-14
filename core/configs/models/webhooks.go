package models


type WebhookToml struct {
	Webhooks struct {
		Enabled 	  bool     				   `toml:"enabled"`
		Token   	  string   				   `toml:"token"`
		Trigger 	  []string 				   `toml:"trigger"`
		Timeout 	  int      				   `toml:"timeout"`
		Colour  	  int      				   `toml:"colour"`		
	} `toml:"webhooks"`
	CustomConfigs map[string]*CustomWebhook `toml:"customConfigs"`
}


type CustomWebhook struct {
	Colour int `toml:"custom_colour"`
	Title  string `toml:"title"`
}