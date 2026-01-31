package lookup

import (
	"fmt"
	"log/slog"
	"net"
)

func LookUp(domains ...string) (IPs []string) {
	for _, domain := range domains {
		if check, net := isIP(domain); check {
			IPs = append(IPs, net)
			continue
		}
		ips, err := net.LookupHost(domain)
		if err != nil {
			slog.Error("Failed to resolve domain", "domain", domain, "err", err)
		}
		for i, v := range ips {
			ips[i] = fmt.Sprintf("%s/32", v)
		}
		IPs = append(IPs, ips...)
	}
	slog.Debug("lookup domains", "domains", domains, "ips", IPs)
	return
}
