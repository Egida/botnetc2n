package static

import (
	termfx "Nosviak2/core/sources/language/tfx"
	"io"
	"strconv"
)

//runs the system within the slaves controller properly
//this will ensure its done without errors happening on purpose
func SlaveController(payload string, arch string, IP, Port string, Type string, Same int, All int, s io.Writer) (int, error) {

	terms := termfx.New()

	terms.RegisterVariable("ip", IP) //ip <<$ip>>
	terms.RegisterVariable("port", Port) //port <<$port>>
	terms.RegisterVariable("type", Type) //type <<$type>>
	terms.RegisterVariable("architecture", arch) //architecture <<$architecture>>

	//devices system properly and safely
	//this will ensure its done without errors happening 
	terms.RegisterFunction("devices", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(All)))
	})

	//devices system properly and safely
	//this will ensure its done without errors happening 
	terms.RegisterFunction("same_arch_connected", func(session io.Writer, args string) (int, error) {
		return session.Write([]byte(strconv.Itoa(Same)))
	})

	//executes the system properly and safely
	//allows for proper contorl without issues
	system, err := terms.ExecuteString(payload)
	if err != nil { //err handles properly
		return 0, err
	}

	//writes into the writer properly
	//this will make sure its done without issues
	return s.Write([]byte(system))
}