package evaluator

import (
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/tools/gradient"

	//"Nosviak2/core/sources/tools"
	//"fmt"
	"regexp"
	"strings"
)

//WithinLine will parse the bodys inside
func WithinLine(line string) string {


	//ranges through all presets set within the configuration
	for key, preset_Fade := range json.GradientColour { //ranges

		if !preset_Fade.Enabled { //disabled
			continue //continues
		}

		//ranges through the bodys information within the system
		for _, text := range regexp.MustCompile(`<`+key+`>(.*)</`+key+`>`).FindAllStringSubmatch(line, -1) {
			if len(text) < 2 { //unknown length has been given
				continue //continues looping
			}

			//performs the gradient on the objects being passed
			ls, err := gradient.NewWithIntArray(text[1], ToRGB(preset_Fade.Colours, make([][]int, 0))...).WorkerWithEscapes()
			if err != nil {
				continue
			}

			//replaces the objective item once within the system
			line = strings.Replace(line, text[0], ls, 1)
		}
	}


	return line //returns the line
}

//converts all to valid colours properly
func ToRGB(C []*models.RGB, src [][]int) [][]int {
	for _, m := range C { //ranges through
		src = append(src, []int{m.Red,m.Green,m.Blue})
	}; return src
}