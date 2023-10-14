package pager

import (
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/views"
	"strings"
	"unicode/utf8"

	"github.com/alexeyco/simpletable"
	tml "github.com/naoina/toml"
)

//stores the information properly
//this will ensure its done without errors happening
type TableConfiguration struct { //allows configuring properly
	Style int `toml:"style"` //stores the ideal table style properly
	NewLine bool `toml:"newline"`
}

//properly parses the information without issues
//this will allow for custom styles without errors happening
func (mtr *MakeTableRender) GetQuery() (*TableConfiguration, error) {
	//properly tries to load the file
	//this will use the branding worker without issues
	subject := views.GetView("tables", mtr.header+".ini")
	var texture TableConfiguration //stores the statement
	//this will properly parse the information without issues
	//this ensures that its done without errors happening on request
	if err := tml.NewDecoder(strings.NewReader(subject.Containing)).Decode(&texture); err != nil {
		return nil, err //returns the error correctly and properly
	}


	//returns the values properly
	//this will ensure its done without errors happening
	return &texture, nil //returns the error properly
}

//this will allow for better content control within the system
//better system controlling without issues happening on request etc
func (mtr *MakeTableRender) TypeControl(t int, has bool) *simpletable.Style {
	if !has && toml.DecorationToml.Gradient.Table.TypeForcedStyle != -1 {
		return mtr.TypeControl(toml.DecorationToml.Gradient.Table.TypeForcedStyle, true)
	}

	switch t { //each numbers will equal a type properly
	case 1, 0: //compact
		return simpletable.StyleCompact
	case 2: //compact lite
		return simpletable.StyleCompactLite
	case 3: //compact classic
		return simpletable.StyleCompactClassic
	case 4: //default
		return simpletable.StyleDefault
	case 5: //markdown
		return simpletable.StyleMarkdown
	case 6: //rounded
		return simpletable.StyleRounded
	default: //unicode/ansi
		return simpletable.StyleUnicode
	}
}

//checks if the table wants gradients properly
//this will allow for better control without issues happening

func (mtr *MakeTableRender) WantGradient() bool {

	if !toml.DecorationToml.Gradient.Status {
		return false
	}
	//ranges through all the gradient tables
	//we will personally check for the table inside the array
	for _, current := range toml.DecorationToml.Gradient.Table.Tables {
		//checks if they are the same and match properly
		//this will now follow the routepath as an true statement
		if strings.ToLower(current) == mtr.header {
			return true //returns true correctly and properly
		}
	}; return false //returns false properly and securely
}

//gets the longest terminal line
//this will ensure its done without any issues
func (mtr *MakeTableRender) GetLongest(dst int) int {
	//ranges through all the lines
	//this will allow us to check each line
	for _, l := range strings.Split(mtr.table.String(), "\n") {
		if utf8.RuneCountInString(l) > dst {
			dst = utf8.RuneCountInString(l)
		}
	}; return dst
}