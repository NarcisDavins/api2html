package engine

import (
	"encoding/json"
	"io"
	"os"
)

// ParseConfigFromFile creates a Config with the contents of the received filepath
func ParseConfigFromFile(path string) (Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	cfg, err := ParseConfig(configFile)
	configFile.Close()
	return cfg, err
}

// ParseConfig parses the content of the reader into a Config
func ParseConfig(r io.Reader) (Config, error) {
	var cfg Config
	err := json.NewDecoder(r).Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	for p, page := range cfg.Pages {
		if len(page.Extra) == 0 {
			cfg.Pages[p].Extra = cfg.Extra
			continue
		}
		for k, v := range cfg.Extra {
			if _, ok := page.Extra[k]; !ok {
				cfg.Pages[p].Extra[k] = v
			}
		}
	}
	return cfg, nil
}
