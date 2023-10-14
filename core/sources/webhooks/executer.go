package webhooks

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"log"
	"strings"
)

//properly gets all the webhooks safely
//this will ensure its done properly without issues
func PerformEventWebhook(events []string, session *sessions.Session, values map[string]string) error {
	//checks to see if the event is classed as a trigger
	//this will ensure its properly done without any errors happening
	if ok, _ := tools.NeedleHaystack(toml.WebhookingToml.Webhooks.Trigger, strings.Join(events, deployment.Runtime())); !ok {
		return nil //ends the function properly
	}

	//stores the header properly
	//this will ensure we have the header
	var header string = "" //stored in the header
	//stores the split event properly and safely
    if len(strings.Join(events, deployment.Runtime())) - 1 == 0 {
		header = "main" //sets the header
	} else { //sets the header properly
		header = events[len(events)-1]
	}

	//gets the text properly without issues
	//this will ensure this is properly done without issues
	text := Webhooking[strings.Join(events, deployment.Runtime())] //gets the text
	text = strings.ReplaceAll(text, "<<$useraction>>", session.User.Username) //replaces the user
	text = strings.ReplaceAll(text, "<<$ip>>", session.Conn.RemoteAddr().String()) //replaces the ip properly
	//ranges through all the values properly
	//this will each object is updated within reason
	for key, value := range values { //ranges through properly
		text = strings.ReplaceAll(text, "<<$"+key+">>", value) //replaces all properly
	}


	//tries to find the webhook file
	//this will ensure its done without errors
	if err := RunWebhookSystem(text, header, toml.WebhookingToml.Webhooks.Colour); err != nil {
		log.Printf("[ERROR WEBHOOK] [Failed to correctly launch webhook for %s] [%s]\r\n", session.User.Username, err.Error())
		return err //returns the error correctly and properly without issues happening
	}
	return nil
}