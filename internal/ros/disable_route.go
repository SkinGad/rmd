package ros

import (
	"fmt"
	"log/slog"
)

func (m *Mikrotik) DisableRoute(address string) {
	reply := m.getRouts()
	for _, v := range reply.Re {
		if v.Map["dst-address"] == address {
			m.changeDisableRoute(v.Map[".id"], "true")
		}
	}
}

func (m *Mikrotik) EnableRoute(address string) {
	reply := m.getRouts()
	for _, v := range reply.Re {
		if v.Map["dst-address"] == address {
			m.changeDisableRoute(v.Map[".id"], "false")
		}
	}
}

func (m *Mikrotik) changeDisableRoute(id, state string) {
	_, err := m.Con.Run(
		"/ip/route/set",
		fmt.Sprintf("=disabled=%s", state),
		fmt.Sprintf("=.id=%s", id),
	)
	if err != nil {
		slog.Error("Error changing route on mikrotik", "address", m.Address, "port", m.Port, "route_id", id, "to_state", state, "err", err)
		return
	}
	slog.Info("Changing route on mikrotik", "address", m.Address, "port", m.Port, "route_id", id, "to_state", state)
}
