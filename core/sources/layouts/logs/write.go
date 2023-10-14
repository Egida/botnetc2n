package logs

import (
	"encoding/json"
	"os"
)

//writes into the logging file properly
//this will ensure its done without any errors happening
func WriteLog(f string, value interface{}) error { //returns error properly
	//tries to convert the value into string
	//this will allow for better control without issues
	values, err := json.Marshal(value) //converts properly
	if err != nil { //error handles properly without issues
		return err //returns the error properly
	}
	
	//tries to open the file for the logs
	//this will ensure its done without any issues happening
	FLoc, err := os.OpenFile(f, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil { //error handles properly without issues
		return err //returns the error properly
	}
	
	//closes once done properly
	//this will ensure its done without errors
	defer FLoc.Close() //closes when function done

	//tries to write to the file
	//this will ensure its done without any errors
	if _, err := FLoc.WriteString(string(values)+"\r\n"); err != nil {
		return err //returns the error properly without issues
	}; return nil //nothing happened
}