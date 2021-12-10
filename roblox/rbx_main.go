// roblox functions package.
// this package contains fixes, commands, and other functions.
package roblox

import (
	"pwd0roblox/console"
	"pwd0roblox/exploits"
	"pwd0roblox/network"

	"github.com/pterm/pterm"
)

// roblox command handler.
// command []string.
func CommandHandler(command []string) {
	if len(command) == 0 {
		return
	}
	switch command[0] {
	case "--exploits", "-ex":
		if len(command) == 1 {
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"Exploit"},
				{"--synapse"},
			}).Render()
		}
		if len(command) == 2 {
			if command[1] == "--synapse" {
				if exploits.IsSynapseUpdated() {
					pterm.Success.Println("Synapse is up to date.")
				} else {
					pterm.Error.Println("Synapse exploit is not updated.")
				}
			}
		}
	case "--fix", "-f":
		if console.IsWindows() {
			Fix_Command_Windows(command)
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--cursor", "-c":
		if console.IsWindows() || console.IsMacOS() {
			Cursor_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--versions", "-v":
		Version_Command_Windows_Mac()
	case "--delete", "-d":
		if console.IsWindows() {
			DeleteRoblox()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--reinstall", "-r":
		if console.IsWindows() {
			if network.IsConnected() {
				DeleteRoblox()
				ver, _ := GetRobloxWindowsVersion()
				InstallRoblox(ver, true)
			} else {
				pterm.Error.Println("No internet connection")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--install", "-i":
		if console.IsWindows() {
			Install_Command_Windows(command)
		} else if console.IsMacOS() {
			Install_Command_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--check", "-C":
		Check_Command_Windows_Mac(command)
	case "--tainted", "-t":
		if console.IsWindows() {
			Tainted_Command_Windows(command)
		} else if console.IsMacOS() {
			IsTaintedLogFiles()
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--set-token", "-s":
		if console.IsWindows() || console.IsMacOS() {
			Set_Token_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--api", "-a":
		if console.IsWindows() || console.IsMacOS() {
			API_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--login", "-l":
		if console.IsWindows() || console.IsMacOS() {
			Login_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--version-bruteforce", "-vb":
		if console.IsWindows() || console.IsMacOS() {
			Version_Bruteforce_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-ip", "-gi":
		if console.IsWindows() {
			Game_IP_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-info", "-go":
		if console.IsWindows() {
			Game_Info_Command_windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--is-connection-lost", "-icl":
		if console.IsWindows() {
			Is_Connection_lost_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--kill", "-k":
		if console.IsWindows() {
			TaskKill("RobloxPlayerBeta.exe")
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--raknet-socket", "-rs":
		if console.IsWindows() {
			RakNet_Socket_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-replicator", "-gr":
		if console.IsWindows() {
			Game_Replicator_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-map", "-gm":
		if console.IsWindows() {
			Game_Map_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-localplayer", "-gl":
		if console.IsWindows() {
			Game_LocalPlayer_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--directory", "-dr":
		/*open roblox directory*/
	case "--botting", "-b":
		if console.IsWindows() {
			if network.IsConnected() {
				/* botting commnnds */
				/*
					bot likes
					bot dislikes
					bot friends
					bot followers
					bot game
				*/
				pterm.Error.Println("Windows is not yet supported")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--cookie-checker", "-cc":
		if console.IsWindows() || console.IsMacOS() {
			pterm.Error.Println("This command is not supported on this OS")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--account-checker", "-ac":
		if console.IsWindows() || console.IsMacOS() {
			pterm.Error.Println("This command is not supported on this OS")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--pin-bruteforce", "-pb":
		if console.IsWindows() || console.IsMacOS() {
			pterm.Error.Println("This command is not supported on this OS")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--ban-account", "-ba":
		if console.IsWindows() || console.IsMacOS() {
			pterm.Error.Println("This command is not supported on this OS")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--crash-local-client", "-clc":
		if console.IsWindows() {

		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--lag-switch", "-ls":
		if console.IsWindows() || console.IsMacOS() {
			pterm.Error.Println("This command is not supported on this OS")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--script-hub", "-sh":
		if console.IsWindows() || console.IsMacOS() {
			Script_Hub_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--global-basic-settings", "-gbs":
		if console.IsWindows() {
			GlobalBasicSettings_Command_Windows()
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--assets-bruteforce", "-ab":
		if console.IsWindows() || console.IsMacOS() {
			Asset_Bruteforce_Command_Windows_Mac()
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--asset-downloader", "-ad":
		if console.IsWindows() || console.IsMacOS() {
			Asset_Downloader_Command_Windows_Mac(command)
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--help", "-h", "?":
		if console.IsWindows() {
			if network.IsConnected() {
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--login", "-l", "Logs into Roblox"},
					{"--set-token", "-t", "Sets the Roblox Security Token"},
					{"--api", "-a", "Gets information about the user"},
					{"--help", "-h", "Shows this help"},
					{"--check", "-C", "Checks/Generate if a username is valid"},
					{"--fix", "-f", "Fixes bugs that happens to Roblox"},
					{"--cursor", "-c", "Installs a custom cursor"},
					{"--delete", "-d", "Deletes Roblox"},
					{"--install", "-i", "Installs Roblox"},
					{"--reinstall", "-r", "Reinstalls Roblox"},
					{"--tainted", "-t", "Checks if user is tainted"},
					{"--versions", "-v", "Prints the latest versions of Roblox and Roblox Studio"},
					{"--version-bruteforce", "-vb", "Bruteforces the Roblox version"},
					{"--kill", "-k", "Kills Roblox process"},
					{"--game-ip", "-gi", "Gets the game ip and port"},
					{"--game-info", "-go", "Gets the game ip, port, and place id and more"},
					{"--game-replicator", "-gr", "Gets the game replicator"},
					{"--game-map", "-gm", "Gets the game map"},
					{"--game-localplayer", "-gl", "Gets the game localplayer"},
					{"--is-connection-lost", "-icl", "Checks if connection is lost on roblox"},
					{"--raknet-socket", "-rs", "Gets the raknet socket"},
					{"--directory", "-dr", "Opens the Roblox directory"},
					{"--botting", "-b", "Starts a botting script"},               // working on it
					{"--cookie-checker", "-cc", "Checks if cookies are valid"},   // working on it
					{"--account-checker", "-ac", "Checks if account is valid"},   // working on it
					{"--pin-bruteforce", "-pb", "Bruteforces the pin"},           // working on it
					{"--ban-account", "-ba", "Bans the account"},                 // working on it
					{"--crash-local-client", "-clc", "Crashes the local client"}, // working on it
					{"--lag-switch", "-ls", "Lags client"},                       // working on it
					{"--script-hub", "-sh", "Opens the script hub"},              // working on it
					{"--global-basic-settings", "-gbs", "Show GlobalBasicSettings"},
					{"--assets-bruteforce", "-ab", "Bruteforces the assets"},
					{"--asset-downloader", "-ad", "Download a asset"},
				}).Render()
			} else {
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--help", "-h", "Shows this help"},
					{"--fix", "-f", "Fixes bugs that happens to Roblox"},
					{"--delete", "-d", "Deletes Roblox"},
					{"--tainted", "-t", "Checks if user is tainted"},
					{"--directory", "-dr", "Opens the Roblox directory"},
					{"--global-basic-settings", "-gbs", "Show GlobalBasicSettings"},
				}).Render()
			}
		} else if console.IsMacOS() {
			if network.IsConnected() {
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--version-bruteforce", "-vb", "Bruteforces the Roblox version"},
					{"--cursor", "-c", "Installs a custom cursor"},
					{"--tainted", "-t", "Checks if user is tainted"},
					{"--login", "-l", "Logs into Roblox"},
					{"--install", "-i", "Installs Roblox"},
					{"--version", "-v", "Prints the latest versions of Roblox and Roblox Studio"},
					{"--api", "-a", "Gets information about the user and more"},
					{"--check", "-C", "Checks/Generate if a username is valid"},
					{"--set-token", "-t", "Sets the Roblox Security Token"},
					{"--assets-bruteforce", "-ab", "Bruteforces the assets"},
					{"--asset-downloader", "-ad", "Download a asset"},
					{"--help", "-h", "Shows this help"},
				}).Render()
			} else {
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--cursor", "-c", "Installs a custom cursor"},
					{"--tainted", "-t", "Checks if user is tainted"},
					{"--help", "-h", "Shows this help"},
				}).Render()
			}
		}
	default:
		pterm.Error.Println("Unknown command: " + command[0])
	}
}
