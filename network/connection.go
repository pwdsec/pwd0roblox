package network

import (
	"net"
)

func IsConnected() bool {
	if _, err := net.Dial("tcp", "google.com:80"); err != nil {
		return false
	}
	return true
}
