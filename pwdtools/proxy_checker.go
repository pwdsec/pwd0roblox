package pwdtools

import (
	"bufio"
	"net/http"
	"net/url"
	"os"

	"github.com/pterm/pterm"
)

// proxy checker
// if proxy is valid, return true
// if proxy is invalid, return false
func CheckProxy(proxy string) bool {
	var (
		url, err = url.Parse(proxy)
		client   = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(url),
			},
		}
	)
	if err != nil {
		return false
	}
	_, err = client.Get("http://www.google.com")
	if err != nil {
		return false
	}
	return true
}

// ReadLines
func ReadLines(path string) []string {
	var (
		lines []string
		line  string
	)
	file, err := os.Open(path)
	if err != nil {
		return lines
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// get proxy list from file
func GetProxyList(file string) []string {
	var (
		proxyList []string
		proxy     string
	)
	proxyList = ReadLines(file)
	p, _ := pterm.DefaultProgressbar.WithTotal(len(proxyList)).WithTitle("Checking proxies").Start()
	for _, proxy = range proxyList {
		if CheckProxy("http://" + proxy) {
			pterm.Success.Println("Good: " + proxy)
			proxyList = append(proxyList, proxy)
			p.Increment()
		} else {
			pterm.Warning.Println("Bad: " + proxy)
		}
	}
	p.Stop()
	return proxyList
}

// write proxy list to file
func WriteProxyList(file string, proxyList []string) {
	var (
		fileObj *os.File
		err     error
	)
	fileObj, err = os.Create(file)
	if err != nil {
		return
	}
	defer fileObj.Close()
	for _, proxy := range proxyList {
		fileObj.WriteString(proxy + "\n")
	}
}
