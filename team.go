package donordrive

type Team struct {
	AvatarImageURL     string  `json:"avatarImageURL"`
	CaptainDisplayName string  `json:"captainDisplayName"`
	CreatedDateUTC     string  `json:"createdDateUTC"`
	EventID            int     `json:"eventID"`
	EventName          string  `json:"eventName"`
	FundraisingGoal    string  `json:"fundraisingGoal"`
	IsInviteOnly       bool    `json:"isInviteOnly"`
	Links              *Links  `json:"links"`
	Name               string  `json:"name"`
	NumAwardedBadges   int     `json:"numAwardedBadges"`
	NumDonations       int     `json:"numDonations"`
	NumParticipants    int     `json:"numParticipants"`
	StreamIsLive       bool    `json:"streamIsLive"`
	SumDonations       float32 `json:"sumDonations"`
	SumPledges         float32 `json:"sumPledges"`
	TeamID             int     `json:"teamID"`
}

type Links struct {
	Donate string `json:"donate"`
	Page   string `json:"page"`
	Stream string `json:"stream"`
}
