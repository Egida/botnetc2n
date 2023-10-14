package models 

type SuggestionMethod struct {
	Enabled bool 							`json:"enabled"`
	Methods map[string]*SuggestionModel		`json:"methods"`
}

type SuggestionModel struct {
	Provider     []string `json:"provider"`
	Asn          []int    `json:"asn"`
	Enabled      bool     `json:"enabled"`
	AccessBypass []string `json:"access_bypass"`
}