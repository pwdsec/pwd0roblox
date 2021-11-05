package roblox

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"pwd0roblox/console"
)

// deletes roblox from pc.
func DeleteRoblox() {
	println("	[+] Deleting Roblox...")
	dir, _ := os.UserCacheDir()
	// check if roblox is installed
	if _, err := os.Stat(dir + "\\Roblox"); err == nil {
		// delete roblox
		os.RemoveAll(dir + "\\Roblox")
	}
	println("	[+] Roblox deleted!")
}

// downloads roblox and installs it.
func InstallRoblox(version string, start bool) {
	println("	[+] Downloading Roblox...")

	var url string
	var file_name string
	if console.IsWindows() {
		url = "https://setup.rbxcdn.com/" + version + "-Roblox.exe"
		file_name = "Roblox.exe"
	} else if console.IsMacOS() {
		url = "http://setup.roblox.com/mac/" + version + "-Roblox.dmg"
		file_name = "Roblox.dmg"
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// write file
	file, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	file.Close()

	println("	[+] Roblox downloaded!")
	if start {
		println("	[+] Installing Roblox...")
		if console.IsWindows() {
			cmd := exec.Command("Roblox.exe")
			cmd.Start()
			cmd.Wait()
		} else if console.IsMacOS() {
			cmd := exec.Command("hdiutil", "attach", "Roblox.dmg")
			cmd.Start()
			cmd.Wait()
		}
		println("	[+] Roblox installed!")
		os.Remove(file_name)
	}
}
