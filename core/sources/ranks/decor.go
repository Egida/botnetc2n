package ranks

import (
	"Nosviak2/core/sources/tools"
	"fmt"
	"sort"
	"strings"
)

//this will properly render each issue without issues
//this allows for better handling without issues happening on execution
func CreateSystem(ranks []string) ([]string) {
	var finish []string = make([]string, 0)

	//stores all possible ranks map properly
	//this will hold all keys gathered from array below
	var handled []string = make([]string, 0)

	//ranges through all possible ranks
	for ranks := range PresetRanks {
		handled = append(handled, ranks)
	}

	//sorts the ranks properly
	sort.Strings(handled)

	//ranges through all the systems properly
	//this will allow us to properly judge information without errors
	for _, storage := range handled { //ranges through properly
		//checks if he can properly access
		//this will make sure its done without issues happening
		if can, _ := tools.NeedleHaystack(ranks, storage); can { //sets the charaters properly without issues happening properly
			finish = append(finish, fmt.Sprintf("\x1b[0m\x1b[%sm\x1b[%sm %s \x1b[0m", strings.Join(Convert(PresetRanks[storage].MainColour), ";"), strings.Join(Convert(PresetRanks[storage].SecondColour), ";"), PresetRanks[storage].SignatureCharater))
		} //returns the values properly without issues happening 
	}
	
	return finish
}