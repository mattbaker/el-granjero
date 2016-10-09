package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const membershipType = "2"

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
	res := InventoryResponse{}
	json.Unmarshal([]byte(responseString), &res)
	if len(res.Response.Data.Items) == 0 {
		return nil, fmt.Errorf("inventory not found for membership ID %s", membershipID)
	}
	return res.Response.Data.Items, nil
}

func getCharacterInventory(membershipID string, characterID string) {
	endPoint := fmt.Sprintf("https://www.bungie.net/Platform/Destiny/%s/Account/%s/Character/%s/Inventory/", membershipType, membershipID, characterID)
	responseString := apiRequest(endPoint)
	fmt.Println(responseString)
}

func main() {
	profile, _ := findPsnMember("malhuevo")
	// inventory, _ := getFullInventory(profile.MembershipID)
	getCharacterInventory(profile.MembershipID, "2305843009323254133")

	// fmt.Printf("%+v\n", inventory)
}
