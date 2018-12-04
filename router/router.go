package router

import (
	"os"

	"github.com/AmyangXYZ/SG_Amyang/config"
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

func loggerSkipper(ctx *sweetygo.Context) bool {
	if (len(ctx.Path()) > 8 && ctx.Path()[0:8] == "/static/") ||
		(len(ctx.Path()) > 15 && ctx.Path()[0:15] == "/uploadsfolder/") {
		return true
	}
	return false
}

func jwtSkipper(ctx *sweetygo.Context) bool {
	if (ctx.Path() == "/api/posts" && ctx.Method() != "GET") ||
		(len(ctx.Path()) > 11 && ctx.Path()[0:11] == "/api/posts/" && ctx.Method() != "GET") ||
		(ctx.Path() == "/api/files" && ctx.Method() == "POST") {
		return false
	}
	return true
}

// SetMiddlewares for SweetyGo APP.
func SetMiddlewares(app *sweetygo.SweetyGo) *sweetygo.SweetyGo {
	f, _ := os.Create(config.RootDir + "sg.log")
	app.USE(middlewares.Logger(f, loggerSkipper))
	app.USE(middlewares.Gzip(6, middlewares.DefaultSkipper))
	app.USE(middlewares.JWT("Header", config.SecretKey, jwtSkipper))
	return app
}

// SetRouter for SweetyGo APP.
func SetRouter(app *sweetygo.SweetyGo) *sweetygo.SweetyGo {

	app.GET("/", controller.Home)
	app.GET("/static/*files", controller.Static)
	app.GET("/uploadsfolder/*files", controller.Uploaded)

	app.GET("/google9c7bdbb18c542f25.html", controller.GoogleVerify)
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
