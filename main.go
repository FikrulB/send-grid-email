package sendgridemail

import (
	"errors"

	"github.com/FikrulB/send-grid-email/domain"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const tryLimit = 5

func SendGridEmail(req domain.RequestSendGrid) (err error) {
	if req.ApiKey == "" {
		err = errors.New("Please provide a api key")
		return
	}

	if req.From.Name == "" || req.From.Address == "" {
		err = errors.New("Error Email From")
		return
	}

	if req.To.Name == "" || req.To.Address == "" {
		err = errors.New("Error Email To")
		return
	}

	from := mail.NewEmail(req.From.Name, req.From.Address)
	to := mail.NewEmail(req.To.Name, req.To.Address)
	subject := req.Subject

	mailInit := mail.NewV3MailInit(from, subject, to)
	mailInit.SetTemplateID(req.TemplateID)

	if req.ReplyTo.Name != "" && req.ReplyTo.Address != "" {
		mailInit.SetReplyTo(mail.NewEmail(req.ReplyTo.Name, req.ReplyTo.Address))
	}

	if len(req.Subs) > 0 {
		for x := 0; x < len(mailInit.Personalizations); x++ {
			for k, v := range req.Subs {
				mailInit.Personalizations[x].SetSubstitution(k, v)
			}
		}
	}

	for i := 0; i < len(req.Attachments); i++ {
		setNewAttachments := mail.NewAttachment()
		setNewAttachments = &req.Attachments[i]
		mailInit.AddAttachment(setNewAttachments)
	}

	for x := 0; x < tryLimit; x++ {
		client := sendgrid.NewSendClient(req.ApiKey)
		_, err = client.Send(mailInit)
		if err != nil {
			continue
		}

		err = nil
		break
	}

	return
}
