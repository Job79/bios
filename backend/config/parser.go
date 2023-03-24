package config

import (
	"encoding/json"
	"os"
)

// ParseConfig parses the config file and returns a Config struct
func ParseConfig(path string) (Config, error) {
	conf := getDefaults()
	file, err := os.Open(path)
	if err == nil {
		err = json.NewDecoder(file).Decode(&conf)
	}
	return conf, err
}
