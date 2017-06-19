package queue

import (
	"container/list"
	"sync"
)


type msgQueue struct {
	l list.List
	lock *sync.Mutex
	size int
}

type MsgHandler struct {
	h Handle
}

type Handle func(data interface{})

var mQueue *msgQueue

var handler *MsgHandler

func InitData(h Handle) *MsgHandler {
	mQueue = new(msgQueue)
	mQueue.l = list.New()
	mQueue.lock = sync.Mutex{}
	mQueue.size = 0
	handler = new(MsgHandler)
	handler.h = h
	for i := 0; i < 20; i++ {
		go loop()
	}
	return handler
}

func (*MsgHandler)PushData(data interface{}) {
	mQueue.lock.Lock()
	defer mQueue.lock.Unlock()
	mQueue.l.PushBack(data)
	mQueue.size = mQueue.l.Len()
}

func readData() interface{} {
	mQueue.lock.Lock()
	defer mQueue.lock.Unlock()
	if mQueue.size  > 0 {
		item := mQueue.l.Front()
		mQueue.l.Remove(item)
		mQueue.size = mQueue.size
		return item
	}
	return nil
}

func loop() {
	for {
		data := readData()
		if data != nil {
			handler.h(data)
		}
	}

}
