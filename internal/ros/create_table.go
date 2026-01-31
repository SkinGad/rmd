package ros

import (
	"fmt"
	"log"
	"log/slog"
)

func (m *Mikrotik) getTable() {
	reply, err := m.Con.Run("/routing/table/print")
	if err != nil {
		log.Fatal(err)
	}

	for _, re := range reply.Re {
		if re.Map["name"] == m.table {
			return
		}
	}
	m.createTable()
}

func (m *Mikrotik) createTable() {
	_, err := m.Con.Run("/routing/table/add", "=fib", fmt.Sprintf("=name=%s", m.table))
	if err != nil {
		slog.Error("Error adding routing table on mikrotik", "address", m.Address, "port", m.Port, "routing_table", m.table, "err", err)
		return
	}
	slog.Info("Adding routing table on mikrotik", "address", m.Address, "port", m.Port, "routing_table", m.table)
}
