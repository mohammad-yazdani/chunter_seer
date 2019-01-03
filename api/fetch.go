package api

import (
	"chunter_seer/shared"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getCourseSchedule(query string) Fetch {
	var jsonCourseSchedule []byte

	data, err := http.Get(query)
	if err != nil {
		shared.LOG(err.Error())
	}

	jsonCourseSchedule, err = ioutil.ReadAll(data.Body)

	var fetched Fetch
	err = json.Unmarshal(jsonCourseSchedule, &fetched)
	if err != nil {
		shared.LOG(err.Error())
	}
	return fetched
}
