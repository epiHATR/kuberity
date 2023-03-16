package helpers

import (
	b64 "encoding/base64"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomString(length int, format string) string {
	b := randSeq(15)

	if len(format) <= 0 {
		return string(b)
	} else {
		if format == "base64" {
			return b64.StdEncoding.EncodeToString([]byte(b))
		} else {
			return string(b)
		}
	}
}
