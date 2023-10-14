package models


//stores the ip rewrite structure
//this will ensure its done without any errors
type IP_Rewrite struct { //stored in type structure
	Ranks     []string `toml:"ranks"`
	Rewritten string   `toml:"rewritten"`
}