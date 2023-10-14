package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

//TODO: get more information about the target!

func main() {
	//checks the length properly
	if len(os.Args) <= 1 { //len checks
		fmt.Printf(" Syntax: cfx <token>\r\n")
	}

	//creates the url properly without issues
	//this will format properly without any errors
	var url string = "https://cfx.re/join/" + url.QueryEscape(os.Args[1])

	//makes the client within the guidelines
	//this will create the client properly without issues
	client := http.Client{ //makes the structure
		//sets the timeout duration
		Timeout: 5 * time.Second, //sets timeout
	}

	//makes the new request properly
	//this will create the structure properly
	req, err := http.NewRequest("GET", url, nil)
	if err != nil { //err handles
		fmt.Printf(" Failed to perform the request towards the server!\r\n")
		return
	}

	//performs the reqeust without issues
	//this will ensure its done without any errors
	res, err := client.Do(req)               //performs the reqeust
	if err != nil || res.StatusCode != 200 { //err handles properly without issues
		fmt.Printf(" Failed to perform the request towards the server!\r\n")
		return
	}

	fmt.Printf(" %s's statistics\r\n", os.Getenv("username"))
	fmt.Print(" IP: ", res.Header.Get("x-citizenfx-url")+"\r\n")
	fmt.Print(" Token: ", os.Args[1]+"\r\n")
}
