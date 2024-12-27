package mail

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, templatePath string, templateData interface{}) error {

	emailConfig := GetEmailConfig()

	body, err := parseTemplate(templatePath, templateData)
	if err != nil {
		log.Printf("Failed to parse email template: %v", err)
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", emailConfig.SenderEmail)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer(emailConfig.SMTPHost, emailConfig.SMTPPort, emailConfig.SenderEmail, emailConfig.SenderPass)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}

func parseTemplate(templatePath string, data interface{}) (string, error) {

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := template.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}

func GetTemplatePath(templateName string) string {
	baseDir, _ := os.Getwd()
	result := filepath.Join(baseDir, "infrastructure", "mail", "templates", templateName)
	return result
}
