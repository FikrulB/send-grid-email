package domain

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
