package cache

import (
	"log/slog"
	"time"
)

func (c *Cache) AddAddresses(rmdTTL time.Duration, domain string, ips ...string) {
	var update bool
	for _, w := range ips {
		var check bool
		for _, v := range c.Addresses {
			if v.Address == w {
				check = true
				break
			}
		}
		if check {
			break
		}
		exp := time.Now().Add(rmdTTL)
		c.Addresses = append(c.Addresses,
			Address{
				Address: w,
				Domain:  domain,
				Exp:     exp,
			},
		)
		slog.Debug("Add address to cache", "address", w, "domain", domain)
		update = true
	}
	if update {
		c.Save()
	}
}
