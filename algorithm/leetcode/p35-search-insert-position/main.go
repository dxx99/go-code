package main

import "fmt"

// p35 搜索并插入元素到数组中
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。
//
// 
//
//示例 1:
//
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//示例 2:
//
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//示例 3:
//
//输入: nums = [1,3,5,6], target = 7
//输出: 4
// 
//
//提示:
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 为 无重复元素 的 升序 排列数组
//-104 <= target <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-insert-position
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 处理找不到的结果
	if nums[left] < target  {
		return left + 1
	}
	return left
}
