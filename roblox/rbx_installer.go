package roblox

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"pwd0roblox/console"

	"github.com/pterm/pterm"
)

// interesting link
// https://clientsettingscdn.roblox.com/v2/client-version/WindowsStudio
// wondering if its future version

var files = []string{
	"ssl.zip",
	"shaders.zip",
	"extracontent-places.zip",
	"extracontent-textures.zip",
	"extracontent-models.zip",
	"extracontent-translations.zip",
	"extracontent-luapackages.zip",
	"content-platform-fonts.zip",
	"content-terrain.zip",
	"content-textures3.zip",
	"content-models.zip",
	"content-textures2.zip",
	"content-sounds.zip",
	"content-sky.zip",
	"content-fonts.zip",
	"content-avatar.zip",
	"content-configs.zip",
	"RobloxApp.zip",
	"RobloxStudioLauncherBeta.exe",
}

// deletes roblox from pc.
func DeleteRoblox() {
	deleting, _ := pterm.DefaultSpinner.Start("Deleting Roblox...")
	dir, _ := os.UserCacheDir()
	// check if roblox is installed
	if _, err := os.Stat(dir + "\\Roblox"); err == nil {
		// delete roblox
		os.RemoveAll(dir + "\\Roblox")
	}
	deleting.Success("Deleted!")
}

// downloads roblox and installs it.
func InstallRoblox(version string, start bool) {
	downloading, _ := pterm.DefaultSpinner.Start("Downloading Roblox...")

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

	downloading.Success("Downloaded!")
	if start {
		installing, _ := pterm.DefaultSpinner.Start("Installing Roblox...")
		if console.IsWindows() {
			cmd := exec.Command("Roblox.exe")
			cmd.Start()
			cmd.Wait()
		} else if console.IsMacOS() {
			cmd := exec.Command("hdiutil", "attach", "Roblox.dmg")
			cmd.Start()
			cmd.Wait()
		}
		installing.Success("Installed!")
		os.Remove(file_name)
	}
}

func ContentInstaller_Ziped(version string) {
	// create directory with version
	os.Mkdir(version, 0777)

	p, _ := pterm.DefaultProgressbar.WithTotal(len(files)).WithTitle("Downloading files...").Start()
	for _, v := range files {
		resp, err := http.Get("https://setup.rbxcdn.com/" + version + "-" + v)
		if err != nil {
			pterm.Warning.Println("Failed: " + v)
		} else {
			defer resp.Body.Close()
			// write file
			file, err := os.Create(version + "\\" + v)
			if err != nil {
				pterm.Warning.Println("Failed: " + v)
			} else {
				_, err = io.Copy(file, resp.Body)
				if err != nil {
					pterm.Warning.Println("Failed: " + v)
				} else {
					file.Close()
					pterm.Success.Println("Downloaded: " + v)
				}
			}
		}
		p.Increment()
	}
	p.Stop()
}
