package commands

import (
	"Nosviak2/core/clients/sessions"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/views"
	"Nosviak2/core/sources/webhooks"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

var (
	//sets default reload value properly
	last_reload time.Time = time.Now()
	first_time  bool      = true
)

func init() {

	MakeCommand(&Command{
		CommandName:        "reload",
		Aliases:            []string{"reload", "update"},
		CommandPermissions: []string{"admin", "moderator"},
		CommandDescription: "concurrently reload all nosviak assets",
		CommandFunction: func(s *sessions.Session, cmd []string) error {

			//only allow reload every 1 minute or ignore request
			if !first_time && toml.ConfigurationToml.Itl.ReloadLimiter && last_reload.Add(1*time.Minute).Unix() > time.Now().Unix() {
				return language.ExecuteLanguage([]string{"commands", "reload", "limited.itl"}, s.Channel, deployment.Engine, s, map[string]string{"until": strconv.FormatFloat(time.Until(last_reload.Add(1*time.Minute)).Seconds(), 'f', 2, 64)})
			}

			last_reload = time.Now() //sets new reloaded time properly
			first_time = false       //filps the first_time properly

			//tries to reload the json information properly
			//this will make sure its done without issues happening on request
			err := json.MakeEngine(deployment.Assets, deployment.JsonHierarchy).RunEngine() //tries to execute the json path
			if err != nil {                                                                 //error handles the statement properly without issues happening on request
				s.Write("Hmmm, it seems there was an error while reloading your json files\r\n")
				s.Write("the issue which was acquired can be viewed inside your Terminal log!\r\n")
				log.Printf("%s acquired error: %s\r\n", s.User.Username, err.Error()) //returns the error properly
				return nil                                                            //ends the function properly without issues happening
			}
			//flushes all the toml ranks properly
			//deletes all the preset ranks from the files correctly
			//this will ensure its done properly without errors happening
			for header := range toml.RanksToml.Ranks { //ranges through the ranks properly
				delete(ranks.PresetRanks, header) //flushes the rank correctly and properyl
			}
			//tries to reload the toml information properly
			//this will make sure its done without issues happening on request
			err = toml.MakeEngine(deployment.Assets, deployment.TomlHierarchy).RunEngine() //tries to execute the toml path
			if err != nil {                                                                //error handles the statement properly without issues happening on request
				s.Write("Hmmm, it seems there was an error while reloading your toml files\r\n")
				s.Write("the issue which was acquired can be viewed inside your Terminal log!\r\n")
				log.Printf("%s acquired error: %s\r\n", s.User.Username, err.Error()) //returns the error properly
				return nil                                                            //ends the function properly without issues happening
			} else { //else statement when correctly done
				//this will properly sync the information loaded without issues
				//this will sync the ranks without issues happening on request
				for header, body := range toml.RanksToml.Ranks { //ranges through the ranks properly
					ranks.PresetRanks[header] = ranks.RankSettings{
						RankDescription: body.RankDescription, //saves into the information
						MainColour:      body.MainColour, SecondColour: body.SecondColour,
						SignatureCharater: body.SignatureCharater, CloseWhenAwarded: false,
						Manage_ranks: body.Manage_ranks, DisplayInTable: true, //show in table
					}
				}
			}

			webhooks.Reset() //clears the webhook map
			//tries to properly and completely reload the webhooks
			//this will allow for the branding to be entirly reloaded
			if err := webhooks.RenderWebhooks(); err != nil { //error handles properly
				s.Write("Hmmm, it seems there was an error while reloading your webhook files\r\n")
				s.Write("the issue which was acquired can be viewed inside your Terminal log!\r\n")
				log.Printf("%s acquired error: %s\r\n", s.User.Username, err.Error()) //returns the error properly
			}

			//clears all the views
			//this will ensure we can properly reload without issues
			views.Reset()                                                          //resets the array of contents properly without issues
			err = views.GatherPeices(filepath.Join(deployment.Assets, "branding")) //tries to load all
			if err != nil {                                                        //error handles the information properly without issues
				s.Write("Hmmm, sorry there was an unforeseen error when reloading branding\r\n")
				s.Write("the issue which was acquired can be viewed inside your terminal log\r\n")
				log.Printf("%s acquired error: %s\r\n", s.User.Username, err.Error()) //returns the error properly
				return nil
			}

			//resets all the custom commands
			//this will ensure its done without any errors
			RemoveObjects()                                               //removes the objects properly without issues
			l, err := EngineLoader(deployment.Assets, "commands", "text") //loads the commands
			if err != nil {                                               //error handles the module without issues happening on reqeust
				s.Write("Hmmm, sorry there was an unforeseen error when custom commands\r\n")
				s.Write("the issue which was acquired can be viewed inside your terminal log\r\n")
				log.Printf("%s acquired error: %s\r\n", s.User.Username, err.Error()) //returns the error properly
				return nil
			}

			//prints a simple success message
			//this will allow us to alert the user about it properly
			return s.Write(fmt.Sprintf("Great news %s, we have reloaded all your assets\r\n%d branding objects have been reinspected and found to be successful\r\n%d custom commands objects have been reinspected and found to be successful\r\n%d webhook objects have been reinspected and found to be successful\r\nyour Json & Toml assets have also been reinspected and found to be correct\r\n", s.User.Username, len(views.Subject), l, len(webhooks.Webhooking)))
		},
	})
}
