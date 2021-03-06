package main

import (
	"bufio"
	"chunter_seer/api"
	"chunter_seer/listen"
	"chunter_seer/notif"
	"chunter_seer/sched"
	"chunter_seer/shared"
	"chunter_seer/store"
	"os"
)

func setup(configArray []string) {
	shared.SetUpLog()
	store.SetUpDb()
	api.SetUpApi(configArray[0])
	sched.SetUpScheduler()
	notif.SetUpMail(configArray[1], configArray[2], configArray[3])
}

func main() {
	keyFile, err := os.Open("config.txt")
	if err != nil {
		shared.LOG(err.Error())
	}

	defer keyFile.Close()

	scanner := bufio.NewScanner(keyFile)

	configArray := make([]string, 0)
	for scanner.Scan() {
		config := scanner.Text()
		configArray = append(configArray, config)
	}

	setup(configArray)

	go sched.PollEndpoint(5)
	listen.Start()

	store.CloseDb()

	shared.LOG("END OF LOG")
	shared.CloseLog()
}
