package models

//stores the information properly
//this will ensure its done without any errors
type BinCommand struct { //stored in structure properly
	Name        string   `toml:"name"`
	Description string   `toml:"description"`
	Runtime     string   `toml:"runtime"`
	Args        []string `toml:"args"`
	MinArgs     int      `toml:"min_args"`
	Timeout     int      `toml:"timeout"`
	Env         []string `toml:"env"`
	Access      []string `toml:"access"`
}