package roblox

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// check username availability.
// returns true if username is available.
// https://api.roblox.com/users/get-by-username.
func CheckUsername(username string) (bool, error) {
	resp, err := http.Get("https://api.roblox.com/users/get-by-username?username=" + username)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if strings.Contains("User not found", string(body)) {
		return true, nil
	}
	return false, nil
}
