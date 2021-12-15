package roblox

import (
	"os"
	"pwd0roblox/network"
	"strconv"

	"github.com/pterm/pterm"
)

func Fix_Command_Windows(command []string) {
	if len(command) == 2 {
		if command[1] == "UBK" {
			Fix_Unexpected_Behavior_Kick_method_1()
		} else {
			pterm.Error.Println("Unknown fix: " + command[1])
		}
	} else {
		pterm.Info.Println("Usage: --fix (-f) [option]")
		pterm.Info.Println("Options:")
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
			{"Command", "Description", "Reliability"},
			{"UBK", "Fixes Unexpected Behavior Kick", "50%"},
		}).Render()
	}
}

func Install_Command_Windows(command []string) {
	if network.IsConnected() {
		if len(command) == 3 {
			if command[1] == "--content" || command[1] == "-c" {
				ContentInstaller_Ziped(command[2])
			}
		} else if len(command) == 2 {
			if command[1] == "--help" || command[1] == "-h" {
				pterm.Info.Println("Usage: --install (-i) [option] ")
				pterm.Info.Println("Options:")
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Command", "Single", "Description"},
					{"--content", "-c", "Install the content"},
					{"--help", "-h", "This help message"},
				}).Render()
			} else if command[1] == "--content" || command[1] == "-c" {
				ver, _ := GetRobloxWindowsVersion()
				ContentInstaller_Ziped(ver)
				DeleteEmptyFiles(ver)
			} else if command[1] == "--versions" || command[1] == "-v" {
				pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
					{"Date", "Version"},
					{"06-16-2021", "version-c864da271a4c44ea"},
					{"06-11-2021", "version-7d96d7dad25d49f1"},
					{"04-29-2021", "version-0658018801724832"},
					{"04-08-2021", "version-278f0258a7224831"},
					{"01-13-2021", "version-d5212926da8e4716"},
					{"11-13-2020", "version-aa7766fcc7cb4906"},
					{"04-10-2019", "version-9f8314ad67c64c0d"},
					{"10-29-2018", "version-e9d1a6c5df10490c"},
					{"12-12-2017", "version-45cc144b134647ea"},
					{"02-25-2016", "version-a1b8c1edf45b4959"},
				}).Render()
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
	} else {
		pterm.Error.Println("No internet connection")
	}
}

func Tainted_Command_Windows(command []string) {
	if len(command) == 2 {
		if command[1] == "-h" {
			pterm.Info.Println("Usage: --tainted (-t) [version]")
		}
	} else {
		var is_tainted bool = false
		var TaintingModule string = ""
		ini_files := GetINIFiles()
		for _, v := range ini_files {
			mapped := ReadINIFile(v)
			if IsTainted(mapped) {
				is_tainted = true
				TaintingModule = GetTaintingModule(mapped)
			} else {
				is_tainted = false
			}
		}
		if is_tainted {
			pterm.Warning.Println("User Tainted, Tainting Module: " + TaintingModule)
		} else {
			pterm.Success.Println("User Not Tainted")
		}
	}
}

func Game_IP_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			pterm.Success.Println(Get_IP_Address(Get_Log()))
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Error.Println("You are not connected to the internet")
	}
}

func Game_Info_Command_windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			pterm.DefaultSection.Println("Game Information")
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"IP Address", "Place ID", "Place Name", "Session ID"},
				{Get_IP_Address(Get_Log()), Get_Place_ID(Get_Log()), Get_Pace_Name(Get_Place_ID(Get_Log())), Get_Session_ID(Get_Log())},
			}).Render()

			username, err := getUserIDInfo(Get_User_ID(Get_Log()))
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}
			pterm.DefaultSection.Println("Player Information")
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"User ID", "Username"},
				{Get_User_ID(Get_Log()), username},
			}).Render()

			pterm.DefaultSection.Println("More Information")
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"Replicator", "RakNet Socket", "Universe ID"},
				{Get_Replicator_ID(Get_Log()), Get_RakNet_IP_Address(Get_Log()), Get_Universe_ID(Get_Log())},
			}).Render()
		}
	} else {
		pterm.Error.Println("You are not connected to the internet")
	}
}

func Is_Connection_lost_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			if IsConnectionLost(Get_Log()) {
				pterm.Warning.Println("Connection Lost")
			} else {
				pterm.Success.Println("Connection is not lost")
			}
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Warning.Println("You are not connected to the internet")
	}
}

func RakNet_Socket_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"RakNet Socket"},
				{Get_RakNet_IP_Address(Get_Log())},
			}).Render()
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Error.Println("You are not connected to the internet")
	}
}

func Game_Replicator_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"Replicator"},
				{Get_Replicator_ID(Get_Log())},
			}).Render()
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Error.Println("You are not connected to the internet")
	}
}

func Game_Map_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"Place ID", "Place Name"},
				{Get_Place_ID(Get_Log()), Get_Pace_Name(Get_Place_ID(Get_Log()))},
			}).Render()
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Warning.Println("Not connected to internet")
	}
}

func Game_LocalPlayer_Command_Windows() {
	if network.IsConnected() {
		if IsProcessRunning("RobloxPlayerBeta.exe") {
			username, err := getUserIDInfo(Get_User_ID(Get_Log()))
			if err != nil {
				pterm.Error.Println(err.Error())
				return
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
				{"User ID", "Username"},
				{Get_User_ID(Get_Log()), username},
			}).Render()
		} else {
			pterm.Warning.Println("Roblox is not running")
		}
	} else {
		pterm.Error.Println("You are not connected to the internet")
	}
}

func GlobalBasicSettings_Command_Windows() {
	appdata_local, b := os.UserCacheDir()
	if b != nil {
		pterm.Error.Println(b.Error())
		return
	}

	xml, err := ReadXML(appdata_local + "\\Roblox\\GlobalBasicSettings_13.xml")
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	bools, err := ReadBoolAll(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	ints, err := ReadIntAll(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	tokens, err := ReadTokenAll(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	binarystring, err := ReadBinaryStringAll(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	int64s, err := ReadInt64All(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	stringmap, err := ReadStringMapAll(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	floats, err := ReadFloatAll(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	pterm.DefaultSection.Println("Global Basic Settings")
	d := pterm.TableData{{"Name", "Value"}}
	for _, v := range bools {
		d = append(d, []string{v.Name, strconv.FormatBool(v.Value)})
	}

	for _, v := range ints {
		d = append(d, []string{v.Name, strconv.Itoa(v.Value)})
	}

	for _, v := range tokens {
		d = append(d, []string{v.Name, v.Value})
	}

	for _, v := range binarystring {
		d = append(d, []string{v.Name, v.Value})
	}

	for _, v := range int64s {
		d = append(d, []string{v.Name, strconv.FormatInt(v.Value, 10)})
	}

	for _, v := range stringmap {
		d = append(d, []string{v.Name, v.Value})
	}

	for _, v := range floats {
		d = append(d, []string{v.Name, strconv.FormatFloat(v.Value, 'f', -1, 64)})
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).WithBoxed().Render()

	vec2, err := ReadVector2All(xml)
	if err != nil {
		pterm.Error.Println(err.Error())
		return
	}

	/* == Working on this == */
	pterm.DefaultSection.Println("Vector2")
	vec := pterm.TableData{{"Name", "X", "Y"}}
	for _, v := range vec2 {
		vec = append(vec, []string{v.Name, strconv.FormatFloat(v.X, 'f', -1, 64), strconv.FormatFloat(v.Y, 'f', -1, 64)})
	}
	pterm.DefaultTable.WithHasHeader().WithData(vec).WithBoxed().Render()
	/* ===================== */
}
