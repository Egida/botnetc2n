package models



//stores the theme information properly 
//this will ensure its done without any errors
type ThemeToml struct { //stored in structure
	ThemeChanger []string `toml:"theme_changer"`
	Theme map[string]*ThemeConfig `toml:"theme"`
}

//stores the configuration for the theme
//this will ensure its done without any errors
type ThemeConfig struct { //stored in structure
	Enabled     bool     `toml:"enabled"`
	Hidden      bool     `toml:"hidden"`
	Description string   `toml:"description"`
	UpdateSesss bool     `toml:"update_all_sessions"`
	Permissions []string `toml:"permissions"`
	Branding    string   `toml:"branding"`
	Decor       *ThemeDecorationConfig `toml:"decoration"`
}

//stores the custom colour configuration if wanted
//this will want a more active information statement
type ThemeDecorationConfig struct { //stored in type structure
	Colours [][]int `toml:"gradient"`
}