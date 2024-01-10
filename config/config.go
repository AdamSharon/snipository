package config

import (
	"os"
	"path/filepath"
)

// later this will be the base of a configuration file that the user can edit
// currently i'll set some default values

type Config struct {
	// the path to the data file
	DataFilePath string

	// the name of the environment variable that holds the history file path
	HistoryEnvVarName string
}

func NewConfig() *Config {
	return &Config{
		DataFilePath: filepath.Join(os.Getenv("HOME"),
			".config/snipository",
			"snipository_data.json",
		),
		HistoryEnvVarName: "HISTFILE", // zsh default
	}
}
