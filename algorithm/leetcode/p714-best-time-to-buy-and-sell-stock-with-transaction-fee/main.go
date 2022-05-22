package main

import "fmt"

//p714 买卖股票的最佳时机含手续费
// 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
//
//你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//
//返回获得利润的最大值。
//
//注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
//
// 
//
//示例 1：
//
//输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出：8
//解释：能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
//示例 2：
//
//输入：prices = [1,3,7,5,10,3], fee = 3
//输出：6
// 
//
//提示：
//
//1 <= prices.length <= 5 * 104
//1 <= prices[i] < 5 * 104
//0 <= fee < 5 * 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxProfit([]int{9,8,7,1,2}, 2))
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1,3,7,5,10,3}, 3))

	fmt.Println("V2-------------")
	fmt.Println(maxProfitV2([]int{9,8,7,1,2}, 2))
	fmt.Println(maxProfitV2([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfitV2([]int{1,3,7,5,10,3}, 3))
}

// 贪心算法
// 情况一：收获利润的这一天并不是收获利润区间里的最后一天（不是真正的卖出，相当于持有股票），所以后面要继续收获利润。
// 情况二：前一天是收获利润区间里的最后一天（相当于真正的卖出了），今天要重新记录最小价格了。
// 情况三：不作操作，保持原有状态（买入，卖出，不买不卖）
func maxProfitV2(prices []int, fee int) int {
	ans := 0

	// 先缩减数组大小，只记录拐点
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {	// 情况二 这个相当于买入操作
			minPrice = prices[i]
		}

		// 情况三, 保持原操作，不作交易
		if prices[i] >= minPrice && prices[i] <= minPrice + fee {
			continue
		}

		// 情况一： 可能多次计算利润，但是最后一次才是真正的利润
		if prices[i] > minPrice + fee {
			ans += prices[i] - minPrice - fee
			minPrice = prices[i] - fee	// 神来之笔
		}
	}
	return ans
}

func maxProfit(prices []int, fee int) int {
	// dp数组定义
	// dp[i][0] 表示第i天，持有股票所得的金额
	// dp[i][1] 表示第i天，不持有股票所得的金额
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	// 遍历
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i] - fee)
	}

	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
