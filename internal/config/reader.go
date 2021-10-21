package config

import (
	os

	"gopkg.in/yaml.v3"
	"github.com/kelseyhightower/envconfig"
)

var (
	cfg *Config 
)


func ReadFile(cfg *Config) (err error) {
	cfgPath := "./build/config/config.yaml"

	file, error := os.Open(cfgPath)

	if error != nil {
		return error
	}

	defer func() {
		closeError := file.Close()
		if closeError != nil {
			error = closeError
		}
	}()

	decoder := yaml.NewDecoder(file)
	if error := decoder.Decode(cfg); error != nil {
		return error
	}
	return nil
}
func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}


func SetConfig(c *Config) {
	cfg = c
}