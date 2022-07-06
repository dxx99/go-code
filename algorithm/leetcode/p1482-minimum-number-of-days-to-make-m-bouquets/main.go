package main

import "fmt"


func main() {
	fmt.Println(minDays([]int{7,7,7,7,12,7,7}, 2,3))
	fmt.Println(minDays([]int{1,10,3,10,2}, 3,1))
}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	f := func(day int) bool {
		s := 0
		c := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] <= day {
				c++
			}else {
				s += c/k
				c = 0
			}
		}
		// 补位
		s += c/k
		return s < m
	}

	left, right := 1, int(1e9)
	for left < right {
		day := int(uint(left+right)>>1)
		if f(day) {
			left = day+1
		}else {
			right = day
		}
	}

	return left
}
