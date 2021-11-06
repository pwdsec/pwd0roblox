package roblox

import (
	"io/ioutil"
	"os"
	"strings"
)

// fixes the Unexpected Behavior Kick that happens when you open Roblox.
// this is a workaround for a bug in Roblox.
// this method is not 100% reliable.
func Fix_Unexpected_Behavior_Kick_method_1() {
	print("	[+] Fixing Unexpected Behavior Kick...\n")
	a, b := os.UserCacheDir()
	if b != nil {
		println("	Failed to get user cache directory")
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
	print("	[+] Successfully fixed Unexpected Behavior Kick\n")
}
