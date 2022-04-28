package main

import "fmt"

// p123 买卖股票的最佳时机III
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 
//
//示例 1:
//
//输入：prices = [3,3,5,0,0,3,1,4]
//输出：6
//解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//示例 2：
//
//输入：prices = [1,2,3,4,5]
//输出：4
//解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。   
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//示例 3：
//
//输入：prices = [7,6,4,3,1] 
//输出：0 
//解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
//示例 4：
//
//输入：prices = [1]
//输出：0
// 
//
//提示：
//
//1 <= prices.length <= 105
//0 <= prices[i] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxProfit([]int{1,2,4,2,5,7,2,4,9,0}))
	fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{1}))
}

// 动态规划求解
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	// dp[i][j] i表示第i天，j为[0-4]五个状态，dp[i][j]表示第i天状态j所剩最大现金
	// 0 表示没有操作
	// 1 第一次买入
	// 2 第一次卖出
	// 3 第二次买入
	// 4 第二次卖出

	// 确定递推公式：
	// dp[i][1] 状态有两个操作：
	//		- 操作一：第i天买入股票了，那么dp[i][1] = dp[i-1][0] - prices[i]
	//		- 操作二：第i天没有操作，那么 dp[i][1] = dp[i-1][1]
	// dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	// dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2])	 卖出了股票/没有操作
	// dp[i][3] = max(dp[i-1][2]-prices[i], dp[i-1][3])
	// dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])

	// dp初始化：
	// dp[0] = []int{0, -prices[i], 0, -prices[i], 0}


	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 5)
	}

	// 初始化
	dp[0] = []int{0,-prices[0],0,-prices[0],0}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])	// 第一次买入
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])	// 第一次卖出
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])	// 第二次买入
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])	// 第二次卖出
	}

	return dp[len(prices)-1][4]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

