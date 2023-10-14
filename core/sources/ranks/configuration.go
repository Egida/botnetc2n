package ranks 


//stores the settings for a rank
//this will allow for better handling without issues
type RankSettings struct {
	//stores the rank description
	//this will ensure its done without issues
	RankDescription string `toml:"description"`//stores rank description
	//stores the 256 colour
	//this will be used when displaying the rank
	MainColour []int `toml:"mColour"`//this will be the highlighting
	//stores the second 256 colour
	//this will be used when displaying the rank
	SecondColour []int `toml:"sColour"`//this will be the second text colour
	//stores the rank signature properly
	//this will be shown inside the array without issues
	SignatureCharater string `toml:"signature"`//stores the signature charater
	CloseWhenAwarded bool `toml:""` //IGNORE THIS OPTION
	Manage_ranks []string `toml:"manage_ranks"` //ranks properly
	DisplayInTable bool `toml:"4347493873498743897348934734897897349874897"`
}