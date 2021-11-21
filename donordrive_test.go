package donordrive

import (
	"strings"
	"testing"
)

func TestSetBaseUrl(t *testing.T) {
	newUrl := "https://test.org"
	expected := "https://test.org/"
	SetBaseUrl(newUrl)

	result := GetBaseUrl()
	if strings.Compare(result,expected) != 0 {
		t.Fatalf("TestSetBaseUrl expected %s but got %s\n", expected, result)
	}

	newUrl2 := "https://test2.org/"
	expected2 := newUrl2
	SetBaseUrl(newUrl2)
	result = GetBaseUrl()
	if strings.Compare(result,expected2) != 0 {
		t.Fatalf("TestSetBaseUrl expected %s but got %s\n", expected2, result)
	}
}

func TestGetEvents(t *testing.T) {
	SetBaseUrl(ExtraLifeUrl)
	_, err := GetEvents()
	if err != nil {
		t.Fatalf("Failed to get resutls: %s", err)
	}
}

func TestGetTeamParticipants(t *testing.T) {
	team := 56078
	SetBaseUrl(ExtraLifeUrl)
	_, err := GetTeamParticipants(team)
	if err != nil {
		t.Fatalf("Failed to get results: %s", err)
	}
}