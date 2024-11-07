package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacSha256(message string, secret string) (string, error) {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	base64Str := base64.StdEncoding.EncodeToString(sum)
	return base64Str, nil
}
