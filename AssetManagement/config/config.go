package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Application defines system listening port
type Application struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

// Repository defines database connection information
type Repository struct {
	Server   string `yaml:"db_server"`
	Name     string `yaml:"db_name"`
	Username string `yaml:"db_username"`
	Password string `yaml:"db_password"`
}

// Config defines system configuration
type Config struct {
	App  *Application `yaml:"application"`
	Repo *Repository  `yaml:"repository"`
}

// NewConfig creates configuration object
func NewConfig(filepath string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
