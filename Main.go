package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"pwd0roblox/auth"
	"pwd0roblox/console"
	"pwd0roblox/network"
	"pwd0roblox/pwdtools"
	"pwd0roblox/roblox"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

var (
	Version         = "1.0.5"
	is_auth_enabled = true
	is_admin        = false
	hashKey         = "046634"
)

var logo string = `
██████╗ ██╗    ██╗██████╗  ██████╗ ██████╗  ██████╗ ██████╗ ██╗      ██████╗ ██╗  ██╗
██╔══██╗██║    ██║██╔══██╗██╔═████╗██╔══██╗██╔═══██╗██╔══██╗██║     ██╔═══██╗╚██╗██╔╝
██████╔╝██║ █╗ ██║██║  ██║██║██╔██║██████╔╝██║   ██║██████╔╝██║     ██║   ██║ ╚███╔╝ 
██╔═══╝ ██║███╗██║██║  ██║████╔╝██║██╔══██╗██║   ██║██╔══██╗██║     ██║   ██║ ██╔██╗ 
██║     ╚███╔███╔╝██████╔╝╚██████╔╝██║  ██║╚██████╔╝██████╔╝███████╗╚██████╔╝██╔╝ ██╗
╚═╝      ╚══╝╚══╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝
 `

func main() {
	is_auth, user := auth.CheckLocalData()
	if user == "pwd0kernel" {
		is_admin = true
	}
	console.ConsoleClear()
	console.SetConsoleTitle("pwd0roblox - " + Version)
	intro()
main_r:
	console.ConsoleClear()
	if network.IsConnected() {
		if is_admin {
			pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
				"pwd0roblox\nVersion: " + Version + "\nWelcome: " + user + " [ADMIN]")
		} else {
			pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
				"pwd0roblox\nVersion: " + Version + "\nWelcome: " + user)
		}
	} else {
		if is_admin {
			pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
				"pwd0roblox\nVersion: " + Version + "\nWelcome: " + user + " [ADMIN]" + "\nNo internet connection, Limited Commands")
		} else {
			pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
				"pwd0roblox\nVersion: " + Version + "\nWelcome: " + "\nNo internet connection, Limited Commands")
		}
	}

	if is_auth || !is_auth_enabled {
		if roblox.NeedUpdate() {
			console.ConsoleClear()
			pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println("\nRoblox Update Available")

			ver, _ := roblox.GetRobloxWindowsVersion()
			roblox.InstallRoblox(ver, true)
			console.ConsoleClear()

			pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
				"pwd0roblox\nVersion: " + Version + "\nWelcome: " + user)
		}

		for {
			reader := bufio.NewReader(os.Stdin)
			color.Print("> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			args := strings.Fields(text)

			switch strings.ToLower(args[0]) {
			case "roblox", "rbx":
				roblox.CommandHandler(args[1:])
			case "cls", "clear":
				console.ConsoleClear()
				goto main_r
			case "information", "info":
				if console.IsWindows() {
					hash := pwdtools.GetHash("pwd0roblox.exe")
					size := pwdtools.GetSize("pwd0roblox.exe")
					// convert to string
					sizeString := strconv.FormatInt(size, 10)

					pterm.DefaultSection.Println("Program Information")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Os", "Name", "Version", "Hash", "Size", "Is Auth Enabled"},
						{"Windows", "pwd0roblox", Version, hash, sizeString + "MB", strconv.FormatBool(is_auth_enabled)},
					}).Render()
				} else if console.IsMacOS() {
					hash := pwdtools.GetHash("pwd0roblox")
					size := pwdtools.GetSize("pwd0roblox")
					// convert to string
					sizeString := strconv.FormatInt(size, 10)

					pterm.DefaultSection.Println("Program Information")
					pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
						{"Os", "Name", "Version", "Hash", "Size", "Is Auth Enabled"},
						{"MacOS", "pwd0roblox", Version, hash, sizeString + "MB", strconv.FormatBool(is_auth_enabled)},
					}).Render()
				}
				pterm.DefaultSection.Println("Developers")
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Username", "Os", "Description"},
					{"calamixy", "MacOS", "MacOS Developer"},
					{"pwd0kernel", "Win", "Windows Developer"},
				}).Render()
			case "proxy", "py":
				if len(args) == 2 {
					if strings.Contains(args[1], ".txt") {
						list := pwdtools.GetProxyList("http://" + args[1])
						if len(list) != 0 {
							pwdtools.WriteProxyList("checked.txt", list)
						} else {
							pterm.Warning.Println("No proxies found")
						}
					} else {
						if pwdtools.CheckProxy(args[1]) {
							pterm.Success.Println("Proxy is valid: " + args[1])
						} else {
							pterm.Warning.Println("Proxy is invalid: " + args[1])
						}
					}
				} else {
					pterm.Info.Println("Usage: proxy [file.txt]")
				}
			case "ex", "quit":
				os.Exit(0)
			case "decode-file", "dec-f":
				if is_admin {
					if len(args) == 2 {
						if strings.Contains(args[1], ".key") {
							// read file from args[1]
							file, err := ioutil.ReadFile(args[1])
							if err != nil {
								pterm.Error.Println(err)
								break
							}

							// remove make it one line
							file = bytes.Replace(file, []byte("\n"), []byte(""), -1)
							println(string(file))

							// decode file
							decoded := auth.Base64Decode(string(file))
							pterm.Info.Println(decoded)
						}
					} else {
						pterm.Info.Println("Usage: decode-file [file.key]")
					}
				} else {
					pterm.Warning.Println("You are not an admin")
				}
			case "help", "?":
				if network.IsConnected() {
					if is_admin {
						pterm.DefaultSection.Println("Admin Commands")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Short", "Description"},
							{"decode", "dec", "decode text, decode (abcd)"},
							{"decode-file", "dec-f", "decode text, decode (path)"},
							{"encode", "enc", "encode text, encode (abcd)"},
							{"encode-file", "enc-f", "encode file, encode (path)"},
						}).Render()

						pterm.DefaultSection.Println("User Commands")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Short", "Description"},
							{"roblox", "rbx", "Run Roblox commands, --help (-h)"},
							{"proxy", "py", "Proxy Checker"},
							{"information", "info", "Information about the program and developers"},
							{"clear", "cls", "Clear the console"},
							{"quit", "ex", "Exit the program"},
							{"help", "?", "Show this help"},
						}).Render()
					} else {
						pterm.DefaultSection.Println("User Commands")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Short", "Description"},
							{"roblox", "rbx", "Run Roblox commands, --help (-h)"},
							{"proxy", "py", "Proxy Checker"},
							{"information", "info", "Information about the program and developers"},
							{"clear", "cls", "Clear the console"},
							{"quit", "ex", "Exit the program"},
							{"help", "?", "Show this help"},
						}).Render()
					}
				} else {
					if is_admin {
						pterm.DefaultSection.Println("Admin Commands")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Short", "Description"},
							{"decode", "dec", "decode text, decode (abcd)"},
							{"encode", "enc", "encode text, decode (abcd)"},
						}).Render()

						pterm.DefaultSection.Println("User Commands")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Short", "Description"},
							{"roblox", "rbx", "Run Roblox commands, --help (-h)"},
							{"information", "info", "Information about the program and developers"},
							{"clear", "cls", "Clear the console"},
							{"quit", "ex", "Exit the program"},
							{"help", "?", "Show this help"},
						}).Render()
					} else {
						pterm.DefaultSection.Println("User Commands")
						pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
							{"Command", "Short", "Description"},
							{"roblox", "rbx", "Run Roblox commands, --help (-h)"},
							{"information", "info", "Information about the program and developers"},
							{"clear", "cls", "Clear the console"},
							{"quit", "ex", "Exit the program"},
							{"help", "?", "Show this help"},
						}).Render()
					}
				}
			default:
				pterm.Error.Println("Unknown command: " + args[0])
			}
		}
	} else {

		console.ConsoleClear()
		reader := bufio.NewReader(os.Stdin)
		color.Print("Please type your username: ")
		username, _ := reader.ReadString('\n')
		username = strings.Replace(username, "\n", "", -1)
		console.ConsoleClear()
		// create folder "auth"
		if _, err := os.Stat("auth"); os.IsNotExist(err) {
			os.Mkdir("auth", 0777)
		}

		// write file to auth folder
		if _, err := os.Stat("auth/auth.key"); os.IsNotExist(err) {
			file, err := os.Create("auth/auth.key")
			if err != nil {
				pterm.Error.Println("Error creating auth file: " + err.Error())
			}
			data_key := auth.Base64Encode("Hash: " + auth.Hash(hashKey, auth.GetHWID()) + " | Username: " + username)
			for i := 0; i < len(data_key); i += 100 {
				if i+100 > len(data_key) {
					file.WriteString(data_key[i:])
				} else {
					file.WriteString(data_key[i:i+100] + "\n")
				}
			}
			file.Close()
		}

		data_key := auth.Base64Encode("Hash: " + auth.Hash(hashKey, auth.GetHWID()) + " | Username: " + username)

		var new_key string

		// make a new line each 50 characters and append new_key
		for i := 0; i < len(data_key); i += 50 {
			if i+50 > len(data_key) {
				new_key += data_key[i:]
			} else {
				new_key += data_key[i:i+50] + "\n"
			}
		}

		// parse username with regex only alphanumeric characters
		username_regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
		username = username_regex.ReplaceAllString(username, "")

		pterm.Warning.Println("You are not whitelisted please send the key file to pwd0kernel\nor you can send the key below")

		pterm.DefaultSection.Println("User Info")
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
			{"Path", "Hash", "Username"},
			{"auth/auth.key", auth.Hash(hashKey, auth.GetHWID()), username},
		}).Render()

		pterm.DefaultSection.Println("Key Info")
		pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
			{"============== Key =============="},
			{new_key},
		}).Render()

		wait_reader := bufio.NewReader(os.Stdin)
		wait_reader.ReadString('\n')
		os.Exit(0)
	}
}

func intro() {
	pterm.DefaultCenter.Print(pterm.NewRGB(249, 159, 255).Sprintf(logo))
	pterm.DefaultCenter.Print(pterm.NewRGB(249, 159, 255).Sprintf("Version: " + Version))
	println()
	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 5 seconds...")
	time.Sleep(1000 * time.Millisecond)
	for i := 4; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(1000 * time.Millisecond)
	}
	introSpinner.Stop()
}
