package main

import (
	GitApi "GitHubTracker/GitAPI"
	"fmt"
	"os"
)

func main() {
	username := os.Args[1]
	url := "https://api.github.com/users/" + username + "/events"

	counted := GitApi.CountEvents(url)
	if counted == nil {
		fmt.Println("No recent events")
		return
	}
	GitApi.EventsHandler(url)
}
