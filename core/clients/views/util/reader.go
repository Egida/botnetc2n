package util

import (
	"strings"

	"golang.org/x/crypto/ssh"
)

//builtin terminal reader properly
//this will ensure its done without any errors happening
func TermReader(channel ssh.Channel, maxVLength int, mask bool, maskChar string) (string, error) {

	var texture []string = make([]string, 0)
	//readers loops inside a loop
	//this will ensure we keep looping reading inputs
	for { //live buffer properly
		buf := make([]byte, 1)

		//reads into from the buffer properly
		//this will take the inputs properly and safely
		if _, err := channel.Read(buf); err != nil { //err handles
			return "", err //returns the error properly
		}


		switch buf[0] {

		//detects random key presses
		//this will ensure its done without any errors
		case 3, 9, 2, 1:
			continue

		//properly ignored within this method
		case 27: //any disallowed charaters
			//absorbs the extra charaters within the reqeust
			if _, err := channel.Read(make([]byte, 2)); err != nil {
				return "", err //returns error properly
			}
			//ignores this charater
			continue


		case 127:
			//stops any futher backspacing
			//this will ensure its stops safely
			if len(texture) <= 0 { //checks properly
				continue //continues looping properly
			}

			//writes the backspace ansi code properly
			//this will remove one charater and move the cursor properly
			if _, err := channel.Write([]byte{127}); err != nil {
				return "", err //returns any error which happens
			}

			//updates the source properly
			//this will ensure its done without any errors
			texture = texture[:len(texture)-1]

		//enter key button press detection
		//this will ensure its done without any errors
		case 13: //detects enter key presses within this system
			if _, err := channel.Write([]byte("\r\n")); err != nil {
				return "", err //returns the error properly
			}

			//returns the product properly
			//this will ensure its done without errors
			return strings.Join(texture, ""), nil
		default:
			//saves into the array properly without issues
			//this will ensure its done without any errors happening
			texture = append(texture, string(buf[0]))
			
			//checks the length properly
			//this will check if we are rending the information
			if len(texture) > maxVLength { //checks the length properly
				continue //continues looping
			} else if mask { //writes the mask char
				channel.Write([]byte(maskChar)) //masking
			} else { //writes the charater given properly without issues
				channel.Write([]byte(string(buf[0]))) //writes properly
			}
		}
	}
}