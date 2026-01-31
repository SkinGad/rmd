package cache

import (
	"log/slog"
	"time"
)

func (c *Cache) FindClear() (addresses []Address) {
	for _, v := range c.Addresses {
		if v.Exp.Unix() <= time.Now().Unix() {
			addresses = append(addresses, v)
		}
	}
	if len(addresses) > 0 {
		slog.Debug("Need deleted addresses", "addresses", addresses)
	}
	return
}

func (c *Cache) Clear(address string) {
	for i, v := range c.Addresses {
		if v.Address == address {
			c.Addresses = append(c.Addresses[:i], c.Addresses[i+1:]...)
			slog.Debug("Deleted address", "address", address)
			break
		}
	}
	c.Save()
}
