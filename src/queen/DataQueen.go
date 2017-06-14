package queen

import (
	"fmt"
)

//var myList *list.List
var mainCh chan interface{} = make(chan interface{}, 5000)

type Handler func(data interface{})

var handler Handler

func CreateDataQueen(h Handler) {
	//myList = list.New()
	handler = h
	go read()
}

//入队
func Push(data interface{}) {
	//myList.PushBack(data)
	mainCh <- data
}

func read() {
	for {
		var d interface{} = <- mainCh
		fmt.Println("我在读", d)
		handler(d)
	}
}



