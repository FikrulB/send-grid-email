package sendgridemail

import (
	"fmt"
	"net/smtp"
)

func SendMail(password string) {
	auth := smtp.PlainAuth(
		"",
		"mfikrulb@gmail.com",
		password,
		"smtp.gmail.com",
	)

	msg := "Subject: This is a Subject\nThis is a Message"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"mfikrulb@gmail.com",
		[]string{"mfikrulb@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}
