package propagation

import (
	"Nosviak2/core/sources/layouts/toml"
	"fmt"
	"net"
)

//makes the propagation server properly
//this allows bot **COUNT** transition properly
func MakePropagation() error { //err handles properly

	//tries to catch the listener
	//this will ensure its done properly
	network, err := net.Listen("tcp", toml.ConfigurationToml.Propagation.Listener)
	if err != nil { //err handles properly
		return err
	} else {
		fmt.Printf("[Propagation] [slave propagation started on %s]\r\n", toml.ConfigurationToml.Propagation.Listener)
	}

	for {
		//accepts the connection from remote
		//allows us to grab information safely
		possibleCatch, err := network.Accept()
		if err != nil { //err handles
			continue
		}

		//runs the count grapper
		//this allows it to run properly
		go GrabCount(possibleCatch)
	}
}