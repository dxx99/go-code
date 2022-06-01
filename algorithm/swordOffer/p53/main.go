package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(search([]int{2,2}, 3))
	fmt.Println(search([]int{5,5,7,7,8,8,8,8,8,8,10}, 6))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	k := sort.SearchInts(nums, target)
	if k < 0 || k >= len(nums) || nums[k] != target {
		return 0
	}

	ans := 1
	k = k+1
	for k < len(nums) && nums[k] == target {
		ans ++
		k++
	}
	return ans
}

func searchV2(nums []int, target int) int {
	left := sort.SearchInts(nums, target)
	if len(nums) == left || nums[left] != target {
		return 0
	}

	right := sort.SearchInts(nums, target+1)-1
	return right - left + 1
}
