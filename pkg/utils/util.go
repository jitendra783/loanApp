package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 6)
	used := make(map[string]bool)

	for {
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}

		randomString := string(b)
		if !used[randomString] {
			used[randomString] = true
			break
		}
	}

	return string(b)
}

