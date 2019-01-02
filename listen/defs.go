package listen

import "chunter_seer/api"

type courseStats struct {
	Listeners 	int `json:"listeners"`
}

func genStats() map[string]map[string]courseStats {
	stats := map[string]map[string]courseStats{}

	for key, val := range *api.GetFetchMap() {
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

	return stats
}
