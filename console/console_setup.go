package console

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

// ismacos
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// change console size on macos
// width: int
// height: int
func SetConsoleSize(width int, height int) {
	if IsMacOS() {
		cmd := exec.Command("/bin/sh", "-c", "printf '\\e[8;"+strconv.Itoa(height)+";"+strconv.Itoa(width)+"t'")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if IsWindows() {
		// untested
		cmd := exec.Command("cmd", "/c", "mode con: cols="+strconv.Itoa(width)+" lines="+strconv.Itoa(height))
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// change set console title name
// title: string
func SetConsoleTitle(title string) {
	if IsWindows() {
		cmd := exec.Command("cmd", "/c", "title "+title)
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if IsMacOS() {
		cmd := exec.Command("/bin/sh", "-c", "echo -ne '\033]0;"+title+"\007'")
		cmd.Stdout = os.Stdout
		cmd.Run()

		SetConsoleSize(100, 30)
	}
}

// clear console function
func ConsoleClear() {
	if IsWindows() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
