package main

import (
	"math"
)

func main() {
	printNumbers(3)
}

func printNumbers(n int) []int {
	ans := make([]int, 0)
	for i := 1; i < int(math.Pow10(n)); i++ {
		ans = append(ans, i)
	}
	return ans
}
