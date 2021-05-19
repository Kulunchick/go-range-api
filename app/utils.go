package util

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
)

func sendEmail(template string, to string, subject string) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body, err := ioutil.ReadFile(template)

	if err != nil {
		fmt.Println(err)
	}

	msg := []byte(subject + mime + string(body))

	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")

	addr := fmt.Sprintf("%s:%s", host, port)

	user := os.Getenv("MAIL_USER")
	password := os.Getenv("MAIL_PASSWORD")
	auth := smtp.PlainAuth("", user, password, host)

	err = smtp.SendMail(addr, auth, user, []string{to}, msg)

	if err != nil {
		fmt.Println(err)
	}
}
