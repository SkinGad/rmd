package ros

import (
	"fmt"
	"log/slog"

	"github.com/go-routeros/routeros/v3"
)

func (m *Mikrotik) GetRouts() (routes []Routes) {
	reply := m.getRouts()
	for _, re := range reply.Re {
		if re.Map["disabled"] == "true" {
			continue
		}
		routes = append(routes, Routes{
			Address: re.Map["dst-address"],
			Comment: re.Map["comment"],
		})
	}
	return
}

func (m *Mikrotik) getRouts() *routeros.Reply {
	reply, err := m.Con.Run(
		"/ip/route/print",
		fmt.Sprintf("?routing-table=%s", m.table),
	)
	if err != nil {
		slog.Error("Error geting routes on mikrotik", "address", m.Address, "port", m.Port, "err", err)
		return nil
	}
	slog.Debug("Adding geting routes on mikrotik", "address", m.Address, "port", m.Port)

	return reply
}
