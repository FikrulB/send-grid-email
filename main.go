package sendgridemail

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/FikrulB/send-grid-email/domain"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	fmt.Println("Success")
}

func SendGridEmail(req domain.RequestSendGrid) {
	from := mail.NewEmail(req.From.Name, req.From.Address)
	to := mail.NewEmail(req.To.Name, req.To.Address)
	subject := req.Subject
	// content := mail.NewContent("text/html", "I'm replacing the <strong>body tag</strong>")

	m := mail.NewV3MailInit(from, subject, to)

	m.Personalizations[0].SetSubstitution("-error-", req.ErrorMessage)
	m.Personalizations[0].SetSubstitution("-name-", req.Username)
	m.SetTemplateID(req.TemplateID)

	client := sendgrid.NewSendClient(req.ApiKey)
	response, err := client.Send(m)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
