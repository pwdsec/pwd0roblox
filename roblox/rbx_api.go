package roblox

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var ROBLOSECURITY string

// get the roblox security token
func getRobloxSecurity(rbx_token string) {
	ROBLOSECURITY = rbx_token
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

// assetids
//type CurrentlyWearing struct {
//	AssetId []int `json:"assetIds"`
//}

// get userdescription : https://accountinformation.roblox.com/v1/description
func getUserDescription() (string, error) {
	url := "https://accountinformation.roblox.com/v1/description"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", ROBLOSECURITY)
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
	url := "https://notifications.roblox.com/v2/stream-notifications/unread-count"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Cookie", ROBLOSECURITY)
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
	url := "https://accountsettings.roblox.com/v1/email"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", false, err
	}
	req.Header.Set("Cookie", ROBLOSECURITY)
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
