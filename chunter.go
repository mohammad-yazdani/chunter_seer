package main

import (
	"bufio"
	"chunter_seer/api"
	"chunter_seer/listen"
	"chunter_seer/notif"
	"chunter_seer/sched"
	"log"
	"os"
)

func setup(configArray []string) {
	api.SetUpApi(configArray[0])
	sched.SetUpScheduler()
	notif.SetUpMail(configArray[1], configArray[2], configArray[3])
}

func main() {
	keyFile, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer keyFile.Close()

	scanner := bufio.NewScanner(keyFile)

	configArray := make([]string, 0)
	for scanner.Scan() {
		config := scanner.Text()
		configArray = append(configArray, config)
	}

	setup(configArray)

	subject := "CS"
	catalogNumber := "450"

	api.AddToFetchList(api.CourseCatalog{Subject:subject, CatalogNumber:catalogNumber})

	go sched.PollEndpoint(5)
	listen.Start()
}
