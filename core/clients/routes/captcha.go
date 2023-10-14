package routes

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language"
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/ranks"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/term"
)

var (
	//invalid answer error
	ErrQuestionUnauth error = errors.New("answer given has been classed as invalid")
)

//allows the capthca maths question
//this will mainly be used for hardcode support
func CaptchaRoute(s *sessions.Session) error { //returns error
	r := ranks.MakeRank(s.User.Username) //gets the ranks
	r.SyncWithString(s.User.Ranks) //syncs with a string format
	//checks if the user can bypass the request
	//this will ensure its done without any errors
	if r.CanAccessArray(toml.CatpchaToml.UserBypass) {
		return nil //returns nil as they can properly
	}
	//randomizes properly
	rand.Seed(time.Now().Unix())
	//loops through the attempt seq
	//this will ensure its done without any errors
	for attempt := 0; attempt < toml.CatpchaToml.MaximumAttempts; attempt++ {

		//generates the numbers properly
		//this will ensure its done without any errors
		paramOne := rand.Intn(toml.CatpchaToml.MaximumGeneration - toml.CatpchaToml.MinimumGeneration) + toml.CatpchaToml.MinimumGeneration //one
		paramTwo := rand.Intn(toml.CatpchaToml.MaximumGeneration - toml.CatpchaToml.MinimumGeneration) + toml.CatpchaToml.MinimumGeneration //two

		//tries to render the banner properly
		//this will ensure its done without any errors
		err := language.ExecuteLanguage([]string{"views", "captcha", "banner.itl"}, s.Channel, deployment.Engine, s, map[string]string{"QueryOne":strconv.Itoa(paramOne), "QueryTwo":strconv.Itoa(paramTwo)})
		if err != nil { //error handles properly
			return err //returns error properly
		}

		//renders the prompt information
		//this will ensure its done without any errors
		err = language.ExecuteLanguage([]string{"views", "captcha", "prompt.itl"}, s.Channel, deployment.Engine, s, map[string]string{"QueryOne":strconv.Itoa(paramOne), "QueryTwo":strconv.Itoa(paramTwo)})
		if err != nil { //error handles properly
			return err //returns error properly
		}

		//reads the input correctly
		//this will ensure its done without any errors
		query, err := term.NewTerminal(s.Channel, "").ReadLine()
		if err != nil { //error handles properly
			return err //returns the error safely
		}

		//converts the query into string
		//this will ensure its done without any errors
		given, err := strconv.Atoi(query) //converts into int
		if err != nil { //error handles properly
			return err //continues looping
		}

		//tries to correctly work the answer
		//this will ensure its done without any errors
		if paramOne + paramTwo == given {
			return nil //nil as its valid
		}

		//executes the branding piece
		//this will ensure its done without any errors
		err = language.ExecuteLanguage([]string{"views", "captcha", "incorrect.itl"}, s.Channel, deployment.Engine, s, map[string]string{"attempts":strconv.Itoa(toml.CatpchaToml.MaximumAttempts-attempt)})
		if err != nil { //error handles properly
			return err //returns error properly
		}

		//continues looping properly
		//this will ensure its done without any errors
		continue
	}

	//returns the max attempts error
	//this will ensure its done without any errors
	return language.ExecuteLanguage([]string{"views","captcha", "incorrect_max.itl"}, s.Channel, deployment.Engine, s, make(map[string]string))
}