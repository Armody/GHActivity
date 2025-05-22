package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getEvents(userName string) (Events, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", userName)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Events{}, fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Events{}, fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		fmt.Println("invalid username")
		os.Exit(0)
	}

	var events Events
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&events); err != nil {
		return Events{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return events, nil
}
