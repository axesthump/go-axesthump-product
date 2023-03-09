package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Postgres struct {
		Url string `yaml:"url"`
	} `yaml:"postgres"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYAML, err := os.ReadFile("config.yml")
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return fmt.Errorf("parsing yaml: %w", err)
	}

	return nil
}
