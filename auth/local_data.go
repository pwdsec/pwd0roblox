package auth

import "net"

var hwids = []string{"98a59fce662c7a694bb591ed0c619f9fc269daebadc37e509805317dd7bdf36b"}

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
func CheckLocalData() bool {
	for _, v := range hwids {
		if v == Hash(hashKey, GetHWID()) {
			return true
		}
	}
	return false
}