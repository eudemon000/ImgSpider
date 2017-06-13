package goqueryStudy

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/henrylee2cn/pholcus/common/mahonia"
)

func Parse(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
	}

	/*for _, item := range doc.Nodes {
		fmt.Println(item.Namespace)
	}*/
	se := doc.Find("meta")
	se.Each(func(i int, s *goquery.Selection) {
		a := s.Nodes
		for _, item := range a {
			fmt.Println(item.Attr[0].Key)
		}

	})

}

func ParseForCharset(url string) {
	doc, _ := goquery.NewDocument(url)
	headTag := doc.Find("head")
	metaTag := headTag.Find("meta")
	metaTag.Each(func(i int, metaSele *goquery.Selection) {
		equivValue, ok := metaSele.Attr("http-equiv")
		content, ok := metaSele.Attr("content")
		if ok {

			fmt.Println(content)
			de := mahonia.NewDecoder("gb2312")
			result := de.ConvertString(content)
			fmt.Println(result)
		}
		if ok {
			fmt.Println(equivValue)
		}
	})

	/*metaBody := doc.Find("body")
	de := mahonia.NewDecoder("gb2312")
	result := de.ConvertString(metaBody.Text())
	fmt.Println(result)*/



}
