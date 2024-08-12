package main
import "fmt"

func main() {
	var arr1 = new([5]int)
	arr1[3] = 100
	var arr2 = arr1 // shallow copy
	arr2[3] = 99
	fmt.Println(arr1[3] , arr2[3])

	var arr3 [5]int = [...]int{1 , 2 , 3 , 4 , 5}
	arr3[3] = 100
	var arr4 = arr3 // deep copy
	arr4[3] = 99
	fmt.Println(arr3[3] , arr4[3])
}