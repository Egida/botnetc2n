package template

import "unicode/utf8"

//properly tries to check for the engine inlet without issues
//this will try to validate the inlet without issues happening on request
func (t *TemplateEngine) takeEngineIN() bool {
	//stores how correct the object is
	//this will ensure we only accept 100 percent correct objects
	var tick int = 0 //set as 0 for default
	//ranges throughout the source from the position
	//this will make sure its properly done without issues happening
	for text, symbol := t.position, 0; text < utf8.RuneCountInString(t.lines) && symbol < utf8.RuneCountInString(t.enginePrefix[0]); text, symbol = text + 1, symbol + 1 {
		//compares the different information properly
		//this will ensure its properly sorted and its valid while being safe
		if t.lines[text] == t.enginePrefix[0][symbol] {
			tick++ //adds an additional object onto without issues
		}
	}
	//returns the information correctly
	//this will ensure its properly configured without errors
	return tick == len(t.enginePrefix)
}

//properly tries to check for the engine outlet without issues
//this will try to validate the outlet without issues happening on request
func (t *TemplateEngine) takeEngineOUT() bool {
	//stores how correct the object is
	//this will ensure we only accept 100 percent correct objects
	var tick int = 0 //set as 0 for default
	//ranges throughout the source from the position
	//this will make sure its properly done without issues happening
	for text, symbol := t.position, 0; text < utf8.RuneCountInString(t.lines) && symbol < utf8.RuneCountInString(t.enginePrefix[1]); text, symbol = text + 1, symbol + 1 {
		//compares the different information properly
		//this will ensure its properly sorted and its valid while being safe
		if t.lines[text] == t.enginePrefix[1][symbol] {
			tick++ //adds an additional object onto without issues
		}
	}
	//returns the information correctly
	//this will ensure its properly configured without errors
	return tick == len(t.enginePrefix)
}