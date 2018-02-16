package controller

import (
	"../model"
	"github.com/AmyangXYZ/sweetygo"
)

// Home Page Handler.
func Home(ctx *sweetygo.Context) {
	ctx.Set("title", "Home")
	posts, err := model.GetPosts("1")
	if err != nil {
		ctx.Text(500, "something error")
	}
	ctx.Set("posts", posts)
	ctx.Render(200, "home")
}

// Upload API Handler.
func Upload(ctx *sweetygo.Context) {
	saveDir := "uploadsfolder/"
	filename, err := ctx.SaveFile("file", saveDir)
	if err != nil {
		ctx.JSON(200, "upload file error", "error")
		return
	}
	filePath := "/" + saveDir + filename
	ctx.JSON(200, filePath, "success")
}
