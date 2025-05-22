package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Expecting a username")
		return
	}

	events, err := getEvents(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Activity in last 24 hours:")
	for _, event := range events {
		fmt.Printf("%v - %v %v\n",
			event.CreatedAt,
			event.Type,
			event.Repo.URL,
		)
	}
}
