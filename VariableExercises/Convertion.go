package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	// var b int64 编译错误，即使 a 和 b 的字节数一样
	// a = 15
	// b = a + a
	fmt.Println(unsafe.Sizeof(a))
}
