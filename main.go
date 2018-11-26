package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/AmyangXYZ/SG_Amyang/config"
	"github.com/AmyangXYZ/SG_Amyang/controller"
	"github.com/AmyangXYZ/SG_Amyang/router"
	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New()

	app.SetTemplates(config.RootDir+"templates", template.FuncMap{"unescaped": unescaped,
		"space2hyphen": space2hyphen, "abstract": abstract, "rmtag": rmtag})
	router.SetMiddlewares(app)
	router.SetRouter(app)

	go http.ListenAndServe(":80", http.HandlerFunc(controller.RedirectQUIC))
	// app.Run(":16311")
	app.RunOverQUIC(":443", "/etc/letsencrypt/live/amyang.xyz/fullchain.pem", "/etc/letsencrypt/live/amyang.xyz/privkey.pem")
}

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
