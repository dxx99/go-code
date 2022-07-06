package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxDistance([]int{1,2,3,4,7}, 3))
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	f := func(mid int) bool {
		cut := 1
		start := position[0]
		for i := 1; i < len(position); i++ {
			if position[i] - start >= mid {
				cut++
				start = position[i]
			}
		}
		return cut >= m
	}

	left, right := 1, position[len(position)-1]
	for left <= right {
		mid := int(uint(left+right)>>1)
		if f(mid) {
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return left-1
}
