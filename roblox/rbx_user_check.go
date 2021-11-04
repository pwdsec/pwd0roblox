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

func NormalUsernameGenerator(how_many string) string {
	resp, err := http.Get("https://story-shack-cdn-v2.glitch.me/generators/username-generator?count=" + how_many)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func ParseUsernames(body []byte) []string {
	var data []string
	var username string
	var i int
	for i = 0; i < len(body); i++ {
		if body[i] == '"' {
			i++
			for body[i] != '"' {
				username += string(body[i])
				i++
			}
			data = append(data, username)
			username = ""
		}
	}
	return data
}
