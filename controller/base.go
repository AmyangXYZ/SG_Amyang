package controller

import (
	"github.com/AmyangXYZ/sweetygo"
)

// Home Page Handler.
func Home(ctx *sweetygo.Context) {
	ctx.Set("title", "Home")
	ctx.Render(200, "home")
}

// About Page Handler.
func About(ctx *sweetygo.Context) {
	ctx.Text(200, "Powered By SweetyGo.\n")
}

// UploadPage Handler.
func UploadPage(ctx *sweetygo.Context) {
	ctx.Text(200, "Powered By SweetyGo.\n")
}

// Upload API Handler.
func Upload(ctx *sweetygo.Context) {
	ctx.Text(200, "Powered By SweetyGo.\n")
}
