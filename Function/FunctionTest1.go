package main

import "fmt"

func main() {
	p1 := Add2(1 , 2)
	// fmt.Println(p2(5, 7))
	fmt.Printf("%T\n" , p1)

	p2 := func(i int) {fmt.Println(i)}
	fmt.Printf("%T" , p2)
}

func Add2(a , b int) int {
	return a + b
}