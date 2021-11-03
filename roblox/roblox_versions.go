package roblox

import (
	"io/ioutil"
	"net/http"
)

// get roblox version from http://setup.roblox.com/version.
// error if not found.
// string roblox version.
func GetRobloxWindowsVersion() (string, error) {
	resp, err := http.Get("http://setup.roblox.com/version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// get roblox version from http://setup.roblox.com/versionStudio.
// error if not found.
// string roblox version.
func GetRobloxStudioWindowsVersion() (string, error) {
	resp, err := http.Get("http://setup.roblox.com/versionStudio")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// get roblox version from http://setup.roblox.com/mac/version.
// error if not found.
// string roblox version.
func GetRobloxMacVersion() (string, error) {
	resp, err := http.Get("http://setup.roblox.com/mac/version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// get roblox version from http://setup.roblox.com/mac/versionStudio.
// error if not found.
// string roblox version.
func GetRobloxStudioMacVersion() (string, error) {
	resp, err := http.Get("http://setup.roblox.com/mac/versionStudio")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
