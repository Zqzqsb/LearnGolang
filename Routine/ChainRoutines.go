package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) { left <- 1 + <-right }

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int) // 左等于右 右等于新
		go f(left, right)                   // 像构造链表一样串起诺干个通道
	}
	right <- 0 // 从右侧开始传递
	x := <-leftmost
	fmt.Println(x)
}
