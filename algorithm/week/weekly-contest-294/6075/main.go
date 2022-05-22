package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBags([]int{2,3,4,5}, []int{1,2,4,4}, 2))
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	// 背包剩余的空间
	leftArr := make([]int, len(capacity))
	for i := 0; i < len(capacity); i++ {
		leftArr[i] = capacity[i] - rocks[i]
	}

	sort.Ints(leftArr)

	ans := 0
	for j := 0; j < len(leftArr); j++ {
		if leftArr[j] == 0 {
			ans++
		}else {
			additionalRocks = additionalRocks - leftArr[j]
			if additionalRocks >= 0 {
				ans++
			}
		}
	}

	return ans
}
