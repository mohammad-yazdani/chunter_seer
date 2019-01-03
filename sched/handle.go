package sched

import "chunter_seer/api"

type EnrollStats struct {
	Capacity int
	Total int
}

type EnrollChange struct {
	course api.CourseCatalog
	change int
}

var courseStats map[api.CourseCatalog]EnrollStats

func SetUpScheduler()  {
	courseStats = make(map[api.CourseCatalog]EnrollStats, 0)
}

func handleChange(change EnrollChange) {
	if change.change != 0 {
		// TODO : Notify user
		println("Handling change ", change.course.Subject, change.course.CatalogNumber, change.change)
	} else {
		println("NO change ", change.course.Subject, change.course.CatalogNumber, change.change)
	}


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
			handleChange(EnrollChange{course:catalog, change:newDiff - oldDiff})
		} else {
			courseStats[catalog] = EnrollStats{Capacity:schedule.EnrollmentCapacity,
				Total:schedule.EnrollmentTotal}
			handleChange(EnrollChange{course:catalog, change:0})
		}
	}
}

