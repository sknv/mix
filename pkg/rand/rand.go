package rand

import (
	"math/rand"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	charset = letters + digits
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset generates a random string with a provided length from a charset.
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Letters generates a random letters-only string with a provided leghth.
func Letters(length int) string {
	return StringWithCharset(length, letters)
}

// Digits generates a random digits-only string with a provided leghth.
func Digits(length int) string {
	return StringWithCharset(length, digits)
}

// String generates a random string with a provided leghth.
func String(length int) string {
	return StringWithCharset(length, charset)
}
