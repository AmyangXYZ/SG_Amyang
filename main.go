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

// show abstract, splited by tag icon.
func abstract(s string) string {
	return strings.Split(s, "<p><i class=\"fa fa-tag fa-emoji\" title=\"tag\"></i></p>")[0]
}

// replace tag icon in content
func rmtag(s string) string {
	return strings.Replace(s, "<p><i class=\"fa fa-tag fa-emoji\" title=\"tag\"></i></p>", "", -1)
}

func main() {
	funcMap := template.FuncMap{"unescaped": unescaped,
		"space2hyphen": space2hyphen, "abstract": abstract, "rmtag": rmtag}
	app := sweetygo.New("/home/amyang/Projects/SG_Amyang/", funcMap)
	router.SetMiddlewares(app)
	router.SetRouter(app)
	app.RunServer(":16311")
}
