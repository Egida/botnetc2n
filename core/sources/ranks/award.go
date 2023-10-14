package ranks

import (
	"errors"
)

var (
	ErrUserAlreadyHas error = errors.New("the user already has this rank properly/")
)

//properly tries to give the user the rank
//this ensures its done without issues happening
func (r *Ranks) GiveRank(name string) error {
	//checks if they have the rank properly
	//this will ensure its done without issues happening
	sock, pos := r.hasRank(name) //returns either position of role of false
	if sock { //the rank already exists properly on profile
		//checks if the user already has the rank
		//this will make sure its not ignored without issues
		if r.ranks[pos].RankStatus { //user already has rank
			return ErrUserAlreadyHas //already has properly
		}

		//sets the rank to active properly
		//this will make sure its not ignored without issues
		r.ranks[pos].RankStatus = true; return nil //and kills function
	}

	//gives the rank properly without issues
	//this will make sure its done without issues happening and kills the function
	r.ranks = append(r.ranks, UserRank{RankName: name, RankStatus: true}); return nil
}

//this will correctly check if the user has the rank
//makes sure its done without errors happening on request
func (r *Ranks) hasRank(a string) (bool, int) {

	//ranges through the ranks user has
	//this will ensure we know if it exits properly
	for pos := range r.ranks {

		//compares the current rank name and one given
		//makes sure its exists without issues happening
		if r.ranks[pos].RankName == a { //compares both rank names
			return true, pos //returns true as it exists properly
		}
	}

	return false, -1
}