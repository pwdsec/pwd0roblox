package auth

import "net"

var hwids = [][]string{{"011bbfcaec0e2fe11f935af43a634d75f2e81300d53e9fe81d6715373e146edb", "pwd0kernel"}}

func GetHWID() string {
	hwid := ""
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.HardwareAddr != nil {
			hwid += iface.HardwareAddr.String()
		}
	}
	return hwid
}

// check all local data for auth
func CheckLocalData() (bool, string) {
	for _, v := range hwids {
		if v[0] == Hash(hashKey, GetHWID()) {
			return true, v[1]
		}
	}
	return false, "Unknown"
}
