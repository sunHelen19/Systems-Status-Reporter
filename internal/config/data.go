package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type DataConfig struct {
	Server struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Mms       string `yaml:"mms"`
		Accendent string `yaml:"accendent"`
		Support   string `yaml:"support"`
	} `yaml:"server"`

	FileName struct {
		Sms     string `yaml:"sms"`
		Billing string `yaml:"billing"`
		Email   string `yaml:"email"`
		Voice   string `yaml:"voice"`
	} `yaml:"fileName"`
}

func NewConfigData(configPath string) (*DataConfig, error) {
	// Create config structure
	config := &DataConfig{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
