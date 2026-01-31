package ros

import (
	"fmt"
	"log/slog"
)

func (m *Mikrotik) addRoute(r Routes) {
	_, err := m.Con.Run(
		"/ip/route/add",
		"=disabled=false",
		fmt.Sprintf("=comment=%s", r.Comment),
		fmt.Sprintf("=dst-address=%s", r.Address),
		fmt.Sprintf("=gateway=%s", m.gateway),
		fmt.Sprintf("=routing-table=%s", m.table),
	)
	if err != nil {
		slog.Error("Error adding route on mikrotik", "address", m.Address, "port", m.Port, "route", r, "err", err)
		return
	}
	slog.Info("Adding route on mikrotik", "address", m.Address, "port", m.Port, "route", r)
}
