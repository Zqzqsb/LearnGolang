package main

import (
	"fmt"
)

// worker 函数，用于计算数组一部分的和，并将结果发送到通道
func worker(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 将计算结果发送到 channel
	ch <- sum
}

func main() {
	// 定义一个大数组
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 创建一个 channel，用于传递部分和
	ch := make(chan int)

	// 启动两个 Goroutines，分别计算数组的前半部分和后半部分的和
	go worker(nums[:len(nums)/2], ch) // 前半部分
	go worker(nums[len(nums)/2:], ch) // 后半部分

	// 从 channel 接收两个部分的结果
	sum1 := <-ch
	sum2 := <-ch

	// 计算总和
	total := sum1 + sum2

	// 打印结果
	fmt.Println("总和:", total)
}
