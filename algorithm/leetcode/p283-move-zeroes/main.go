package main

import (
	"fmt"
)

// p283 移动零
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
//示例 1:
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//示例 2:
//
//输入: nums = [0]
//输出: [0]
//
//
//提示:
//
//1 <= nums.length <= 104
//-231 <= nums[i] <= 231 - 1
//
//
//进阶：你能尽量减少完成的操作次数吗？
func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)


	nums2 := []int{0,0}
	moveZeroes(nums2)
	fmt.Println(nums2)

	nums3 := []int{2,1}
	moveZeroes(nums3)
	fmt.Println(nums3)
}
// 双指针求解
// left指向第一个零的位置，也就是左边都为非零数
// right指针左边直到左指针均为零
// 也就是左指针与右指针中间全部都是零，每次右指针发现不为零的元素，然后交换数据即可
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}


