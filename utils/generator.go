package utils

import (
	"math/rand"
	"time"
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func EncodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	base := int64(len(base62Chars))
	var result []byte

	for num > 0 {
		remainder := num % base
		result = append([]byte{base62Chars[remainder]}, result...)
		num = num / base
	}

	// 🔥 Pad to 8 characters with random characters
	for len(result) < 8 {
		randomChar := base62Chars[rand.Intn(len(base62Chars))]
		result = append([]byte{randomChar}, result...)
	}

	return string(result)
}
