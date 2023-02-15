package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	dsn := "root:inipassmysql@tcp(127.0.0.1:3306)/mhs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	return db
}
