package api

import (
	"encoding/json"
)

type courseStats struct {
	Listeners 	int `json:"listeners"`
}

// Request Handler
func GetStats(_ string) (string, error) {
	stats := map[string]map[string]courseStats{}

	for key, val := range fetchList {
		if key.IsEmpty() {
			continue
		}

		subject := key.Subject
		catalogNumber := key.CatalogNumber
		listeners := val

		if stats[subject] == nil {
			stats[subject] = map[string]courseStats{}
		}

		stats[subject][catalogNumber] = courseStats{Listeners:listeners}
	}

	jsonBody, err := json.Marshal(stats)

	return string(jsonBody), err
}

