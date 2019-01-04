package api

import (
	"chunter_seer/shared"
	"encoding/json"
	"strconv"
)

type Stats map[string]map[string]string

// Request Handler
func GetStats(_ string) (string, error) {

	shared.LOG("STATS QUERY")

	stats := make(Stats)

	for key, val := range fetchList {
		if key.IsEmpty() {
			continue
		}

		subject := key.Subject
		catalogNumber := key.CatalogNumber
		listeners := val

		if stats[subject] == nil {
			stats[subject] = map[string]string{}
		}

		stats[subject][catalogNumber] = strconv.FormatInt(int64(listeners), 10)
	}

	jsonBody, err := json.Marshal(stats)
	jsonString := string(jsonBody)

	return jsonString, err
}

