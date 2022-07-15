package main

import "math"

//162
//154
//528
//1508
//1574
//1292
//1498
//981
//1300
//1802
//1901
//1146
//1488
func main() {

}

// 162 寻找峰值
// https://leetcode.cn/problems/find-peak-element/
// 思路：
//	1. 寻找最大值, 遍历取最大值即可
//  2. 根据 nums[i-1], nums[i], nums[i+1] 三者的关系决定向哪个方向走
//		- 如果nums[i]大于左右两边的值，直接返回
//		- 如果上坡, 也就是 nums[i-1]<nums[i]<nums[i+1], 最大值在右边，left=mid+1
//		- 如果下坡，也就是 nums[i-1]>nums[i]>nums[i+1], 最大值在左边，right=mid-1
//		- 如果低谷，也就是 nums[i-1]>nums[i]<nums[i+1], 也就是两侧都可以走
//	- 总结：
//		- nums[i] < nums[i+1], 右边
//		- nums[i] > nums[i+1], 左边
func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]
	}

	left, right := 0, len(nums)-1
	for{
		mid := int(uint(left+right)>>1)
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}else if get(mid) < get(mid+1) {
			left = mid+1
		}else {
			right = mid-1
		}
	}

}

// 1901 矩阵中寻找峰值
// https://leetcode.cn/problems/find-a-peak-element-ii/
// 思路：
//	1. 左右找最大的
//	2. 上下找峰值
func findPeakGrid(mat [][]int) []int {
	getMaxIndex := func(arr []int) int {
		m, index := arr[0], 0
		for i := 1; i < len(arr); i++ {
			if arr[i] > m {
				m = arr[i]
				index = 0
			}
		}
		return index
	}

	m := len(mat)
	left, right := 0, m-1
	for left < right {
		mid := int(uint(left+right)>>1)
		// 左右判断
		maxIndex := getMaxIndex(mat[mid])

		// 上下判断
		if mat[mid][maxIndex] > mat[mid+1][maxIndex]  {
			right = mid
		}else {	// 相等情况要向下移
			left = mid+1
		}
	}

	maxIndex := getMaxIndex(mat[left])
	return []int{left, maxIndex}
}


