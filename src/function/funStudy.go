package function

import (
	"fmt"
)

//可变参数
func MutiFunc(args ...int) {
	for index, item := range args {
		fmt.Println("index:", index, "item:", item)
	}
}

//可变参数，接口类型，更加灵活
func MutiInterFunc(args ...interface{}) {
	for index, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(index, arg, "int")
		case string:
			fmt.Println(index, arg, "string")
		case int64:
			fmt.Println(index, arg, "int64")
		default:
			fmt.Println(index, arg, "unknown")
			
		}
	}
}