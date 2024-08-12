package main
import "fmt"

func main() {
	var arr1 [5]int

	for i , _ := range arr1 {
		fmt.Printf("index: %d , value: %d\n" , i , arr1[i]);
	}
	
}