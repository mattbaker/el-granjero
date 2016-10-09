package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// BaseResponse is a general API response from bungie
type BaseResponse struct {
	ErrorCode       int
	ThrottleSeconds int
	ErrorStatus     string
	Message         string
}

// ProfileResponse is a response to a profile request from bungie
type ProfileResponse struct {
	*BaseResponse
	Response []Profile
}

type InventoryResponse struct {
	*BaseResponse
	Response struct {
		Data struct {
			Items []InventoryItem
		}
	}
}

// Profile represents a profile
type Profile struct {
	IconPath       string
	MembershipType int
	MembershipID   string
	DisplayName    string
}

type InventoryItem struct {
	ItemHash       int
	ItemID         string
	Quantity       int
	TransferStatus int
	State          int
	BucketHash     int
}

var client = &http.Client{}
var apiKey = os.Getenv("DESTINY_KEY")

func apiRequest(endpoint string) string {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("X-API-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func findPsnMember(username string) (*Profile, error) {
	endPoint := fmt.Sprintf("https://www.bungie.net/Platform/Destiny/SearchDestinyPlayer/2/%s/", username)
	responseString := apiRequest(endPoint)
	// fmt.Printf("Member Response:\n%s\n", responseString)
	res := ProfileResponse{}
	json.Unmarshal([]byte(responseString), &res)
	if len(res.Response) == 0 {
		return nil, fmt.Errorf("0 search results returned for %s", username)
	}
	return &res.Response[0], nil
}

func getCharacters() {

}

func getFullInventory(membershipID string) ([]InventoryItem, error) {
	endPoint := fmt.Sprintf("https://www.bungie.net/Platform/Destiny/2/Account/%s/Items/", membershipID)
	responseString := apiRequest(endPoint)
	// fmt.Printf("Inventory Response:\n%s\n", responseString)
	res := InventoryResponse{}
	json.Unmarshal([]byte(responseString), &res)
	if len(res.Response.Data.Items) == 0 {
		return nil, fmt.Errorf("inventory not found for membership ID %s", membershipID)
	}
	return res.Response.Data.Items, nil
}

func getCharacterInventory() {
	// http://www.bungie.net/platform/Destiny/[membershipType]/Account/[d]/Character/[g]/Inventory/
}

func main() {
	profile, _ := findPsnMember("swigwillis")
	fmt.Printf("MembershipID: %s\n", profile.MembershipID)
	inventory, _ := getFullInventory(profile.MembershipID)
	fmt.Printf("%+v\n", inventory)
}
