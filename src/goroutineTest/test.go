package goroutineTest

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

var count int = 0

func RunTest() {
	now := time.Now()
	l := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go add(i, l)
	}

	for {
		l.Lock()
		c := count
		l.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}

	}

	result := time.Since(now)
	fmt.Println("花费时间：", result)
}

func RunTest_1() {
	now := time.Now()
	chs := make([]chan int, 10000)
	for i := 0; i < 10000; i++ {
		chs[i] = make(chan int)
		go add_1(i, chs[i])
	}

	for _, ch := range chs {
		a := <- ch
		fmt.Println("结果为：", a)
	}

	result := time.Since(now)
	fmt.Println("花费时间：", result)
	fmt.Println("aaa")
}

func add(params int, l *sync.Mutex) {
	for i := 1; i < 1000000; i++ {
		fmt.Println(params + i)
	}
	l.Lock()
	count++
	defer l.Unlock()

}

func add_1(params int, ch chan int) {
	var a int
	for i := 1; i < 1000000; i++ {
		//fmt.Println(params + i)
		a = params + i
	}
	fmt.Println(params, "执行结束")
	ch <- a

}

func Test() {
	t1 := time.Now()
	var a int
	for i:= 0; i < 10000; i++ {
		for j := 0; j < 1000000; j++ {
			a = i + j
		}
	}
	fmt.Println(a)
	fmt.Println(time.Since(t1))
}

