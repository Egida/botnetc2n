package fakes

import (
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/layouts/toml"
	"fmt"
)

//starts the fake slaves properly
//this will create all the routes within the system
func Start() { //ranges through all possible objects within safely
	for Header, Slave := range toml.FakeToml.FakeSlaves { //ranges


		//enabled?
		if Slave.Act {
			continue
		}


		if deployment.DebugMode { //little debug option within the system
			fmt.Printf("[FAKE_SLAVES].(%s): has been triggered and started properly\r\n", Header)
		} //makes the fake slave killer properly
		go MakeFakeSlaveWorker(Header, Slave) //creates header
	}
}