// roblox functions package.
// this package contains fixes, commands, and other functions.
package roblox

import (
	"io/ioutil"
	"os"
	"strings"

	wapi "github.com/jcollie/w32"
)

// fixes the Unexpected Behavior Kick that happens when you open Roblox.
// this is a workaround for a bug in Roblox.
// this method is not 100% reliable.
func Fix_Unexpected_Behavior_Kick_method_1() {
	print("[+] Fixing Unexpected Behavior Kick...\n")
	a, b := os.UserCacheDir()
	if b != nil {
		println("Failed to get user cache directory")
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
	print("[+] Successfully fixed Unexpected Behavior Kick\n")
}

func IS_Open() {
	hwid := wapi.FindWindow("", "Roblox")

	if hwid == 0 {
		print("Roblox is not open\n")
		return
	}
}

// roblox command handler.
// command []string.
func CommandHandler(command []string) {
	switch command[0] {
	case "--fix", "-f":
		if len(command) == 2 {
			if command[1] == "UBK" {
				Fix_Unexpected_Behavior_Kick_method_1()
			} else {
				println("Unknown fix: " + command[1])
			}
		} else {
			println("Usage: --fix (-f) [option]")
			println("Options:")
			println("UBK - Fixes Unexpected Behavior Kick in Roblox")
		}
	case "--cursor", "-c":
		if len(command) == 2 {
			err := CursorsInstaller(command[1])
			if err != nil {
				println("Invalid command")
			}
		} else {
			println("Usage: --cursor (-c) [option]")
			println("Options:")
			for _, v := range CursorsList() {
				println("	" + v)
			}
		}
	case "--versions", "-v":
		roblox_windows_version, _ := GetRobloxWindowsVersion()
		roblox_studio_windows_version, _ := GetRobloxStudioWindowsVersion()
		roblox_mac_version, _ := GetRobloxMacVersion()
		roblox_studio_mac_version, _ := GetRobloxStudioMacVersion()

		println("Roblox Windows Version: " + roblox_windows_version)
		println("Roblox Studio Windows Version: " + roblox_studio_windows_version)
		println("Roblox Mac Version: " + roblox_mac_version)
		println("Roblox Studio Mac Version: " + roblox_studio_mac_version)
	case "--help", "-h":
		print("--fix, -f ~ Fixes the Unexpected Behavior Kick that happens when you open Roblox\n")
		print("--cursor, -c ~ Installs a custom cursor\n")
		print("--versions, -v ~ Prints the latest versions of Roblox and Roblox Studio\n")
	default:
		print("Unknown command: " + command[0] + "\n")
	}
}
