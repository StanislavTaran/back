package main

import (
	"back/internal/server"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const configPath = "config/config.json"

func main() {
	serverCfg := server.NewConfig()
	if err := server.ReadConfig(configPath, serverCfg); err != nil {
		log.Fatal(fmt.Sprintf("Config read failed...%s", err.Error()))
	}

	//logger := logger.New(serverCfg.Server.LogLevel)

	s := server.NewServer(serverCfg)

	err := s.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("Server start failed... %s", err.Error()))
	}
}
