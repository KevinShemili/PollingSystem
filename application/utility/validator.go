package utility

import (
	"net/mail"
	"regexp"
)

// ValidateEmail checks if the provided email is in valid format
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

// ValidatePassword checks if the provided password is in valid format. 8+ Characters, At least one uppercase, one lowercase & one number.
func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	return hasUpper && hasLower && hasDigit
}
