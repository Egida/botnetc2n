package events

import (
	deployment "Nosviak2/core/configs"
	"log"
	"strings"
)

//DebugLaunch will check for debug mode and log error
func DebugLaunch(code int, event string, subevent string, args []string) {
	if deployment.DebugMode { //checks for debug mode
		log.Printf("[DEBUG] [%d] [%s] [%s] ["+strings.Join(args, "] [")+"]\r\n", code, event, subevent)
	}
}