package roblox

import (
	"io/ioutil"
	"net/http"
	"os"
)

// get roblox version from http://setup.roblox.com/version.
// error if not found.
// string roblox version.
func GetRobloxVersion() (string, error) {
	resp, err := http.Get("http://setup.roblox.com/version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

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
	println("[+] Installing " + cursor + " cursor")
	var ArrowCursor string
	var ArrowFarCursor string
	for i, v := range data {
		if v[0] == cursor {
			ArrowCursor = data[i][1]
			ArrowFarCursor = data[i][2]
		}
	}
	a, b := os.UserCacheDir()
	version, err := GetRobloxVersion()
	if err != nil {
		return err
	}

	if b != nil {
		println("Failed to get user cache directory")
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
	println("[+] Done")
	// return
	return nil
}
