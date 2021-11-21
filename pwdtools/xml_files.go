package pwdtools

import (
	"io/ioutil"
	"os"
)

// read xml file
func ReadXML(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(byteValue), nil
}
