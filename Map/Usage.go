package main

import "fmt"

func main() {
	m := make(map[string]int) // create
	m["one"] = 1              // insert
	m["two"] = 2

	fmt.Println(m)                  // formart printer
	fmt.Println(m["one"], m["two"]) // retrieve
	fmt.Println(m["unknown"])       // 0

	r, ok := m["unknown"]
	fmt.Println(r, ok) // false

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)

	// traverse
	for key, value := range m2 {
		fmt.Println(key, value)
	}
}
