package apis

import (
	deployment "Nosviak2/core/configs"
	"Nosviak2/core/sources/events"
	"Nosviak2/core/sources/layouts/toml"
	"fmt"
	"net/http"
)

//Edition will allow you to see the version & app name
func Edition(w http.ResponseWriter, r *http.Request) { //writes information about the system without issues
	//launchs the debug message properly
	events.DebugLaunch(events.Functioning, "API", "REQUEST", []string{r.RemoteAddr, r.URL.String(), r.URL.Path})

	//renders the information about the edition etc
	w.Write([]byte(fmt.Sprintf("Nosviak2 (%s) - %s", deployment.Version, toml.ConfigurationToml.AppSettings.AppName)))
}
