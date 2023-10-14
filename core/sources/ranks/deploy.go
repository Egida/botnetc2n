package ranks

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//this will properly render each issue without issues
//this allows for better handling without issues happening on execution
func (r *Ranks) DeployRanks(haveSpace bool) ([]string, error) {
	var complete []string = make([]string, 0)

	
	//stores all possible ranks map properly
	//this will hold all keys gathered from array below
	var handled []string = make([]string, 0)

	//ranges through all possible ranks
	for ranks := range PresetRanks {
		handled = append(handled, ranks)
	}

	//sorts the ranks properly
	sort.Strings(handled)

	//ranges through every array given properly
	//this ensures its done without errors happening
	for _, project := range handled {
		//checks if the system can access without issues
		//makes sure its done better without issues happening
		if r.CanAccess(project) { //checks if we can access it without issues
			//saves into the array correctly and properly without issues happening
			complete = append(complete, fmt.Sprintf("\x1b[0m\x1b[%sm\x1b[%sm %s \x1b[0m", strings.Join(Convert(PresetRanks[project].MainColour), ";"), strings.Join(Convert(PresetRanks[project].SecondColour), ";"), PresetRanks[project].SignatureCharater)); continue
		}
	}

	if haveSpace {
		//ranges through the ranks missed
		//we will add a simple space here propetly
		for p := 0; p < len(PresetRanks)-len(complete); p++ {
			complete = append(complete, "\x1b[0m   \x1b[0m"); continue //continues the loop properly
		}
	}
	//returns the information properly
	//this will make sure its done without issues happening
	return complete, nil
}


func Convert(i []int) []string {
	var src []string = make([]string, 0)

	for pos := range i {
		src = append(src, strconv.Itoa(i[pos]))
	}

	return src
}