package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AllConfig struct {
	Debug  bool   `yaml:"debug"`
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Data struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Driver       string `yaml:"driver"`
	Connection   string `yaml:"connection"`
	MaxOpenConns int    `yaml:"max_open_connections"`
	MaxIdleConns int    `yaml:"max_idle_connections"`
	MaxLifeSecs  int    `yaml:"max_life_seconds"`
}

func LoadConfig(path string) (*AllConfig, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config AllConfig
	err = yaml.Unmarshal(yamlFile, &config)
	return &config, err
}
