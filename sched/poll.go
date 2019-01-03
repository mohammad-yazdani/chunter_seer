package sched

import (
	"chunter_seer/api"
	"time"
)

func PollEndpoint(interval int)  {
	for t := range time.NewTicker(time.Duration(interval) * time.Second).C {
		fetchCourses(t)
	}
}

func fetchCourses(_ time.Time) {
	catalogs := api.GetFetchList()
	for _, catalog := range catalogs {
		if catalog.IsEmpty() {
			continue
		}
		schedule := api.CourseScheduleQuery(catalog)
		hasChanged(schedule)
	}
}
