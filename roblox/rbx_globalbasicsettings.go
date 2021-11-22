package roblox

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

// read xml file
func ReadXML(file string) (string, error) {
	xml, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(xml), nil
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func atoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

type Bool struct {
	Name  string
	Value bool
}

type Int struct {
	Name  string
	Value int
}

type Token struct {
	Name  string
	Value string
}

type BinaryString struct {
	Name  string
	Value string
}

type Int64 struct {
	Name  string
	Value int64
}

type String struct {
	Name  string
	Value string
}

type Float struct {
	Name  string
	Value float64
}

type Vector2 struct {
	Name string
	X    float64
	Y    float64
}

// read all bools with regex
func ReadBoolAll(xml string) (map[string]Bool, error) {
	re := regexp.MustCompile(`<bool name="(.*?)">(.*?)</bool>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	boolMap := make(map[string]Bool)
	for _, v := range all {
		boolMap[v[1]] = Bool{v[1], v[2] == "true"}
	}
	return boolMap, nil
}

// read all int with regex
func ReadIntAll(xml string) (map[string]Int, error) {
	re := regexp.MustCompile(`<int name="(.*?)">(.*?)</int>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	intMap := make(map[string]Int)
	for _, v := range all {
		intMap[v[1]] = Int{v[1], atoi(v[2])}
	}
	return intMap, nil
}

// read all token with regex
func ReadTokenAll(xml string) (map[string]Token, error) {
	re := regexp.MustCompile(`<token name="(.*?)">(.*?)</token>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	tokenMap := make(map[string]Token)
	for _, v := range all {
		tokenMap[v[1]] = Token{v[1], v[2]}
	}
	return tokenMap, nil
}

// read all BinaryString with regex
func ReadBinaryStringAll(xml string) (map[string]BinaryString, error) {
	re := regexp.MustCompile(`<BinaryString name="(.*?)">(.*?)</BinaryString>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	binaryStringMap := make(map[string]BinaryString)
	for _, v := range all {
		binaryStringMap[v[1]] = BinaryString{v[1], v[2]}
	}
	return binaryStringMap, nil
}

// read all int64 with regex
func ReadInt64All(xml string) (map[string]Int64, error) {
	re := regexp.MustCompile(`<int64 name="(.*?)">(.*?)</int64>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	int64Map := make(map[string]Int64)
	for _, v := range all {
		int64Map[v[1]] = Int64{v[1], atoi64(v[2])}
	}
	return int64Map, nil
}

// read all string with regex
func ReadStringMapAll(xml string) (map[string]String, error) {
	re := regexp.MustCompile(`<string name="(.*?)">(.*?)</string>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	stringMap := make(map[string]String)
	for _, v := range all {
		stringMap[v[1]] = String{v[1], v[2]}
	}
	return stringMap, nil
}

// read all float with regex
func ReadFloatAll(xml string) (map[string]Float, error) {
	re := regexp.MustCompile(`<float name="(.*?)">(.*?)</float>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	floatMap := make(map[string]Float)
	for _, v := range all {
		floatMap[v[1]] = Float{v[1], atof(v[2])}
	}
	return floatMap, nil
}

// read Vector 2
func ReadVector2All(xml string) (map[string]Vector2, error) {
	re := regexp.MustCompile(`<Vector2 name="(.*?)">(.*?)</Vector2>`)
	all := re.FindAllStringSubmatch(xml, -1)
	if len(all) == 0 {
		return nil, nil
	}
	vector2Map := make(map[string]Vector2)
	for _, v := range all {
		vector2Map[v[1]] = Vector2{
			X: atof(v[2]),
			Y: atof(v[3]),
		}
	}
	return vector2Map, nil
}
