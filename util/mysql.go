package util

import (
	"51job/cons"
	"51job/log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, err := sql.Open("mysql", cons.Db)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}
