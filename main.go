package main

import (
	GitApi "GitHubTracker/GitAPI"
	"fmt"
)

func main() {
	counted := GitApi.CountEvents()
	if counted == nil {
		fmt.Println("No recent events")
	}
	fmt.Println(counted)
}
