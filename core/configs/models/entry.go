package models


type CatpchaToml struct {
	Enabled           bool     `toml:"enabled"`
	UserBypass        []string `toml:"user_bypass"`
	MinimumGeneration int      `toml:"minimum_generation"`
	MaximumGeneration int      `toml:"maximum_generation"`
	MaximumAttempts   int      `toml:"maximum_attempts"`
	Mfa               struct {
		Forced       bool     `toml:"forced"`
		BypassForced []string `toml:"bypass_forced"`
		App          string   `toml:"app"`
	} `toml:"mfa"`
}