package ranks

import (
	"errors"
)

var (
	ErrUserAlreadyCant error = errors.New("the user already doesnt have this rank properly")
)

// this will remove the rank from the users array
// makes sure they no longer have access to the permissions without issues
func (r *Ranks) RemoveRank(name string) error {
	//checks if the user can access
	//this will check inside the dups
	Has, Pos := r.hasRank(name) //checks if they can access the rank
	if !Has {                   //checks if they has the rank
		return ErrUserAlreadyCant //returns error
	}

	//compares the name properly
	//this will ensure its done without issues
	if r.ranks[Pos].RankName == name { //compares the names
		r.ranks[Pos].RankStatus = false //removes the status properly
	} else { //returns the error correctly and properly
		return ErrUserAlreadyCant
	}

	//returns nil as it worked properly
	//this will ensure its done without errors happening
	return nil
}
