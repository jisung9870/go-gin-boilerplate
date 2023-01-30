package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func LoadConfig(path, envPrefix string) (*Config, error) {
	cfg := &Config{}
	if path != "" {
		err := LoadFile(path, cfg)
		if err != nil {
			return cfg, errors.Wrap(err, "error loading config from file")
		}
	}
	// TODO: 동작 수정 필요
	err := envconfig.Process(envPrefix, cfg)
	return cfg, errors.Wrap(err, "error loading config from env")
}

func LoadFile(path string, cfg *Config) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "failed to open config file")
	}
	if err = yaml.Unmarshal(buf, &cfg); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}

	return nil
}

type Config struct {
	DatabaseConfig DatabaseConfig `yaml:"database"`
	AuthConfig     AuthConfig     `yaml:"auth"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"databaseName"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
}

type AuthConfig struct {
	AccessSecret  string `yaml:"accessSecret"`
	RefreshSecret string `yaml:"refreshSecret"`
}
