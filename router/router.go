package router

import (
	"os"

	"github.com/AmyangXYZ/SG_Amyang/controller"
	"github.com/AmyangXYZ/sweetygo"
	"github.com/AmyangXYZ/sweetygo/middlewares"
)

var (
	requireJWTMap = map[string]string{
		"/api/posts":   "!GET",
		"/api/posts/*": "!GET",
		"/api/files":   "POST",
	}
)

// SetMiddlewares for SweetyGo APP.
func SetMiddlewares(app *sweetygo.SweetyGo) *sweetygo.SweetyGo {
	logfile, _ := os.Create("access.log")
	app.USE(middlewares.Logger(logfile))
	app.USE(middlewares.JWT("Header", controller.SecretKey, requireJWTMap))
	return app
}

// SetRouter for SweetyGo APP.
func SetRouter(app *sweetygo.SweetyGo) *sweetygo.SweetyGo {

	app.GET("/", controller.Home)

	app.GET("/posts/new", controller.NewPage)
	app.GET("/posts/category/:cat", controller.Cat)
	app.GET("/posts/:title", controller.Show)
	app.GET("/posts/:title/edit", controller.EditPage)

	app.POST("/api/posts", controller.New)
	app.GET("/api/posts/page/:n", controller.PaginationHome)
	app.GET("/api/posts/category/:cat/page/:n", controller.PaginationCat)
	app.PUT("/api/posts/:title", controller.Update)

	app.POST("/api/files", controller.Upload)

	app.POST("/api/token", controller.Login)
	return app
}
