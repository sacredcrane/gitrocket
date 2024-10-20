package config

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Env      `toml:"env"`
	Server   `toml:"server"`
	Database `toml:"database"`
}

type Env struct {
	Type      string `toml:"type"`
	Workspace string `toml:"workspace"`
}

type Server struct {
	Host          string `toml:"host"`
	Port          int    `toml:"port"`
	ServerTimeout `toml:"timeout"`
}

type ServerTimeout struct {
	Server time.Duration `toml:"server"`
	Read   time.Duration `toml:"read"`
	Write  time.Duration `toml:"write"`
	Idle   time.Duration `toml:"idle"`
}

type Database struct {
	Driver   string `toml:"driver"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Hostname string `toml:"hostname"`
	Port     int    `toml:"port"`
	Name     string `toml:"db_name"`
	Schema   string `toml:"schema"`
}

func LoadConfigFromFile() (*Config, error) {
	var cfg Config

	path := os.Getenv("CFG_PATH")
	if path == "" {
		return nil, fmt.Errorf("unable to get CFG_PATH environment variable")
	}

	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
