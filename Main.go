package main

import (
	"bufio"
	"os"
	"pwd0roblox/console"
	"pwd0roblox/network"
	"pwd0roblox/pwdtools"
	"pwd0roblox/roblox"
	"strconv"
	"strings"
	"time"

	gc "github.com/gookit/color"
	"github.com/pterm/pterm"
)

var (
	Version = "1.0.6"
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
	console.ConsoleClear()
	console.SetConsoleTitle("pwd0roblox - " + Version)
	intro()
main_r:
	Header_func()

	for {
		reader := bufio.NewReader(os.Stdin)
		gc.Print("> ")
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
			Information_Command()
		case "proxy", "py":
			Proxy_Command(args)
		case "ex", "quit":
			os.Exit(0)
		case "help", "?":
			Help_Command()
		default:
			pterm.Error.Println("Unknown command: " + args[0])
		}
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

func Header_func() {
	console.ConsoleClear()
	if network.IsConnected() {
		pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
			"pwd0roblox\nVersion: " + Version)
	} else {

		pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
			"pwd0roblox\nVersion: " + Version + "\nNo internet connection, Limited Commands")
	}

	if roblox.NeedUpdate() {
		console.ConsoleClear()
		pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println("\nRoblox Update Available")

		ver, _ := roblox.GetRobloxWindowsVersion()
		roblox.InstallRoblox(ver, true)
		console.ConsoleClear()

		pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
			"pwd0roblox\nVersion: " + Version)
	}
}

func Help_Command() {
	if network.IsConnected() {
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
			{"information", "info", "Information about the program and developers"},
			{"clear", "cls", "Clear the console"},
			{"quit", "ex", "Exit the program"},
			{"help", "?", "Show this help"},
		}).Render()

	}
}

func Proxy_Command(args []string) {
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
}

func Information_Command() {
	if console.IsWindows() {
		hash := pwdtools.GetHash("pwd0roblox.exe")
		size := pwdtools.GetSize("pwd0roblox.exe")
		// convert to string
		sizeString := strconv.FormatInt(size, 10)

		pterm.DefaultSection.Println("Program Information")
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
			{"Os", "Name", "Version", "Hash", "Size"},
			{"Windows", "pwd0roblox", Version, hash, sizeString + "MB"},
		}).Render()
	} else if console.IsMacOS() {
		hash := pwdtools.GetHash("pwd0roblox")
		size := pwdtools.GetSize("pwd0roblox")
		// convert to string
		sizeString := strconv.FormatInt(size, 10)

		pterm.DefaultSection.Println("Program Information")
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
			{"Os", "Name", "Version", "Hash", "Size"},
			{"MacOS", "pwd0roblox", Version, hash, sizeString + "MB"},
		}).Render()
	}
	pterm.DefaultSection.Println("Developers")
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
		{"Username", "Os", "Description"},
		{"calamixy", "MacOS", "MacOS Developer"},
		{"pwd0kernel", "Win", "Windows Developer"},
	}).Render()
}
