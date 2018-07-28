package controller

import (
	"fmt"

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

func GoogleVerify(ctx *sweetygo.Context) {
	ctx.Text(200, `google-site-verification: google9c7bdbb18c542f25.html`)
}
