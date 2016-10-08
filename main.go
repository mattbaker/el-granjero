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

// Profile represents a profile
type Profile struct {
	IconPath       string
	MembershipType int
	MembershipID   string
	DisplayName    string
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
	res := ProfileResponse{}
	json.Unmarshal([]byte(responseString), &res)
	if len(res.Response) == 0 {
		return nil, fmt.Errorf("0 search results returned for %s", username)
	}
	return &res.Response[0], nil
}

func getCharacters() {

}

func getInventory() {
	// http://www.bungie.net/platform/Destiny/[membershipType]/Account/[d]/Character/[g]/Inventory/
}

func main() {
	profile, _ := findPsnMember("swigwillis")
	fmt.Printf("%+v\n", profile)
}
