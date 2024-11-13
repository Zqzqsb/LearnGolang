package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 定义一个 int 类型的变量
	var num int = 42

	// 使用 reflect.ValueOf 获取变量的 reflect.Value
	value := reflect.ValueOf(num)

	// 调用 Type() 方法获取变量的 reflect.Type
	typ := value.Type()

	// 打印变量的值和类型信息
	fmt.Printf("值: %v\n", value)
	fmt.Printf("类型: %s\n", typ)

	// 进一步演示 reflect.Type 的功能
	fmt.Printf("类型名: %s\n", typ.Name())
	fmt.Printf("种类: %s\n", typ.Kind())
}
