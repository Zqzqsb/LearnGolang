package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string // 结构中只有以大写开头的导出字段是可设置
	Age  int
	money int // 可以通过反射获得 但不可以设置
}

// 方法：介绍自己
func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s, and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	// 获取结构体的反射值
	v := reflect.ValueOf(p)

	// 获取结构体的类型
	t := reflect.TypeOf(p)

	// 获取字段的数量
	numFields := v.NumField()
	fmt.Println("Number of fields:", numFields)

	// 遍历所有字段并打印字段的名字、类型和值
	for i := 0; i < numFields; i++ {
		field := v.Field(i)          // 获取第 i 个字段的值
		fieldName := t.Field(i).Name // 获取第 i 个字段的名字
		fieldType := t.Field(i).Type // 获取第 i 个字段的类型
		fmt.Printf("Field %d: Name=%s, Type=%s, Value=%v\n", i, fieldName, fieldType, field)
	}

	// 调用结构体的第 0 个方法（Introduce 方法是第 0 个方法）
	method := v.Method(0) // 获取第 0 个方法
	method.Call(nil)      // 调用方法（没有参数）
}
