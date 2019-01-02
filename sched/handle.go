package sched

import (
	"chunter_seer/api"
	"chunter_seer/notif"
)

type EnrollStats struct {
	Capacity int
	Total int
}

type EnrollChange struct {
	Course api.CourseCatalog
	Change int
}

var courseStats map[api.CourseCatalog]EnrollStats

func SetUpScheduler()  {
	courseStats = make(map[api.CourseCatalog]EnrollStats, 0)
}

func handleChange(change EnrollChange) {
	notif.MailChange(change.Course.Subject + change.Course.CatalogNumber, change.Change)
}

// TODO : Thread runnable
// TODO : Make atomic
func hasChanged(schedules []api.CourseSchedule) {
	for _, schedule := range schedules {
		catalog := api.CourseCatalog{Subject:schedule.Subject,
			CatalogNumber:schedule.CatalogNumber, Section:schedule.Section}

		if stats, exists := courseStats[catalog]; exists {
			oldDiff := stats.Capacity - stats.Total
			newDiff := schedule.EnrollmentCapacity - schedule.EnrollmentTotal
			handleChange(EnrollChange{Course:catalog, Change:newDiff - oldDiff})
		} else {
			courseStats[catalog] = EnrollStats{Capacity:schedule.EnrollmentCapacity,
				Total:schedule.EnrollmentTotal}
			handleChange(EnrollChange{Course:catalog, Change:0})
		}
	}
}

