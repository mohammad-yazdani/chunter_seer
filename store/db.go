package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// TODO : ALL TEMP

var store * sql.DB

func SetUpDb() {
	db, err := sql.Open("sqlite3", "./store.db")
	store = db
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table if not exists emails (id integer not null primary key, email text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	sqlStmt = `
	create table if not exists courses 
		(id integer not null primary key, subject text, catalog_number text, listeners integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CloseDb() {
	store.Close()
}

func AddEmail(email string) {
	tx, err := store.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into emails(email) values(?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(email)
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func AddCourse(subject string, catalogNumber string) {
	tx, err := store.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into courses(subject, catalog_number) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(subject, catalogNumber)
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func GetEmails() []string {
	emails := make([]string, 0)

	rows, err := store.Query("select email from emails")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)

		emails = append(emails, email)

		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return emails
}

func GetCourses()[][]string  {
	courses := make([][]string, 0)

	rows, err := store.Query("select subject, catalog_number from courses")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var subject string
		var catalogNumber string

		err = rows.Scan(&subject, &catalogNumber)

		courses = append(courses, []string{subject, catalogNumber})

		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return courses
}

