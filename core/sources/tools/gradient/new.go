package gradient

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Gradient struct {
	Colours []Colour // Stores all colours for the gradient target
	Text    string   // Stores all the Text for the gradient target
}

type Colour struct {
	Red   int // Red value
	Green int // Green value
	Blue  int // Blue value
}

// New will create the new gradient object
func New(text string, colours ...Colour) *Gradient {
	return &Gradient{
		Text:    text,    // Gradient text
		Colours: colours, // Colours for gradient
	}
}

func NewWithIntArray(text string, colours ...[]int) *Gradient {
	var new *Gradient = new(Gradient)
	new.Colours = make([]Colour, 0)
	new.Text = text

	for _, colour := range colours {
		new.Colours = append(new.Colours, *int_array_to_rgb(colour))
	}

	return new
}

// Worker will just create the gradient object without any ansi attributes
func (G *Gradient) Worker(setLen ...int) (string, error) {
	var output []string = make([]string, 0)

	var length int = utf8.RuneCountInString(G.Text)
	if setLen[0] > 0 {
		length = setLen[0]
	}

	Red, Green, Blue := G.gradient(length) // Produces the curve parameters
	for pos := 0; pos < length; pos++ { // Loops through all the different selections inside the text
		RC, GC, BC := strconv.Itoa(int(Red[pos])), strconv.Itoa(int(Green[pos])), strconv.Itoa(int(Blue[pos])) // Converts colour types
		output = append(output,  fmt.Sprintf("\x1b[38;2;%s;%s;%sm%s\x1b[0m", RC, GC, BC, strings.Split(G.Text, "")[pos])) // Saves into the array
	}

	// Returns the objects given
	return strings.Join(output, ""), nil
}

type escape struct {
	charater string		// Stores the charater
	escape   bool	 	// Stores is there is an escape
	esc      string     // Uses the escape colour used
}

// WorkerWithEscapes will work the gradient as normal but use ansi escapes sequences
func (G *Gradient) WorkerWithEscapes(setLen ...int) (string, error) {
	var output []string = make([]string, 0)		// Stores the text output from the string
	var system []escape = make([]escape, 0)		// Stores all the charaters and if they are escapes
	

	G.Text = strings.ReplaceAll(G.Text, "\x1b", "\\x1b") 	// replaces all the strings
	var current bool = false								// Stores if gradient

	// Loops through the text given ensures its done through all charaters
	for position := 0; position < utf8.RuneCountInString(G.Text); position++ {

		// Detects the ansi escape sequence
		if strings.Split(G.Text, "")[position] == "\\" {
			var captured_case string = ""
			for attempts := position; attempts < utf8.RuneCountInString(G.Text); attempts++ {
				position = attempts // Sets the position attempt
				

				captured_case += strings.Split(G.Text, "")[attempts] // Addons onto the array given
				if strings.Split(G.Text, "")[attempts - 1] == "m" && strings.Split(G.Text, "")[attempts] != "\\" {
					break
				}
			}

	

			// Checks for the reset ansi escape
			if strings.Contains(strings.Join(strings.Split(captured_case, "")[:len(strings.Split(captured_case, ""))-1], ""), "\\x1b[0m") {
				current = !current

				if !current {
					system = append(system, escape{charater: captured_case, escape: true})
					continue
				}
			}
		
			system = append(system, escape{charater: captured_case, escape: current})

			continue
		}

		system = append(system, escape{charater: strings.Split(G.Text, "")[position], escape: current})
	}

	var length int = len(system)
	if len(setLen) > 0 &&setLen[0] > 0 {
		length = setLen[0]
	}

	RED, GREEN, BLUE := G.gradient(length)				// Performs the gradient
	for position := range make([]string, len(system)) {		// Loops through the different object
		current := system[position]

		if current.escape {
			output = append(output, fmt.Sprintf("%s%s", strings.ReplaceAll(current.esc, "\\x1b", "\x1b"), strings.ReplaceAll(current.charater, "\\x1b", "\x1b")))
			continue
		}


		Red, Green, Blue := strconv.Itoa(int(RED[position])), strconv.Itoa(int(GREEN[position])), strconv.Itoa(int(BLUE[position])) // Converts colour types

		output = append(output, fmt.Sprintf("\x1b[38;2;%s;%s;%sm%s\x1b[0m", Red, Green, Blue, current.charater))
	}

	// Returns the objects given
	return strings.Join(output, ""), nil
}

// int_array_to_rgb converts an array of ints into an array
func int_array_to_rgb(i []int) *Colour {
	if len(i) != 3 {
		return &Colour{0,0,0}
	}
	return &Colour{i[0], i[1], i[2]}
}