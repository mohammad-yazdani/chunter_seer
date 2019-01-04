package api

import (
	"chunter_seer/listen"
	"chunter_seer/store"
	"strings"
)

const uwBaseApiV2Courses = "https://api.uwaterloo.ca/v2/courses"

var apiKey string

type MetaMethod struct {
	Disclaimer  string `json:"disclaimer"`
	License 	string `json:"license"`
}

type FetchMeta struct {
	Requests  int `json:"requests"`
	Timestamp int `json:"timestamp"`
	Status    int `json:"status"`
	Message   string `json:"message"`
	MethodId  int `json:"method_id"`
	Version   string `json:"version"`
	Method    MetaMethod `json:"method"`
}

type Fetch struct {
	Meta FetchMeta `json:"meta"`
	Data []CourseSchedule `json:"data"`
}

func SetUpApi(key string) {
	apiKey = key
	fetchListMutex.Lock()
	fetchList = make(map[CourseCatalog]int, 0)
	fromDb := store.GetCourses()
	for _, course := range fromDb {
		catalog := strings.Split(course, " ")
		fetchList[CourseCatalog{Subject:catalog[0], CatalogNumber:catalog[1]}] = 0
	}
	fetchListMutex.Unlock()

	listen.AddHandler("add_course", AddToFetchList)
	listen.AddHandler("stats", GetStats)
}

func formQuery(subQueries ...string) string {
	var fullQuery strings.Builder

	fullQuery.WriteString(uwBaseApiV2Courses)
	for _, subQuery := range subQueries {
		fullQuery.WriteString("/")
		fullQuery.WriteString(subQuery)
	}

	return fullQuery.String()
}

func addUriArgs(query string, args map[string]string) string {
	var fullQuery strings.Builder

	fullQuery.WriteString(query)

	fullQuery.WriteString("?")
	for arg, val := range args {
		fullQuery.WriteString(arg)
		fullQuery.WriteString("=")
		fullQuery.WriteString(val)
		fullQuery.WriteString("&")
	}

	fullQuery.WriteString("key=")
	fullQuery.WriteString(apiKey)

	return fullQuery.String()
}


