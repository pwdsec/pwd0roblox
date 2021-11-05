package console

import (
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"unsafe"
)

// ismacos
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// change set console title name using windows api
// title: string
func SetConsoleTitle(title string) {
	if IsWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTitleW")
		proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
	} else if IsMacOS() {
		cmd := exec.Command("/bin/sh", "-c", "echo -ne '\033]0;"+title+"\007'")
		cmd.Stdout = os.Stdout
		cmd.Run()
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
