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

func getFullInventory(membershipID string) ([]InventoryItem, error) {
	endPoint := fmt.Sprintf("/%s/Account/%s/Items/", membershipType, membershipID)
	response := InventoryResponse{}
	apiClient.RequestAsJSON(endPoint, &response)
	return response.Response.Data.Items, nil
}

// func getCharacterInventory(membershipID string, characterID string) {
// 	endPoint := fmt.Sprintf("/%s/Account/%s/Character/%s/Inventory/", membershipType, membershipID, characterID)
// 	response := apiClient.RequestAsJson(endPoint, &InventoryResponse{}).(InventoryResponse)
// 	return response.Response.Data.Items, nil
// }

func main() {
	profile, _ := findPsnMember("malhuevo")
	fmt.Printf("%+v\n", profile)
	inventory, _ := getFullInventory(profile.MembershipID)
	fmt.Printf("%+v\n", inventory)
}
