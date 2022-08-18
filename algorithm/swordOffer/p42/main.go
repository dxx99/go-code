package main

import "fmt"

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//
//要求时间复杂度为O(n)。
//
// 
//
//示例1:
//
//输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
//提示：
//
//1 <= arr.length <= 10^5
//-100 <= arr[i] <= 100
//注意：本题与主站 53 题相同：https://leetcode-cn.com/problems/maximum-subarray/
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{-1}))
}

// 贪心算法
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans, cur := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = 0
		}
		cur += nums[i]
		if cur > ans {
			ans = cur
		}
	}
	return ans
}

// dynamicProgramming
func maxSubArrayV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]

	max := func (x, y int) int {
		if x > y {
		return x
	}
		return y
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

