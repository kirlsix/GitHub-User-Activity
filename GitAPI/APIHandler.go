package GitApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

var User string

func GetEvents() []Event {
	var result []Event
	resp, err := http.Get("https://api.github.com/users/sindresorhus/events?per_page=100")

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Fail to get answer from GitHub. Status:", resp.Status)
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Fail to decode response", err)
		return nil
	}

	if len(result) == 0 {
		return nil
	}
	User = result[0].Actor.Login
	return result
}
