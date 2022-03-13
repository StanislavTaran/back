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

//func main() {
//	db, err := sqlx.Connect("mysql", "root:test@tcp(localhost:3306)/testDb?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	if err = db.Ping(); err != nil {
//		log.Fatalln(err)
//	}
//
//	db.DB.SetMaxOpenConns(40)
//	db.DB.SetMaxIdleConns(5)
//
//}
