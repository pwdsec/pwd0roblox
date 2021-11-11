package roblox

import (
	"encoding/json"
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

func Get_Place_ID(log_line []string) string {

	for _, line := range log_line {
		if strings.Contains(line, "{") {
			var json_data map[string]interface{}
			json.Unmarshal([]byte(line), &json_data)
			job_id := json_data["jobId"]
			if job_id != nil {
				if strings.Contains(job_id.(string), "JoinPlace=") {
					// remove "JoinPlace=" and ";"
					return strings.Replace(strings.Replace(job_id.(string), "JoinPlace=", "", -1), ";", "", -1)
				}
			}
		}
	}
	return ""
}

func Get_Session_ID(log_line []string) string {
	pattern := `sid:([a-z0-9-]+)`
	re := regexp.MustCompile(pattern)

	for _, line := range log_line {
		if strings.Contains(line, "sid:") {
			if re.MatchString(line) {
				sid := re.FindString(line)
				return strings.Replace(sid, "sid:", "", -1)
			}
		}
	}
	return ""
}

func Get_User_ID(log_line []string) string {
	pattern := `userid:([0-9]+)`
	re := regexp.MustCompile(pattern)

	for _, line := range log_line {
		if strings.Contains(line, "userid:") {
			if re.MatchString(line) {
				userid := re.FindString(line)
				return strings.Replace(userid, "userid:", "", -1)
			}
		}
	}
	return ""
}

// IsConnectionLost
func IsConnectionLost(log_line []string) bool {
	for _, line := range log_line {
		if strings.Contains(line, "Connection lost") {
			return true
		}
	}
	return false
}
