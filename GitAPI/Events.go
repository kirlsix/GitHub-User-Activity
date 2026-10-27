package GitApi

import "fmt"

type RepoEvent struct {
	Type   string
	Name   string
	Action string
}

func count(eventType string, name string, action string) {
	key := RepoEvent{
		Type:   eventType,
		Name:   name,
		Action: action,
	}
	eventsCount[key] += 1
}

var eventsCount map[RepoEvent]int

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
