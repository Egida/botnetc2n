package ranks



//this will allow the system to access structures
//makes sure its done properly without issues happening
type Ranks struct { //stored inside structures properly
	//stores the username properly first
	//this will make sure its done without errors
	username string //stored in type string
	//stores all their ranks properly
	//this will ensure its done without errors happening
	ranks []UserRank //stored in array of userRanks
}

//creates the rank structure
//this will ensure its done properly
func MakeRank(user string) *Ranks {
	return &Ranks{ //returns the structure properly
		username: user, //sets the username properly
		ranks: make([]UserRank, 0), //creates the array properly
	}
}

func (r *Ranks) Ranks() []UserRank {
	return r.ranks
}

//stores the rank settings
//this will ensure its done without issues
type UserRank struct { //stored in type structure
	//stores the Rank information properly
	//this will make sure its done without errors
	RankName string `json:"name"` //stored in type string
	//stores the rank status properly
	//this will ensure its done without errors
	RankStatus bool `json:"status"` //stored in type boolean
}

//stores more information about the user
//this will ensure it stores the array without issues happening
type UserRanked struct { //stored in type structure properly and safely
	//stores the username for this structure
	//this will make sure its done without issues happening
	Username string `json:"username"`
	//this will properly allow for better settings
	//this will store every single role the user has
	Ranks []UserRank `json:"ranks"`
}