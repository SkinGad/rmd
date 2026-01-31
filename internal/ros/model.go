package ros

import (
	"github.com/go-routeros/routeros/v3"
)

type Mikrotik struct {
	Con     *routeros.Client
	Address string
	Port    string
	table   string
	gateway string
}

type Routes struct {
	Address string
	Comment string
}
