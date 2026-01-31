package cache

import (
	"time"
)

func (c *Cache) UpdateAddresses(rmdTTL time.Duration, domain string, ips ...string) {
	var update bool
	for _, ip := range ips {
		var needAdded bool = true
		for i, v := range c.Addresses {
			if v.Address == ip && v.Domain == domain {
				c.Addresses[i] = Address{
					Address: ip,
					Domain:  domain,
					Exp:     time.Now().Add(rmdTTL),
				}
				needAdded = false
				update = true
				break
			}
		}
		if needAdded {
			c.AddAddresses(rmdTTL, domain, ip)
		}
	}
	if update {
		c.Save()
	}
}
