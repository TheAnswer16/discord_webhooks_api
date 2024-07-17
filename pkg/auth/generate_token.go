package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandBase64Token(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
