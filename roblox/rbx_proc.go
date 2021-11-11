package roblox

import (
	"os/exec"
	"strings"
)

// is process running function
func IsProcessRunning(name string) bool {
	out, err := exec.Command("tasklist", "/FO", "CSV", "/NH", "/FI", "imagename eq "+name).Output()
	if err != nil {
		return false
	}
	if len(out) > 0 {
		if strings.Contains(string(out), "RobloxPlayerBeta") {
			return true
		}
	}
	return false
}

// taskkill function
func TaskKill(name string) {
	_, err := exec.Command("taskkill", "/F", "/IM", name).Output()
	if err != nil {
		return
	}
}
