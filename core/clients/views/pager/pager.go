package pager

import (
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools/gradient"
	"fmt"

	//"fmt"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// renders the table with proper information
// allows for better handling without issues happening
func (mtr *MakeTableRender) Pager(screenshot string) error {

	var (
		//stores our system properly and safely
		//this will ensure its done without issues happening
		system []string = strings.Split(mtr.table.String(), "\n")

		//this will ensure its done safely
		position int = 1 //posiiton holder properly

		searched int = 0 //limited to 7
	)

	//checks for gradient enabled
	//this will make sure its done
	if mtr.WantGradient() { //gradient table
		var lines []string = make([]string, 0)
		guide := GetLongestLineWithSTRIP(strings.Split(mtr.table.String(), "\n"))

		//ranges through the table lines properly
		//this will enable gradient on each selected
		for _, col := range strings.Split(mtr.table.String(), "\n") {

			liner, err := gradient.NewWithIntArray(col, mtr.session.Colours...).WorkerWithEscapes(guide)
			if err != nil {
				return err
			}

			//saves into properly
			lines = append(lines, liner)
		}

		system = lines
	} else { //TEMP: update later with ansi code support
		system = strings.Split(strings.ReplaceAll(mtr.table.String(), "*", " "), "\n")
	}

	system = append([]string{"\x1b[47m\x1b[30m" + Centre("use W,S,UP,DOWN,SCROLL to navigate", mtr.session.Length, "") + "\x1b[0m"}, system...)
	system = append(system, "\x1b[47m\x1b[30m"+Centre(toml.ConfigurationToml.AppSettings.AppName+" - "+deployment.Version, mtr.session.Length, "")+"\x1b[0m")

	captured := mtr.session.Capture() //before pager enabled

	//cursor blink disabled properly
	//this will ensure its ignore within the system
	mtr.session.Channel.Write([]byte("\033[?25h\033[?0c\033c"))

	//renders everything the current pager is showing
	//this will ensure its done properly and safely within it
	//for proc := 0; proc < mtr.session.Height; proc++ {
	//	mtr.session.Write(system[proc]) //renders properly
	//	fmt.Println(system[proc])
	//}

	//displays the current chunk needed
	//this will display the first amount of commands
	mtr.session.Write(strings.Join(DisplayChunk(system, 0, mtr.session.Height), "\r\n"))

	//moves cursor to position one properly
	//this will make sure its done without issues
	mtr.session.Write("\033[0;0f") //resets cursor position

	for {
		//enables the cursor feedback information
		//allows us to detect cursor position within it
		mtr.session.Channel.Write([]byte("\033[?1000h\033[?25l"))

		buffer := make([]byte, 1)
		//reads from the channel properly
		//this will ensure its done properly
		if _, err := mtr.session.Channel.Read(buffer); err != nil {
			return err
		}

		switch buffer[0] {

		case 27: //arrow keys and cursor

			buf := make([]byte, 5)
			if _, err := mtr.session.Channel.Read(buf); err != nil {
				return err
			}

			if buf[1] == 77 { //cursor support here
				if buf[2] == 96 { //up
					position = mtr.UP(position, system, 1)
					continue
				} else if buf[2] == 97 { //down
					position = mtr.DOWN(position, system, 1)
					continue
				}

				// else if buf[2] == 32 { //click detect
				//fmt.Println("CLICKED AT:",buf[3]-32, buf[4]-32)
				//}
			} else if buf[1] == 66 { //down
				position = mtr.DOWN(position, system, 1)
				continue
			} else if buf[1] == 65 { //up
				position = mtr.UP(position, system, 1)
				continue
			}

		case 13, 113, 81:
			//executes the welcome.itl screen properly
			//this will ensure its done without issues happening
			return mtr.session.Write("\033c" + captured)
		case 87, 119: //up: w, W
			position = mtr.UP(position, system, 1)

			//ignores header cols properly
			//ensures its done without issues
			if position <= 2 { //length checks
				continue //continues
			}

		case 102, 70:

			if searched >= len(toml.ConfigurationToml.Pager.Colours) {
				searched = 0
			}

			//moves cursor to the bottom properly
			//this will be used within the future proeprly
			mtr.session.Write(fmt.Sprintf("\033[s\033[%d;%df\r\033[K", mtr.session.Height, mtr.session.Length))

			//takes the input correctly and properly
			//used within the system to try to locate the feature
			lookin, err := term.NewTerminal(mtr.session.Channel, "\x1b[0m Command>").ReadLine()
			if err != nil {
				return err
			}

			var commandPosition int = 0

			//ranges through all
			for position, line := range system {

				//checks current line
				//ensures we can find properly
				if strings.Contains(line, lookin) {
					commandPosition = position //looks
					break                      //breaks looping
				}
			}

			if commandPosition <= 1 {
				continue
			}

			//highlights the option wanted properly
			system[commandPosition] = strings.ReplaceAll(toml.ConfigurationToml.Pager.Colours[searched], "\\x1b", "\x1b") + " \x1b[0m" + strings.Replace(system[commandPosition], " ", "", 1)

			mtr.session.Write(strings.Join(DisplayChunk(system, 0, mtr.session.Height), "\r\n"))
			mtr.session.Write("\033[" + strconv.Itoa(position) + ";0f")

			searched++
		case 83, 115: //down: d, D
			position = mtr.DOWN(position, system, 1)
		}
	}
}

func (mtr *MakeTableRender) UP(position int, system []string, move int) int {
	//if position == 1 {
	//	mtr.session.Write("\033[0;0F\x1b[48;5;15m\x1b[38;5;16m" + Centre("Use W,S,UP,DOWN,SCROLL to navigate", mtr.session.Length, "") + "\x1b[0m\r")
	//}

	if position-move <= 0 {
		return position
	}

	position -= move
	if position > mtr.session.Height {
		mtr.session.Write("\033c" + strings.Join(DisplayChunk(system, position-mtr.session.Height, position), "\r\n"))
	}

	if position+1 > mtr.session.Height {
		mtr.session.Write("\033c" + strings.Join(DisplayChunk(system, 0, position), "\r\n"))
	}

	mtr.session.Write("\033[" + strconv.Itoa(position) + ";0f")
	return position
}

func (mtr *MakeTableRender) DOWN(position int, system []string, move int) int {
	if position+move > len(system) {
		return position
	}

	position += move
	if position > mtr.session.Height {
		mtr.session.Write("\033c" + strings.Join(DisplayChunk(system, position-mtr.session.Height, position), "\r\n"))
	}

	if position+1 > mtr.session.Height {
		mtr.session.Write("\033c" + strings.Join(DisplayChunk(system, 0, position), "\r\n"))
	}

	mtr.session.Write("\033[" + strconv.Itoa(position) + ";0f")
	return position
}

// displays a chunk with given length as a perfect position given
// this will ensure it only shows the chunk wanted within the args
func DisplayChunk(array []string, position int, at int) []string {

	//for loop
	return array[position:at]
}

// centres the text properly
// this will return the string properly
func Centre(s string, c int, dst string) string { //string
	//for loops through properly
	//this will ensure its done without any issues
	for p := 0; p < c; p++ { //loops through properly
		if p == c/2-len(s)/2 { //compares to middle
			dst += s        //saves in properly
			p += len(s) - 1 //skips chars properly
		} else {
			dst += " "
		}
	}
	return dst
}
