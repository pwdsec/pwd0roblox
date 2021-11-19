package roblox

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/pterm/pterm"
)

/* working on this */

func GetScriptsHref(html string) ([]string, error) {
	re := regexp.MustCompile(`<a href="(.*?)"`)
	href := re.FindStringSubmatch(html)
	if len(href) == 0 {
		return nil, nil
	}
	return href, nil
}

func CleanHrefs(hrefs []string) []string {
	cleaned := make([]string, 0)
	for _, v := range hrefs {
		if !strings.Contains(v, "author/admin/") || strings.Contains(v, "https://robloxscripts.com/") {
			cleaned = append(cleaned, v)
		}
	}
	return cleaned
}

func GetScripts(search_data string) {
	req, err := http.NewRequest("GET", "https://robloxscripts.com/?s="+search_data, nil)
	if err != nil {
		println(err.Error())
		return
	}
	req.AddCookie(&http.Cookie{Name: "SirMemeNetworks", Value: "U2FsdGVkX19Lfj/GmxESHfkUFmwmxtooFB3EFAx0+mh8CTrvAVsGZYbS0FjDu2RECwNBiATHiKxO9YLDjLWVMu991UZlg+e7COZGpN6OOk/gS61b8u25BAc7yAQTPWCxRhWbJv8gY9NBdKaVfPuDPlmv0nh4xMDzJNTN7khn+h3Wcm8VbDIlSFS3lauWM2oidx4wSwDYyS+xQ/T/O4iOWWszbHpYmp04vZCNSk1OzRLkckKlexbdJT6ehp77h9D2VaT6yHnyiNNFCprs1mK9arIbB06uCktdUsjoIH4V+JxTgcUh9BMT/7hzEK5ZosNDPMVuzFr538KbyFnrEz2lFTTH/W1ZOK2nWEhXktJiNXE="})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		pterm.Error.Println("Failed to read html")
		return
	}

	hrefs, err := GetScriptsHref(string(body))
	if err != nil {
		pterm.Error.Println("Failed to get hrefs")
		return
	}

	for _, v := range hrefs {
		pterm.Info.Println(v)
	}
}
