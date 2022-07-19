package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 2
	fmt.Println(m)

	v, ok := m[1]
	if ok {
		fmt.Println(v)
	}
}
