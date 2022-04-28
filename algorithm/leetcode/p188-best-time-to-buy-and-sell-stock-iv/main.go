package main

import "fmt"

// p188 买股票的最佳时机iv
// 给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 
//
//示例 1：
//
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//示例 2：
//
//输入：k = 2, prices = [3,2,6,5,0,3]
//输出：7
//解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。
// 
//
//提示：
//
//0 <= k <= 100
//0 <= prices.length <= 1000
//0 <= prices[i] <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxProfit(1, []int{2,1}))
	fmt.Println(maxProfit(2, []int{2,4,1}))
	fmt.Println(maxProfit(2, []int{3,2,6,5,0,3}))
}

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// 可以存在的操作策略 1次不操作，2k次买入卖出操作
	op := 2*k+1

	// dp数组的定义
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, op)
	}

	// 初始化操作
	for i := 1; i < op; i++ {
		if i % 2 == 1 {
			dp[0][i] = -prices[0]
		}
	}

	// 递推函数
	//dp[i][j] = max(dp[i-1][j], dp[i-1][j-1] +/- prices[i])
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < op; j++ {	// 对应op种操作
			if j % 2 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1] - prices[i])
			}else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1] + prices[i])
			}
		}
	}

	return dp[len(prices)-1][op-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
