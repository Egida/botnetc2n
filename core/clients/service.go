package clients

import (
	"Nosviak2/core/clients/routes"
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/logs"
	"encoding/binary"
	"path/filepath"
	"time"

	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

//handles the new socket connections
//this will make sure its done without issues happening
//this will make the system safer without issues happening
func serveIncomingConnection(c net.Conn, s ssh.ServerConfig) {
	//creates the new server connection
	//this will allow us to sort the different connection types
	//allows for proper and safer connection handling without errors
	conns, chans, _, err := ssh.NewServerConn(c, &s)
	if err != nil { //error handles the new server statement
		//this will print the information properly about it
		//this will helpfully print the reason to the terminal screen
		log.Printf("[SSH SERVER] HANDSHAKE FAULT: (%s) [%s]", c.RemoteAddr().String(), err.Error())
		return
	}

	//checks for normal auth properly
	//if this is enabled we will print the normal
	if !json.ConfigSettings.Masters.Server.DynamicAuth {
		//prints a simple message saying new connection
		//this will help with handling without issues happening
		log.Printf("[SSH SERVER] NEW SSH CONNECTION (%s) (%s) (%s)\r\n", c.RemoteAddr().String(), conns.User(), string(conns.ClientVersion()))
		//tries to correctly write the log into the file
		//this will ensure its done without any errors happening
		if err := logs.WriteLog(filepath.Join(deployment.Assets, "logs", "connections.json"), logs.ConnectionLog{Type: "ssh", Address: conns.RemoteAddr().String(), Username: conns.User(), Time: time.Now()}); err != nil {
			log.Printf("logging fault: %s\r\n", err.Error()) //alerts the main terminal properly
		}
	}

	//ranges througout the channels
	//this will ensure its done properly without issues
	for channel := range chans { //ranges througout channels
		//detects invalid sessions properly
		//this will make sure its not ignored error
		if channel.ChannelType() != "session" { //rejects the connection properly
			channel.Reject(ssh.UnknownChannelType, "UnknownChannelType")
			return
		}

		//accepts the incoming channel properly
		//this will allow it to interface with the network
		Channel, request, err := channel.Accept() //accepts the channel
		if err != nil {
			//this will print the information properly about it
			//this will helpfully print the reason to the terminal screen
			log.Printf("[SSH SERVER] HANDSHAKE FAULT: (%s) [%s]", c.RemoteAddr().String(), err.Error())
			return
		}

		//used inside the session name
		//this will ensure its done without any errors
		created := time.Now()

		//handles the ssh requests properly
		//this will make sure they are all replyed to safely
		go func(in <-chan *ssh.Request) {
			for req := range in { //ranges through the reqs
				switch req.Type { //switchs the req type properly
				case "pty-req": //PTY interfaces
					req.Reply(true, nil) //replys true
					continue             //continues the loop
				case "shell": //basic shell interfaces
					req.Reply(true, nil) //replys true
					continue             //continues the loop
				case "window-change": //different window change
					//tries to properly change the different
					//this will ensure its done without any errors
					w, h := ParseWinDifferent(req.Payload) //parses
					//tries to find the value inside the map
					//this will ensure its done without any issues
					s := sessions.Sessions[created.Unix()] //tries to grab the session
					if s == nil {                          //removes the instance properly
						return //returns the function properly
					}

					//sets the new values properly without issues happening
					//this will ensure its done without any errors happening
					sessions.Sessions[created.Unix()].Height = int(h) //sets height
					sessions.Sessions[created.Unix()].Length = int(w) //sets length properly
				}
			}
		}(request)
		//creates a new handling function correctly
		//this will help for safer controlling without issues
		err = routes.NewHandlerFunction(conns, Channel, created)
		if err != nil { //error handles the main interface properly
			log.Printf("connection from %s unlawfully closed because %s\r\n", c.RemoteAddr().String(), err.Error())
		}

		//force closes the session
		//this will make sure its done without errors
		Channel.Close() //anything which is returned should be closest without issues

		//checks to close the session
		//this wil allow for better handling without issues happening
		delete(sessions.Sessions, created.Unix()) //deletes from the map correctly
	}
}

//parses the byte value into the sections
//this will ensure its done without any error
func ParseWinDifferent(b []byte) (uint32, uint32) {
	w := binary.BigEndian.Uint32(b)     //detects the windows length of the binary
	h := binary.BigEndian.Uint32(b[4:]) //detects how tall the window is
	return w, h
}
