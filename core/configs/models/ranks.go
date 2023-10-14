package models

import "Nosviak2/core/sources/ranks"

//this will store information about the ranks
//allows for better customizing without issues happening
type RanksToml struct { //stored in type structure correctly
	//stores the rank map properly
	//this will be used inside without issues happening
	Ranks map[string]ranks.RankSettings `toml:"ranks"`
}