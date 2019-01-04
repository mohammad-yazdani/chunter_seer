package sched

import (
	"chunter_seer/api"
	"chunter_seer/notif"
	"chunter_seer/shared"
	"chunter_seer/store"
)

var forceFlushInterval = 360

var courseStats map[int]store.EnrollStats
var forceFlushCounter int

func SetUpScheduler()  {
	courseStats = make(map[int]store.EnrollStats, 0)

	fromDb := store.GetEnrollments()

	for _, course := range fromDb {
		courseStats[course.Class] = course
	}

	forceFlushCounter = 0
}

func hasChanged(schedules []api.CourseSchedule) {
	if forceFlushCounter == forceFlushInterval {
		forceFlushCounter = 0
	}

	changeBatch := make([]notif.ChangeNotification, 0)
	processBatch := make([]store.EnrollStats, 0)

	for _, schedule := range schedules {
		catalog := schedule.ClassNumber

		stat := store.EnrollStats{Class:schedule.ClassNumber, Subject:schedule.Subject, CatalogNumber:schedule.CatalogNumber,
			Section:schedule.Section, Total:schedule.EnrollmentTotal, Capacity:schedule.EnrollmentCapacity}

		if stats, exists := courseStats[catalog]; exists {
			oldDiff := stats.Capacity - stats.Total
			newDiff := schedule.EnrollmentCapacity - schedule.EnrollmentTotal

			if oldDiff != newDiff {
				course := schedule.Subject + " " + schedule.CatalogNumber + " " + schedule.Subject
				changeBatch = append(changeBatch, notif.ChangeNotification{Catalog:course, Change:newDiff - oldDiff})
			}
		} else {
			courseStats[catalog] = stat
		}

		processBatch = append(processBatch, stat)
	}

	if forceFlushCounter == 0 {
		shared.LOG("FORCE FLUSH")

		mailFlush := make([]notif.ChangeNotification, 0)
		for _, c := range processBatch {
			store.SaveEnrollment(c)
			catalog := c.Subject + " " + c.CatalogNumber + " " + c.Section
			change := notif.ChangeNotification{Catalog:catalog, Change: c.Capacity - c.Total}
			mailFlush = append(mailFlush, change)
		}

		notif.MailChange(changeBatch)
	}

	if len(changeBatch) > 0 && forceFlushCounter != 0 {
		notif.MailChange(changeBatch)
	}

	forceFlushCounter += 1
}

