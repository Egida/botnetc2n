package models 


type DecorationToml struct {
	Gradient struct {
		Status            bool    `toml:"status"`
		Colours           [][]int `toml:"colours"`
		EnableWithCredits bool    `toml:"enable_with_credits"`
		Table             struct {
			Tables          []string `toml:"tables"`
			TypeForcedStyle int      `toml:"type_forced_style"`
		} `toml:"table"`
	} `toml:"gradient"`
}