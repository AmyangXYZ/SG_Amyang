package controller

import (
	"github.com/AmyangXYZ/sweetygo"
)

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
