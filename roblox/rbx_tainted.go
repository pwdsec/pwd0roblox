package roblox

import "os"

func GetINIFiles() []string {
	a, b := os.UserCacheDir()
	if b != nil {
		println(b.Error())
	}
	// get each file from Roblox\logs\archive

}
