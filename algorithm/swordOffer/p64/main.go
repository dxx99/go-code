package main

import "fmt"

func main() {
	fmt.Println(sumNums(3))
}

func sumNums(n int) int {
	if n == 1 {
		return 1
	}

	return sumNums(n-1) + n
}
