package main

import (
	"net/http",
	"gitlab.com/dcajio/patreon-go/schemas"
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
func GetCurrentUser() (error, User) {
	url := API_URL + "current_user/"
}

func GetCampaigns() (err, []Campaign) {
}

func GetCampaign(ID int) (err, Campaign) {
}

func GetPledgesByCampaign(ID int) (err, []Pledge) {
}
