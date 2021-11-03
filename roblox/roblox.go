package roblox

import (
	"io/ioutil"
	"os"
	"strings"

	wapi "github.com/jcollie/w32"
)

func fixcrash_method1() {
	a, b := os.UserCacheDir()
	if b != nil {
		println("Failed to get user cache directory")
		return
	}

	files, err := ioutil.ReadDir(a + "\\Roblox\\")
	if err != nil {
		print(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "GlobalBasicSettings") {
			os.Remove(a + "\\Roblox\\" + file.Name())
		}
	}
}

func IS_Open() {
	hwid := wapi.FindWindow("", "Roblox")

	if hwid == 0 {
		print("Roblox is not open\n")
		return
	}
}

func CommandHandler(command []string) {
	switch command[0] {
	case "--fix":
		if len(command) == 2 {
			if command[1] == "UBK" {
				fixcrash_method1()
			} else {
				println("Unknown fix: " + command[1])
			}
		} else {
			println("Usage: --fix [option]")
		}
	case "--help":
		print("--fix: Fixes the Unexpected Behavior Kick that happens when you open Roblox\n")
	default:
		print("Unknown command: " + command[0] + "\n")
	}
}
