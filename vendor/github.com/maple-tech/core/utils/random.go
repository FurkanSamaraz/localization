package utils

import (
	"math/rand"
	"time"
)

const codeLetters = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

//GenerateRandomCode makes a base32 code of given length
func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i := 0; i < length; i++ {
		b[i] = codeLetters[rand.Intn(len(codeLetters))]
	}

	return string(b)
}
