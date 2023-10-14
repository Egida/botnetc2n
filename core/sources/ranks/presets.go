package ranks





//stores all the different ranks which we made
//these will include banner, admin, reseller, mod etc
var PresetRanks map[string]RankSettings = map[string]RankSettings{
	"admin" : {
		RankDescription: "highest level of permissions",
		MainColour: []int{48,5,10}, //grey
		SecondColour: []int{38,5,16}, //black
		SignatureCharater: "A",
		CloseWhenAwarded: false,
		DisplayInTable: false,
		Manage_ranks: []string{"admin"},
	},
	"moderator" : {
		RankDescription: "manage all users but admin",
		MainColour: []int{48,5,11}, // blue
		SecondColour: []int{38,5,16},
		SignatureCharater: "M",
		CloseWhenAwarded: false,
		DisplayInTable: false,
		Manage_ranks: []string{"admin", "moderator"},
	},
	"reseller" : {
		RankDescription: "can resell accounts",
		MainColour: []int{48,5,105}, // yellow
		SecondColour: []int{38,5,16},
		SignatureCharater: "R",
		CloseWhenAwarded: false,
		DisplayInTable: true,
		Manage_ranks: []string{"admin", "moderator"},
	},

	"banned" : {
		RankDescription: "is a banned user",
		MainColour: []int{48,5,9},
		SecondColour: []int{38,5,16},
		SignatureCharater: "B",
		CloseWhenAwarded: true,
		DisplayInTable: true,
		Manage_ranks: []string{"admin"},
	},

	"bypass-bl" : {
		RankDescription: "can bypass attack blacklist",
		MainColour: []int{48,2,209, 88, 207},
		SecondColour: []int{38,5,16},
		SignatureCharater: "BL",
		CloseWhenAwarded: false,
		DisplayInTable: true,
		Manage_ranks: []string{"admin", "moderator"},
	},

	"bypass_ps" : {
		RankDescription: "can bypass powersaving",
		MainColour: []int{48,2,111, 217, 100},
		SecondColour: []int{38,5,16},
		SignatureCharater: "PS",
		CloseWhenAwarded: false,
		DisplayInTable: true,
		Manage_ranks: []string{"admin", "moderator"},
	},

	"api" : {
		RankDescription: "has api access",
		MainColour: []int{48,2,47, 47, 54},
		SecondColour: []int{38,5,15},
		SignatureCharater: "A",
		CloseWhenAwarded: false,
		DisplayInTable: true,
		Manage_ranks: []string{"admin", "moderator", "reseller"},
	},
}