package main

import (
	"bufio"
	"chunter_seer/api"
	"log"
	"os"
)

func setup(key string) {
	api.SetApiKey(key)
}

func main() {
	keyFile, err := os.Open("apiKey.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer keyFile.Close()

	scanner := bufio.NewScanner(keyFile)
	scanner.Scan()
	key := scanner.Text()
	setup(key)

	subject := "CS"
	catalogNumber := "450"

	fetched := api.CourseScheduleQuery(api.CourseCatalog{Subject:subject, CatalogNumber:catalogNumber})
	for _, course := range fetched {
		println(course.ToString())
	}
}
