package main
import "fmt"

func main() {
	var arr1 [10]int
	var slice1 []int = arr1[2:5]
	
	for i:= 0; i < len(arr1); i++{
		arr1[i] = i
	}
	
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n" , i , slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n" , len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow th slice slice = slice[a:b]
	// a 标志原切片的起始位置 可以溢出增长 b要小于 切片的容量
	slice1 = slice1[0:4]
	for i:= 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i , slice1[i])
	}
	
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	
	// grow th slice slice = slice[a:b]
	// a 标志原切片的起始位置 可以溢出增长 b要小于 切片的容量
	slice1 = slice1[1:len(slice1)+1]
	for i:= 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i , slice1[i])
	}
	
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}