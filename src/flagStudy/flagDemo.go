package flagStudy

import (
	"flag"
	"fmt"
)

var firstFlag = flag.String("url", "http://www.qq.com/", "用法示例")

var isChange = flag.Bool("isChange", false, "Bool示例")

func TestFlag() {
	flag.Parse()
	fmt.Println("isChange", *isChange)
	fmt.Println("firstFlag", *firstFlag)
	fmt.Println(flag.Args())
}

