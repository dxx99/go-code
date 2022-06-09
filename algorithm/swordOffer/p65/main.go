package main

import "fmt"

func main() {
	fmt.Println(add(111,899))
}

func add(a int, b int) int {
	for b != 0 {	// 当b进位为0时跳出
		c := (a & b) << 1	// 进位
		a ^= b	// a = 非进位和
		b = c	// b = 进位
	}
	return a
}
