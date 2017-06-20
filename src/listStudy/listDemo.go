package listStudy

import (
	"container/list"
	"fmt"
)

//list遍历，请参考http://www.cnblogs.com/ziyouchutuwenwu/p/3780800.html
func Demo1() {
	l := list.New()
	l.PushBack("1=")
	l.PushBack("2=")
	l.PushBack("3=")
	l.PushBack("4=")

	fmt.Println("start")
	var n *list.Element
	for e := l.Front(); e != nil; e = n {
		fmt.Println(e.Value)
		n = e.Next()
		l.Remove(e)
	}
	fmt.Println("End")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

