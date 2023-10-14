package logs

import "time"

//stores information about the log
//this will store informaiton about the user
//like there ip and username etc without issues
type ConnectionLog struct { //stored in type structure
	Type string `json:"type"` //stores the type: ssh, tcp
	Address string `json:"address"` //stores the address: 1*.1*.1*.1
	Username string `json:"username"` //this might be a option
	Time time.Time `json:"time"`
}

//stores the command structure properly
//this will ensure its done without any errors
type CommandLog struct { //stored in type structure properly
	Command string `json:"command"`
	Args []string `json:"args"`
	Username string `json:"username"`
	Time time.Time `json:"time"`
}

//stores the attack command structure properly
//this will ensure its done without any errors happening
type AttackLog struct { //stores the command information properly
	Target string `json:"target"`
	Duration int `json:"duration"`
	Port int `json:"port"`
	Method string `json:"method"`
	Created int64 `json:"created"`
	End int64 `json:"finish"`
	Username string `json:"username"`
}