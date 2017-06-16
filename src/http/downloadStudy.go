package http

import (
	_"fmt"
	"net/http"
	"os"
	"io"
	"fmt"
	"strings"
	"ImgSpider/src/utils"
)

func DownloadFile(url string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//fileName := formatUrl(url)
	fileName := formatFileName(url)
	osErr := os.MkdirAll("img", os.ModePerm)
	if osErr != nil {
		fmt.Println(osErr)
	}
	f, err := os.Create("img/" + fileName)
	if err != nil {
		panic(err)
	}
	//fmt.Println("body===>", resp.Body)
	_, wErr := io.Copy(f, resp.Body)
	if wErr != nil {
		fmt.Println("保存失败", wErr)
	}
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

func formatFileName(url string) string {
	index := strings.LastIndex(url, ".")
	temp := url[index:]
	m, _ := utils.Md5(url)

	return m + temp
}

