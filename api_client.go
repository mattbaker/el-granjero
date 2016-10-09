package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//TODO: check for 404s and other bad responses, send error

const server = "https://www.bungie.net/Platform/Destiny"

// APIClient is a destiny-api specific wrapper around http.Client
type APIClient struct {
	HTTPClient *http.Client
	APIKey     string
}

// NewAPIClient creates a new APIClient with the given key
func NewAPIClient(key string) *APIClient {
	var client = &http.Client{}
	return &APIClient{HTTPClient: client, APIKey: key}
}

func buildEndpointURI(endpoint string) string {
	return server + endpoint
}

// Request makes an APIRequest to Destiny
func (apiClient *APIClient) Request(endpoint string) string {
	req, _ := http.NewRequest("GET", buildEndpointURI(endpoint), nil)
	req.Header.Add("X-API-Key", apiClient.APIKey)

	resp, err := apiClient.HTTPClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

// RequestAsJSON makes a request and unmarshals the result into responseStruct
func (apiClient *APIClient) RequestAsJSON(endpoint string, responseStruct interface{}) {
	jsonString := apiClient.Request(endpoint)
	json.Unmarshal([]byte(jsonString), responseStruct)
}
