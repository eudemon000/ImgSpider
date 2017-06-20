package channelStudy

import (
	"fmt"
	"container/list"
	"time"
	"net/http"
	"io/ioutil"
)

type queue struct {
	l *list.List
	ch chan int
}

var q *queue

func init() {
	//初始化队列
	q = new(queue)
	q.l = list.New()
	q.ch = make(chan int)
	go readMsg()
	time.Sleep(time.Second)
}

func readMsg() {
	for {
		var n *list.Element
		for e := q.l.Front(); e != nil; e = n {
			n = e.Next()
			q.l.Remove(e)
			fmt.Println("队列读取", e.Value)
			test()
		}
		fmt.Println("循环执行完毕")
		q.ch <- 1
	}

}

func push(data string) {
	fmt.Println("向队列插入数据", data)
	q.l.PushBack(data)
	c := <- q.ch
	fmt.Println("读取到的通道值", c)
}

func Demo1() {
	//time.Sleep(time.Second * 3)
	push("1")
}

func test() {
	push("2")
}

func Demo2() {
	for i := 0; i < 1000000; i++ {
		go getWeb()
	}
}

func getWeb() {
	//resp, err := http.Get("http://219.136.249.21:8080/cancerKnowledge/user/v1/authorization/2c908184582d58680158995615e80042")
	resp, err := http.Get("http://www.sanpotel.com/")
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	defer resp.Body.Close()
}

