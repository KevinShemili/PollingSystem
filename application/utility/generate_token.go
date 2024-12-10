package utility

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strconv"
	"time"
)

func GenerateRefreshToken() (string, time.Time, error) {

	token := make([]byte, 16)

	_, err := rand.Read(token)
	if err != nil {
		return "", time.Time{}, err
	}

	refreshToken := hex.EncodeToString(token)

	expiry := os.Getenv("EXPIRY_REFRESH")
	expiryDays, err := strconv.Atoi(expiry)

	if err != nil {
		return "", time.Time{}, err
	}

	// expiry after 7 days
	expirationTime := time.Now().UTC().Add(time.Duration(expiryDays) * 24 * time.Hour)

	return refreshToken, expirationTime, nil
}
