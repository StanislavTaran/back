package main

import (
	"back/internal/server"
	"back/pkg/logger"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const configPath = "config/config.json"

// @title User API documentation
// @version 1.0.0
func main() {
	serverCfg := server.NewConfig()
	if err := server.ReadConfig(configPath, serverCfg); err != nil {
		log.Fatal(fmt.Sprintf("Config read failed...%s", err.Error()))
	}

	logger := logger.New(serverCfg.LogLevel)

	s := server.NewServer(serverCfg, logger)

	err := s.Run()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Server start failed... %s", err.Error()))
	}
}
