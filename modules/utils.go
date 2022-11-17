package modules

import (
	"fmt"
	"mime"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)


type metadata struct{
	fname string
	contentLength int64
	done int64
	

}

func DownloadFileHandler(c echo.Context) error {

}

func GetNameFromURL(url string) string {
	ls := strings.Split(url, "/")
	return ls[len(ls)-1]
}

func GetNameFromHeder(url string) (string, http.Header, error) {
	request, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return "", nil, err
	}
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return "", response.Header, err
	}

	cd := response.Header.Get("Content-Disposition")
	if cd == "" {
		return "", response.Header, fmt.Errorf("Failed To GET FROM Header")
	}
	_, params, err := mime.ParseMediaType(cd)
	if err != nil {
		return "", response.Header, err
	}

	return params["filename"], response.Header, nil
}
