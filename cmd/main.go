package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:test@tcp(localhost:3306)/testDb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	db.DB.SetMaxOpenConns(40)
	db.DB.SetMaxIdleConns(5)

}
