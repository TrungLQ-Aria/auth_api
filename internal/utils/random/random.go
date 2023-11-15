package random

import (
	"math/rand"
	"time"
)

const letters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandInt(max, min int) int {
	rand.Seed(time.Now().UnixNano())
	max++
	return rand.Intn(max-min) + min
}
func String(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
