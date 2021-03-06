package roblox

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

// gets each ini file path
func GetINIFiles() []string {
	appdata_local, b := os.UserCacheDir()
	if b != nil {
		pterm.Error.Println("Failed to get user cache directory")
	}
	// get each .ini file path from Roblox\logs\archive\
	files, err := ioutil.ReadDir(appdata_local + "\\Roblox\\logs\\archive")
	if err != nil {
		println(err.Error())
	}

	var ini_files []string
	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".ini" {
			ini_files = append(ini_files, appdata_local+"\\Roblox\\logs\\archive\\"+file.Name())
		}
	}
	return ini_files
}

// clear all logs from Roblox\logs\archive\ folder
func ClearLogs() {
	appdata_local, b := os.UserCacheDir()
	if b != nil {
		pterm.Error.Println("Failed to get user cache directory")
	}
	files, err := ioutil.ReadDir(appdata_local + "\\Roblox\\logs\\archive")
	if err != nil {
		println(err.Error())
	}

	for _, file := range files {
		if file.Name()[len(file.Name())-4:] == ".ini" {
			os.Remove(appdata_local + "\\Roblox\\logs\\archive\\" + file.Name())
		}
	}
}

// MacOS
// get logs files
func GetLOGFiles() []string {
	user := os.Getenv("USER")
	files, err := ioutil.ReadDir("/Users/" + user + "/Library/Logs/Roblox")
	if err != nil {
		println(err.Error())
	}

	var log_files []string
	for _, file := range files {
		if strings.Contains(file.Name(), "bootstrapper") {
			log_files = append(log_files, "/Users/"+user+"/Library/Logs/Roblox/"+file.Name())
		}
	}
	return log_files
}

// MacOS
// is tainted function
func IsTaintedLogFiles() {
	var is_tainted bool = false
	for _, v := range GetLOGFiles() {
		log_file, err := ioutil.ReadFile(v)
		if err != nil {
			println(err.Error())
		}
		if strings.Contains(string(log_file), "IsProcessTainted: true") {
			is_tainted = true
		}
	}
	if is_tainted {
		pterm.Warning.Println("User is tainted")
	} else {
		pterm.Success.Println("User is not tainted")
	}
}

// read each ini file function
func ReadINIFile(file string) map[string]string {
	ini_file, err := ioutil.ReadFile(file)
	if err != nil {
		println(err.Error())
	}

	ini_lines := strings.Split(string(ini_file), "\n")
	ini_map := make(map[string]string)

	for _, line := range ini_lines {
		key_value := strings.Split(line, "=")
		if len(key_value) > 1 {
			ini_map[key_value[0]] = key_value[1]
		}
	}
	return ini_map
}

func IsTainted(ini_map map[string]string) bool {
	return ini_map["IsTainted"] == "true"
}

// get TaintingModule
func GetTaintingModule(ini_map map[string]string) string {
	return ini_map["TaintingModule"]
}