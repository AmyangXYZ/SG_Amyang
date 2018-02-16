package main

import (
	"html/template"
	"strings"

	"./router"
	"github.com/AmyangXYZ/sweetygo"
)

func unescaped(s string) interface{} {
	return template.HTML(s)
}

// for title in url, Hello World -> Hello-World
func space2hyphen(s string) string {
	return strings.Replace(s, " ", "-", -1)
}

func main() {
	funcMap := template.FuncMap{"unescaped": unescaped, "space2hyphen": space2hyphen}
	app := sweetygo.New("/home/amyang/Projects/SG_Amyang/", funcMap)
	router.SetMiddlewares(app)
	router.SetRouter(app)
	app.RunServer(":16311")
}
