package domain

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type (
	User struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	RequestSendGrid struct {
		From        User
		To          User
		ReplyTo     User
		Subject     string
		TemplateID  string
		Username    string
		Subs        map[string]string
		Attachments []mail.Attachment
		ApiKey      string
	}
)
