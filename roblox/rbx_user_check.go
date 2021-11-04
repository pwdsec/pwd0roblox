package roblox

import (
	"io/ioutil"
	"math/rand"
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
	if strings.Contains(string(body), "User not found") {
		return true, nil
	}
	return false, nil
}

// generates a random username.
// returns the username.
func GenerateUsername(lenght int) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, lenght)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
