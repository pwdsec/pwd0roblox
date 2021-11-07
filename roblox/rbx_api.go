package roblox

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Cookies struct {
	ROBLOSECURITY string
	RBXID         string
}

// UserDescription
type UserDescription struct {
	Description string `json:"description"`
}

// UnreadMessages
type UnreadMessages struct {
	UnreadCount int `json:"unreadNotifications"`
}

// EmailInfo
type EmailInfo struct {
	Email    string `json:"emailAddress"`
	Verified bool   `json:"verified"`
}

type UserIDInfo struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	IsOnline bool   `json:"IsOnline"`
}

// assetids
//type CurrentlyWearing struct {
//	AssetId []int `json:"assetIds"`
//}

// get the roblox security token
func getRobloxSecurity(rbx_token string) {
	var cookies Cookies

	if strings.Contains(rbx_token, ".ROBLOSECURITY=") {
		cookies.ROBLOSECURITY = rbx_token
	} else {
		cookies.ROBLOSECURITY = ".ROBLOSECURITY=" + rbx_token
	}
}

// not finished working on it \\
func postRequestLogin(username, password string) error {
	var cookies Cookies

	url := "https://auth.roblox.com/v2/login"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	params := map[string]interface{}{
		"ctype":    "Username",
		"cvalue":   username,
		"password": password,
	}
	params_bytes, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(strings.NewReader(string(params_bytes)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	print(string(body))

	for _, cookie := range resp.Cookies() {
		print("cookies" + cookie.Name)
		if cookie.Name == ".ROBLOSECURITY" {
			cookies.ROBLOSECURITY = cookie.Value
		} else if cookie.Name == ".RBXID" {
			cookies.RBXID = cookie.Value
		}
	}

	if cookies.ROBLOSECURITY == "" {
		return errors.New("ROBLOSECURITY cookie not found")
	}

	if cookies.RBXID == "" {
		return errors.New("RBXID cookie not found")
	}

	return nil
}

// get userdescription : https://accountinformation.roblox.com/v1/description
func getUserDescription() (string, error) {
	var cookies Cookies
	url := "https://accountinformation.roblox.com/v1/description"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", cookies.ROBLOSECURITY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var user_description UserDescription
	err = json.Unmarshal(body, &user_description)
	if err != nil {
		return "", err
	}

	return user_description.Description, nil
}

// get unread messages https://notifications.roblox.com/v2/stream-notifications/unread-count
func getUnreadMessages() (int, error) {
	var cookies Cookies
	url := "https://notifications.roblox.com/v2/stream-notifications/unread-count"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Cookie", cookies.ROBLOSECURITY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var unread_messages UnreadMessages
	err = json.Unmarshal(body, &unread_messages)
	if err != nil {
		return 0, err
	}

	return unread_messages.UnreadCount, nil
}

// get email info https://accountsettings.roblox.com/v1/email
func getEmailInfo() (string, bool, error) {
	var cookies Cookies
	url := "https://accountsettings.roblox.com/v1/email"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", false, err
	}
	req.Header.Set("Cookie", cookies.ROBLOSECURITY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false, err
	}

	var email_info EmailInfo
	err = json.Unmarshal(body, &email_info)
	if err != nil {
		return "", false, err
	}

	return email_info.Email, email_info.Verified, nil
}

func getUserIDInfo(id string) (int, string, bool, error) {
	var cookies Cookies
	url := "https://api.roblox.com/users/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, "", false, err
	}
	req.Header.Set("Cookie", cookies.ROBLOSECURITY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", false, err
	}

	var user_id_info UserIDInfo
	err = json.Unmarshal(body, &user_id_info)
	if err != nil {
		return 0, "", false, err
	}

	return user_id_info.Id, user_id_info.Username, user_id_info.IsOnline, nil
}

// get cookie .RBXID
func getRBXID() (string, error) {
	var cookies Cookies
	url := "https://www.roblox.com/home"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", cookies.ROBLOSECURITY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// get response cookie
	for _, cookie := range resp.Cookies() {
		print(cookie.Name)
		if cookie.Name == ".RBXID" {
			return cookie.Value, nil
		}
	}

	return "", errors.New("cookie not found")
}
