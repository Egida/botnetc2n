package events

import (
	"Nosviak2/core/clients/commands"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/views"
	"Nosviak2/core/sources/webhooks"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
)

//this will properly detect a live update within the assets
//allows for proper control without issues happening on reqeust
func LiveRenderUpdate() error { //returns an error if one happened properly

	//this will make sure that its ignored
	//follows the configs wishes of live being disabled
	if !toml.ConfigurationToml.Itl.LiveEnabled { //disabled
		return nil //detects that live is disabled
	}

	//creates the watcher controller properly
	//this will ensure its done without any errors
	watchers := watcher.New() //makes the watcher event
	watchers.SetMaxEvents(1)

	//adds the recursive reader for the assets
	//this will follow through properly without issues happening
	if err := watchers.AddRecursive(deployment.Assets); err != nil {
		return err //returns the error properly without issues
	}
	go func() {
		//loops through events
		//this will ensure its done without any errors
		for { //loops forever through properly without issues
			select { //selects the information properly
			//displays the new event information properly
			case event := <- watchers.Event: //new event
				//gets the file which has been edited
				//this will ensure we ignore the path upto
				internal := strings.Split(event.Path, deployment.Runtime())[len(strings.Split(event.Path, deployment.Runtime()))-1]

				//detects the different file types
				//we will only reload the given file type
				switch filepath.Ext(internal) { //switches through properly

				//detects a json shift properly
				//this will now follow the only json route
				case ".json": //we will also ignore the log information
					//detects if we are allowing live config updates
					//this will ensure we don't ignore it without issues
					if !toml.ConfigurationToml.Itl.JSONBrandingRefresh || strings.Contains(event.Path, "logs") {
						continue //continues looping properly
					}
					//runs with the json update properly
					//this will enforce its done without any errors
					//this will follow the json engine properly without issues happening
					if err := json.MakeEngine(deployment.Assets, deployment.JsonHierarchy).RunEngine(); err != nil {
						log.Printf("[LiveReload] [Json] [%s]\r\n", err.Error()); continue //continues looping properly
					}; 
					
					//displays the correct information properly
					//this will print to the terminal without issues
					log.Printf("[LiveReload] [Json] [@%s] [reloaded all of your json assets properly]", deployment.Assets)
					//adds support for the sleep between properly without issues
					time.Sleep(time.Duration(toml.ConfigurationToml.Itl.TimeoutBetween) * time.Millisecond); continue

				//detects a toml shift properly
				//this will now follow the only toml route
				case ".toml": //we will also ignore the log information
					//detects if we are allowing live config updates
					//this will ensure we don't ignore it without issues
					if !toml.ConfigurationToml.Itl.TomlBrandingRefresh {
						continue //continues looping properly
					}
					//flushes all the toml ranks properly
					//deletes all the preset ranks from the files correctly
					//this will ensure its done properly without errors happening
					for header := range toml.RanksToml.Ranks { //ranges through the ranks properly
						delete(ranks.PresetRanks, header) //flushes the rank correctly and properyl
					}

					//runs with the toml update properly
					//this will enforce its done without any errors
					//this will follow the toml engine properly without issues happening
					if err := toml.MakeEngine(deployment.Assets, deployment.TomlHierarchy).RunEngine(); err != nil {
						log.Printf("[LiveReload] [Toml] [%s]\r\n", err.Error()); continue //continues looping properly
					}; 

					//ranges through and adds properly
					//this will ensure its done without errors happening
					for header, body := range toml.RanksToml.Ranks { //ranges through
						ranks.PresetRanks[header] = ranks.RankSettings{
							RankDescription: body.RankDescription, //saves into the information
							MainColour: body.MainColour, SecondColour: body.SecondColour,
							SignatureCharater: body.SignatureCharater, CloseWhenAwarded: false,
							Manage_ranks: body.Manage_ranks, DisplayInTable: true, //show in table
						}
					}
					
					//displays the correct information properly
					//this will print to the terminal without issues
					log.Printf("[LiveReload] [Toml] [%s] [reloaded all of your toml assets properly]", deployment.Assets)
					//adds support for the sleep between properly without issues
					time.Sleep(time.Duration(toml.ConfigurationToml.Itl.TimeoutBetween) * time.Millisecond); continue
				}

				//checks if branding reload is active
				//this will ensure its only done if asked for
				if !toml.ConfigurationToml.Itl.LiveBrandingRefresh {
					continue //continues looping properly without issues
				}
				//reloads all branding and commands
				//this will reload all branding, commands & webhooking properly
				//this will reset all the structures properly without any errors happening
				views.Reset(); webhooks.Reset(); commands.RemoveObjects() //clears all properly
				//reloads all webhooks properly
				//this will reload every single webhooks
				if err := webhooks.RenderWebhooks(); err != nil { //err handles properly without issues
					log.Printf("[LiveReload] [Webhooks] [%s] [%s]", deployment.Assets, err.Error()); continue //err
				}
				//reloads all branding properly
				//this will reload every single branding
				if err := views.GatherPeices(filepath.Join(deployment.Assets, "branding")); err != nil { //err handles properly without issues
					log.Printf("[LiveReload] [Views] [%s] [%s]", deployment.Assets, err.Error()); continue //err
				}
				//reloads all text commands properly
				//this will ensure its done without any errors happening
				LR, err := commands.EngineLoader(deployment.Assets, "commands", "text")
				if err != nil { //err handles properly without any issues happening
					log.Printf("[LiveReload] [CustomCommands@Text] [%s] [%s]", deployment.Assets, err.Error()); continue //err
				}
				//prints all of the success messages properly
				//this will ensure its done without any issues happening
				log.Printf("[LiveReload] [Webhooks] [%d] [reloaded all of your webhook items properly]", len(webhooks.Webhooking)) //webhooks
				log.Printf("[LiveReload] [Views] [%d] [reloaded all of your view items properly]", len(views.Subject)) //views
				log.Printf("[LiveReload] [CustomCommands@Text] [%d] [reloaded all of your custom text commands items properly]", LR) //custom text commands
			}
		}
	}()

	//adds the watcher start properly
	//this will for loop through properly
	if err := watchers.Start(time.Duration(toml.ConfigurationToml.Itl.TimeoutBetween) * time.Millisecond); err != nil {
		return err //returns the error which happened properly
	}; return nil //we wont return any errors properly
}