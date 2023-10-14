package animations

import (
	"Nosviak2/core/sources/layouts/toml"
	"errors"
	"sync"
	"time"
)

//gets that current spinners frame
//this will ensure its done without issues happening
func AccessCurrentFrame(spinner string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	//tries to get the current spinner
	//this will ensure the frame is gotton without issues
	if core, ok := SpinnersPool[spinner]; !ok { //checks if there
		return "", errors.New("spinner could not be located") //error handles
	} else { //returns the frame properly
		//this will ensure its done without issues happening on reqeust
		return core.Frames[core.Position], nil
	}
}

//stores the spinners config properly
//this will allow for proper control within the system
type SpinnerConfig struct {
	Frames []string //stores all frames
	Position int //stores the current position
}

var (
	//stores all the spinners and there current frame
	//this will ensure its done without issues happening on request
	SpinnersPool map[string]*SpinnerConfig = make(map[string]*SpinnerConfig) //stored in array
	mutex sync.Mutex //mutex for the function correctly without issues happening
)

//starts the spinner runtime properly
//this will ensure its done without issues
func WorkersRuntime() { //returns no values properly

	//ranges through all the toml spinners
	//this will create the instance properly and safely
	for tag, startup := range toml.Spinners.Spins { //ranges through
		mutex.Lock() //locks the mutex properly
		SpinnersPool[tag] = &SpinnerConfig{Frames: startup.Frames, Position: 0}
		mutex.Unlock() //unlocks the mutex properly
	}

	for { //loops independently properly

		//ranges through all spinners active properly
		//this will ensure information is properly gathered
		for _, CurrentPoolItem := range SpinnersPool { //ranges through

			//fmt.Println(pool, CurrentPoolItem.Position + 1 >= len(CurrentPoolItem.Frames))
			//checks the current positions length within the system
			//this will allow for the frame current to shift across properly
			if CurrentPoolItem.Position + 1 >= len(CurrentPoolItem.Frames) {
				CurrentPoolItem.Position = 0; continue //resets position properly
			}

			//enforces the frame shift
			//this will display the next frame on call
			CurrentPoolItem.Position++ //shifts frame properly
		}

		//sleeps for the duration until next frame is called onwards
		time.Sleep(time.Duration(toml.ConfigurationToml.TitleWorker.Duration-20) * time.Millisecond)
	}
}