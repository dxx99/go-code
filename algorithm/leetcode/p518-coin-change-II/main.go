package main

import "fmt"

// p518 零钱兑换
// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
//
//请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
//假设每一种面额的硬币有无限个。 
//
//题目数据保证结果符合 32 位带符号整数。
//
// 
//
//示例 1：
//
//输入：amount = 5, coins = [1, 2, 5]
//输出：4
//解释：有四种方式可以凑成总金额：
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
//示例 2：
//
//输入：amount = 3, coins = [2]
//输出：0
//解释：只用面额 2 的硬币不能凑成总金额 3 。
//示例 3：
//
//输入：amount = 10, coins = [10]
//输出：1
// 
//
//提示：
//
//1 <= coins.length <= 300
//1 <= coins[i] <= 5000
//coins 中的所有值 互不相同
//0 <= amount <= 5000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/coin-change-2
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(change(5, []int{1,2,5}))
	fmt.Println(changeV2(5, []int{1,2,5}))
}

func change(amount int, coins []int) int {
	// 硬币只有取与不取，定义dp数组
	dp := make([]int, amount+1)

	// 初始化，总价值为零的时候，也就是都不选，只有1种方法
	dp[0] = 1

	//递推公式 dp[j] += dp[j - coins[i]] + 1
	//遍历， 先遍历背包，在遍历容量
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j - coins[i]]
			fmt.Println(dp)
		}
	}

	return dp[amount]
}

// 这里是不通的，不能先遍历背包，再遍历物品
// 本题是求凑出来的方案个数，且每个方案个数是为组合数。
func changeV2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for j := 0; j <= amount; j++ {
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] {
				dp[j] += dp[j - coins[i]]
				fmt.Println(dp)
			}
		}
	}
	return dp[amount]
}
