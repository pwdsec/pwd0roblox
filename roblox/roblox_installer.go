package roblox

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func DeleteRoblox() {
	println("[+] Deleting Roblox...")
	dir, _ := os.UserCacheDir()
	// check if roblox is installed
	if _, err := os.Stat(dir + "\\Roblox"); err == nil {
		// delete roblox
		os.RemoveAll(dir + "\\Roblox")
	}
	println("[+] Roblox deleted!")
}

func InstallRoblox() {
	println("[+] Downloading Roblox...")
	ver, _ := GetRobloxWindowsVersion()

	resp, err := http.Get("https://setup.rbxcdn.com/" + ver + "-Roblox.exe")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// write file
	file, err := os.Create("Roblox.exe")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	file.Close()

	println("[+] Installing Roblox...")
	// run file
	cmd := exec.Command("Roblox.exe")
	cmd.Start()
	cmd.Wait()

	println("[+] Roblox installed!")
}
