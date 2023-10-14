package models


//stores the configuration settings
type FakeSlaves struct { //stored in structure properly
	FakeSlaves map[string]*FakeConfig `toml:"fake_slaves"`
}

type FakeConfig struct {
	Act bool `toml:"enabled"`
	Max int `toml:"maximum"` //max duration
	Min int `toml:"minimum"` //min duration
	Equ int `toml:"equilibrium"` //middle ground
	Ticks int `toml:"ticker"` //divides by 4 for ticker properly
}