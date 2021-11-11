package roblox

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func Get_Log() []string {
	// get all logs file in path: "C:\Users\HellFire\AppData\Local\Roblox\logs"
	path := "C:\\Users\\HellFire\\AppData\\Local\\Roblox\\logs"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// check for the latest log file in the path
	var latest_log_file string
	var latest_log_file_time int64
	for _, file := range files {
		if file.ModTime().Unix() > latest_log_file_time {
			latest_log_file = file.Name()
			latest_log_file_time = file.ModTime().Unix()
		}
	}

	// read the latest log file
	log_file_path := path + "\\" + latest_log_file
	log_file, err := ioutil.ReadFile(log_file_path)
	if err != nil {
		log.Fatal(err)
	}

	// return each line of the log file
	return strings.Split(string(log_file), "\n")
}

// function that will get ip address from string
func Get_IP_Address(log_line []string) string {
	pattern := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}`
	re := regexp.MustCompile(pattern)

	// loop through each log line
	for _, line := range log_line {
		if strings.Contains(line, "Connecting to") {
			if re.MatchString(line) {
				// return ip address
				return re.FindString(line)
			}
		}
	}

	// return empty string if no ip address found
	return ""
}
