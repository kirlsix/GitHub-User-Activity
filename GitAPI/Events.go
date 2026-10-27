package GitApi

import "fmt"

// RepoEvent is used as a map key to count the number of events
// of each event type performed on each repository.
type RepoEvent struct {
	Type   string
	Name   string
	Action string
}

// A local function that counts the number of events of each type per repository and stores the results in a map.
func count(eventType string, name string, action string) {
	key := RepoEvent{
		Type:   eventType,
		Name:   name,
		Action: action,
	}
	eventsCount[key] += 1
}

var eventsCount map[RepoEvent]int

// A function that retrieves all events using GetEvents(),
// counts the number of events of each type per repository,
// and returns a map where the key is a RepoEvent struct and the value is the event count.
func CountEvents(url string) map[RepoEvent]int {
	eventsCount = make(map[RepoEvent]int)
	events := GetEvents(url)
	if events == nil {
		return nil
	}
	for _, v := range events {
		count(v.Type, v.Repo.Name, v.Payload.Action)
	}
	return eventsCount
}

// A function that processes the map returned by CountEvents(),
// checks the event type, and prints the number of events of that type for each repository.
// A function that processes the map returned by CountEvents(), checks the event type, and prints the number of events of that type for each repository
// Additional event types can be handled by extending the switch statement.
func EventsHandler(url string) {
	CountedEvents := CountEvents(url)
	for k, v := range CountedEvents {
		switch {
		case k.Type == "PushEvent":
			fmt.Println("Pushed", v, "commits to", k.Name)
		case k.Type == "IssuesEvent":
			if k.Action == "opened" {
				fmt.Println("Opened", v, "issues in", k.Name)
			} else if k.Action == "closed" {
				fmt.Println("Closed", v, "issues in", k.Name)
			} else if k.Action == "reopened" {
				fmt.Println("Reopened", v, "issues in", k.Name)
			}
		case k.Type == "WatchEvent":
			fmt.Println("Starred", k.Name)
		case k.Type == "CreateEvent":
			fmt.Println("Created", k.Name)
		}

	}
}
