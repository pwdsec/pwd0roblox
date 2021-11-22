package roblox

import (
	"fmt"

	goxymemmory "github.com/Xustyx/goxymemory"
)

// edit memory of a process
func Crash() {
	dm := goxymemmory.DataManager("RobloxPlayerBeta.exe")
	if !dm.IsOpen {
		fmt.Printf("Failed opening process.\n")
		return
	}

	err := dm.Write(0x001,
		goxymemmory.Data{25600, goxymemmory.BYTE})
	if err != nil { //Check if not failed.
		fmt.Printf("Failed writing memory. %s", err)
	}
}
