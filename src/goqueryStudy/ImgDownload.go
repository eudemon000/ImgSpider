package goqueryStudy

import (
	myqueen "ImgSpider/src/queen"
	"github.com/PuerkitoBio/goquery"
	_"fmt"
	myhttp "ImgSpider/src/http"
	"fmt"
	"strings"
	otqueue "github.com/otium/queue"
)

func ImgSp(url string) {
	/*handler := myqueen.MsgQueenHandler(func(data interface{}){
		fmt.Println("回调数据", data.(string))
	})*/
	var h myqueen.Handler  = myqueen.Handler(func(data interface{}){
		fmt.Println("===", data)
		if data.(string) == "" || strings.Index(data.(string), "java") != -1 {
			return
		}
		document, _ := goquery.NewDocument(data.(string))
		//下载图片
		bodySelect := document.Find("body").Find("img")
		bodySelect.Each(func(i int, imgSelect *goquery.Selection) {
			imgUrl, _ := imgSelect.Attr("src")
			//fmt.Println(imgUrl)
			if imgUrl != "" {
				myhttp.DownloadFile(imgUrl)
			}
		})

		//获取新URL
		aSelect := document.Find("body").Find("a")
		aSelect.Each(func(i int, aTag *goquery.Selection) {
			aUrl, _ := aTag.Attr("href")
			//fmt.Println("待爬取：", aUrl)
			if strings.Index(aUrl, "/") != -1 {
				//myqueen.Push("https://www.4493.com" + aUrl)
				myqueen.Push(aUrl)
			}
		})
	})
	myqueen.CreateDataQueen(h)
	myqueen.Push(url)

}

func OtherImgSp(url string) {
	n := 10
	var q *otqueue.Queue
	q = otqueue.NewQueue(func(val interface{}) {
		if val.(string) == "" || strings.Index(val.(string), "java") != -1 {
			return
		}
		fmt.Println(val.(string))
		document, _ := goquery.NewDocument(val.(string))
		//下载图片
		bodySelect := document.Find("body").Find("img")
		bodySelect.Each(func(i int, imgSelect *goquery.Selection) {
			imgUrl, _ := imgSelect.Attr("src")
			//fmt.Println(imgUrl)
			if imgUrl != "" {
				myhttp.DownloadFile(imgUrl)
			}
		})

		//获取新URL
		aSelect := document.Find("body").Find("a")
		aSelect.Each(func(i int, aTag *goquery.Selection) {
			aUrl, _ := aTag.Attr("href")
			//fmt.Println("待爬取：", aUrl)
			if strings.Index(aUrl, "/") != -1 {
				//myqueen.Push("https://www.4493.com" + aUrl)
				q.Push(aUrl)
			}
		})
	}, n)
	q.Push(url)
	q.Wait()
}
