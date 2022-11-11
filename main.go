package main

import (
	"FileServer/modules"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	var port int = 4000
	e := echo.New()
	os.Mkdir("public", os.ModePerm)

	e.GET("/", modules.IndexPage)
	e.GET("/upload", modules.UploadPage)
	e.POST("/api/upload", modules.HandleUpload)
	e.GET("/dl/name/:name", modules.DownloadFile)
	e.GET("/files", modules.GetFiles)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
