package main

import (
	"fmt"
)

func main() {
	fmt.Println(1<<0)
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
}

func hammingWeight(num uint32) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if (num & (1<<i)) == (1<<i) {
			ans++
		}
	}
	return ans
}
