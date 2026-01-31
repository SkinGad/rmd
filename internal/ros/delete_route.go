package ros

import (
	"fmt"
	"log/slog"
)

func (m *Mikrotik) deleteRoute(id string) {
	_, err := m.Con.Run(
		"/ip/route/remove",
		fmt.Sprintf("=.id=%s", id),
	)
	if err != nil {
		slog.Error("Error deleting route on mikrotik", "address", m.Address, "port", m.Port, "id", id, "err", err)
		return
	}
	slog.Info("Deleting route on mikrotik", "address", m.Address, "port", m.Port, "id", id)
}
