package ros

func (m *Mikrotik) SyncRoute(routes []Routes) {
	reply := m.getRouts()
	for _, v := range routes {
		check := true
		for _, w := range reply.Re {
			if w.Map["dst-address"] == v.Address {
				check = false
				if w.Map["disabled"] == "true" {
					m.EnableRoute(v.Address)
				}
				break
			}
		}
		if check {
			m.addRoute(v)
		}
	}
}
