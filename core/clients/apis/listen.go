package apis

import (
	"Nosviak2/core/sources/layouts/toml"
	"Nosviak2/core/sources/tools"
	"net/http"
)

//ListenAndServe will run the api
func ListenAndServe() error {

	//checks if the attack route is blocked inside api.toml
	if !tools.NeedleHaystackOne(toml.ApiToml.API.BlockRoute, "attack") {
		http.HandleFunc(toml.ApiToml.API.Path+"/attack", Attack) //attack launch feature
	}

	//checks if the ongoing route is blocked inside api.toml
	if !tools.NeedleHaystackOne(toml.ApiToml.API.BlockRoute, "ongoing") {
		http.HandleFunc(toml.ApiToml.API.Path+"/ongoing", Ongoing) //list ongoing attacks feature
	}
	
	//checks if the edition route is blocked inside api.toml
	if !tools.NeedleHaystackOne(toml.ApiToml.API.BlockRoute, "edition") {
		http.HandleFunc(toml.ApiToml.API.Path+"/edition", Edition)
	}

	//checks if the method route is blocked inside api.toml
	if !tools.NeedleHaystackOne(toml.ApiToml.API.BlockRoute, "method") {
		http.HandleFunc(toml.ApiToml.API.Path+"/method", Methods)
	}


	//TLS mode enabled on api
	if toml.ApiToml.TLS.TLS { //starts listening on the tls side
		return http.ListenAndServeTLS(toml.ApiToml.API.Host, toml.ApiToml.TLS.Certification, toml.ApiToml.TLS.Key, nil)
	} else { //non tls driver will start here
		return http.ListenAndServe(toml.ApiToml.API.Host, nil)
	}
}