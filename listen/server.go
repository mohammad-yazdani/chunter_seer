package listen

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func requestDispatcher(w http.ResponseWriter, r *http.Request)  {
	jsonBody, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	request := make(Request)
	err = json.Unmarshal(jsonBody, &request)
	if err != nil {
		log.Fatal(err)
	}

	response := handleRequest(request)

	jsonBody, err = json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonBody)
	if err != nil {
		log.Fatal(err)
	}
}

func Start()  {
	http.HandleFunc("/", requestDispatcher)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
