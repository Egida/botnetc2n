package attacks

import (
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/tools"
	"strconv"
	"strings"
)

//MatchSuggestion will suggest the closest model inside the array
func MatchSuggestion(isp string, asn string, method string) (string, *models.SuggestionModel) {


	//ranges through all possible entrys within the system
	for entry, information := range json.Suggestions.Methods {

		//ignores the current method
		if entry == method || !information.Enabled {
			continue
		}

		//ISP suggestion has been found here and we will prompt
		if tools.NeedleHaystackContains(information.Provider, isp) {
			return entry, information
		}

		//converts properly and safely
		convert, err := strconv.Atoi(strings.ReplaceAll(strings.Split(asn, " ")[0], "AS", ""))
		if err != nil {
			continue
		}

		//asn detection inside the string aswell properly
		if tools.NeedleHaystackOne(Quick(information.Asn, make([]string, 0)), string(convert)) {
			return entry, information
		}

	} 

	return "", nil
}

func Quick(i []int, src []string) []string {
	for _, c := range i {
		src = append(src, string(c))
	}; return src
}