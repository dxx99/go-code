package main

import (
	"fmt"
	"sort"
)

// p34 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：
//
//你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
// 
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]
// 
//
//提示：
//
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//nums 是一个非递减数组
//-109 <= target <= 109
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//[-1 -1]
	//[3 4]
	//[5 5]
	//[-1 -1]
	//[-1 -1]
	//[0 3]
	fmt.Println(searchRange([]int{1,3}, 1))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 10))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1,1,1,1}, 1))
}
// binarySearch。。。。
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)
	lIndex, rIndex := -1, -1
	for left < right {
		mid := (left + right)>>1
		if nums[mid] > target {
			right = mid
		}else if nums[mid] < target {
			left = mid+1
		}else { // 找到，向两边扩展
			lIndex, rIndex = mid, mid
			for lIndex >=0 && nums[lIndex] == target {
				lIndex--
			}
			for rIndex <len(nums) && nums[rIndex] == target {
				rIndex++
			}
			return []int{lIndex+1, rIndex-1}
		}
	}
	return []int{lIndex, rIndex}
}

//
func searchRangeV2(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	// 没找到就返回长度
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}

	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}
