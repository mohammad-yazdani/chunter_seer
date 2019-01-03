package listen

import (
	"chunter_seer/api"
	"chunter_seer/notif"
	"encoding/json"
	"log"
	"net/http"
)

func handleAddCourse(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL.Path)

	subject := r.URL.Query().Get("subject")
	catalogNumber := r.URL.Query().Get("catalog_number")

	if subject == "" || catalogNumber == "" {
		_, err := w.Write([]byte("No subject or catalog number provided."))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("No subject or catalog number provided.")
		return
	}

	catalog := api.CourseCatalog{Subject:subject, CatalogNumber:catalogNumber}

	api.AddToFetchList(catalog)

	w.Header().Add("subject", catalog.Subject)
	w.Header().Add("catalog_number", catalog.CatalogNumber)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal(err)
	}
}

func handleAddMail(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL.Path)

	server := r.URL.Query().Get("server")
	email := r.URL.Query().Get("email")

	if server == "" || email == "" {
		_, err := w.Write([]byte("No server or email provided."))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("No server or email provided.")
		return
	}

	mail := notif.UserMail{Server:server, Username:email}

	notif.AddToMailingList(mail)

	w.Header().Add("server", server)
	w.Header().Add("email", email)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal(err)
	}
}

func handleStats(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL.Path)

	stats := genStats()
	statsJSON, err := json.Marshal(stats)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(statsJSON)
	if err != nil {
		log.Fatal(err)
	}
}

func Start()  {
	http.HandleFunc("/add/course", handleAddCourse)
	http.HandleFunc("/add/mail", handleAddMail)
	http.HandleFunc("/stats", handleStats)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
