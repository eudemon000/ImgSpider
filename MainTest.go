package main

import (
	"ImgSpider/src/queen"
	"fmt"
	"ImgSpider/src/function"
	myHttp "ImgSpider/src/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
	//"ImgSpider/src/goqueryStudy"
	_"ImgSpider/src/goroutineTest"
	"runtime"
	_"ImgSpider/src/goqueryStudy"
	"ImgSpider/src/goqueryStudy"
)

type Abc struct {
	a int
	b int
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	function.MutiFunc(1, 2, 3, 4, 5)
	var a int = 1
	var b int64 = 64
	var c string = "Hello world"
	var d = [4]byte{1, 2, 3, 4}
	//结构体的初始化
	var e Abc = Abc{1, 1}
	//结果体的另种初始化方式，类型为指针类型
	var f *Abc = new(Abc)
	f.a = 3
	f.b = 4
	function.MutiInterFunc(a, b, c, d, e, *f)
	/*myHttp.TestGet()
	myHttp.TestPost()*/
	//myHttp.DownloadFile("http://106.2.184.227:9999/dldir1.qq.com/qqfile/qq/QQ8.9.3/21149/QQ8.9.3.exe")
	//downloadPic()
	//goqueryStudy.ParseForCharset("http://www.qq.com/")
	//goroutineTest.RunTest()
	//goroutineTest.RunTest_1()
	//goroutineTest.Test()
	//testQueen()
	//goqueryStudy.ImgSp("https://www.4493.com/")
	//goqueryStudy.ImgSp("http://slide.games.sina.com.cn/t/slide_21_2248_426871.html")
	goqueryStudy.OtherImgSp("http://slide.games.sina.com.cn/t/slide_21_2248_426871.html")
	/*var h queen.Handler  = queen.Handler(func(data interface{}){
		fmt.Println("===", data)
	})
	queen.CreateDataQueen(h)
	queen.Push("啊啊啊")
	queen.Push(1)
	queen.Push(1)*/
	var ch chan int
	ch <- 1

}

func downloadPic() {
	url := "http://www.tooopen.com/img/88_878.aspx"
	var doc *goquery.Document
	doc, _ = goquery.NewDocument(url)
	body := doc.Find("body")
	body.Each(func(i int, sele *goquery.Selection) {
		fmt.Println("111")
		imgTag := sele.Find("img")
		fmt.Println(imgTag.Length())
		imgTag.Each(func(index int, imgSele *goquery.Selection) {
			imgUrl, ok := imgSele.Attr("src")
			if ok {
				//fmt.Println(imgUrl)
				r := checkUrl(url, imgUrl)
				myHttp.DownloadFile(r)
			}

		})
	})
}

func checkUrl(baseUrl, url string) string {
	var result string
	index := strings.Index(url, "/")
	if index == 0 {
		fmt.Println("不是正经URL")
		b := strings.Split(baseUrl, "")
		var in int = 0
		for q, a := range b {
			if a == "/" {
				in++
			}
			if in == 3 {
				result = baseUrl[0:q]
				result += url
				fmt.Println(result)
				break
			}
		}
	} else {
		fmt.Println("正经URL")
		result = url
	}
	return result
}

func testQueen() {
	handler := queen.MsgQueenHandler(func(data interface{}){
		//fmt.Println("从队列里取出", data)
		for i := 0; i < 1000000; i++ {
			fmt.Println(data.(int) + i)
		}
	})
	queen.CreateQueen(handler)
	for i := 0; i < 10; i++ {
		queen.PushData(1)
	}

}



