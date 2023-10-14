package pointer

import (
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"fmt"
	"net"
	"strconv"
)

//makes the pointers towards the system
//this will ensure its done without issues happening
func MakePointer() error { //returns an error properly

	for { //keeps looping properly
		system, err := net.Dial("tcp", toml.ConfigurationToml.Pointer.Pointer)
		if err != nil { //err handles properly
			if deployment.DebugMode { //debug mode
				fmt.Printf("[POINTER] %s\r\n", err.Error())
			}; continue
		} else if deployment.DebugMode {
			fmt.Printf("[POINTER] [Connection made with pointer server]")
		}

		var cache int = -1
		for { //loops properly and safely
			current := BuildString()

			//gets the current
			if current == cache {
				continue //continues
			}

			//writes the count into the file properly
			system.Write([]byte(strconv.Itoa(current)))
			cache = current //sets cache properly
		}
	}
}