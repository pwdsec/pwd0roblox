package roblox

import (
	"os"

	"github.com/pterm/pterm"
)

func NeedUpdate() bool {
	appdata_local, b := os.UserCacheDir()
	if b != nil {
		pterm.Error.Println("Failed to get user cache directory")
	}

	version, err := GetRobloxWindowsVersion()
	if err != nil {
		pterm.Error.Println("Failed to get roblox version")
		return false
	}

	if _, err := os.Stat(appdata_local + "\\Roblox\\Versions\\" + version); os.IsNotExist(err) {
		return true
	}

	return false
}
