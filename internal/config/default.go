package config

func setDefaultValue(cfg *Config) {
	for i, v := range cfg.Routers {
		if v.Address == "" {
			cfg.Routers[i].Address = "192.168.88.1"
		}
		if v.Port == "0" {
			cfg.Routers[i].Port = "8728"
		}
		if v.Login == "" {
			cfg.Routers[i].Login = "admin"
		}
		if v.RoutingTable == "" {
			cfg.Routers[i].RoutingTable = "rmd"
		}
	}
}
