package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestVariance("aababbb"))
	fmt.Println(largestVariance("abcde"))
}

func largestVariance(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		used := [26]int{}
		for j := i; j < len(s); j++ {
			used[s[j] - 'a']++
			v := getSub(used)
			if v > res {
				res = v
			}
		}
	}
	return res
}

func getSub(used [26]int) int {
	min, max := math.MaxInt, math.MinInt
	for i := 0; i < len(used); i++ {
		if used[i] != 0 {
			if min > used[i] {
				min = used[i]
			}
			if max < used[i] {
				max = used[i]
			}
		}
	}

	if min == math.MaxInt || max == math.MinInt {
		return 0
	}
	return max-min
}
