package roblox

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

// fixes the Unexpected Behavior Kick that happens when you open Roblox.
// this is a workaround for a bug in Roblox.
// this method is not 100% reliable.
func Fix_Unexpected_Behavior_Kick_method_1() {
	fixing, _ := pterm.DefaultSpinner.Start("Fixing Unexpected Behavior Kick...")
	a, b := os.UserCacheDir()
	if b != nil {
		pterm.Error.Println("Failed to get user cache directory")
		return
	}

	files, err := ioutil.ReadDir(a + "\\Roblox\\")
	if err != nil {
		print(err)
		return
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "GlobalBasicSettings") {
			os.Remove(a + "\\Roblox\\" + file.Name())
		}
	}
	fixing.Success("Fixed Unexpected Behavior Kick")
}
