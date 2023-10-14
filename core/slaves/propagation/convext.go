package propagation

import (
	"Nosviak2/core/slaves"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"log"
	"net"
	"strconv"
	"strings"
)

var Reps map[string]int = make(map[string]int)

//grabs the count from the connection properly
//this will ensure its done without any errors happening
func GrabCount(connection net.Conn) error { //err happening

	//checks if server is inside the whitelist properly
	//only accepts whitelisted server properly and ensures safely
	if !tools.NeedleHaystackOne(toml.ConfigurationToml.Propagation.Whitelist, strings.Split(connection.RemoteAddr().String(), ":")[0]) {
		log.Printf("[PROPAGATION] (unwhitelisted server has tried to connection) FROM: (%s)\r\n", connection.RemoteAddr().String())
		connection.Close(); return nil //drops connection properly and safely
	}

	//lets the main server know about the connection
	log.Printf("[PROPAGATION] (New propagated count introduced) (%s)\r\n", connection.RemoteAddr().String())

	for {
		Incoming := make([]byte, 10)
		if _, err := connection.Read(Incoming); err != nil {
			return err //err handles properly
		}

		out := slaves.Sanatize(string(Incoming)) //sanatizes properly

		//converts into int properly
		//this will ensure its done without errors
		raw, err := strconv.Atoi(out) //converts properly
		if err != nil {
			return err
		}

		//saves into the map properly and safely
		//allows us to build a relative count safely
		Reps[connection.RemoteAddr().String()] = raw
		defer delete(Reps, connection.RemoteAddr().String())
	}
}