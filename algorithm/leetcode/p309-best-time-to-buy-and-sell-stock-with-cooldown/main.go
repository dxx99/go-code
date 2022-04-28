package main

import "fmt"

// p309 最佳买卖股票时机含冷冻期
// 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 
//
//示例 1:
//
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//示例 2:
//
//输入: prices = [1]
//输出: 0
// 
//
//提示：
//
//1 <= prices.length <= 5000
//0 <= prices[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxProfit([]int{1,2,3,0,2}))
}

func maxProfit(prices []int) int {
	// dp 数组的定义
	// dp[i][j]  第i天，对应的第j种状态
	// 状态1 ->  达到买入股票的状态
	//		1. 前一天就买入了该股票
	//		2. 今天买入了股票(1. 前一天是冷冻期状态才能买入股票，2. 前一天保持着卖出股票的状态)
	// 状态2 -> 卖出股票状态(以前卖出的)
	//		1. 前一天就是卖出股票的状态
	//		2. 前一天是冷冻期
	// 状态3 -> 今天卖出股票（这里要分开，因为这里买上今天冷冻期，今天卖出的不能买）
	//		1. 前一天一定是买入股票的状态
	// 状态4 -> 冷冻期状态
	//		1. 前一天刚卖出股票
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 4)
	}

	//初始化
	dp[0][0] = -prices[0]

	// 遍历
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3], dp[i-1][1]) - prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])	//
		dp[i][2] = dp[i-1][0] + prices[i]	//
		dp[i][3] = dp[i-1][2]	// 前一天刚卖出
	}

	return max(dp[len(prices)-1][1], max(dp[len(prices)-1][2], dp[len(prices)-1][3]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
