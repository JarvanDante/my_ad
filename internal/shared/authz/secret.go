package authz

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"strings"
)

func HashSecret(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(sum[:])
}

func MatchSecret(plain, stored string, hashed bool) bool {
	plain = strings.TrimSpace(plain)
	stored = strings.TrimSpace(stored)
	if plain == "" || stored == "" {
		return false
	}
	if hashed || looksLikeSHA256Hex(stored) {
		want := HashSecret(plain)
		a, b := []byte(want), []byte(strings.ToLower(stored))
		if len(a) != len(b) {
			return false
		}
		return subtle.ConstantTimeCompare(a, b) == 1
	}
	a, b := []byte(plain), []byte(stored)
	if len(a) != len(b) {
		return false
	}
	return subtle.ConstantTimeCompare(a, b) == 1
}

func looksLikeSHA256Hex(s string) bool {
	if len(s) != 64 {
		return false
	}
	for _, c := range strings.ToLower(s) {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}
