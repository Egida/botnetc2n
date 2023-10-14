package models 

type Plans struct {
	Plans map[string]*Plan `toml:"plans"`
}

type Plan struct {
	Newuser            bool     `toml:"newuser"`
	Maxtime            int      `toml:"maxtime"`
	Cooldown           int      `toml:"cooldown"`
	Concurrents        int      `toml:"concurrents"`
	MaxSessions        int      `toml:"max_sessions"`
	AccessArrangements []string `toml:"access_arrangements"`
	DefaultTheme       string   `toml:"default_theme"`
	PlanLength         int      `toml:"plan_length"`
	Description        string   `toml:"description"`
	MaxSlaves		   int 		`toml:"max_slaves"`
}