package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("DESTINY_KEY")
	endpointURL := "http://www.bungie.net/Platform/Destiny/Explorer/Items/"
	client := &http.Client{}

	req, _ := http.NewRequest("GET", endpointURL, nil)
	req.Header["X-API-Key"] = []string{key}
	fmt.Print(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to make request")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
