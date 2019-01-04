package api

import (
	"chunter_seer/shared"
	"chunter_seer/store"
	"encoding/json"
	"strings"
	"sync"
)

type CourseCatalog struct {
	Subject       	string
	CatalogNumber 	string
	Section 		string
}

func (c * CourseCatalog) IsEmpty() bool {
	return c.Subject == "" || c.CatalogNumber == ""
}

type CourseReserves struct {
	ReserveGroup       string `json:"reserve_group"`
	EnrollmentCapacity int `json:"enrollment_capacity"`
	EnrollmentTotal    int `json:"enrollment_total"`
}

type ClassDate struct {
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Weekdays    string `json:"weekdays"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsTba       bool `json:"is_tba"`
	IsCancelled bool `json:"is_cancelled"`
	IsClosed    bool `json:"is_closed"`
}

type ClassLocation struct {
	Building string `json:"building"`
	Room     string `json:"room"`
}

type ClassSchedule struct {
	Date 		ClassDate `json:"date"`
	Location 	ClassLocation `json:"location"`
	Instructors []string `json:"instructors"`
}

type CourseSchedule struct {
	Subject            string `json:"subject"`
	CatalogNumber      string `json:"catalog_number"`
	Units              float32 `json:"units"`
	Title              string `json:"title"`
	Note               string `json:"note"`
	ClassNumber        int `json:"class_number"`
	Section            string `json:"section"`
	Campus             string `json:"campus"`
	AssociatedClass    int `json:"associated_class"`
	RelatedComponent1  string `json:"related_component_1"`
	RelatedComponent2  string `json:"related_component_2"`
	EnrollmentCapacity int `json:"enrollment_capacity"`
	EnrollmentTotal    int `json:"enrollment_total"`
	WaitingCapacity    int `json:"waiting_capacity"`
	WaitingTotal       int `json:"waiting_total"`
	Topic              string `json:"topic"`
	Reserves           []CourseReserves `json:"reserves"`
	Classes            []ClassSchedule `json:"classes"`
	HeldWith           []string `json:"held_with"`
	Term               int `json:"term"`
	AcademicLevel      string `json:"academic_level"`
	LastUpdated        string `json:"last_updated"`
}

var fetchList map[CourseCatalog]int
var fetchListMutex = &sync.Mutex{}

func AddToFetchList(catalog string) (string, error) {

	trim := strings.Split(catalog, " ")

	c := CourseCatalog{Subject:trim[0], CatalogNumber:trim[1], Section:""}

	fetchListMutex.Lock()
	fetchList[c] += 1
	store.AddCourse(c.Subject + " " + c.CatalogNumber)
	fetchListMutex.Unlock()

	return "OK", nil
}

func GetFetchList() []CourseCatalog {
	catalogs := make([]CourseCatalog, len(fetchList))

	fetchListMutex.Lock()
	for key := range fetchList {
		catalogs = append(catalogs, key)
	}
	fetchListMutex.Unlock()

	return catalogs
}

func (c * CourseSchedule) ToString () string {
	str, err := json.Marshal(c)
	if err != nil {
		shared.LOG(err.Error())
	}
	return string(str)
}

func CourseScheduleQuery(catalog CourseCatalog) []CourseSchedule {
	query := formQuery(catalog.Subject, catalog.CatalogNumber, "schedule.json")
	var argMap map[string]string
	query = addUriArgs(query, argMap)
	fetched := getCourseSchedule(query)

	return fetched.Data
}
