package tools

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

//generates a strong password properly
//this will return the strong password string
func CreateStrongPassword(l int) string { //returns string
	//everyone one of these charaters can be used inside the system
	var charaters string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	var Pass []string = make([]string, 0) //creates the password array
	//ranges through the length of the suggested password
	//this will make sure its done without issues happening on request
	for pass := 0; pass < l; pass++ { //saves into the array correctly and properly
		Pass = append(Pass, strings.Split(charaters, "")[rand.Intn(len(charaters))])
	}; return strings.Join(Pass, "") //returns the password correctly and properly
}