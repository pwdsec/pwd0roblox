package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

var hashKey = "046634"

var letter_table = [][]string{
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

	{"a", "------=-----="},
	{"b", "=--==--===-=-"},
	{"c", "=--=-==-=--=-"},
	{"b", "=-=-=-=-==-=-"},
	{"d", "=--=-==-=---="},
	{"e", "==-=-==-=--=="},
	{"f", "=--=-==-==-=="},
	{"g", "=--=-=--=--=="},
	{"h", "=----==-=--=="},
	{"i", "=-----=-=--=="},
	{"j", "---=-==-=--=-"},
	{"k", "---=-==-=---="},
	{"l", "=--=-==-=----"},
	{"m", "=--=-==----=="},
	{"n", "---=-==-=--=="},
	{"o", "=-------=--=="},
	{"p", "=--=-==-----="},
	{"q", "=--=-=--=--=-"},
	{"r", "=----==-=--=-"},
	{"s", "---=--=-=--=="},
	{"t", "=--=--=-=---="},
	{"u", "---=----=--=="},
	{"v", "=--=-=--=---="},
	{"w", "=--=-=--=----"},
	{"x", "==-=-==-==-=="},
	{"y", "=--=-==-====="},
	{"z", "====-==-=--=="},
}

// hash a string with key, custom hash function
func Hash(key, value string) string {
	hash := sha256.New()
	hash.Write([]byte(key + value))
	return hex.EncodeToString(hash.Sum(nil))
}

func PwdEncoder(text string) string {
	for _, letter := range text {
		for i := 0; i < len(letter_table); i++ {
			if string(letter) == letter_table[i][0] {
				text = strings.Replace(text, string(letter), letter_table[i][1], -1)
			} else if string(letter) == " " {
				text = strings.Replace(text, string(letter), "/", -1)
			}
		}
	}
	return text
}

func PwdDecoder(text string) string {
	for _, letter := range strings.Split(text, "/") {
		for i := 0; i < len(letter_table); i++ {
			if string(letter) == letter_table[i][1] {
				text = strings.Replace(text, string(letter), letter_table[i][0], -1)
			}
		}
	}

	for _, letter := range text {
		if string(letter) == "/" {
			text = strings.Replace(text, string(letter), " ", -1)
		}
	}

	return text
}
