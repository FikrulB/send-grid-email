package sendgridemail

import (
	"errors"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type (
	RequestSendGrid struct {
		From        User
		To          User
		ReplyTo     User
		Subject     string
		TemplateID  string
		Username    string
		Subs        map[string]string
		Attachments []Attachment
		ApiKey      string
	}

	User struct {
		Name    string
		Address string
	}

	Attachment struct {
		Content     string
		Type        string
		Name        string
		Filename    string
		Disposition string
		ContentID   string
	}
)

const tryLimit = 5

func SendGridEmail(req RequestSendGrid) (err error) {
	if req.ApiKey == "" {
		err = errors.New("Please provide a api key")
		return
	}

	if req.From.Address == "" {
		err = errors.New("Error Email From")
		return
	}

	if req.To.Address == "" {
		err = errors.New("Error Email To")
		return
	}

	from := mail.NewEmail(req.From.Name, req.From.Address)
	to := mail.NewEmail(req.To.Name, req.To.Address)
	subject := req.Subject

	mailInit := mail.NewV3MailInit(from, subject, to)
	mailInit.SetTemplateID(req.TemplateID)

	if (req.ReplyTo.Name != "" && req.ReplyTo.Address != "") || req.ReplyTo.Address != "" {
		mailInit.SetReplyTo(mail.NewEmail(req.ReplyTo.Name, req.ReplyTo.Address))
	}

	for x := 0; x < len(mailInit.Personalizations); x++ {
		for k, v := range req.Subs {
			mailInit.Personalizations[x].SetSubstitution(k, v)
		}
	}

	for i := 0; i < len(req.Attachments); i++ {
		setNewAttachments := mail.NewAttachment()
		setNewAttachments = &mail.Attachment{
			Content:     req.Attachments[i].Content,
			Type:        req.Attachments[i].Type,
			Name:        req.Attachments[i].Name,
			Filename:    req.Attachments[i].Filename,
			Disposition: req.Attachments[i].Disposition,
			ContentID:   req.Attachments[i].ContentID,
		}
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
