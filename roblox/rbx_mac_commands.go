package roblox

import (
	"pwd0roblox/network"

	"github.com/pterm/pterm"
)

func Install_Command_Mac(command []string) {
	if network.IsConnected() {
		if len(command) == 2 {
			if command[1] == "-h" {
				pterm.Info.Println("Usage: --install (-i) [version] (-s)[start] ")
			} else {
				var start bool = false
				for _, v := range command {
					if v == "-s" {
						start = true
					}
				}
				InstallRoblox(command[1], start)
			}
		} else {
			ver, _ := GetRobloxMacVersion()
			InstallRoblox(ver, true)
		}
	} else {
		pterm.Error.Println("No internet connection")
	}
}
