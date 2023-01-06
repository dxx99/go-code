package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{8,8,7,7,7}))

	//fmt.Println(majorityElement([]int{1}))
	//fmt.Println(majorityElement([]int{3,2,3}))
	//fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}

// 分治求解
func majorityElement(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := (lo + hi)>>1
	left := helper(nums, lo, mid)
	right := helper(nums, mid+1, hi)

	if left == right {
		return left
	}

	// 比较左右数的多数
	leftCount := countInRange(nums, left, lo, mid)
	rightCount := countInRange(nums, right, mid+1, hi)
	if leftCount > rightCount {
		return left
	}
	return right
}

func countInRange(nums []int, left, lo, hi int) int {
	count := 0
	for i := lo; i <= hi; i++ {
		if nums[i] == left {
			count++
		}
	}
	return count
}
