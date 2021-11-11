package roblox

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func Get_Log() []string {
	a, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}
	path := a + "\\Roblox\\logs"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var latest_log_file string
	var latest_log_file_time int64
	for _, file := range files {
		if file.ModTime().Unix() > latest_log_file_time {
			latest_log_file = file.Name()
			latest_log_file_time = file.ModTime().Unix()
		}
	}

	log_file_path := path + "\\" + latest_log_file
	log_file, err := ioutil.ReadFile(log_file_path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(log_file), "\n")
}

func Get_IP_Address(log_line []string) string {
	pattern := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}`
	re := regexp.MustCompile(pattern)

	for _, line := range log_line {
		if strings.Contains(line, "Connecting to") {
			if re.MatchString(line) {
				return re.FindString(line)
			}
		}
	}

	return ""
}
