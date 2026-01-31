package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ParsingConfig(configFileName string) (cfg Config, err error) {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	setDefaultValue(&cfg)
	return
}
