package roblox

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/pterm/pterm"
)

// http://assetdelivery.roblox.com/v1/asset?id=1018966

// random number generator function
func RandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

// create asset directory
func AssetDirectory() {
	if _, err := os.Stat("./assets"); os.IsNotExist(err) {
		os.Mkdir("./assets", 0777)
	}
}

func AssetDownload(id string) {
	AssetDirectory()
	downloading, _ := pterm.DefaultSpinner.Start("Downloading Asset...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://assetdelivery.roblox.com/v1/asset?id="+id, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", ROBLOSECURITY)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("assets\\"+id, body, 0644)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	downloading.Success("Downloaded!")
}

func AssetBruteforce() {
	AssetDirectory()
	downloading, _ := pterm.DefaultSpinner.Start("BruteForce Assets...")
	client := &http.Client{}
	// convert RandomNumber(1, 1) to string
	id := strconv.Itoa(RandomNumber(1000, 1000000))
	req, err := http.NewRequest("GET", "https://assetdelivery.roblox.com/v1/asset?id="+id, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", ROBLOSECURITY)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if len(body) == 0 || string(body) == `{"errors":[{"code":404,"message":"Request asset was not found"}]}` || string(body) == `{"errors":[{"code":409,"message":"User is not authorized to access Asset."}]}` || string(body) == `{"errors":[{"code":403,"message":"Asset is not approved for the requester"}]}` {
		downloading.Warning("Failed! :" + id)
	} else {
		err = ioutil.WriteFile("assets\\"+id, body, 0644)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		downloading.Success("Downloaded! :" + id)
	}
}
