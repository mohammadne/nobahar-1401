package config

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/mohammadne/nobahar-1401/internal/db"
	"github.com/mohammadne/nobahar-1401/internal/http"
	"github.com/mohammadne/nobahar-1401/internal/jwt"
)

type Config struct {
	HTTP *http.Config `koanf:"http"`
	JWT  *jwt.Config  `koanf:"jwt"`
	DB   *db.Config   `koanf:"db"`
}

const delimeter = "."

func Load() (*Config, error) {
	k := koanf.New(delimeter)

	if err := LoadValues(k); err != nil {
		return nil, fmt.Errorf("error loading values: %v", err)
	}

	config := Config{}
	if err := k.Unmarshal("", &config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}

	return &config, nil
}

//go:embed values
var res embed.FS

func LoadValues(k *koanf.Koanf) error {
	files, err := fs.ReadDir(res, "values")
	if err != nil {
		return fmt.Errorf("error reading configs directory: %s", err)
	}

	for _, file := range files {
		path := "values/" + file.Name()
		data, err := fs.ReadFile(res, path)
		if err != nil {
			return fmt.Errorf("error reading configs directory: %s", err)
		}

		var parser koanf.Parser
		switch filepath.Ext(path) {
		case ".yml":
			parser = yaml.Parser()
		case ".yaml":
			parser = yaml.Parser()
		case ".toml":
			parser = toml.Parser()
		default:
			continue
		}

		if err := k.Load(rawbytes.Provider(data), parser); err != nil {
			return fmt.Errorf("error loading values: %s", err)
		}
	}

	return nil
}
