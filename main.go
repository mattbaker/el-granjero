package main

import (
	"fmt"
	"os"
	"time"
)

const membershipType = "2"
const timeParseTemplate = "2006-01-02T15:04:05Z"

var destinyAPIClient = NewAPIClient("Destiny", os.Getenv("DESTINY_KEY"))
var userAPIClient = NewAPIClient("User", os.Getenv("DESTINY_KEY"))

func findPsnMember(username string) (*Profile, error) {
	endPoint := fmt.Sprintf("/SearchDestinyPlayer/%s/%s/", membershipType, username)
	response := ProfileResponse{}
	destinyAPIClient.RequestAsJSON(endPoint, &response)
	if len(response.Response) == 0 {
		return nil, fmt.Errorf("0 search results returned for %s", username)
	}
	return &response.Response[0], nil
}

func getCharacterInventory(membershipID string, characterID string) {
	endPoint := fmt.Sprintf("/%s/Account/%s/Character/%s/Inventory/Summary/?definitions=true", membershipType, membershipID, characterID)
	responseString := destinyAPIClient.Request(endPoint)
	fmt.Println(responseString)
}

func getAccountSummary(membershipID string) ([]Character, error) {
	endPoint := fmt.Sprintf("/%s/Account/%s/Summary/", membershipType, membershipID)
	response := SummaryResponse{}
	destinyAPIClient.RequestAsJSON(endPoint, &response)
	if len(response.Response.Data.Characters) == 0 {
		return nil, fmt.Errorf("0 characters returned for membership ID %s", membershipID)
	}
	return response.Response.Data.Characters, nil
}

func getMostRecentCharacterID(characters []Character) (characterID string) {
	var mostRecentTimestamp int64
	var parsedTime time.Time
	var mostRecentCharacterID string

	for i := 0; i < len(characters); i++ {
		parsedTime, _ = time.Parse(timeParseTemplate, characters[i].CharacterBase.DateLastPlayed)
		if parsedTime.Unix() > mostRecentTimestamp {
			mostRecentTimestamp = parsedTime.Unix()
			mostRecentCharacterID = characters[i].CharacterBase.CharacterID
		}
	}
	return mostRecentCharacterID
}

func main() {
	profile, _ := findPsnMember("malhuevo")
	characters, _ := getAccountSummary(profile.MembershipID)
	mostRecentCharacterID := getMostRecentCharacterID(characters)
	getCharacterInventory(profile.MembershipID, mostRecentCharacterID)
}
