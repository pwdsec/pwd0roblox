// roblox functions package.
// this package contains fixes, commands, and other functions.
package roblox

import (
	"os"
	"pwd0roblox/console"
	"pwd0roblox/network"
	"strconv"
	"sync"

	"github.com/pterm/pterm"
)

// roblox command handler.
// command []string.
func CommandHandler(command []string) {
	if len(command) == 0 {
		return
	}
	switch command[0] {
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
			if network.IsConnected() {
				if len(command) == 2 {
					if command[1] == "-h" {
						pterm.Info.Println("Usage: --version-bruteforce (-vb) [how many]")
					} else {
						// multi threading
						howMany, err := strconv.Atoi(command[1])

						var wg sync.Pool
						wg.New = func() interface{} {
							return new(sync.WaitGroup)
						}

						for i := 0; i < howMany; i++ {
							wg.Get().(*sync.WaitGroup).Add(1)
							go func(i int) {
								defer wg.Get().(*sync.WaitGroup).Done()
								if err != nil {
									pterm.Error.Println("Failed to convert to int")
									return
								}
								VersionBruteForce(howMany)
							}(i)
						}
						wg.Get().(*sync.WaitGroup).Wait()
					}
				} else {
					pterm.Info.Println("Usage: --version-bruteforce (-vb) [how many]")
				}
			} else {
				pterm.Error.Println("Not connected to internet")
			}
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-ip", "-gi":
		if console.IsWindows() {
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					pterm.Success.Println(Get_IP_Address(Get_Log()))
				} else {
					pterm.Warning.Println("Roblox is not running")
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-info", "-go":
		if console.IsWindows() {
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					pterm.DefaultSection.Println("Game Information")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"IP Address", "Place ID", "Place Name", "Session ID"},
						{Get_IP_Address(Get_Log()), Get_Place_ID(Get_Log()), Get_Pace_Name(Get_Place_ID(Get_Log())), Get_Session_ID(Get_Log())},
					}).Render()

					username, err := getUserIDInfo(Get_User_ID(Get_Log()))
					if err != nil {
						pterm.Error.Println(err.Error())
						return
					}
					pterm.DefaultSection.Println("Player Information")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"User ID", "Username"},
						{Get_User_ID(Get_Log()), username},
					}).Render()

					pterm.DefaultSection.Println("More Information")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Replicator", "RakNet Socket", "Universe ID"},
						{Get_Replicator_ID(Get_Log()), Get_RakNet_IP_Address(Get_Log()), Get_Universe_ID(Get_Log())},
					}).Render()
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--is-connection-lost", "-icl":
		if console.IsWindows() {
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					if IsConnectionLost(Get_Log()) {
						pterm.Warning.Println("Connection Lost")
					} else {
						pterm.Success.Println("Connection is not lost")
					}
				} else {
					pterm.Warning.Println("Roblox is not running")
				}
			} else {
				pterm.Warning.Println("You are not connected to the internet")
			}
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
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"RakNet Socket"},
						{Get_RakNet_IP_Address(Get_Log())},
					}).Render()
				} else {
					pterm.Warning.Println("Roblox is not running")
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-replicator", "-gr":
		if console.IsWindows() {
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Replicator"},
						{Get_Replicator_ID(Get_Log())},
					}).Render()
				} else {
					pterm.Warning.Println("Roblox is not running")
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-map", "-gm":
		if console.IsWindows() {
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Place ID", "Place Name"},
						{Get_Place_ID(Get_Log()), Get_Pace_Name(Get_Place_ID(Get_Log()))},
					}).Render()
				} else {
					pterm.Warning.Println("Roblox is not running")
				}
			} else {
				pterm.Warning.Println("Not connected to internet")
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--game-localplayer", "-gl":
		if console.IsWindows() {
			if network.IsConnected() {
				if IsProcessRunning("RobloxPlayerBeta.exe") {
					username, err := getUserIDInfo(Get_User_ID(Get_Log()))
					if err != nil {
						pterm.Error.Println(err.Error())
						return
					}
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"User ID", "Username"},
						{Get_User_ID(Get_Log()), username},
					}).Render()
				} else {
					pterm.Warning.Println("Roblox is not running")
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
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
			if network.IsConnected() {
				if len(command) == 2 {
					if command[1] == "--list" || command[1] == "-l" {
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Script ID", "Script Name"},
						}).Render()
					} else if command[1] == "--help" || command[1] == "-h" {
						pterm.Info.Println("Script Hub Help")
						pterm.Info.Println("--list -l - Lists all scripts")
						pterm.Info.Println("--help -h - Displays this help")
						pterm.Info.Println("--download -d - Download the script")
					}
				} else if len(command) == 3 {
					if command[1] == "--download" || command[1] == "-d" {
						println("Downloading script")
					} else if command[1] == "--search" || command[1] == "-s" {
						GetScripts(command[2])
					}
				} else {
					pterm.Info.Println("Script Hub Help")
					pterm.Info.Println("--list -l - Lists all scripts")
					pterm.Info.Println("--help -h - Displays this help")
					pterm.Info.Println("--download -d - Download the script")
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--global-basic-settings", "-gbs":
		if console.IsWindows() {
			appdata_local, b := os.UserCacheDir()
			if b != nil {
				pterm.Error.Println(b.Error())
				return
			}

			xml, err := ReadXML(appdata_local + "\\Roblox\\GlobalBasicSettings_13.xml")
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			bools, err := ReadBoolAll(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			ints, err := ReadIntAll(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			tokens, err := ReadTokenAll(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			binarystring, err := ReadBinaryStringAll(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			int64s, err := ReadInt64All(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			stringmap, err := ReadStringMapAll(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			floats, err := ReadFloatAll(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			pterm.DefaultSection.Println("Global Basic Settings")
			d := pterm.TableData{{"Name", "Value"}}
			for _, v := range bools {
				d = append(d, []string{v.Name, strconv.FormatBool(v.Value)})
			}

			for _, v := range ints {
				d = append(d, []string{v.Name, strconv.Itoa(v.Value)})
			}

			for _, v := range tokens {
				d = append(d, []string{v.Name, v.Value})
			}

			for _, v := range binarystring {
				d = append(d, []string{v.Name, v.Value})
			}

			for _, v := range int64s {
				d = append(d, []string{v.Name, strconv.FormatInt(v.Value, 10)})
			}

			for _, v := range stringmap {
				d = append(d, []string{v.Name, v.Value})
			}

			for _, v := range floats {
				d = append(d, []string{v.Name, strconv.FormatFloat(v.Value, 'f', -1, 64)})
			}
			pterm.DefaultTable.WithHasHeader().WithData(d).WithBoxed().Render()

			vec2, err := ReadVector2All(xml)
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}

			/* == Working on this == */
			pterm.DefaultSection.Println("Vector2")
			vec := pterm.TableData{{"Name", "X", "Y"}}
			for _, v := range vec2 {
				vec = append(vec, []string{v.Name, strconv.FormatFloat(v.X, 'f', -1, 64), strconv.FormatFloat(v.Y, 'f', -1, 64)})
			}
			pterm.DefaultTable.WithHasHeader().WithData(vec).WithBoxed().Render()
			/* ===================== */

		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--assets-bruteforce", "-ab":
		if console.IsWindows() || console.IsMacOS() {
			if len(ROBLOSECURITY) > 0 {
				if network.IsConnected() {
					for {
						AssetBruteforce()
					}
				} else {
					pterm.Error.Println("You are not connected to the internet")
				}
			} else {
				pterm.Error.Println("ROBLOSECURITY is not set")
			}
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--asset-downloader", "-ad":
		if console.IsWindows() || console.IsMacOS() {
			if network.IsConnected() {
				if len(ROBLOSECURITY) == 0 {
					pterm.Error.Println("ROBLOSECURITY is not set")
					return
				}
				if len(command) == 1 {
					pterm.Error.Println("Please specify an id")
				} else if len(command) == 2 {
					AssetDownload(command[1])
				} else {
					pterm.Error.Println("Too many arguments")
				}
			} else {
				pterm.Error.Println("You are not connected to the internet")
			}
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
