package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0}))
	fmt.Println(missingNumber([]int{0,1,2,3}))
	fmt.Println(missingNumber([]int{0,1,3}))
	fmt.Println(missingNumber([]int{0,1,2,3,4,5,6,7,9}))
}

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1

	// 缺失最右边的数据
	if len(nums) != nums[right] {
		return len(nums)
	}

	for left < right {
		mid := (right+left)>>1
		if nums[mid] == mid {
			left = mid+1
		}else {
			right = mid
		}
	}

	return right
}
