package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateRunning
	StateStopped
	StateRetrying
)

var stateNames = map[ServerState]string{
	StateIdle:     "Idle",
	StateRunning:  "Running",
	StateStopped:  "Stopped",
	StateRetrying: "Retrying",
}

func main() {
	var mystate ServerState
	fmt.Println("Server State:", verificate(mystate))

	mystate = StateRunning

	fmt.Println("Server State:", verificate(mystate))

	mystate = StateStopped
	fmt.Println("Server State:", verificate(mystate))

	fmt.Println("Server State:", verificate(1))
	fmt.Println("Server State:", verificate(2))
	fmt.Println("Server State:", verificate(99))

	fmt.Println("--------------")
	fmt.Println("Server State:", verificate(mystate))
	updateState(&mystate)
	fmt.Println("Server State:", verificate(mystate))

}

func updateState(s *ServerState) {
	switch *s {
	case StateIdle:
		*s = StateRunning
	case StateRunning, StateRetrying:
		*s = StateStopped
	case StateStopped:
		*s = StateRetrying
	default:
		fmt.Println("Unknown state")
	}

}

func verificate(s ServerState) string {
	if name, ok := stateNames[s]; ok {
		// fmt.Println("State found:", name)
		// fmt.Println("State found:", ok)
		return name
	}
	return "Unknown"
}
