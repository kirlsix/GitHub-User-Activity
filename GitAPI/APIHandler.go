package GitApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Event represents a GitHub event returned by the GitHub API.
// It contains information about the actor, event type, repository, and payload.
type Event struct {
	Actor struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
	} `json:"actor"`

	Type string `json:"type"`

	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`

	Payload struct {
		Action string `json:"action"`
	} `json:"payload"`
}

// A function that retrieves all events by making a GET request to the GitHub API. It takes a URL as an argument and returns a slice of Event structs.
func GetEvents(url string) []Event {
	var result []Event
	// Executing the GET request and decoding its response below, with checks for potential errors.
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Fail to get answer from GitHub. Status:", resp.Status)
		return nil
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Fail to decode response", err)
		return nil
	}
	// Checking the length of the result. If the length is 0, there are no recent events, so nil is returned.
	if len(result) == 0 {
		return nil
	}
	return result
}
