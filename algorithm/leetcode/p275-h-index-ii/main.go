package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(hIndexV2([]int{0,1,3,5,6}))
}

func hIndex(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
}


func hIndexV2(citations []int) int {
	left, right := 0, len(citations)
	for left < right {
		mid := int(uint(left+right)>>1)
		if citations[mid] < len(citations)-mid {
			left = mid+1
		}else {
			right = mid
		}
	}
	return len(citations) - left
}
