package utility

import (
	"net/mail"
	"regexp"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	return hasUpper && hasLower && hasDigit
}
