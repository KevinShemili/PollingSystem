package mail

import (
	"os"
	"strconv"
)

type EmailConfig struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
}

func GetEmailConfig() *EmailConfig {
	return &EmailConfig{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: func() int {
			port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
			return port
		}(),
		SenderEmail: os.Getenv("SENDER_EMAIL"),
		SenderPass:  os.Getenv("SENDER_PASS"),
	}
}
