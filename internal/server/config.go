package server

import (
	"back/pkg/mysqlClient"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BindAddr string              `json:"port"`
	LogLevel string              `json:"logLevel"`
	Mysql    *mysqlClient.Config `json:"mysql"`
}

// NewConfig - initialize new config with default values for 'outRunner' server.
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3030",
		LogLevel: "debug",
	}
}

// ReadConfig
// Reads file by path specified in first argument and write values into target config
func ReadConfig(filePath string, target *Config) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, target); err != nil {
		return err
	}
	return nil
}
