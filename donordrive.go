package donordrive

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)
const DefautlBaseUrl  = "https://donordrive.com/"
const ExtraLifeUrl    = "https://www.extra-life.org"
const TryBaseUrl = "https://try.donordrive.com/"

var baseUrl = DefautlBaseUrl
var client = &http.Client{}
const (
	apiEvents = "api/events"
	apiEventsTeam = "api/events/%d/teams"
	apiTeamParticipants = "api/teams/%d/participants"
	apiParticipantDetails = "api/participants/%d"
	apiParticipantDonations = "api/participants/%d/donations"
	apiParticipantsTopDonor = "api/participants/%d/donors"
)

func GetBaseUrl() string {
	return baseUrl
}

func SetBaseUrl(url string) {
	baseUrl = strings.TrimSpace(url)

	if strings.HasSuffix(baseUrl, "/") {
		return
	}

	baseUrl += "/"
}

func GetEvents() ([]Event, error){
	res, err := client.Get(fmt.Sprintf("%s%s", baseUrl, apiEvents))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%d returned", res.StatusCode))
	}

	var results []Event

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func GetTeamParticipants(team int) ([]Participant, error) {
	teamPath := fmt.Sprintf(apiTeamParticipants, team)
	res, err := client.Get(fmt.Sprintf("%s%s", baseUrl, teamPath))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%d returned", res.StatusCode))
	}

	var results []Participant

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&results); err != nil {
		return nil, err
	}
	return results, nil
}