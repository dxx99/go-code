package main

import (
	"fmt"
)

func main() {
	fmt.Println(4^1)
	fmt.Println(singleNonDuplicate([]int{1,1,2,2,4,4,5,8,8}))
}

// 540 有序数组中的单一元素
// https://leetcode.cn/problems/single-element-in-a-sorted-array/
// 分析:
//	如果 mid 是偶数，则比较nums[mid] 和 nums[mid+1] 是否相等；
//	如果 mid 是奇数，则比较nums[mid−1] 和 nums[mid] 是否相等。
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right)>>1)
		// 如果mid为奇数，mid^1 表示减1
		// 如果mid为偶数，mid^1 表示加1
		if nums[mid] == nums[mid^1] {
			left = mid+1
		} else {
			right = mid
		}
	}
	return nums[left]
}


