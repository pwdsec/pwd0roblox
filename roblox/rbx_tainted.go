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

// read each ini file function
func ReadINIFile(file string) map[string]string {
	// read the ini file
	ini_file, err := ioutil.ReadFile(file)
	if err != nil {
		println(err.Error())
	}

	// split the ini file into lines
	ini_lines := strings.Split(string(ini_file), "\n")

	// create a map to store the ini file
	ini_map := make(map[string]string)

	// loop through each line
	for _, line := range ini_lines {
		// split the line into key and value
		key_value := strings.Split(line, "=")
		// if the line is not empty
		if len(key_value) > 1 {
			// add the key and value to the map
			ini_map[key_value[0]] = key_value[1]
		}
	}
	return ini_map
}

// if "IsTainted=true" is in the map, return true
func IsTainted(ini_map map[string]string) bool {
	return ini_map["IsTainted"] == "true"
}
