package queen

import (
	_"fmt"
	"container/list"
	"sync"
	"fmt"
)


type queen struct {
	list *list.List
	lock *sync.Mutex
	size int
	data interface{}
	manager MsgQueenManager
}

type MsgQueenHandler func(data interface{})

type MsgQueenManager struct {
	Callback MsgQueenHandler
}

var q *queen

func CreateQueen(handler MsgQueenHandler, size int) {
	q = new(queen)
	q.list = list.New()
	q.lock = &sync.Mutex{}
	q.size = 0
	q.manager = MsgQueenManager{handler}
	for i := 0; i < size; i++ {
		go readData()
	}
}

//入队
func PushData(data interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.list.PushBack(data)
	q.size = q.list.Len()
	q.data = data
}

//出队
func readData() {
	for {
		q.lock.Lock()
		var element *list.Element
		if q.size > 0 {
			element = q.list.Front()
			q.list.Remove(element)
			q.size = q.list.Len()
			q.lock.Unlock()
			fmt.Println(element)
			q.manager.Callback(element.Value)
		}

	}
}



