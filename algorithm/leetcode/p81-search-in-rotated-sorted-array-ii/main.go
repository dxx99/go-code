package main

import "fmt"

func main() {
	fmt.Println(searchII([]int{1,0,1,1,1}, 0))
	//fmt.Println(search([]int{2,5,6,0,0,1,2}, 0))
}

// 81. 搜索旋转排序数组 II
// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/
// 题意：nums中的元素有可能相同
// 分析：
// 	  1. 看数组元素的递增曲线
//	  2. 每次判断递增区间，这样就可以排除到无序的区间
func searchII(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := int(uint(left+right)>>1)

		if target == nums[mid] {
			return true
		}

		// target [left, mid)
		if nums[mid] == nums[left] && nums[right] == nums[mid] {	// 如果三个数相等就没有办法判断
			left++
			right--
		}else if nums[mid] >= nums[left] {	// 左边有序
			if target < nums[mid] && nums[left] <= target {
				right = mid-1
			}else {
				left = mid+1
			}
		} else {	// 右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid+1
			}else {
				right = mid-1
			}
		}
	}

	if left < len(nums)-1 && nums[left] == target {
		return true
	}
	return false
}


// 33 搜索旋转数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/
func searchI(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}

	left, right := 0, l-1	// 细节1: 容易越界
	for left <= right {		// 细节2: 需要判断等于计算
		mid := int(uint(left+right)>>1)
		if nums[mid] == target {
			return mid
		}

		// 细节3: 查看划分的数组，那边有序
		if nums[left] <= nums[mid] { //左边有序, 则判断左右
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { 	// 右边有序，则判断右边的数据
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}