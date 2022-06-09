package main

func main() {

}

func maxProfit(prices []int) int {
	dp := make([][]int, 0)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp[0][0], dp[0][1] = -prices[0], 0	// 持有股票的金额，不持有股票的金额

	for i := 0; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])	// 上一天持有，或者当天买入
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])	// 上一天卖出，或者当天卖出的最高金额
	}

	return dp[len(prices)-1][1]
}
