package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//TODO: check for 404s and other bad responses, send error

const server = "https://www.bungie.net/Platform"

// APIClient is a destiny-api specific wrapper around http.Client
type APIClient struct {
	HTTPClient *http.Client
	APIKey     string
	Service    string
}

// NewAPIClient creates a new APIClient with the given key
func NewAPIClient(service string, key string) *APIClient {
	var client = &http.Client{}
	return &APIClient{HTTPClient: client, APIKey: key, Service: service}
}

func (apiClient *APIClient) buildEndpointURI(endpoint string) string {
	return fmt.Sprintf("%s/%s%s", server, apiClient.Service, endpoint)
}

// Request makes an APIRequest to Destiny
func (apiClient *APIClient) Request(endpoint string) string {
	req, _ := http.NewRequest("GET", apiClient.buildEndpointURI(endpoint), nil)
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
