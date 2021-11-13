package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

var hashKey = "046634"

// hash a string with key, custom hash function
func Hash(key, value string) string {
	hash := sha256.New()
	hash.Write([]byte(key + value))
	return hex.EncodeToString(hash.Sum(nil))
}
