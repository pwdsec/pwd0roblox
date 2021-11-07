// roblox functions package.
// this package contains fixes, commands, and other functions.
package roblox

import (
	"math/rand"
	"pwd0roblox/console"
	"strconv"
)

// roblox command handler.
// command []string.
func CommandHandler(command []string) {
	switch command[0] {
	case "--fix", "-f":
		if console.IsWindows() {
			if len(command) == 2 {
				if command[1] == "UBK" {
					Fix_Unexpected_Behavior_Kick_method_1()
				} else {
					println("	Unknown fix: " + command[1])
				}
			} else {
				println("	Usage: --fix (-f) [option]")
				println("	Options:")
				println("		UBK - Fixes Unexpected Behavior Kick in Roblox")
			}
		} else if console.IsMacOS() {
			println("	MacOS is not yet supported")
		} else {
			println("	Unknown OS")
		}
	case "--cursor", "-c":
		if console.IsWindows() {
			if len(command) == 2 {
				err := CursorsInstaller(command[1])
				if err != nil {
					println("Invalid command")
				}
			} else {
				println("	Usage: --cursor (-c) [option]")
				println("	Options:")
				for _, v := range CursorsList() {
					println("		" + v)
				}
			}
		} else if console.IsMacOS() {
			println("	MacOS is not yet supported")
		} else {
			println("	Unknown OS")
		}
	case "--versions", "-v":
		roblox_windows_version, _ := GetRobloxWindowsVersion()
		roblox_studio_windows_version, _ := GetRobloxStudioWindowsVersion()
		roblox_mac_version, _ := GetRobloxMacVersion()
		roblox_studio_mac_version, _ := GetRobloxStudioMacVersion()
		roblox_studio_qt_version, _ := GetRobloxStudioQTVersion()

		println("	Roblox Windows Version: " + roblox_windows_version)
		println("	Roblox Studio Windows Version: " + roblox_studio_windows_version)
		println("	Roblox Mac Version: " + roblox_mac_version)
		println("	Roblox Studio Mac Version: " + roblox_studio_mac_version)
		println("	Roblox Studio Qt Version: " + roblox_studio_qt_version)
	case "--delete", "-d":
		if console.IsWindows() {
			DeleteRoblox()
		} else if console.IsMacOS() {
			println("	MacOS is not yet supported")
		} else {
			println("	Unknown OS")
		}
	case "--reinstall", "-r":
		if console.IsWindows() {
			DeleteRoblox()
			ver, _ := GetRobloxWindowsVersion()
			InstallRoblox(ver, true)
		} else if console.IsMacOS() {
			println("	MacOS is not yet supported")
		} else {
			println("	Unknown OS")
		}
	case "--install", "-i":
		if console.IsWindows() {
			if len(command) == 2 {
				if command[1] == "-h" {
					println("	Usage: --install (-i) [version] (-s)[start] ")
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
		} else if console.IsMacOS() {
			if len(command) == 2 {
				if command[1] == "-h" {
					println("	Usage: --install (-i) [version] (-s)[start] ")
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
			println("	Unknown OS")
		}
	case "--check", "-C":
		if len(command) == 3 {
			if command[1] == "--generate" || command[1] == "-g" {
				var i int
				i, _ = strconv.Atoi(command[2])
				username := GenerateUsername(i)
				check, err := CheckUsername(username)
				if err != nil {
					println("	Failed to check username")
					return
				}
				if check {
					println("	Username is valid: " + username)
				} else {
					println("	Username is invalid: " + username)
				}
			} else if command[1] == "--username" || command[1] == "-u" {
				check, err := CheckUsername(command[2])
				if err != nil {
					println("	Failed to check username")
					return
				}
				if check {
					println("	Username is valid: " + command[2])
				} else {
					println("	Username is invalid: " + command[2])
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
							println("	Failed to check username")
							return
						}
						if check {
							println("	Username is valid: " + v)
						} else {
							println("	Username is invalid: " + v)
						}
					}
				}
			} else {
				println("	Usage: --check (-c) [option]")
				println("	Options:")
				println("		--generate (-g) [lenght] - generate random username and check if valid")
				println("		--username (-u) [username] - Checks if a username is valid")
				println("		--normal (-n) [how many] - generates usernames and check if valid")
			}
		} else {
			println("	Usage: --check (-c) [option]")
			println("	Options:")
			println("		--generate (-g) [lenght] - generate random username and check if valid")
			println("		--username (-u) [username] - Checks if a username is valid")
			println("		--normal (-n) [how many] - generates usernames and check if valid")
		}
	case "--tainted", "-t":
		if console.IsWindows() {
			if len(command) == 2 {
				if command[1] == "-h" {
					println("	Usage: --tainted (-t) [version]")
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
					println("	User Tainted")
				} else {
					println("	User Not tainted")
				}
			}
		} else if console.IsMacOS() {
			println("	MacOS is not yet supported")
		} else {
			println("	Unknown OS")
		}
	case "--set-token", "-s":
		if console.IsWindows() || console.IsMacOS() {
			if len(command) == 2 {
				if command[1] == "-h" {
					println("	Usage: --set-token (-t) [token]")
				} else {
					getRobloxSecurity(command[1])
				}
			} else {
				println("	Usage: --set-token (-t) [token]")
			}
		} else {
			println("	Unknown OS")
		}
	case "--api", "-a":
		if console.IsWindows() || console.IsMacOS() {
			if len(ROBLOSECURITY) == 0 {
				println("	Roblox Security Token not set (rbx --set-token [token])")
				return
			}
			if len(command) == 2 {
				if command[1] == "--description" || command[1] == "-d" {
					description, err := getUserDescription()
					if err != nil {
						println("	Failed to get user description")
						return
					}
					println("	User Description: " + description)
				} else if command[1] == "--messages" || command[1] == "-m" {
					ms, err := getUnreadMessages()
					if err != nil {
						println("	Failed to get unread messages")
						return
					}
					// convert ms to string
					println("	Unread Messages: " + strconv.Itoa(ms))
				} else if command[1] == "--email" || command[1] == "-e" {
					email, verified, err := getEmailInfo()
					if err != nil {
						println("	Failed to get email info")
						return
					}
					if verified {
						println("	Email: " + email)
					} else {
						println("	Email: " + email + " (Not Verified)")
					}
				}
			} else if len(command) == 3 {
				if command[1] == "--userid" || command[1] == "-u" {
					_, user, online, err := getUserIDInfo(command[2])
					if err != nil {
						println("	Failed to get user id info")
						return
					}
					if len(user) == 0 {
						println("	User not found")
						return
					}
					if online {
						println("	User: " + user + " (Online)")
					} else {
						println("	User: " + user + " (Offline)")
					}
				} else if command[1] == "--get" || command[1] == "-g" {
					if command[2] == "--rbxid" || command[2] == "-r" {
						rbxid, err := getRBXID()
						if err != nil {
							println("	Failed to get rbxid")
							return
						}
						println("	RBXID: " + rbxid)
					}
				}

			} else {
				println("	Usage: --api (-a) [option]")
				println("	Options:")
				println("		--description (-d) - Gets the user description")
				println("		--messages (-m) - Gets the unread messages")
				println("		--email (-e) - Gets the email info")
				println("		--help (-h) - Shows this help")
				println("		--userid (-u) [user] - Gets the user id info")
				println("		--get (-g) - Gets option")
				println("			--rbxid (-r) - Gets the rbxid")
			}
		} else {
			println("	Unknown OS")
		}
	case "--help", "-h", "?":
		if console.IsWindows() {
			print("	--fix, -f ~ Fixes bugs that happens to Roblox\n")
			print("	--cursor, -c ~ Installs a custom cursor\n")
			print("	--delete, -d ~ Deletes Roblox\n")
			print("	--reinstall, -r ~ Reinstalls Roblox\n")
			print("	--install, -i ~ Installs Roblox\n")
			print("	--tainted, -t ~ Checks if user is tainted\n")
			print("	--versions, -v ~ Prints the latest versions of Roblox and Roblox Studio\n")
			print("	--check, -C ~ Checks/Generate if a username is valid\n")
			print("	--set-token, -s ~ Sets the security token\n")
			print("	--api, -a ~ Gets info about the user\n")
		} else if console.IsMacOS() {
			print("	--install, -i ~ Installs Roblox\n")
			print("	--versions, -v ~ Prints the latest versions of Roblox and Roblox Studio\n")
			print("	--check, -C ~ Checks/Generate if a username is valid\n")
			print("	--set-token, -s ~ Sets the security token\n")
			print("	--api, -a ~ Gets info about the user\n")
		}
	default:
		print("	Unknown command: " + command[0] + "\n")
	}
}
