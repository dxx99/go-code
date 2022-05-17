package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {

	hash := make(map[int]int)

	getSqrtNum := func(n int) int {
		res := 0
		for n >= 10 {
			k := n % 10
			n = n / 10
			res += int(math.Pow(float64(k), 2))
		}
		res += int(math.Pow(float64(n), 2))
		return res
	}

	for {
		hash[n]++
		if hash[n] > 1 {
			return false
		}

		n = getSqrtNum(n)

		if n == 1 {
			return true
		}
	}
}

