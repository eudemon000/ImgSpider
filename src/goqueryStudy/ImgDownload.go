package goqueryStudy

import (
	myqueen "ImgSpider/src/queen"
	"github.com/PuerkitoBio/goquery"
	_"fmt"
	myhttp "ImgSpider/src/http"
	"fmt"
	"strings"
	otqueue "github.com/otium/queue"
	"regexp"
)

var cUrl string

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
		cUrl = document.Url.String()
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
	if len(url) == 0 {
		return
	}
	if len(cUrl) == 0 {
		cUrl = url
	}
	n := 10
	var q *otqueue.Queue
	q = otqueue.NewQueue(func(val interface{}) {
		if val.(string) == "" || strings.Index(val.(string), "java") != -1 {
			return
		}
		fmt.Println(val.(string))
		document, _ := goquery.NewDocument(val.(string))
		//下载图片
		b := document.Find("body")
		if b == nil {
			return
		}
		bodySelect := b.Find("img")
		bodySelect.Each(func(i int, imgSelect *goquery.Selection) {
			imgUrl, _ := imgSelect.Attr("src")
			r, ok := Format(imgUrl)
			if ok {
				myhttp.DownloadFile(r)
			}

			//fmt.Println(imgUrl)
			/*if imgUrl != "" {
				myhttp.DownloadFile(imgUrl)
			}*/
		})
		cUrl = document.Url.String()

		//获取新URL
		aSelect := document.Find("body").Find("a")
		aSelect.Each(func(i int, aTag *goquery.Selection) {
			aUrl, _ := aTag.Attr("href")
			fmt.Println("待爬取：", aUrl)
			r, ok := Format(aUrl);
				if ok {
					q.Push(r)

			}
		})
	}, n)

	q.Push(url)
	q.Wait()
}

func Format(str string) (result string, ok bool) {
	//fmt.Println("接口方法调用", str)
	//首先判断是不是是不是javascript，#或*开头的,如果是代表不是合法URL
	ok, err := regexp.MatchString("^javascript|^#|^\\*", str)
	if err !=nil {
		fmt.Println(err)
		return "", false
	}
	if ok {
		return "", false
	}

	//判断是不是http开头的，http和https均可判断
	ok, err = regexp.MatchString("^http", str)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	if ok {
		return str, true
	}

	//还要一种是相对路径，分两种情况，1、"/"开头；2、非"/"开头
	ok, err = regexp.MatchString("^/{1}[a-zA-Z0-9]{1,}?", str)
	if ok {
		//需要找路径根
		strs := strings.Split(cUrl, "/")
		re := strs[0] + "//" + strs[2] + str
		return re, false
	}

	ok, err = regexp.MatchString("[a-zA-Z0-9]{1,}?", str)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	if ok {
		postion := strings.LastIndex(cUrl, "/")
		postion += 1
		a := cUrl[0:postion]
		re := a + str
		return re, true
	}
	return "", false
}
