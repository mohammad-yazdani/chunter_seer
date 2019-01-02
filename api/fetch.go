package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getCourseSchedule(query string) Fetch {
	var jsonCourseSchedule []byte

	data, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}

	jsonCourseSchedule, err = ioutil.ReadAll(data.Body)

	jsonString := string(jsonCourseSchedule)
	println(jsonString)

	var fetched Fetch
	err = json.Unmarshal(jsonCourseSchedule, &fetched)
	if err != nil {
		log.Fatal(err)
	}
	return fetched
}
