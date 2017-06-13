package http

import (
	_"fmt"
	"net/http"
	"os"
	"io"
	"fmt"
	"strings"
)

func DownloadFile(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fileName := formatUrl(url)
	osErr := os.MkdirAll("img", os.ModePerm)
	if osErr != nil {
		fmt.Println(osErr)
	}
	f, err := os.Create("img/" + fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("body===>", resp.Body)
	io.Copy(f, resp.Body)
	defer resp.Body.Close()
}

func formatUrl(url string) string {
	strs := strings.Split(url, "/")
	/*for index, item := range strs {
		fmt.Println(index, item)
	}*/
	result := strs[len(strs) - 1]
	return result
}
