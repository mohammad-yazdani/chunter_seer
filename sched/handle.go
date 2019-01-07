package sched

import (
	"chunter_seer/api"
	"chunter_seer/listen"
	"chunter_seer/notif"
	"chunter_seer/shared"
	"chunter_seer/store"
	"strconv"
)

var forceFlushInterval = 4320

var courseStats map[int]store.EnrollStats
var forceFlushCounter int

func SetUpScheduler() {
	courseStats = make(map[int]store.EnrollStats, 0)

	fromDb := store.GetEnrollments()

	for _, course := range fromDb {
		courseStats[course.Class] = course
	}

	forceFlushCounter = 0

	listen.AddHandler("set_interval", SetFlushInterval)
}

func SetFlushInterval(interval string) (string, error)  {
	intInterval, err := strconv.Atoi(interval)
	if err != nil {
		return "COULD NOT PARSE", err
	}

	forceFlushCounter = intInterval

	return "FORCE FLUSH INTERVAL IS SET AT: " + strconv.FormatInt(int64(forceFlushCounter), 10), nil
}

func hasChanged(schedules []api.CourseSchedule) {
	if forceFlushCounter == forceFlushInterval {
		forceFlushCounter = 0
	}

	changeBatch := make([]notif.ChangeNotification, 0)
	processBatch := make([]store.EnrollStats, 0)
	oldDiffs := make([]int, 0)

	for _, schedule := range schedules {
		catalog := schedule.ClassNumber

		stat := store.EnrollStats{Class: schedule.ClassNumber, Subject: schedule.Subject, CatalogNumber: schedule.CatalogNumber,
			Section: schedule.Section, Total: schedule.EnrollmentTotal, Capacity: schedule.EnrollmentCapacity}

		if stats, exists := courseStats[catalog]; exists {
			oldDiff := stats.Capacity - stats.Total
			newDiff := schedule.EnrollmentCapacity - schedule.EnrollmentTotal

			if oldDiff != newDiff {
				course := schedule.Subject + " " + schedule.CatalogNumber + " " + schedule.Subject
				changeBatch = append(changeBatch,
					notif.ChangeNotification{
						Catalog:  course,
						Total:    schedule.EnrollmentTotal,
						Capacity: schedule.EnrollmentCapacity,
						Change:   newDiff - oldDiff})
			}
			oldDiffs = append(oldDiffs, oldDiff)
		} else {
			oldDiffs = append(oldDiffs, 0)
		}

		processBatch = append(processBatch, stat)
		store.SaveEnrollment(stat)
		courseStats[catalog] = stat
	}

	if forceFlushCounter == 0 {
		shared.LOG("FORCE FLUSH")

		mailFlush := make([]notif.ChangeNotification, 0)
		for diffIndex, c := range processBatch {
			catalog := c.Subject + " " + c.CatalogNumber + " " + c.Section
			change := notif.ChangeNotification{
				Catalog:  catalog,
				Total:    c.Total,
				Capacity: c.Capacity,
				Change:   (c.Capacity - c.Total) - (oldDiffs[diffIndex])}
			mailFlush = append(mailFlush, change)
		}

		notif.MailChange(mailFlush, true)
	}

	if len(changeBatch) > 0 && forceFlushCounter != 0 {
		notif.MailChange(changeBatch, false)
	}

	forceFlushCounter += 1
}
