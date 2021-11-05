package roblox

import (
	"io/ioutil"
	"net/http"
	"os"
)

// cursor data array.
//
// string name.
// string url.
// string url.
var data = [][]string{
	{
		"default",
		"https://github.com/pwd0kernel/NormalCursor/blob/main/ArrowCursor.png?raw=true",
		"https://github.com/pwd0kernel/NormalCursor/blob/main/ArrowFarCursor.png?raw=true",
	},
	{
		"3dpixel",
		"https://cdn.custom-cursor.com/db/4827/32/arrow2747.png",
		"https://cdn.custom-cursor.com/db/4828/32/arrow2747.png",
	},
	{
		"glitch",
		"https://cdn.custom-cursor.com/db/8625/32/starter-glitch-pointer.png",
		"https://cdn.custom-cursor.com/db/8626/32/starter-glitch-cursor.png",
	},
	{
		"red",
		"https://cdn.custom-cursor.com/db/pointer/32/Red_pointer.png",
		"https://cdn.custom-cursor.com/db/cursor/32/Red_curson.png",
	},
	{
		"blue",
		"https://cdn.custom-cursor.com/db/pointer/32/Blue_pointer.png",
		"https://cdn.custom-cursor.com/db/cursor/32/Blue_cursor.png",
	},
	{
		"green",
		"https://cdn.custom-cursor.com/db/pointer/32/Green_pointer.png",
		"https://cdn.custom-cursor.com/db/cursor/32/Green_cursor.png",
	},
	{
		"hexa",
		"https://cdn.custom-cursor.com/db/9349/32/starter-hexagonal-modern-technology-pointer.png",
		"https://cdn.custom-cursor.com/db/9350/32/starter-hexagonal-modern-technology-cursor.png",
	},
	{
		"3dgeo",
		"https://cdn.custom-cursor.com/db/9231/32/starter-3d-geometric-figures-pointer.png",
		"https://cdn.custom-cursor.com/db/9232/32/starter-3d-geometric-figures-cursor.png",
	},
	{
		"slime",
		"https://cdn.custom-cursor.com/db/8125/32/starter-green-slime-pointer.png",
		"https://cdn.custom-cursor.com/db/8126/32/starter-green-slime-cursor.png",
	},
	{
		"abstract",
		"https://cdn.custom-cursor.com/db/7457/32/starter-abstract-space-pointer.png",
		"https://cdn.custom-cursor.com/db/7458/32/starter-abstract-space-cursor.png",
	},
	{
		"dark-irdescent",
		"https://cdn.custom-cursor.com/db/6617/32/starter-dark-iridescent-pointer-a.png",
		"https://cdn.custom-cursor.com/db/6618/32/starter-dark-iridescent-cursor-a.png",
	},
	{
		"modern-green",
		"https://cdn.custom-cursor.com/db/pointer/32/Modern_Green.png",
		"https://cdn.custom-cursor.com/db/cursor/32/Modern_Green.png",
	},
	{
		"neon-mushroom",
		"https://cdn.custom-cursor.com/db/10133/32/neon-mushroom-pointer.png",
		"https://cdn.custom-cursor.com/db/10134/32/neon-mushroom-cursor.png",
	},
}

// cursor list function.
func CursorsList() []string {
	var list []string
	for _, v := range data {
		list = append(list, v[0])
	}
	return list
}

// CursorInstaller function.
// error if not found.
// string cursor name.
func CursorsInstaller(cursor string) error {
	println("	[+] Installing " + cursor + " cursor")
	var ArrowCursor string
	var ArrowFarCursor string
	for i, v := range data {
		if v[0] == cursor {
			ArrowCursor = data[i][1]
			ArrowFarCursor = data[i][2]
		}
	}
	a, b := os.UserCacheDir()
	version, err := GetRobloxWindowsVersion()
	if err != nil {
		return err
	}

	if b != nil {
		println("	Failed to get user cache directory")
		return nil
	}

	files, err := ioutil.ReadDir(a + "\\Roblox\\Versions\\" + version + "\\content\\textures\\Cursors\\KeyboardMouse\\")
	if err != nil {
		print(err)
	}

	for _, f := range files {
		if f.Name() == "ArrowCursor.png" || f.Name() == "ArrowFarCursor.png" {
			os.Remove(a + "\\Roblox\\Versions\\" + version + "\\content\\textures\\Cursors\\KeyboardMouse\\" + f.Name())
		}
	}

	resp, err := http.Get(ArrowCursor)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(a+"\\Roblox\\Versions\\"+version+"\\content\\textures\\Cursors\\KeyboardMouse\\ArrowCursor.png", body, 0644)
	if err != nil {
		return err
	}
	// download another cursor
	resp, err = http.Get(ArrowFarCursor)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(a+"\\Roblox\\Versions\\"+version+"\\content\\textures\\Cursors\\KeyboardMouse\\ArrowFarCursor.png", body, 0644)
	if err != nil {
		return err
	}
	println("	[+] Done")
	// return
	return nil
}
