package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AmyangXYZ/SG_Amyang/config"
	//
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", config.DB)
	for {
		if err := db.Ping(); err != nil {
			fmt.Println(err, ", retry in 10s...")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	// https://github.com/go-sql-driver/mysql/issues/674
	db.SetMaxIdleConns(0)

	db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT(10) NOT NULL AUTO_INCREMENT,
			name VARCHAR(16) NULL DEFAULT NULL,
			passwd VARCHAR(64) NULL DEFAULT NULL,
			PRIMARY KEY (id)
		);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS posts ( 
		id INT(10) NOT NULL AUTO_INCREMENT,
	 	title VARCHAR(64) NULL DEFAULT NULL,
		cat VARCHAR(16) NULL DEFAULT NULL,
		html TEXT NULL DEFAULT NULL,
		md TEXT NULL DEFAULT NULL,
		time DATETIME NULL DEFAULT NULL,
		PRIMARY KEY (id)
		);`)
}
