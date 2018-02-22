package config

import (
	"fmt"
	"os"
)

var (
	// RootDir of your app
	RootDir = "/"

	// SecretKey computes sg_token
	SecretKey string

	// DB addr and passwd.
	DB string
)

func init() {
	DB = fmt.Sprintf("SuperSweetie:%v@tcp(db:3306)/sg_amyang?charset=utf8", os.Getenv("DB_Passwd"))
	SecretKey = os.Getenv("SecretKey")
}
