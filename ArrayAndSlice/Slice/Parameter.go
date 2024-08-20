package main
import "fmt"

// 接受切片作为参数
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4}
	// 传递一个切片
	fmt.Println(sum(arr[:]))
}