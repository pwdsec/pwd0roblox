package roblox

import (
	"math/rand"
	"pwd0roblox/network"
	"strconv"
	"sync"

	"github.com/pterm/pterm"
)

func Fix_Command_Windows(command []string) {
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
}

func Cursor_Command_Windows_Mac(command []string) {
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
}

func Version_Command_Windows_Mac() {
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
}

func Install_Command_Windows(command []string) {
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
}

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

func Check_Command_Windows_Mac(command []string) {
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
}

func Tainted_Command_Windows(command []string) {
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
}

func Set_Token_Command_Windows_Mac(command []string) {
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
}

func API_Command_Windows_Mac(command []string) {
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
}

func Login_Command_Windows_Mac(command []string) {
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
}

func Version_Bruteforce_Command_Windows_Mac(command []string) {
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
}

func Game_IP_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			pterm.Success.Println(Get_IP_Address(Get_Log()))
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Error.Println("You are not connected to the internet")
	}
}
