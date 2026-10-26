package GitApi

type RepoEvent struct {
	Type string
	Name string
}

func count(eventType string, name string) {
	key := RepoEvent{
		Type: eventType,
		Name: name,
	}
	eventsCount[key] += 1
}

var eventsCount map[RepoEvent]int

func CountEvents() map[RepoEvent]int {
	eventsCount = make(map[RepoEvent]int)
	events := GetEvents()
	if events == nil {
		return nil
	}
	for _, v := range events {
		count(v.Type, v.Repo.Name)
	}
	return eventsCount
}
