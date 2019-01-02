package notif

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"strings"
)

type UserMail struct {
	Username string
	Server string
}

var mailingList []UserMail

var serverUsername string
var serverPassword string
var serverHostName string

func SetUpMail(username string, password string, host string) {
	serverUsername = username
	serverPassword = password
	serverHostName = host

	mailingList = make([]UserMail, 0)
}

func AddToMailingList(mail UserMail)  {
	mailingList = append(mailingList, mail)
}

// TODO : TEMP solution
func MailChange(course string, change int) {

	if len(mailingList) == 0 {
		return
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         serverHostName,
	}

	auth := smtp.PlainAuth("", serverUsername, serverPassword, serverHostName)

	conn, err := tls.Dial("tcp", serverHostName + ":465", tlsConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, serverHostName)
	if err != nil {
		log.Panic(err)
	}

	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(serverUsername); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(mailingList[0].Username); err != nil {
		log.Panic(err)
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	msgString := strings.Builder{}
	msgString.WriteString("Course: ")
	msgString.WriteString(course)
	msgString.WriteString("\n")

	msgString.WriteString("Change: ")
	msgString.WriteString(string(change))
	msgString.WriteString("\n")

	msg := msgString.String()
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	err = client.Quit()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Mail sent successfully")
}
