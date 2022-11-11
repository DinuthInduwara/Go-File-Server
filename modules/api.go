package modules

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func HandleUpload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	log.Println("Upload:", file.Filename)
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("public/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// urlPath := "http://" + c.Request().Host + "/dl"

	longLink := "/dl/name/" + file.Filename

	return c.JSON(http.StatusOK, map[string]string{"fileName": file.Filename, "longLink": longLink})
}

func DownloadFile(c echo.Context) error {
	link := c.Request().URL.Path

	if strings.Contains(link, "dl/name") {
		fileId := c.Param("name")
		filePath := fmt.Sprintf("public/%s", fileId)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "File not found")
		}
		return c.Attachment(filePath, fileId)
	}

	return c.String(http.StatusNotFound, "File not found")
}
