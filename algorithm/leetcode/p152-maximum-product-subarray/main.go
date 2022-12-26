package main

import "fmt"

// p152 乘积最大子数组
// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//测试用例的答案是一个 32-位 整数。
//
//子数组 是数组的连续子序列。
//
// 
//
//示例 1:
//
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//示例 2:
//
//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
// 
//
//提示:
//
//1 <= nums.length <= 2 * 104
//-10 <= nums[i] <= 10
//nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/maximum-product-subarray
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxProductV2([]int{2,3,-2,4}))
	//fmt.Println(maxProductV2([]int{-2,0,-1}))
}

// 动态规划 优化存储空间, 使用滚动数组
func maxProductV2(nums []int) int {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 初始化
	minF, maxF, ans:= nums[0], nums[0], nums[0]
	// 状态转移方程
	for i := 1; i < len(nums); i++ {
		mn, mx := minF, maxF
		minF = min(min(mx*nums[i], mn*nums[i]), nums[i])
		maxF = max(max(mx*nums[i], mn*nums[i]), nums[i])
		ans = max(ans, maxF)
	}
	return ans
}


// 动态规划
func maxProduct(nums []int) int {
	// dp数组的定义， dp[i][0] 表示选择当前点的最小值
	// 				dp[i][1] 表示选择当前点的最大值
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	// 初始化
	ans := nums[0]
	dp[0][0], dp[0][1] = nums[0], nums[0]

	// 状态转移方程
	for i := 1; i < len(nums); i++ {
		dp[i][0] = min(min(dp[i-1][1]*nums[i], dp[i-1][0]*nums[i]), nums[i])
		dp[i][1] = max(max(dp[i-1][1]*nums[i], dp[i-1][0]*nums[i]), nums[i])
		ans = max(ans, dp[i][1])
	}

	return ans
}
