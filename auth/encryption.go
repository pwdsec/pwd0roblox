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

func PwdEncoder(text string) {
	var enc_1 = [][]string{
		{"A", "+=+=+++=+=++="},
		{"B", "=++==++===+=+"},
		{"C", "=++=+==+=++=+"},
		{"D", "=+=+=+=+==+=+"},
		{"E", "=++==+=+==+=+"},
		{"F", "+++++=+=+=+=+"},
		{"F", "++=+=+=+==+=+"},
		{"G", "+===++++++++="},
		{"H", "+===++======="},
		{"H", "+++=======+++"},
		{"I", "+++==++===+++"},
		{"J", "+==++==++=+++"},
		{"K", "+==++++==++=+"},
		{"L", "+=++=++++==++"},
		{"M", "+===++=++++=="},
		{"N", "+===++=++++++"},
		{"O", "++++++==++=++"},
		{"P", "+====+++==++="},
		{"Q", "+==+=++++=+++"},
		{"R", "==+++==++=+++"},
		{"S", "+==+++++++=++"},
		{"T", "+==++==+++=++"},
		{"U", "+==+===++=+++"},
		{"V", "=++=+++++++++"},
		{"W", "+==++++++=+++"},
		{"X", "+==+===++==++"},
		{"Y", "+++++++==++=+"},
		{"Z", "+==+++++==+++"},
		{"0", "+++++==++=+++"},
		{"1", "+++=+==++=++="},
		{"2", "+++++==++=++="},
		{"3", "+++++==++=+=+"},
		{"4", "+++++==++==++"},
		{"5", "+++++==++++++"},
		{"6", "+++++==+++++="},
		{"7", "+++++==++++=="},
		{"8", "++++++=++++++"},
		{"9", "++++++=+++++="},
	}
}
