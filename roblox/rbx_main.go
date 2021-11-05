// roblox functions package.
// this package contains fixes, commands, and other functions.
package roblox

import (
	"io/ioutil"
	"math/rand"
	"os"
	"pwd0roblox/console"
	"strconv"
	"strings"
)

// fixes the Unexpected Behavior Kick that happens when you open Roblox.
// this is a workaround for a bug in Roblox.
// this method is not 100% reliable.
func Fix_Unexpected_Behavior_Kick_method_1() {
	print("	[+] Fixing Unexpected Behavior Kick...\n")
	a, b := os.UserCacheDir()
	if b != nil {
		println("	Failed to get user cache directory")
		return
	}

	files, err := ioutil.ReadDir(a + "\\Roblox\\")
	if err != nil {
		print(err)
		return
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "GlobalBasicSettings") {
			os.Remove(a + "\\Roblox\\" + file.Name())
		}
	}
	print("	[+] Successfully fixed Unexpected Behavior Kick\n")
}

// checks if roblox is running.

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

		println("Roblox Windows Version: " + roblox_windows_version)
		println("Roblox Studio Windows Version: " + roblox_studio_windows_version)
		println("Roblox Mac Version: " + roblox_mac_version)
		println("Roblox Studio Mac Version: " + roblox_studio_mac_version)
		println("Roblox Studio Qt Version: " + roblox_studio_qt_version)
	case "--delete", "-d":
		if console.IsWindows() {
			DeleteRoblox()
		} else if console.IsMacOS() {
			println("	MacOS is not yet supported")
		} else {
			println("	Unknown OS")
		}
	case "--install", "-i":
		if console.IsWindows() {
			if len(command) == 2 {
				if command[1] == "-h" {
					println("Usage: --install (-i) [version] (-s)[start] ")
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
					println("Usage: --install (-i) [version] (-s)[start] ")
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
				InstallRobloxMac(ver, true)
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
				println("		--generate (-g) [lenght] - Checks if a username is valid")
				println("		--username (-u) [username] - Checks if a username is valid")
				println("		--normal (-n) [how many] - Checks if a username is valid")
			}
		} else {
			println("	Usage: --check (-c) [option]")
			println("	Options:")
			println("		--generate (-g) [lenght] - Checks if a username is valid")
			println("		--username (-u) [username] - Checks if a username is valid")
			println("		--normal (-n) [how many] - Checks if a username is valid")
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
	case "--help", "-h", "?":
		if console.IsWindows() {
			print("	--fix, -f ~ Fixes stuff that happens when you open Roblox etc\n")
			print("	--cursor, -c ~ Installs a custom cursor\n")
			print("	--delete, -d ~ Deletes Roblox\n")
			print("	--install, -i ~ Installs Roblox\n")
			print("	--tainted, -t ~ Checks if user is tainted\n")
			print("	--versions, -v ~ Prints the latest versions of Roblox and Roblox Studio\n")
			print("	--check, -C ~ Checks if a username is valid\n")
		} else if console.IsMacOS() {
			print("	--versions, -v ~ Prints the latest versions of Roblox and Roblox Studio\n")
			print("	--check, -C ~ Checks if a username is valid\n")
		}
	default:
		print("	Unknown command: " + command[0] + "\n")
	}
}
