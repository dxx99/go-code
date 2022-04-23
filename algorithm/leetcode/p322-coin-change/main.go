package main

import (
	"fmt"
	"math"
)

// p322 零钱兑换
// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
//你可以认为每种硬币的数量是无限的。
//
// 
//
//示例 1：
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//示例 2：
//
//输入：coins = [2], amount = 3
//输出：-1
//示例 3：
//
//输入：coins = [1], amount = 0
//输出：0
// 
//
//提示：
//
//1 <= coins.length <= 12
//1 <= coins[i] <= 231 - 1
//0 <= amount <= 10
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/coin-change
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(coinChangeV2([]int{2,5,10,1}, 27))
	fmt.Println(coinChangeV2([]int{2}, 3))
	fmt.Println(coinChangeV2([]int{1,2,5}, 11))
	fmt.Println(coinChangeV2([]int{1}, 0))
}

func coinChangeV2(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// 初始化
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	// 递推公式, 尽可能少的元素
	// dp[j] = min(dp[j], dp[j-coins[i]]+1)
	// 先遍历背包，再遍历物品
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount ; j++ {
			if dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}


func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// 初始化
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}

	// 递推公式, 尽可能少的元素
	// dp[j] = min(dp[j], dp[j-coins[i]]+1)

	// 遍历
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount ; j++ {
			if dp[j] == -1 || dp[j - coins[i]] == -1 {
				if dp[j-coins[i]] >= 0 {	// 只有不等于-1才能借助向上加一
					dp[j] = dp[j-coins[i]] +1
				}
			} else {
				if dp[j] > dp[j-coins[i]]+1 {
					dp[j] = dp[j-coins[i]]+1
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[amount]
}
