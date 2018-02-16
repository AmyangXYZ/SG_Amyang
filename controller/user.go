package controller

import (
	"time"

	"../model"

	"github.com/AmyangXYZ/sweetygo"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// SecretKey is for JWT.
	SecretKey = "secret"
)

// Login API Handler.
func Login(ctx *sweetygo.Context) {
	if ctx.Param("name") != "" && ctx.Param("passwd") != "" {
		name, passwd, err := model.GetAdmin()
		if err != nil {
			ctx.JSON(500, "Database Error", "error")
			return
		}
		if name == ctx.Param("name") && passwd == ctx.Param("passwd") {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "Amyang"
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
			t, _ := token.SignedString([]byte(SecretKey))
			ctx.JSON(200, map[string]string{"SG_Token": t}, "success")
			return
		}
		ctx.JSON(200, "Username or Password Error.", "fail")
	}
}
