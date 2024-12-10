package utility

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func GenerateJWTWithID(UserID uint) (string, error) {

	jwtExpiryHour, _ := strconv.Atoi(os.Getenv("EXPIRY_JWT"))
	jwtSigningKey := os.Getenv("SECRET_JWT")

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": UserID,
		"exp": time.Now().Add(time.Duration(jwtExpiryHour) * time.Hour).Unix(),
	})

	signedToken, err := jwtToken.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateJWTWithClaims(claims jwt.MapClaims) (string, error) {

	jwtExpiryHour, _ := strconv.Atoi(os.Getenv("EXPIRY_JWT"))
	jwtSigningKey := os.Getenv("SECRET_JWT")

	// update only expiry
	claims["exp"] = time.Now().Add(time.Duration(jwtExpiryHour) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	signedToken, err := token.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
