package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return []Events{}, fmt.Errorf("error reading response: %w", err)
	// }

	var events Events
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&events); err != nil {
		return Events{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	// if err := json.Unmarshal(body, &events); err != nil {
	// 	return []Events{}, fmt.Errorf("error unmarshalling response: %w", err)
	// }

	return events, nil
}
