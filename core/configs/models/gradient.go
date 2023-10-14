package models 

type GradientBodys struct {
	Enabled bool `json:"enabled"`
	Colours []*RGB `json:"colours"`
}

type RGB struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}