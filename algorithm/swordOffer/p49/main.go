package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumberV2(10))
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() interface{} {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return x
}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	hash := make(map[int]bool)
	hash[1] = true

	factors := []int{2,3,5}

	count := 1
	for count <= 1690 {
		x := heap.Pop(h).(int)
		if count == n {
			return x
		}

		for i := 0; i < len(factors); i++ {
			next := x * factors[i]
			if _, ok := hash[next]; !ok {
				heap.Push(h, next)
				hash[next] = true
			}
		}
		count++
	}
	return -1
}

func nthUglyNumberV2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
