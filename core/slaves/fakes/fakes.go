package fakes

import (
	"Nosviak2/core/configs/models"
	"math/rand"
	"sync"
	"time"
)

var (
	Mutex	   sync.Mutex
	FakeSlaves map[string]int = make(map[string]int)
)

//creates the fake slave worker
//this will run inside a goroutine properly
//creates a realish fake slave guide using time tickers and the configuration
func MakeFakeSlaveWorker(key string, config *models.FakeConfig) { //stores the body

	//randomizes the seed
	//allows for better protection from dups
	rand.Seed(time.Now().Unix()) //randomizes seed

	//stores the ticker information
	//this will allow for proper control without issues
	TickerMain := time.NewTicker(time.Duration(config.Ticks) * time.Millisecond) //stores the main ticker properly
	TickerSecondary := time.NewTicker(time.Duration(config.Ticks / 4) * time.Millisecond) //different by divide 4

	//stores the different counters properly
	//this will ensure its done without any errors
	CounterMain := rand.Intn(config.Max - config.Equ) + config.Equ //1st

	//for loops through the counter
	//this will ensure its done without any errors
	for { //loops through for every hopefully and safely

		//selects the channels properly for timer
		//this will allow for proper adjustments within
		select { //selects properly

		//main tickers channel
		//properly adjusts the system
		case <- TickerMain.C: //channel access

			//makes sure the counter isnt on 0
			//allows for proper controlling without issues
			if CounterMain > 0 { //checks that its not 0 properly
				CounterMain = rand.Intn(config.Max - config.Equ) + config.Equ //makes new counter
			} else if CounterMain < config.Min { //sets the new lowest properly
				CounterMain = rand.Intn(config.Max) //sets the random slaves properly
			}

			//creates the random object
			//this will ensure its done without issues
			r := rand.Intn(config.Ticks) //sleeps for the random duration
			time.Sleep(time.Duration(r) * time.Millisecond) //sleeps for the duration
		//secondary ticker channel
		//properly adjusts the count properly
		case <- TickerSecondary.C: //channel access
			Mutex.Lock()
			FakeSlaves[key] = CounterMain //sets the new number properly
			Mutex.Unlock()
		}
	}
}