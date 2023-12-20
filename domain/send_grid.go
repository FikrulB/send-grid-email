package domain

// from := mail.NewEmail("Example User", "test@example.com")
// subject := "Sending with Twilio SendGrid is Fun"
// to := mail.NewEmail("Example User", "test@example.com")
// plainTextContent := "and easy to do anywhere, even with Go"
// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// response, err := client.Send(message)
// if err != nil {
// 	log.Println(err)
// } else {
// 	fmt.Println(response.StatusCode)
// 	fmt.Println(response.Body)
// 	fmt.Println(response.Headers)
// }

type (
	User struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	RequestSendGrid struct {
		From         User   `json:"from"`
		To           User   `json:"to"`
		Subject      string `json:"subject"`
		TemplateID   string `json:"template_id"`
		Username     string `json:"username"`
		ErrorMessage string `json:"error_message"`
		ApiKey       string `json:"api_key"`
	}
)
