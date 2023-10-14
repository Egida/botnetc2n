package models


//stores the spinner properly
//this will ensure its done without issues
type SpinnerConfig struct { //stored in type structure
	Spins map[string]*Spinner `toml:"spinners"`
}

//stores the configuration for spinners
type Spinner struct { //this will allow for better handling
	Launch bool     `toml:"launchITL"`
	Frames []string `toml:"frames"`
	Ticks  int      `toml:"ticks"`
}