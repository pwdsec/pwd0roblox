package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
)

var hashKey = "046634"

var letter_table = [][]string{
	{"A", "+#+#+++#+#++#"}, {"B", "#++##++###+#+"},
	{"C", "#++#+##+#++#+"}, {"D", "#+#+#+#+##+#+"},
	{"E", "#++##+#+##+#+"}, {"F", "+++++#+#+#+#+"},
	{"F", "++#+#+#+##+#+"}, {"G", "+###++++++++#"},
	{"H", "+###++#######"}, {"H", "+++#######+++"},
	{"I", "+++##++###+++"}, {"J", "+##++##++#+++"},
	{"K", "+##++++##++#+"}, {"L", "+#++#++++##++"},
	{"M", "+###++#++++##"}, {"N", "+###++#++++++"},
	{"O", "++++++##++#++"}, {"P", "+####+++##++#"},
	{"Q", "+##+#++++#+++"}, {"R", "##+++##++#+++"},
	{"S", "+##+++++++#++"}, {"T", "+##++##+++#++"},
	{"U", "+##+###++#+++"}, {"V", "#++#+++++++++"},
	{"W", "+##++++++#+++"}, {"X", "+##+###++##++"},
	{"Y", "+++++++##++#+"}, {"Z", "+##+++++##+++"},
	{"0", "+++++##++#+++"}, {"1", "+++#+##++#++#"},
	{"2", "+++++##++#++#"}, {"3", "+++++##++#+#+"},
	{"4", "+++++##++##++"}, {"5", "+++++##++++++"},
	{"6", "+++++##+++++#"}, {"7", "+++++##++++##"},
	{"8", "++++++#++++++"}, {"9", "++++++#+++++#"},

	{"a", "------#-----#"}, {"b", "#--##--###-#-"},
	{"c", "#--#-##-#--#-"}, {"b", "#-#-#-#-##-#-"},
	{"d", "#--#-##-#---#"}, {"e", "##-#-##-#--##"},
	{"f", "#--#-##-##-##"}, {"g", "#--#-#--#--##"},
	{"h", "#----##-#--##"}, {"i", "#-----#-#--##"},
	{"j", "---#-##-#--#-"}, {"k", "---#-##-#---#"},
	{"l", "#--#-##-#----"}, {"m", "#--#-##----##"},
	{"n", "---#-##-#--##"}, {"o", "#-------#--##"},
	{"p", "#--#-##-----#"}, {"q", "#--#-#--#--#-"},
	{"r", "#----##-#--#-"}, {"s", "---#--#-#--##"},
	{"t", "#--#--#-#---#"}, {"u", "---#----#--##"},
	{"v", "#--#-#--#---#"}, {"w", "#--#-#--#----"},
	{"x", "##-#-##-##-##"}, {"y", "#--#-##-#####"},
	{"z", "####-##-#--##"}, {"#", "#^^^^^#-----#"},
	{"-", "#^^^^^#^^^^^#"}, {"+", "#^^*^^#^^^^^#"},
	{"/", "#^^*^^#^^*^^#"}, {"\\", "#*^*^#^^*^^#"},
	{"*", "#*^**#^^*^^##"}, {"@", "#*^**#^^#^^^#"},
}

// hash a string with key, custom hash function
func Hash(key, value string) string {
	hash := sha256.New()
	hash.Write([]byte(key + value))
	return hex.EncodeToString(hash.Sum(nil))
}

// Base64Encode encodes a string to base64
func Base64Encode(text string) string {
	data := []byte(text)
	return base64.StdEncoding.EncodeToString(data)
}

// base64 decode a string
func Base64Decode(text string) string {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	return string(data)
}

func PwdEncoder(text string) string {
	text = Base64Encode(text)
	for _, letter := range text {
		for i := 0; i < len(letter_table); i++ {
			if string(letter) == letter_table[i][0] {
				text = strings.Replace(text, string(letter), letter_table[i][1]+";", -1)
			}
		}
	}

	var result string
	for _, letter := range text {
		min := rand.Intn(99 - 0)
		// convert min to string
		minStr := strconv.Itoa(min)
		result += minStr + "a(" + minStr + "a){a" + minStr + string(letter) + minStr + "}a*" + minStr
	}

	result = "[" + result + "]"
	return result
}

func PwdDecoder(text string) string {
	var data = []string{
		"[", "]", "(", ")", "{", "}", "*", "1", "2",
		"3", "4", "5", "6", "7", "8", "9", "0", "a",
	}

	for _, letter := range data {
		text = strings.Replace(text, letter, "", -1)
	}

	TableData := strings.Split(text, ";")

	for i := 0; i < len(TableData); i++ {
		for j := 0; j < len(letter_table); j++ {
			if TableData[i] == letter_table[j][1] {
				text = strings.Replace(text, TableData[i], letter_table[j][0], -1)
			}
		}
	}
	text = strings.Replace(text, ";", "", -1)

	text = Base64Decode(text)
	return text
}
