package main

import (
	"fmt"
	"os"
)

const membershipType = "2"

var apiClient = NewAPIClient(os.Getenv("DESTINY_KEY"))

func findPsnMember(username string) (*Profile, error) {
	endPoint := fmt.Sprintf("/SearchDestinyPlayer/%s/%s/", membershipType, username)
	response := ProfileResponse{}
	apiClient.RequestAsJSON(endPoint, &response)
	if len(response.Response) == 0 {
		return nil, fmt.Errorf("0 search results returned for %s", username)
	}
	return &response.Response[0], nil
}

func getCharacterInventory(membershipID string, characterID string) {
	endPoint := fmt.Sprintf("/%s/Account/%s/Character/%s/Inventory/", membershipType, membershipID, characterID)
	responseString := apiClient.Request(endPoint)
	fmt.Println(responseString)
}

func getAccountSummary(membershipID string) ([]Character, error) {
	endPoint := fmt.Sprintf("/%s/Account/%s/Summary/", membershipType, membershipID)
	// fmt.Println(responseString)
	response := SummaryResponse{}
	apiClient.RequestAsJSON(endPoint, &response)
	if len(response.Response.Data.Characters) == 0 {
		return nil, fmt.Errorf("0 characters returned for membership ID %s", membershipID)
	}
	return response.Response.Data.Characters, nil
}

func main() {
	profile, _ := findPsnMember("malhuevo")
	fmt.Printf("%+v\n", profile)
	characters, _ := getAccountSummary(profile.MembershipID)
	fmt.Printf("%+v\n", characters)
}
