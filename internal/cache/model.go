package cache

import (
	"time"
)

type Cache struct {
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Address string    `json:"address"`
	Domain  string    `json:"domain"`
	Exp     time.Time `json:"exp,omitempty"`
}
