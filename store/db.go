package store

import (
	"chunter_seer/shared"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// TODO : ALL TEMP

var store * sql.DB

func SetUpDb() {
	db, err := sql.Open("sqlite3", "./store.db")
	store = db
	if err != nil {
		shared.LOG(err.Error())
	}

	sqlStmt := `
	create table if not exists emails (id integer not null primary key, email text unique);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		shared.LOG(err.Error())
		return
	}

	sqlStmt = `
	create table if not exists courses 
		(id integer not null primary key, catalog text unique, listeners integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		shared.LOG(err.Error())
		return
	}

	sqlStmt = `
	create table if not exists enrollments 
		(class integer not null primary key, subject text, catalog_number text, section text,
		 total integer, capacity integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		shared.LOG(err.Error())
		return
	}
}

func CloseDb() {
	store.Close()
}

func AddEmail(email string) {
	tx, err := store.Begin()
	if err != nil {
		shared.LOG(err.Error())
	}
	stmt, err := tx.Prepare("insert into emails(email) values(?)")
	if err != nil {
		shared.LOG(err.Error())
	}

	_, err = stmt.Exec(email)
	if err != nil {
		shared.LOG(err.Error())
	}

	err = stmt.Close()
	if err != nil {
		shared.LOG(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		shared.LOG(err.Error())
	}
}

func AddCourse(catalog string) {
	tx, err := store.Begin()
	if err != nil {
		shared.LOG(err.Error())
	}
	stmt, err := tx.Prepare("insert into courses(catalog) values(?)")
	if err != nil {
		shared.LOG(err.Error())
	}

	_, err = stmt.Exec(catalog)
	if err != nil {
		shared.LOG(err.Error())
	}

	err = stmt.Close()
	if err != nil {
		shared.LOG(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		shared.LOG(err.Error())
	}
}

func SaveEnrollment(e EnrollStats) {
	tx, err := store.Begin()
	if err != nil {
		shared.LOG(err.Error())
	}

	stmt, err := tx.Prepare(
		"insert or replace into " +
			"enrollments(class, subject, catalog_number, section, total, capacity) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		shared.LOG(err.Error())
	}

	_, err = stmt.Exec(e.Class, e.Subject, e.CatalogNumber, e.Section, e.Total, e.Capacity)
	if err != nil {
		shared.LOG(err.Error())
	}

	err = stmt.Close()
	if err != nil {
		shared.LOG(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		shared.LOG(err.Error())
	}
}

func GetEmails() []string {
	emails := make([]string, 0)

	rows, err := store.Query("select email from emails")
	if err != nil {
		shared.LOG(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)

		emails = append(emails, email)

		if err != nil {
			shared.LOG(err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		shared.LOG(err.Error())
	}

	return emails
}

func GetCourses() []string  {
	courses := make([]string, 0)

	rows, err := store.Query("select catalog from courses")
	if err != nil {
		shared.LOG(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var catalog string

		err = rows.Scan(&catalog)

		courses = append(courses, catalog)

		if err != nil {
			shared.LOG(err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		shared.LOG(err.Error())
	}

	return courses
}

func GetEnrollments() []EnrollStats  {
	enrolls := make([]EnrollStats, 0)

	rows, err := store.Query("select class, subject, catalog_number, section, total, capacity from enrollments")
	if err != nil {
		shared.LOG(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var classId int
		var subject string
		var catalogNumber string
		var section string
		var total int
		var capacity int

		err = rows.Scan(&classId, &subject, &catalogNumber, &section, &total, &capacity)

		enrolls = append(enrolls, EnrollStats{Class:classId, Subject:subject, CatalogNumber:catalogNumber,
				Section:section, Total:total, Capacity:capacity})

		if err != nil {
			shared.LOG(err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		shared.LOG(err.Error())
	}

	return enrolls
}
