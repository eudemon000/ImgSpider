package http

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/henrylee2cn/pholcus/common/mahonia"
	"strings"
)

func TestGet() {
	var resp *http.Response
	var err error
	resp, err = http.Get("http://www.qq.com")
	defer resp.Body.Close()
	if err !=nil {
		fmt.Println(err)
	}
	var body []byte
	var respErr error
	body, respErr = ioutil.ReadAll(resp.Body)
	if respErr != nil {
		fmt.Println(respErr)
	}
	var de mahonia.Decoder
	de = mahonia.NewDecoder("gb2312")
	b := string(body)
	result := de.ConvertString(b)
	fmt.Println(result)
	var headers http.Header
	headers = resp.Header
	for index, item := range headers {
		fmt.Printf("k=%v, v=%v\n", index, item)
	}

	fmt.Printf("resp status %v, status code %v\n", resp.Status, resp.StatusCode)
	fmt.Printf("resp proto %s\n", resp.Proto)
	var cookies []*http.Cookie
	cookies = resp.Cookies()
	for index, item := range cookies {
		fmt.Printf("index=%v, item=%v\n", index, item)
	}
}

func TestPost() {
	resp, err := http.Post("http://219.136.249.24:8022/keeper/login_login.action", "application/x-www-form-urlencoded", strings.NewReader("userName=aaa&passWord=aaa"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(body))

}