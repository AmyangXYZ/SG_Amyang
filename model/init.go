package model

import (
	"database/sql"

	//
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:0311@/amyang_xyz?charset=utf8")
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
