package sessions

import (
	"Nosviak2/core/database"
	"Nosviak2/core/sources/ranks"
	"time"

	"golang.org/x/crypto/ssh"
)

//creates the new session properly
//this will ensure its done correctly without issues
func MakeCorrectiveSession(usr *database.User, c *ssh.ServerConn, channel ssh.Channel, dura int64, themePath []string, colours [][]int, rewrite string, tracer []int, ranks ranks.Ranks) *Session {
	Mutex.Lock() //locks the maps mutex properly
	defer Mutex.Unlock() //unlocks once function completed
	//this will hold the information without issues happening
	var s *Session = &Session{User: usr, Conn: c, Channel: channel, ID: dura, Connected: time.Now(), Idle: time.Now(), WindowSize: &WindowSize{Length: 80, Height: 24}, BrandingPath: themePath, Colours: colours, Target: rewrite, TracerMatrix: tracer, Ranks: &ranks, Chat: false, Title: "", Viewers: 0} //sets the default information correctly and properly
	Sessions[s.ID] = s //saves into the map correctly and properly without issues happening on requests
	return s
}