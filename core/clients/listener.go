package clients

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
	"time"

	"golang.org/x/crypto/ssh"

	"Nosviak2/core/clients/routes"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/logs"
)

//starts the listener for the server properly
//this will make sure its done without issues happening
func CreateListener(configuration ssh.ServerConfig) error {

	//tries to request the server to start
	//this will ensure its done without issues happening
	network, err := net.Listen(json.ConfigSettings.Masters.Server.Protocol, json.ConfigSettings.Masters.Server.Listener)
	if err != nil { //error handles the new startup request
		return err //returns the error which happened properly
	} else {
		//prints a basic message to say the server was started
		//this will make sure its not ignored correctly and properly
		fmt.Printf("[SSH Server] [the ssh server conection watcher has been started] [%s%s]\r\n", json.ConfigSettings.Masters.Server.Protocol, json.ConfigSettings.Masters.Server.Listener)
	}

	//starts the title worker
	//this will broadcast the title to each session without issues
	go routes.TitleWorkerFunction()

	//for loops througout the system connections
	//this will accept any new connections without issues
	for { //for loops for connections
		//accepts the new request properly
		//this will make sure its accepted properly
		request, err := network.Accept()
		if err != nil { //tries to accept the request
			//prints the error with the connection location properly
			log.Printf("connection from %s has failed, reason: %s\r\n", request.RemoteAddr().String(), err.Error()); continue
		}

		//tries to correctly write the log into the file
		//this will ensure its done without any errors happening
		if err := logs.WriteLog(filepath.Join(deployment.Assets, "logs", "connections.json"), logs.ConnectionLog{Type: "tcp", Address: request.RemoteAddr().String(), Username: "tbd", Time: time.Now()}); err != nil {
			log.Printf("logging fault: %s\r\n", err.Error()) //alerts the main terminal properly
		}

		//prints the new connection information
		//this will allow for virtual errors without errors
		log.Printf("new tcp connection has been recorded from %s\r\n", request.RemoteAddr().String())


		//serves the new connection properly without issues
		//this will ensure its happened properly without issues
		go serveIncomingConnection(request, configuration)
	}
}