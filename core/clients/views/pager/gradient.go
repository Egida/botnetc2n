package pager

import (
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/tools/gradient"
	"strings"
	"unicode/utf8"
)

//this will properly execute the standard table
//ensures its done without issues happening on reqeust
func (mtr *MakeTableRender) GradientTable() ([]string, error) { //returns error
	var layers []string = strings.Split(mtr.table.String(), "\n")
	guide := GetLongestLineWithSTRIP(layers)
	
	for pos, text := range layers  {
		performed, err := gradient.NewWithIntArray(text, mtr.session.Colours...).WorkerWithEscapes(guide)
		if err != nil {
			return layers, err
		}

		layers[pos] = performed
	}

	//returns the output properly without issues happening
	//this makes the system safer without issues happening on request
	return layers, nil
}


// GetLongestLineWithSTRIP will strip each line and count the longest line
func GetLongestLineWithSTRIP(str []string) int {
	var current int = utf8.RuneCountInString(tools.Strip(str[0])) // sets default to 0

	for _, newLine := range str[1:] {
		if utf8.RuneCountInString(tools.Strip(newLine)) > current {
			current = utf8.RuneCountInString(tools.Strip(newLine))
		}
	}

	return current
}