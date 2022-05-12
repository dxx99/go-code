package main

import "fmt"

// p674 最长连续递增序列
// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
//连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
//
// 
//
//示例 1：
//
//输入：nums = [1,3,5,4,7]
//输出：3
//解释：最长连续递增序列是 [1,3,5], 长度为3。
//尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
//示例 2：
//
//输入：nums = [2,2,2,2,2]
//输出：1
//解释：最长连续递增序列是 [2], 长度为1。
// 
//
//提示：
//
//1 <= nums.length <= 104
//-109 <= nums[i] <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findLengthOfLCIS([]int{1,3,5,7}))
	fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7}))
	fmt.Println(findLengthOfLCIS([]int{2,2,2,2}))
	fmt.Println(findLengthOfLCIS([]int{1}))
	fmt.Println("V2.......................")

	fmt.Println(findLengthOfLCISV2([]int{1,3,5,7}))
	fmt.Println(findLengthOfLCISV2([]int{1,3,5,4,7}))
	fmt.Println(findLengthOfLCISV2([]int{2,2,2,2}))
	fmt.Println(findLengthOfLCISV2([]int{1}))
}

// 动态规划
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {return 0}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		if dp[i+1] > res {
			res = dp[i+1]
		}
	}
	return res
}

// 贪心算法
func findLengthOfLCISV3(nums []int) int {
	if len(nums) == 0 {return 0}
	res, count := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			count++
		}else {
			count = 1
		}
		if count > res {
			res = count
		}
	}
	return res
}



// 双指针求解
func findLengthOfLCISV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, ans := 0, 1
	for right := 1; right < len(nums); right++ {
		if nums[right] <= nums[right-1] {
			if right - left > ans {
				ans = right - left
			}
			left = right
		}
	}
	if len(nums) - left > ans {
		ans = len(nums) - left
	}
	return ans
}
