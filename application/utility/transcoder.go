package utility

import (
	"encoding/base64"
)

func Encode(content string) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(content))
}

func Decode(content string) (string, error) {
	decoded, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(content)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
