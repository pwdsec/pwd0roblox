// roblox functions package.
// this package contains fixes, commands, and other functions.
package roblox

import (
	"math/rand"
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
			if len(command) == 2 {
				if command[1] == "UBK" {
					Fix_Unexpected_Behavior_Kick_method_1()
				} else {
					pterm.Error.Println("Unknown fix: " + command[1])
				}
			} else {
				pterm.Info.Println("Usage: --fix (-f) [option]")
				pterm.Info.Println("Options:")
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Description", "Reliability"},
					{"UBK", "Fixes Unexpected Behavior Kick", "50%"},
				}).Render()
			}
		} else if console.IsMacOS() {
			pterm.Error.Println("MacOS is not yet supported")
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--cursor", "-c":
		if console.IsWindows() || console.IsMacOS() {
			if network.IsConnected() {
				if len(command) == 2 {
					if command[1] == "--reset" || command[1] == "-r" {
						CursorsInstaller("default")
					}
					err := CursorsInstaller(command[1])
					if err != nil {
						println(err)
					}
				} else {
					pterm.Info.Println("Usage: --cursor (-c) [option]")
					pterm.Info.Println("Options:")
					d := pterm.TableData{{"Cursor"}}
					for _, v := range CursorsList() {
						d = append(d, []string{v})
					}
					pterm.DefaultTable.WithHasHeader().WithData(d).WithBoxed().Render()

				}
			} else {
				pterm.Error.Println("Not connected to internet")
			}
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--versions", "-v":
		if network.IsConnected() {
			roblox_windows_version, _ := GetRobloxWindowsVersion()
			roblox_studio_windows_version, _ := GetRobloxStudioWindowsVersion()
			roblox_mac_version, _ := GetRobloxMacVersion()
			roblox_studio_mac_version, _ := GetRobloxStudioMacVersion()
			roblox_studio_qt_version, _ := GetRobloxStudioQTVersion()

			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"Os", "Name", "Version"},
				{"Windows", "Roblox", roblox_windows_version},
				{"Windows", "Studio", roblox_studio_windows_version},
				{"Mac", "Roblox", roblox_mac_version},
				{"Mac", "Studio", roblox_studio_mac_version},
				{"Unknown", "Studio Qt", roblox_studio_qt_version},
			}).Render()
		} else {
			pterm.Error.Println("Not connected to internet")
		}
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
			if network.IsConnected() {
				if len(command) == 3 {
					if command[1] == "--content" || command[1] == "-c" {
						ContentInstaller_Ziped(command[2])
					}
				} else if len(command) == 2 {
					if command[1] == "--help" || command[1] == "-h" {
						pterm.Info.Println("Usage: --install (-i) [option] ")
						pterm.Info.Println("Options:")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Single", "Description"},
							{"--content", "-c", "Install the content"},
							{"--help", "-h", "This help message"},
						}).Render()
					} else if command[1] == "--content" || command[1] == "-c" {
						ver, _ := GetRobloxWindowsVersion()
						ContentInstaller_Ziped(ver)
						DeleteEmptyFiles(ver)
					} else if command[1] == "--versions" || command[1] == "-v" {
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Date", "Version"},
							{"06-16-2021", "version-c864da271a4c44ea"},
							{"06-11-2021", "version-7d96d7dad25d49f1"},
							{"04-29-2021", "version-0658018801724832"},
							{"04-08-2021", "version-278f0258a7224831"},
							{"01-13-2021", "version-d5212926da8e4716"},
							{"11-13-2020", "version-aa7766fcc7cb4906"},
							{"04-10-2019", "version-9f8314ad67c64c0d"},
							{"10-29-2018", "version-e9d1a6c5df10490c"},
							{"12-12-2017", "version-45cc144b134647ea"},
							{"02-25-2016", "version-a1b8c1edf45b4959"},
						}).Render()
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
					ver, _ := GetRobloxWindowsVersion()
					InstallRoblox(ver, true)
				}
			} else {
				pterm.Error.Println("No internet connection")
			}
		} else if console.IsMacOS() {
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
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--check", "-C":
		if network.IsConnected() {
			if len(command) == 3 {
				if command[1] == "--generate" || command[1] == "-g" {
					var i int
					i, _ = strconv.Atoi(command[2])
					username := GenerateUsername(i)
					check, err := CheckUsername(username)
					if err != nil {
						pterm.Error.Println("Failed to check username")
						return
					}
					if check {
						pterm.Success.Println("Username is available: " + username)
					} else {
						pterm.Warning.Println("Username is not available: " + username)
					}
				} else if command[1] == "--username" || command[1] == "-u" {
					check, err := CheckUsername(command[2])
					if err != nil {
						pterm.Error.Println("Failed to check username")
						return
					}
					if check {
						pterm.Success.Println("Username is available: " + command[2])
					} else {
						pterm.Warning.Println("Username is not available: " + command[2])
					}
				} else if command[1] == "--normal" || command[1] == "-n" {
					normal_username := NormalUsernameGenerator(command[2])
					usernames := ParseUsernames([]byte(normal_username))
					for _, v := range usernames {
						if v == "data" || v == "name" {
							continue
						} else {
							num := rand.Intn(10)
							num_str := strconv.Itoa(num)
							v += num_str
							check, err := CheckUsername(v)
							if err != nil {
								pterm.Error.Println("Failed to check username")
								return
							}
							if check {
								pterm.Success.Println("Username is available: " + v)
							} else {
								pterm.Warning.Println("Username is not available: " + v)
							}
						}
					}
				} else {
					pterm.Info.Println("Usage: --check (-c) [option]")
					pterm.Info.Println("Options:")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Command", "Single", "Description"},
						{"--generate", "-g", "[lenght] - generate random username and check if valid"},
						{"--username", "-u", "[username] - Checks if a username is valid"},
						{"--normal", "-n", "[how many] - generates usernames and check if valid"},
					}).Render()
				}
			} else {
				pterm.Info.Println("Usage: --check (-c) [option]")
				pterm.Info.Println("Options:")
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--generate", "-g", "[lenght] - generate random username and check if valid"},
					{"--username", "-u", "[username] - Checks if a username is valid"},
					{"--normal", "-n", "[how many] - generates usernames and check if valid"},
				}).Render()
			}
		} else {
			pterm.Error.Println("Not connected to internet")
		}
	case "--tainted", "-t":
		if console.IsWindows() {
			if len(command) == 2 {
				if command[1] == "-h" {
					pterm.Info.Println("Usage: --tainted (-t) [version]")
				}
			} else {
				var is_tainted bool = false
				ini_files := GetINIFiles()
				for _, v := range ini_files {
					mapped := ReadINIFile(v)
					if IsTainted(mapped) {
						is_tainted = true
					} else {
						is_tainted = false
					}
				}
				if is_tainted {
					pterm.Warning.Println("User Tainted")
				} else {
					pterm.Success.Println("User Not Tainted")
				}
			}
		} else if console.IsMacOS() {
			IsTaintedLogFiles()
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--set-token", "-s":
		if console.IsWindows() || console.IsMacOS() {
			if network.IsConnected() {
				if len(command) == 2 {
					if command[1] == "-h" {
						pterm.Info.Println("Usage: --set-token (-t) [token]")
					} else {
						getRobloxSecurity(command[1])
					}
				} else {
					pterm.Info.Println("Usage: --set-token (-t) [token]")
				}
			} else {
				pterm.Error.Println("Not connected to internet")
			}
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--api", "-a":
		if console.IsWindows() || console.IsMacOS() {
			if network.IsConnected() {
				if len(ROBLOSECURITY) == 0 {
					pterm.Info.Println("Roblox Security Token not set (rbx --set-token [token]) or (rbx --login [username] [password])")
					return
				}
				if len(command) == 2 {
					if command[1] == "--description" || command[1] == "-d" {
						description, err := getUserDescription()
						if err != nil {
							pterm.Error.Println("Failed to get user description")
							return
						}
						pterm.Success.Println("User Description: " + description)
					} else if command[1] == "--messages" || command[1] == "-m" {
						ms, err := getUnreadMessages()
						if err != nil {
							pterm.Error.Println("Failed to get unread messages")
							return
						}
						// convert ms to string
						pterm.Success.Println("Unread Messages: " + strconv.Itoa(ms))
					} else if command[1] == "--email" || command[1] == "-e" {
						email, verified, err := getEmailInfo()
						if err != nil {
							pterm.Error.Println("Failed to get email info")
							return
						}
						if verified {
							pterm.Success.Println("Email: " + email)
						} else {
							pterm.Warning.Println("Email: " + email + " (Not Verified)")
						}
					}
				} else if len(command) == 3 {
					/*if command[1] == "--userid" || command[1] == "-u" {
						_, user, online, err := getUserIDInfo(command[2])
						if err != nil {
							pterm.Error.Println("Failed to get user id info")
							return
						}
						if len(user) == 0 {
							pterm.Error.Println("User not found")
							return
						}
						if online {
							pterm.Success.Println("User: " + user + " (Online)")
						} else {
							pterm.Success.Println("User: " + user + " (Offline)")
						}
					} else if command[1] == "--get" || command[1] == "-g" {
						if command[2] == "--rbxid" || command[2] == "-r" {
							rbxid, err := getRBXID()
							if err != nil {
								pterm.Error.Println("Failed to get rbxid")
								return
							}
							pterm.Success.Println("RBXID: " + rbxid)
						}
					}*/
				} else {
					pterm.Info.Println("Usage: --api (-a) [option]")
					pterm.Info.Println("Options:")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Command", "Single", "Description"},
						{"--description", "-d", "Gets the user description"},
						{"--messages", "-m", "Gets the unread messages"},
						{"--email", "-e", "Gets the email info"},
						{"--help", "-h", "Shows this help"},
						{"--userid", "-u", "Gets the user id info"},
						{"--get", "-g", "Gets option"},
					}).Render()
					pterm.Info.Println("Usage: --api (-a) --get (-g) [option]")
					pterm.Info.Println("Options:")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Command", "Single", "Description"},
						{"--rbxid", "-r", "Gets the rbxid"},
					}).Render()
				}
			} else {
				pterm.Error.Println("Not connected to the internet")
			}
		} else {
			pterm.Error.Println("Unknown OS")
		}
	case "--login", "-l":
		if console.IsWindows() || console.IsMacOS() {
			if network.IsConnected() {
				if len(command) == 3 {
					if command[1] == "-h" {
						pterm.Info.Println("Usage: --login (-l) [username] [password]")
					} else {
						err := postRequestLogin(command[1], command[2])
						if err != nil {
							println(err.Error())
							return
						} else {
							pterm.Success.Println("Login Successful")
						}
					}
				} else {
					pterm.Info.Println("Usage: --login (-l) [username] [password]")
				}
			} else {
				pterm.Error.Println("Not connected to the internet")
			}
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
					{"--botting", "-b", "Starts a botting script"},
					{"--cookie-checker", "-cc", "Checks if cookies are valid"},
					{"--account-checker", "-ac", "Checks if account is valid"},
					{"--pin-bruteforce", "-pb", "Bruteforces the pin"},
					{"--ban-account", "-ba", "Bans the account"},
				}).Render()
			} else {
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--help", "-h", "Shows this help"},
					{"--fix", "-f", "Fixes bugs that happens to Roblox"},
					{"--delete", "-d", "Deletes Roblox"},
					{"--tainted", "-t", "Checks if user is tainted"},
					{"--directory", "-dr", "Opens the Roblox directory"},
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
