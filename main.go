package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/AmyangXYZ/SG_Amyang/config"
	"github.com/AmyangXYZ/SG_Amyang/router"
	"github.com/AmyangXYZ/sweetygo"
)

// Host is for subdomain control
type Host struct {
	SG *sweetygo.SweetyGo
}

func main() {
	hosts := map[string]*Host{}

	// blog
	// the blog
	blog := sweetygo.New()
	blog.SetTemplates(config.RootDir+"templates", template.FuncMap{
		"unescaped":    unescaped,
		"space2hyphen": space2hyphen,
		"abstract":     abstract,
		"rmtag":        rmtag,
	})
	router.SetMiddlewares(blog)
	router.SetRouter(blog)
	hosts["amyang.xyz"] = &Host{blog}

	// BirdSong Recg
	// convey the requests to birdsong-recg app powered by flask.
	bird := sweetygo.New()
	u, _ := url.Parse("http://172.17.0.3:80/")
	bird.Any("/*", proxyHandler(httputil.NewSingleHostReverseProxy(u)))
	hosts["birdsong.amyang.xyz"] = &Host{bird}

	// hacking
	// some evil script :)
	hacking := sweetygo.New()
	hacking.GET("/*", func(ctx *sweetygo.Context) error {
		staticHandle := http.StripPrefix("/",
			http.FileServer(http.Dir(config.RootDir+"/hacking.amyang.xyz")))
		staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
		return nil
	})
	hosts["hacking.amyang.xyz"] = &Host{hacking}

	// server
	// distribute requests
	server := sweetygo.New()
	server.Any("/*", func(ctx *sweetygo.Context) error {
		if host := hosts[ctx.Req.Host]; host != nil {
			host.SG.ServeHTTP(ctx.Resp, ctx.Req)
			return nil
		}
		return ctx.Text(404, "404 not found")

	})

	// force redirect http to https
	redirector := sweetygo.New()
	redirector.Any("/*", func(ctx *sweetygo.Context) error {
		ctx.Redirect(301, fmt.Sprintf("https://%s:443", ctx.Req.Host)+ctx.Path())
		return nil
	})
	go redirector.Run("amyang.xyz:80")

	server.RunOverQUIC(":443", "/etc/letsencrypt/live/amyang.xyz/fullchain.pem", "/etc/letsencrypt/live/amyang.xyz/privkey.pem")
}

func proxyHandler(p *httputil.ReverseProxy) func(ctx *sweetygo.Context) error {
	return func(ctx *sweetygo.Context) error {
		p.ServeHTTP(ctx.Resp, ctx.Req)
		return nil
	}
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

