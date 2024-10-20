package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func init() {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "gitrocket") {
		wd = filepath.Dir(wd)
	}
	os.Setenv("CFG_PATH", filepath.Join(wd, "cfg/gitrocket.example.toml"))
}

func TestLoadFile(t *testing.T) {
	cfg, err := LoadConfigFromFile()
	if err != nil {
		t.Fatalf(`Failed
		Via error: %v 
		Config is: %v`,
			err,
			*cfg,
		)
	}
	// Based on gitrocket.example.toml
	config := Config{
		Env: Env{
			Type:      "local",
			Workspace: "/tmp/gitrocket",
		},
		Server: Server{
			Host: "127.0.0.1",
			Port: 8080,
			ServerTimeout: ServerTimeout{
				Server: 30,
				Read:   15,
				Write:  10,
				Idle:   5,
			},
		},
		Database: Database{
			Driver:   "postgres",
			User:     "postgres",
			Password: "postgres",
			Hostname: "localhost",
			Port:     5432,
			Name:     "postgres",
			Schema:   "public",
		},
	}
	if config != *cfg {
		t.Fatalf(`Failed
		Config is: %v`,
			*cfg,
		)
	}
}

func TestLoadFileWhichNotExists(t *testing.T) {
	env := os.Getenv("CFG_PATH")
	t.Setenv("CFG_PATH", env[:len(env)-1])

	cfg, err := LoadConfigFromFile()
	if err == nil {
		t.Fatalf(`Failed
		Via error: %v 
		Config is: %v`,
			err,
			*cfg,
		)
	}
}

func TestLoadFileWithoutEnvvar(t *testing.T) {
	os.Unsetenv("CFG_PATH")

	cfg, err := LoadConfigFromFile()
	if err == nil || err.Error() != "unable to get CFG_PATH environment variable" {
		t.Fatalf(`Failed
		Via error: %v 
		Config is: %v`,
			err,
			*cfg,
		)
	}
}
