package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"pwd0roblox/console"
	"pwd0roblox/roblox"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

var (
	Version = "1.0.5"
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
	console.ConsoleClear()
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.Color(0))).Println(
		"pwd0roblox\nVersion: " + Version)
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
		case "information", "info":
			if console.IsWindows() {
				hash := getHash("pwd0roblox.exe")
				size := getSize("pwd0roblox.exe")
				// convert to string
				sizeString := strconv.FormatInt(size, 10)

				pterm.DefaultSection.Println("Program Information")
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Os", "Name", "Version", "Hash", "Size"},
					{"Windows", "pwd0roblox", Version, hash, sizeString + "MB"},
				}).Render()
			} else if console.IsMacOS() {
				hash := getHash("pwd0roblox")
				size := getSize("pwd0roblox")
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
		case "ex", "quit":
			os.Exit(0)
		case "help", "?":
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"Command", "Short", "Description"},
				{"roblox", "rbx", "Run Roblox commands, --help (-h)"},
				{"information", "info", "Information about the program and developers"},
				{"clear", "cls", "Clear the console"},
				{"quit", "ex", "Exit the program"},
				{"help", "?", "Show this help"},
			}).Render()
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
	for i := 5; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(1000 * time.Millisecond)
	}
	introSpinner.Stop()
}

// get md5 hash of file
func getHash(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// get size of file in mb
func getSize(file string) int64 {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0
	}
	return fileInfo.Size() / 1024 / 1024
}
