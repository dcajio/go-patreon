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

type Patreon struct {
	AccessKey string
}

// Fetches the current logged-in user
func (p Patreon) GetCurrentUser() (error, User) {
	url := API_URL + "current_user/"
}

func (p Patreon) GetCampaigns() (err, []Campaign) {
}

func (p Patreon) GetCampaign(ID int) (err, Campaign) {
}

func (p Patreon) GetPledgesByCampaign(Id int) (err, []Pledge) {
}
