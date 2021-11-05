package roblox

import (
	"io"
	"net/http"
	"os"
	"os/exec"
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

	resp, err := http.Get("https://setup.rbxcdn.com/" + version + "-Roblox.exe")
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

	println("	[+] Roblox downloaded!")
	if start {
		println("	[+] Installing Roblox...")
		// run file
		cmd := exec.Command("Roblox.exe")
		cmd.Start()
		cmd.Wait()

		println("	[+] Roblox installed!")
	}
}
