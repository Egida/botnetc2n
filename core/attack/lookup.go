package attacks

import (
	goip "github.com/jpiontek/go-ip-api"
)

/*
	this system will use an iplookup system to verify and recommend methods
	for each target allowing for true moderation and management within the
	system

	this system will use the nordvpn database tool for it's lookup
*/



//Lookup will lookup the given ip address
func Lookup(ip string) (*goip.Location, error) {
	return goip.NewClient().GetLocationForIp(ip)
}