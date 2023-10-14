package tools

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Position struct {
	Horizontal, Vertical int
}

var (
	//err when invalid position happens
	ErrPositionOutput error = errors.New("invalid position output has been given properly")
)

//gets the current window position on the screen
//returns the structure properly and safely within it
func GetWindow(channel ssh.Channel) (*Position, error) {

	//enables the window reporting properly
	//this will allow us to completely run within it
	if _, err := channel.Write([]byte("\x1b[13t")); err != nil {
		return nil, err
	}

	buf := make([]byte, 56)
	//reads from the channel properly
	//this will allow us to parse and grab values
	if _, err := channel.Read(buf); err != nil {
		return nil, err
	}

	//parses the incomign system properly
	//this will make sure its done properly
	position := string(LoopTillZero(buf[4:], make([]byte, 0)))

	//validates the information properly and safely
	//this will ensure its done without errors happening
	if strings.Count(position, ";") <= 0 || strings.ToLower(strings.Split(position, "")[len(strings.Split(position, ""))-1]) != "t" {
		return nil, ErrPositionOutput
	}

	//gets positions properly and safely
	//vert, hori := strings.ReplaceAll(strings.Split(position, ";")[1], "t", ""), strings.Split(position, ";")[0]

	//tries to convert into en int properly
	//this will ensure its done without errors happening
	horizontal, err := strconv.Atoi(strings.Split(position, ";")[0])
	if err != nil { //err handles
		return nil, err
	}

	//tries to convert into en int properly
	//this will ensure its done without errors happening
	vertical, err := strconv.Atoi(strings.ReplaceAll(strings.Split(position, ";")[1], "t", ""))
	if err != nil { //err handles
		return nil, err
	}

	//returns the values properly and safely
	//this will ensure its done without errors happening
	return &Position{Horizontal: horizontal, Vertical: vertical}, nil
}

//loops until the byte value is equal to 0
//this will ensure its only has proper values
func LoopTillZero(src []byte, dst []byte) []byte {
	//ranges through src
	for p := range src {

		//if byte value is 0
		//we ignore them properly
		if src[p] == 0 { //checks
			return dst //ends loop
		}

		//appends into the dst properly
		dst = append(dst, src[p])
	}
	
	//returns the values
	return dst
}