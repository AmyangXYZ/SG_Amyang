package router

import (
	"os"

	"../controller"
	"github.com/AmyangXYZ/sweetygo"
	"github.com/AmyangXYZ/sweetygo/middlewares"
)

var (
	requireJWTMap = map[string]string{
		"/api/post": "!GET",
	}
)

// SetMiddlewares for SweetyGo APP.
func SetMiddlewares(app *sweetygo.SweetyGo) *sweetygo.SweetyGo {
	app.USE(middlewares.Logger(os.Stdout))
	app.USE(middlewares.JWT("Header", controller.SecretKey, requireJWTMap))
	return app
}

// SetRouter for SweetyGo APP.
func SetRouter(app *sweetygo.SweetyGo) *sweetygo.SweetyGo {
	app.GET("/", controller.Home)
	app.GET("/about", controller.About)
	app.GET("/upload", controller.UploadPage)
	app.Static("/uploads", "/home/amyang/Projects/SG_Amyang/uploads")

	app.GET("/posts/:title", controller.Show)
	app.GET("/login", controller.LoginPage)

	app.GET("/api/post", controller.Get)
	app.POST("/api/post", controller.New)
	app.PUT("/api/post", controller.Update)

	app.POST("/api/file", controller.Upload)

	app.POST("/api/token", controller.Login)
	return app
}
