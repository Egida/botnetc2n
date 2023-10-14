package ranks

//checks if the user can access the rank given
//this makes sure its properly done without issues happening
func (r *Ranks) CanAccess(rank string) bool {
	//ranges through all the ranks
	//this will ensure its done without errors
	for ro := range r.ranks { //compares the different names etc
	
		//checks inside the toml file properly
		//this will ensure the rank is active without issues
		if _, ok := PresetRanks[r.ranks[ro].RankName]; !ok {
			continue //continues looping as it wasnt found
		}

		if r.ranks[ro].RankName == rank && r.ranks[ro].RankStatus {
			return true //returns true as its active properly without isuses
		}
	}
	//this will point as they dont have it
	//this ensures its done without errors happening
	return false
}

//checks if the user can access the array of objects
//it takes one trigger to properly trigger the output properly
func (r *Ranks) CanAccessArray(ranks []string) bool {
	//ranges through all the ranks without issues
	//this will ensure its done without errors happening
	for pos := range ranks { //ranges through the ranks array
		if r.CanAccess(ranks[pos]) { //checks if the user can access it
			return true //they can access at this point properly without issues
		}
	}
	//returns false as they cant access properly
	//this will allow better handling without issues happening
	return false
}