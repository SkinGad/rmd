package cache

import (
	"encoding/json"
	"log/slog"
	"os"
)

const cacheFile = "cache.json"

func (c *Cache) WriteToFile(filename string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (c *Cache) ReadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}

func (c *Cache) Save() error {
	return c.WriteToFile(cacheFile)
}

func (c *Cache) Load() error {
	err := c.ReadFromFile(cacheFile)
	if err != nil {
		slog.Info("Write new cache file")
		err = c.WriteToFile(cacheFile)
	}
	return err
}
