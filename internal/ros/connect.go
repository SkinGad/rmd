package ros

import (
	"fmt"
	"log/slog"

	"github.com/go-routeros/routeros/v3"
)

func Connect(address, login, password, table, gateway, port string) (*Mikrotik, error) {
	conn, err := routeros.Dial(
		fmt.Sprintf("%s:%s", address, port),
		login, password)
	if err != nil {
		slog.Error("Error connected to mikrotik", "error", err)
		return nil, err
	}
	slog.Info("Connected to mikrotik", "address", address, "port", port)
	mik := Mikrotik{
		Con:     conn,
		table:   table,
		gateway: gateway,
		Address: address,
		Port:    port,
	}
	mik.getTable()
	return &mik, err
}

func (m *Mikrotik) HethCheck() (res bool) {
	_, err := m.Con.Run("/interface/print")
	if err == nil {
		res = true
	}
	slog.Info("HelthCheck to mikrotik", "address", m.Address, "port", m.Port, "helthcheck", res)
	return
}
