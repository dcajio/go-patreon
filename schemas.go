package main

type UserAttributes struct {
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	FullName    string            `json:"full_name"`
	Gender      string            `json:"gender"`
	Vanity      string            `json:"vanity"`
	About       string            `json:"about"`
	FacebookID  string            `json:"facebook_id"`
	ImageURL    string            `json:"image_url"`
	ThumbURL    string            `json:"thumb_url"`
	Thumbnails  map[string]string `json:"thumbnails"`
	Youtube     string            `json:"youtube"`
	Twitter     string            `json:"twitter"`
	Facebook    string            `json:"facebook"`
	IsSuspended bool              `json:"is_suspended"`
	IsDeleted   bool              `json:"is_deleted"`
	IsNuked     bool              `json:"is_nuked"`
	Created     int               `json:"created"`
	URL         string            `json:"url"`
}

type User struct {
	Type          string         `json:"type"`
	ID            string         `json:"id"`
	Attributes    UserAttributes `json:"attributes"`
	Relationships struct {
		Campaign Campaign `json:"campaign"`
	} `json:"relationships"`
}

type CampaignAttributes struct {
	Summary                       string `json:"summary"`
	CreationName                  string `json:"creation_name"`
	PayPerName                    string `json:"pay_per_name"`
	OneLiner                      string `json:"one_liner"`
	MainVideoEmbed                string `json:"main_video_embed"`
	MainVideoURL                  string `json:"main_video_url"`
	ImageSmallURL                 string `json:"image_small_url"`
	ImageURL                      string `json:"image_url"`
	ThanksVideoURL                string `json:"thanks_video_url"`
	ThanksEmbed                   string `json:"thanks_embed"`
	ThanksMsg                     string `json:"thanks_msg"`
	IsMonthly                     bool   `json:"is_monthly"`
	IsNSFW                        bool   `json:"is_nsfw"`
	CreatedAt                     int    `json:"created_at"`
	PublishedAt                   int    `json:"published_at"`
	PledgeURL                     string `json:"pledge_url"`
	PledgeSum                     int    `json:"pledge_sum"`
	PatronCount                   int    `json:"patron_count"`
	CreationCount                 int    `json:"creation_count"`
	OutstandingPaymentAmountCents int    `json:"outstanding_payment_amount_cents"`
}

type Campaign struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Attributes CampaignAttributes `json:"attributes"`
}

type PledgeAttributes struct {
	AmountCents    int  `json:"amount_cents"`
	CreatedAt      int  `json:"created_at"`
	PledgeCapCents int  `json:"pledge_cap_Cents"`
	PatronPaysFees bool `json:"patron_pays_ffes"`
}

type Pledge struct {
	Type       string           `json:"type"`
	ID         string           `json:id"`
	Attributes PledgeAttributes `json:"attributes"`
}
