package modules

import (
	"io"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func DownloadFileHandler(c echo.Context) error {
	var url string = c.Param("url")
	downloadHandler("/public", url)
	return nil
}

func downloadHandler(path string, url string) error {
	resp, err := http.Get(url)
	contentType := resp.Header.Get("Content-type")
	arr := strings.Split(url, "/")

	ext, err := mime.ExtensionsByType(contentType)

	out, err := os.Create(path + "/" + arr[len(arr)-1] + ext[0])
	defer out.Close()

	defer resp.Body.Close()
	io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}
	return nil

}

func download(destinationPath, downloadUrl string) error {
	tempDestinationPath := destinationPath + ".tmp"
	req, _ := http.NewRequest("GET", downloadUrl, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(tempDestinationPath, os.O_CREATE|os.O_WRONLY, 0644)


	io.Copy(io.MultiWriter(f, bar), resp.Body)
	os.Rename(tempDestinationPath, destinationPath)
	return nil

}
