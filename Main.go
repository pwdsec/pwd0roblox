package main

import (
	"bufio"
	"os"
	"pwd0roblox/console"
	"strings"

	"github.com/gookit/color"
)

func main() {
	console.SetConsoleTitle("pwd0roblox - 1.0.0")
	color.Print("<red>  ________________________</>\n")
	color.Print("  <red>|</>----------------------<red>|</><red> ___/|</>		\n")
	color.Print("  <red>|</>------[ <red>PWD</>.<red>RX</> ]------<red>|</><red> \\o.O|</>		\n")
	color.Print("  <red>|</>--[ <red>Version</>: <red>1.0.0</> ]--<red>|</><red> (___)</>\n")
	color.Print("  <red>|</>----------------------<red>|</><red>   U</>\n")
	color.Print("  <red>|</>-------[ <red>Help</> ]-------<red>|</>\n")
	color.Print("<red>  ●●●●●●●●●●●●●●●●●●●●●●●●</>\n\n")

	for {
		reader := bufio.NewReader(os.Stdin)
		color.Print("[<red>●</>]> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		args := strings.Fields(text)

		switch strings.ToLower(args[0]) {
		case "cls", "clear":
			console.ConsoleClear()
		default:
			color.Println("<red>Unknown command</>")
		}
	}
}
