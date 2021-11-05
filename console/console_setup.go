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
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTitleW")
	proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
}

// clear console function
func ConsoleClear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
