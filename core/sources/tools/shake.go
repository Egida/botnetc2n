package tools

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

//purely shakes the terminal properly safely
//this will ensure its shakes in certain ways properly
func ShakeTerminal(rotations int, timer time.Duration, channel ssh.Channel) error {


	//gets the current window pos
	//this will ensure its done properly
	current, err := GetWindow(channel)
	if err != nil { //err handles properly
		return err //err handles
	}

	//loops through the amount of times wanted properly
	for times := 0; times < rotations; times++ {

		channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal+3, current.Vertical+3)))
		time.Sleep(timer) //sleeps for the time given inbetween //sleeps for the time given inbetween

		channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal, current.Vertical+3)))
		time.Sleep(timer) //sleeps for the time given inbetween

		channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal+3, current.Vertical)))
		time.Sleep(timer) //sleeps for the time given inbetween

		channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal-3, current.Vertical-3)))
		time.Sleep(timer) //sleeps for the time given inbetween

		channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal-3, current.Vertical)))
		time.Sleep(timer) //sleeps for the time given inbetween

		channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal, current.Vertical-3)))
		time.Sleep(timer) //sleeps for the time given inbetween
	}

	channel.Write([]byte(fmt.Sprintf("\x1b[3;%d;%dt", current.Horizontal, current.Vertical)))

	return nil
}