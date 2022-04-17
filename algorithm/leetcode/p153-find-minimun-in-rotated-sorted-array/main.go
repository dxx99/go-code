package main

import (
	"fmt"
	"math"
)

// p153 寻找旋转排序数组中的最小值
// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
//若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
//若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
//给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
//
//你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
//
// 
//
//示例 1：
//
//输入：nums = [3,4,5,1,2]
//输出：1
//解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
//示例 2：
//
//输入：nums = [4,5,6,7,0,1,2]
//输出：0
//解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
//示例 3：
//
//输入：nums = [11,13,15,17]
//输出：11
//解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。
// 
//
//提示：
//
//n == nums.length
//1 <= n <= 5000
//-5000 <= nums[i] <= 5000
//nums 中的所有整数 互不相同
//nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findMin([]int{2,1}))
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
	fmt.Println(findMin([]int{1}))
}

func findMin(nums []int) int {
	l := len(nums)
	minVal := math.MaxInt32

	if l == 0 {
		return minVal
	}

	left, right := 0, l-1
	for left <= right{
		mid := (left+right) >> 1

		if minVal > nums[mid] {
			minVal = nums[mid]
		}

		// 要精确的判断最小的元素在那一边
		if nums[left] <= nums[mid] && nums[mid] >= nums[right] {
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return minVal
}
