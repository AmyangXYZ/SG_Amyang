package config

import (
	"fmt"
	"os"
)

var (
	// RootDir of your app
	RootDir = "/root/go/src/github.com/AmyangXYZ/SG_Amyang/"

	// SecretKey computes sg_token
	SecretKey string

	// DB addr and passwd.
	DB string
)

func init() {
	DB = fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/sg_amyang?charset=utf8", os.Getenv("DB_Passwd"))
	SecretKey = os.Getenv("SecretKey")
}
