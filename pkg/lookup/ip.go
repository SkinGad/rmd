package lookup

import (
	"net"
)

func isIP(value string) (bool, string) {
	ip, net, _ := net.ParseCIDR(value)
	return ip != nil, net.String()
}
