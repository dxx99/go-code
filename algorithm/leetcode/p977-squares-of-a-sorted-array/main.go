package main

import (
	"fmt"
	"math"
	"sort"
)

// p977 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
// 
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]
// 
//
//提示：
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums 已按 非递减顺序 排序
// 
//
//进阶：
//
//请你设计时间复杂度为 O(n) 的算法解决本问题
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(sortedSquares([]int{-4,-1,1,3,10}))
	fmt.Println(sortedSquares([]int{-7,-3,2,3,11}))
	fmt.Println("V2...........................")
	fmt.Println(sortedSquaresV2([]int{-4,-1,1,3,10}))
	fmt.Println(sortedSquaresV2([]int{-7,-3,2,3,11}))
}

func sortedSquares(nums []int) []int {
	for k := 0; k < len(nums); k++ {
		nums[k] = int(math.Pow(float64(nums[k]), 2))
	}

	sort.Sort(sort.IntSlice(nums))
	return nums
}

// 双指针求解
func sortedSquaresV2(nums []int) []int {
	l := len(nums)
	left, right := 0, l-1

	ans := make([]int, l)
	for left <= right {
		if nums[right]*nums[right] > nums[left]*nums[left] {
			ans[l-1] = nums[right]*nums[right]
			right--
		} else {
			ans[l-1] = nums[left]*nums[left]
			left++
		}
		l--
	}

	return ans
}