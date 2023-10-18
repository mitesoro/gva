package utils

import (
	"math/rand"
)

func GenerateRandomCode(length int) string {
	charSet := "0123456789"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charSet))
		code[i] = charSet[randomIndex]
	}

	return string(code)
}
