package models


type ApiTomlModel struct {
	API struct {
		Enabled 	bool   `toml:"enabled"`
		Host    	string `toml:"host"`
		KeyLen  	int    `toml:"default_keylen"`
		DType   	string `toml:"default_type"`
		Path		string `toml:"path"`
		BlockRoute []string `toml:"blocked_routes"`
	} `toml:"api"`
	TLS struct {
		TLS           bool   `toml:"tls"`
		Certification string `toml:"certification"`
		Key           string `toml:"key"`
	} `toml:"tls"`
}