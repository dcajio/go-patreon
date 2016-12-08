package main

import (
	"net/http"
	"strings"
)

func main() {
}

const (
	API_URL string = "https://api.patreon.com/oauth2/api/"
)

var (
	AccessKey string
)

// Fetches the current logged-in user
func GetCurrentUser(inc ...string) (error, User) {
	url := API_URL + "current_user/"

	if len(inc) > 0 {
		include_string := strings.Join(inc, ",")
		url = url + "?include=" + include_string
	}

	return nil, User{}
}

func GetCampaigns(inc ...string) (error, []Campaign) {
	url := API_URL + "current_user/campaigns"

	if len(inc) > 0 {
		include_string := strings.Join(inc, ",")
		url = url + "?include=" + include_string
	}
}

func GetCampaign(ID int) (error, Campaign) {
}

func GetPledgesByCampaign(ID int) (error, []Pledge) {
}
