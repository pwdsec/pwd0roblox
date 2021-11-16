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
	{"A", "+#+#x+++?#+#++#"}, {"B", "#++##++##?#X+#+"},
	{"C", "#++#+?##x+#++#+"}, {"D", "#+#+?x#+#+##+#+"},
	{"E", "#++##+#+#?x#+#+"}, {"F", "+++++#+#+#?+x#+"},
	{"F", "++#+#?+#+##x+#+"}, {"G", "+###?++X++++++#"},
	{"H", "+###++##?##x###"}, {"H", "+++##x####?X+++"},
	{"I", "++X+##+?+###+++"}, {"J", "+#?#+x+##++#+++"},
	{"K", "+##+++?+##X++#+"}, {"L", "+#++#++++X#?#++"},
	{"M", "+###++#?++x++##"}, {"N", "+###?++#++x++++"},
	{"O", "+++?+++##+X+#++"}, {"P", "+###X#+++#?#++#"},
	{"Q", "+##+#+?+++X#+++"}, {"R", "##+X+?##x++#+++"},
	{"S", "+##X?+++++++#++"}, {"T", "+##++##+?++#X++"},
	{"U", "+##X+###++#+?++"}, {"V", "#++#++++?+X++++"},
	{"W", "+##x++?++++#+++"}, {"X", "+#?#+X###++##++"},
	{"Y", "++++x++?+##++#+"}, {"Z", "+##+++++#x#?+++"},
	{"0", "+++++x##?++#+++"}, {"1", "+++#+##+?+#+X+#"},
	{"2", "+++++#x#?++#++#"}, {"3", "+++?+x+##++#+#+"},
	{"4", "+++++##?x++##++"}, {"5", "+++++##+x+?++++"},
	{"6", "++++?+##+x++++#"}, {"7", "+++++?##++x++##"},
	{"8", "++++++#++x++++?"}, {"9", "++++?+x+#+++++#"},

	{"a", "---?---#-x----#"}, {"b", "#--##?X--###-#-"},
	{"c", "#--#x-##-#-?-#-"}, {"b", "#-#-#-#-#?#-x#-"},
	{"d", "#--?#-##x-#---#"}, {"e", "##-#-x#?#-#--##"},
	{"f", "#--#-##?-X##-##"}, {"g", "#--#-x#--#?--##"},
	{"h", "#-x-?--##-#--##"}, {"i", "#-----#-x#-?-##"},
	{"j", "--x-#-?##-#--#-"}, {"k", "--?-#x-##-#---#"},
	{"l", "#--x#-#?#-#----"}, {"m", "#--#-##----?x##"},
	{"n", "---#x-##-?#--##"}, {"o", "#-----x-?-#--##"},
	{"p", "#--#x-##-?----#"}, {"q", "#--#-?#--#x--#-"},
	{"r", "#----##x-#-?-#-"}, {"s", "---#--x#-#--?##"},
	{"t", "#--#--#-#x--?-#"}, {"u", "---?#--x--#--##"},
	{"v", "#--?#-x#--#---#"}, {"w", "?#--#-#--#--x--"},
	{"x", "##-#?-#x#-##-##"}, {"y", "#?--#-#x#-#####"},
	{"z", "####-##-#x--#?#"}, {"#", "#^^^^^#--?x---#"},
	{"-", "#^^^x^^#^^^?^^#"}, {"+", "#^^*?^^x#^^^^^#"},
	{"/", "#^^x*^^#^^*^^#?"}, {"\\", "#*^*^#^x^?*^^#"},
	{"*", "#*^?**#^^*x^^##"}, {"@", "#*^*?*#^x^#^^^#"},
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
		result += minStr + "a(zv" + minStr + "a)z{a" + minStr + string(letter) + minStr + "zv}a*" + minStr
	}

	result = "[" + result + "]"
	return result
}

func PwdDecoder(text string) string {
	var data = []string{
		"[", "]", "(", ")", "{", "}", "*", "1", "2", "z",
		"3", "4", "5", "6", "7", "8", "9", "0", "a", "v",
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
