package main

import "fmt"

func main() {
	fmt.Println(waysToSplitArray([]int{10,4,-8,7}))
	fmt.Println(waysToSplitArray([]int{2,3,1,0}))
	fmt.Println(waysToSplitArray([]int{-100000,100000}))
}

func waysToSplitArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	left := 0
	res := 0
	for j := 0; j < len(nums)-1; j++ {
		left += nums[j]
		if left >= sum - left {
			res++
		}
	}

	return res
}
