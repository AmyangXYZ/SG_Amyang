package main

import (
	"./router"
	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New("/home/amyang/Projects/SG_Amyang/")
	router.SetMiddlewares(app)
	router.SetRouter(app)
	app.RunServer(":16311")
}
