package config

import "time"

type Config struct {
	RMD     RMD      `yaml:"rmd"`
	Routers []Router `yaml:"routers"`
	Domains []string `yaml:"domain"`
}

type RMD struct {
	TTL    time.Duration `yaml:"ttl"`
	LookUp time.Duration `yaml:"lookup"`
}
type Router struct {
	Address      string `yaml:"address"`
	Port         string `yaml:"port"`
	Login        string `yaml:"login"`
	Password     string `yaml:"password"`
	RoutingTable string `yaml:"table"`
	Gateway      string `yaml:"gateway"`
}
