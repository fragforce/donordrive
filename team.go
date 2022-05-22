package donordrive

type Team struct {
	AvatarImageURL       *string  `json:"avatarImageURL,omitempty"`
	CaptainDisplayName   *string  `json:"captainDisplayName,omitempty"`
	CreatedDateUTC       *string  `json:"createdDateUTC,omitempty"`
	EventID              *int     `json:"eventID,omitempty"`
	EventName            *string  `json:"eventName,omitempty"`
	FundraisingGoal      *float64 `json:"fundraisingGoal,omitempty"`
	IsInviteOnly         *bool    `json:"isInviteOnly,omitempty"`
	Links                *Links   `json:"links,omitempty"`
	Name                 *string  `json:"name,omitempty"`
	NumAwardedBadges     *int     `json:"numAwardedBadges,omitempty"`
	NumDonations         *int     `json:"numDonations,omitempty"`
	NumParticipants      *int     `json:"numParticipants,omitempty"`
	StreamIsLive         *bool    `json:"streamIsLive,omitempty"`
	SumDonations         *float64 `json:"sumDonations,omitempty"`
	SumPledges           *float64 `json:"sumPledges,omitempty"`
	TeamID               *int     `json:"teamID,omitempty"`
	HasActivityTracking  *bool    `json:"hasActivityTracking,omitempty"`
	SourceTeamID         *int     `json:"sourceTeamID,omitempty"`
	HasTeamOnlyDonations *bool    `json:"hasTeamOnlyDonations,omitempty"`
	StreamIsEnabled      *bool    `json:"streamIsEnabled,omitempty"`
	StreamingPlatform    *string  `json:"streamingPlatform,omitempty"`
	StreamingChannel     *string  `json:"streamingChannel,omitempty"`
	IsCustomAvatarImage  *bool    `json:"isCustomAvatarImage,omitempty"`
}

type Links struct {
	Donate *string `json:"donate,omitempty"`
	Page   *string `json:"page,omitempty"`
	Stream *string `json:"stream,omitempty"`
}
