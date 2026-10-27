package main

import (
	GitApi "GitHubTracker/GitAPI"
	"fmt"
	"os"
)

func main() {
	// Retrieves the username from the first command-line argument and inserts it into the URL.
	username := os.Args[1]
	url := "https://api.github.com/users/" + username + "/events"

	// Checks whether there are any recent activities. If there are none, prints No recent events.
	counted := GitApi.CountEvents(url)
	if counted == nil {
		fmt.Println("No recent events")
		return
	}

	// Runs the event handler.
	GitApi.EventsHandler(url)
}
