package webhooks

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//runs the webhook system properly
//this will ensure its done without issues
func RunWebhookSystem(text string, header string, colour int) error { //returns error properly
	//checks for the enabled option properly
	//this will ensure its done without any errors
	if !toml.WebhookingToml.Webhooks.Enabled { //checks properly
		return nil //returns nil as its disabled
	}
	

	//checks inside the custom map for the object
	//this will ensure its done without any errors
	custom := toml.WebhookingToml.CustomConfigs[strings.Split(strings.Split(header, deployment.Runtime())[len(strings.Split(header, deployment.Runtime()))-1], ".")[0]]
	if custom != nil { //decides that a custom config was found
		colour = custom.Colour //sets the custom colour properly
		header = strings.ReplaceAll(custom.Title, "<<$cnc>>", toml.ConfigurationToml.AppSettings.AppName) //enforces the custom header properly
	}

	//properly sets the webhook model
	//this will ensure its configured without errors
	ByteValue, err := json.Marshal(Model{Embeds: []Embeds{{Title: header,Description: text,Color: colour,Footer: Footer{Text: toml.ConfigurationToml.AppSettings.AppName + " - " + strconv.Itoa(time.Now().Year())}}}})
	if err != nil { //error handles the request without issues
		return err //returns the error which happened
	}

	//sets the client timeout properly
	//this will ensure its done without any errors happening
	cli := http.Client{ //stores the information, sets the timeout properlu
		Timeout: time.Duration(toml.WebhookingToml.Webhooks.Timeout) * time.Second,
	}

	//creates the request information properly
	//this will allow for better system handling without errors
	req, err := http.NewRequest("POST", toml.WebhookingToml.Webhooks.Token, bytes.NewBuffer(ByteValue))
	if err != nil { //error handles the reqeust properly without issues
		return err //returns the error which happened on reqeust
	}

	//sets the content type properly
	//this will ensure its done without any errors
	req.Header.Set("Content-Type", "application/json")

	//performs the reqeust
	//this will ensure its done without any errors
	res, err := cli.Do(req) //performs the reqeust without any issues
	if err != nil { //error handles the request without issues
		return err //returns the error properly
	}

	//checks the value length properly
	//this will ensure its properly done
	if res.StatusCode > 204 { //checks for safe
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(body), err)
		return errors.New("invalid status code of "+strconv.Itoa(res.StatusCode))
	}

	//returns nil as it was completed properly
	//this will ensure its done without any errors
	return nil
}