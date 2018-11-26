package controller

import (
	"fmt"
	"net/http"

	"github.com/AmyangXYZ/SG_Amyang/config"
	"github.com/AmyangXYZ/sweetygo"
)

// Upload API Handler.
func Upload(ctx *sweetygo.Context) {
	saveDir := "uploadsfolder/"
	filename, err := ctx.SaveFile("file", config.RootDir+saveDir)
	fmt.Println(filename)
	if err != nil {
		ctx.JSON(200, 0, "upload file error", nil)
		return
	}
	filePath := "/" + saveDir + filename
	ctx.JSON(200, 1, "success", filePath)
}

// GoogleVerify .
func GoogleVerify(ctx *sweetygo.Context) {
	ctx.Text(200, `google-site-verification: google9c7bdbb18c542f25.html`)
}

// Static files handler
func Static(ctx *sweetygo.Context) {
	staticHandle := http.StripPrefix("/static",
		http.FileServer(http.Dir(config.RootDir+"/static")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
}

// Uploaded files handler
func Uploaded(ctx *sweetygo.Context) {
	staticHandle := http.StripPrefix("/uploadsfolder",
		http.FileServer(http.Dir(config.RootDir+"/uploadsfolder")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
}

// RedirectQUIC .
func RedirectQUIC(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://amyang.xyz:443"+r.RequestURI, http.StatusMovedPermanently)
}
