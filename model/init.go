package model

import (
	"database/sql"

	//
	_ "github.com/go-sql-driver/mysql"
)

// CREATE TABLE `users` (
// 	`id` INT(10) NOT NULL AUTO_INCREMENT,
// 	`name` VARCHAR(16) NULL DEFAULT NULL,
// 	`passwd` VARCHAR(64) NULL DEFAULT NULL,
// 	PRIMARY KEY (`id`)
// );
//
// CREATE TABLE `posts` (
// 	`id` INT(10) NOT NULL AUTO_INCREMENT,
//  `title` VARCHAR(64) NULL DEFAULT NULL,
// 	`cat` VARCHAR(16) NULL DEFAULT NULL,
// 	`html` TEXT NULL DEFAULT NULL,
// 	`md` TEXT NULL DEFAULT NULL,
// 	`time` DATETIME NULL DEFAULT NULL,
// 	PRIMARY KEY (`id`)
// );

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:0311@/sg_amyang?charset=utf8")
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
