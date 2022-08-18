package sw

import "math"

// dp相关题目

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = amount+1
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i>=coins[j] && dp[i-coins[j]] != amount+1 {
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
// 先遍历物品，再遍历背包
func coinChangeV2(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {	// 上一个元素有值可以找到
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
// 递归遍历求解【会超时】
func coinChangeV3(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	min := func(x, y int) int {
		if x>y {
			return y
		}
		return x
	}

	// 树的遍历[N叉树的遍历]
	res := math.MaxInt
	for i := 0; i < len(coins); i++ {
		cur := coinChangeV3(coins, amount-coins[i])	// 这里会有重叠子问题需要优化
		if cur == -1 {
			continue
		}
		res = min(res, cur+1)
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
//
func coinChangeV4(coins []int, amount int) int {
	min := func(x, y int) int {
		if x>y {
			return y
		}
		return x
	}

	// 使用额外的空间进行存储
	used := make([]int, amount+1)
	for i := 0; i < len(used); i++ {
		used[i] = math.MinInt
	}

	var helper func(amount int) int
	helper = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		// 存储已经计算的值，做剪枝操作
		if used[amount] != math.MinInt {
			return used[amount]
		}

		res := math.MaxInt	// 前序位置
		for i := 0; i < len(coins); i++ {
			cur := helper(amount-coins[i])
			if cur == -1 {
				continue
			}
			res = min(res, cur+1)	// 后序位置
		}
		if res == math.MaxInt {
			return -1
		}
		used[amount] = res
		return res
	}

	return helper(amount)
}