package roblox

import (
	"io/ioutil"
	"os"
	"strings"

	wapi "github.com/jcollie/w32"
)

func Fix_Unexpected_Behavior_Kick_method1() {
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
	case "--fix", "-f":
		if len(command) == 2 {
			if command[1] == "UBK" {
				Fix_Unexpected_Behavior_Kick_method1()
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
	case "--help", "-h":
		print("--fix, -f ~ Fixes the Unexpected Behavior Kick that happens when you open Roblox\n")
		print("--cursor, -c ~ Installs a custom cursor\n")
	default:
		print("Unknown command: " + command[0] + "\n")
	}
}
