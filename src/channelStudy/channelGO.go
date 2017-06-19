package channelStudy

import (
	"fmt"
	"time"
)

func getMessagesChannel(msg string, delay time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println("goroutine")
			a := fmt.Sprintf("%s %d", msg, i)
			c <- a
			//c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * delay) /** 仅仅起到，下一次的 c 在何时输入 */
		}
	}()
	return c
}

func Test16()  {
	/** 编译通过 */
	/** 复杂的演示例子 */
	/** 多channel模式 */
	c1 := getMessagesChannel("第一", 600 )
	c2 := getMessagesChannel("第二", 500 )
	c3 := getMessagesChannel("第三", 5000)
	//time.Sleep(time.Second)

	/** 层层限制阻塞 */
	/** 这个 for 里面会造成等待输入，c1 会阻塞 c2 ,c2 阻塞 c3 */
	/** 所以它总是，先输出 c1 然后是 c2 最后是 c3 */
	for i := 1; i <= 3; i++ {
		/** 每次循环提取一轮，共三轮 */
		println(<-c1)  /** 除非 c1 有输入值，否则就阻塞下面的 c2,c3 */
		println(<-c2)  /** 除非 c2 有输入值，否则就阻塞下面的 c3 */
		println(<-c3)  /** 除非 c3 有输入值，否则就阻塞进入下一轮循环，反复如此 */
	}
	/**
	 *  这个程序的运行结果，首轮的，第一，第二，第三 很快输出，因为
	 *  getMessagesChannel 函数的延时 在 输入值之后，在第二轮及其之后
	 *  因为下一个 c3 要等到 5秒后才能输入，所以会阻塞第二轮循环的开始5秒，如此反复。
	 */
	/** 修改：如果把 getMessagesChannel 里面的延时，放在输入值之前，那么 c3 总是等待 5秒 后输出 */
}
