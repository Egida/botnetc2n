package models


type AttackToml struct {
	Attacks struct {
		Prefix      string 	`toml:"prefix"`
		KVPrefix    string 	`toml:"kv_prefix"`
		Timeoutunit string 	`toml:"timeoutunit"`
		Timeout     int    	`toml:"timeout"`
		Useragent   string 	`toml:"useragent"`
		Groups      []Group `toml:"attack_groups"`
		Callback	bool	`toml:"attack_callback"`
		DNS			*DNS 	`toml:"dns"`
		
	} `toml:"attacks"`
}


type DNS struct {
	Enabled bool `toml:"enabled"`
	Routes []string `toml:"routes"`
}

type Group struct {
	Name string `toml:"group"`
	MaxTime int `toml:"maxtime"`
	Conns   int `toml:"conns"`
}

type Blacklists struct {
	Ips     []string `toml:"ips"`
	Domains []string `toml:"domains"`
}