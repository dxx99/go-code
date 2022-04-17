package main

import "fmt"

// p33 搜索旋转排序数组
// 整数数组 nums 按升序排列，数组中的值 互不相同 。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 
//
//示例 1：
//
//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4
//示例 2：
//
//输入：nums = [4,5,6,7,0,1,2], target = 3
//输出：-1
//示例 3：
//
//输入：nums = [1], target = 0
//输出：-1
// 
//
//提示：
//
//1 <= nums.length <= 5000
//-10^4 <= nums[i] <= 10^4
//nums 中的每个值都 独一无二
//题目数据保证 nums 在预先未知的某个下标上进行了旋转
//-10^4 <= target <= 10^4
// 
//
//进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int{1}, 0))
}

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}

	left, right := 0, l-1	// 细节1: 容易越界
	for left <= right {		// 细节2: 需要判断等于计算
		mid := (left+right) >> 1
		if nums[mid] == target {
			return mid
		}

		// 细节3: 查看划分的数组，那边有序
		if nums[0] <= nums[mid] { //左边有序, 则判断左右
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { 	// 右边有序，则判断右边的数据
			if nums[mid] < target && target <= nums[l-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

