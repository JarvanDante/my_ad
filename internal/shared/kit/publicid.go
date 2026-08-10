package kit

import (
	"crypto/rand"
	"fmt"
)

const alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

func NewPublicID() (string, error) {
	const n = 16
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("rand: %w", err)
	}
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		out[i] = alphabet[int(buf[i])%len(alphabet)]
	}
	return string(out), nil
}
