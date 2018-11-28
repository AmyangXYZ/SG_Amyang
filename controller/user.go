package controller

import (
	"time"

	"github.com/AmyangXYZ/SG_Amyang/config"
	"github.com/AmyangXYZ/SG_Amyang/model"
	"github.com/AmyangXYZ/sweetygo"
	jwt "github.com/dgrijalva/jwt-go"
)

// Login API Handler.
func Login(ctx *sweetygo.Context) error {
	if ctx.Param("name") != "" && ctx.Param("passwd") != "" {
		name, passwd, err := model.GetAdmin()
		if err != nil {
			return ctx.JSON(500, 0, "Database Error", nil)
		}
		if name == ctx.Param("name") && passwd == ctx.Param("passwd") {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = "Amyang"
			claims["admin"] = true
			claims["exp"] = time.Now().Add(time.Hour * 4).Unix()
			t, _ := token.SignedString([]byte(config.SecretKey))
			return ctx.JSON(200, 1, "success", map[string]string{"SG_Token": t})

		}
		return ctx.JSON(200, 0, "Username or Password Error.", nil)
	}
	return ctx.JSON(406, 0, "I can't understand what u want", nil)
}
