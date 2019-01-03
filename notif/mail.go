package notif

import (
	"chunter_seer/listen"
	"chunter_seer/shared"
	"chunter_seer/store"
	"crypto/tls"
	"net/smtp"
	"strconv"
	"strings"
)

var mailingList []string

var serverUsername string
var serverPassword string
var serverHostName string

func SetUpMail(username string, password string, host string) {
	serverUsername = username
	serverPassword = password
	serverHostName = host

	mailingList = store.GetEmails()

	listen.AddHandler("add_mail", AddToMailingList)
}

func AddToMailingList(mail string) (string, error) {
	mailingList = append(mailingList, mail)
	store.AddEmail(mail)
	return "OK", nil
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
		shared.LOG(err.Error())
	}

	client, err := smtp.NewClient(conn, serverHostName)
	if err != nil {
		shared.LOG(err.Error())
	}

	if err = client.Auth(auth); err != nil {
		shared.LOG(err.Error())
	}

	if err = client.Mail(serverUsername); err != nil {
		shared.LOG(err.Error())
	}

	u0 := mailingList[0]

	if err = client.Rcpt(u0); err != nil {
		shared.LOG(err.Error())
	}

	w, err := client.Data()
	if err != nil {
		shared.LOG(err.Error())
	}

	msg := "From: " + serverUsername + "\n" +
		"To: " + u0 + "\n" +
		"Subject: Chunter UPDATE\n\n"

	msgString := strings.Builder{}
	msgString.WriteString(msg)

	msgString.WriteString("Course: ")
	msgString.WriteString(course)
	msgString.WriteString("\n")

	changeStr := strconv.FormatInt(int64(change), 10)
	msgString.WriteString("Change: ")
	msgString.WriteString(changeStr)
	msgString.WriteString("\n")

	msg = msgString.String()
	_, err = w.Write([]byte(msg))
	if err != nil {
		shared.LOG(err.Error())
	}

	err = w.Close()
	if err != nil {
		shared.LOG(err.Error())
	}

	err = client.Quit()
	if err != nil {
		shared.LOG(err.Error())
	}

	shared.LOG("Mail sent successfully")
}
