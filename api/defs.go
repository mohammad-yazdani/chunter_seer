package api

import "strings"

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

func SetApiKey(key string) {
	apiKey = key
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


