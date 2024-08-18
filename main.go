package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tng/model"
	"tng/readtext"
)

// fetchJSONData retrieves the JSON data from the API endpoint
func fetchJSONData(url string, printToTerm bool) []model.Episode {
	// Create a custom request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	// Start getting the JSON
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Decode the body to an object
	var pluto model.Pluto
	err = json.Unmarshal(body, &pluto)
	if err != nil {
		panic(err)
	}

	// Get the episodes and print to terminal
	episodes := pluto.EPG[0].Episodes
	if printToTerm {
		prettyJSON, err := json.MarshalIndent(episodes, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(prettyJSON))
	}

	return episodes
}

func main() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	urls := readtext.OpenTextFile("url.txt")
	fetchJSONData(urls[0], true)
}
