package roblox

import (
	"io/ioutil"
	"os"
)

func fixcrash_method1() {
	files, err := ioutil.ReadDir(os.Getenv("APPDATA") + "Local\\Roblox\\")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.Name() == "GlobalBasicSettings" {
			os.Remove(os.Getenv("APPDATA") + "Local\\Roblox\\" + file.Name())
		}
	}
}

func CommandHandler(command string) {
	switch command {
	case "--fix-UBK":
		fixcrash_method1()
	case "--help":
		print("--fix-UBK: Fixes the Unexpected Behavior Kick that happens when you open Roblox [static/jayyy#8941]\n")
	default:
		print("Unknown command: " + command + "\n")
	}
}
