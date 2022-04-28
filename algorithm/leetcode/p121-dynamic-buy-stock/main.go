package main

import "fmt"

// 121. 买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//
// 
//
//示例 1：
//
//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//示例 2：
//
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
// 
//
//提示：
//
//1 <= prices.length <= 105
//0 <= prices[i] <= 104
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{2,1,2,1,0,1,2}))   // 2
	fmt.Println("V2...............")
	fmt.Println(maxProfitV2([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfitV2([]int{7,6,4,3,1}))
	fmt.Println(maxProfitV2([]int{2,1,2,1,0,1,2}))   // 2
	fmt.Println("V4...............")
	fmt.Println(maxProfitV4([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfitV4([]int{7,6,4,3,1}))
	fmt.Println(maxProfitV4([]int{2,1,2,1,0,1,2}))   // 2

}

// 双指针求解
func maxProfitV4(prices []int) int {
	max := 0
	if len(prices) == 0 {
		return max
	}

	left := 0
	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
			continue
		}

		 cur := prices[right] - prices[left]
		if cur > max {
			max = cur
		}
	}
	return max
}

// 动规五部曲求解【求的累积最大值】
func maxProfitV3(prices []int) int {
	dp := make([]int, len(prices))

	// 递推公式 dp[i] = dp[i-1] + profit(prices[i], prices[i-1])

	// 初始化, 当天买卖没有收益
	dp[0] = 0

	//遍历
	for i := 1; i < len(prices); i++ {
		dp[i] = dp[i-1] + profit(prices[i], prices[i-1])
	}
	return dp[len(prices)-1]
}

func profit(x, y int) int {
	if x > y {
		return x-y
	}
	return 0
}


// 暴利解法，双循环
func maxProfit(prices []int) int {
	profit := 0
	l := len(prices)
	for i := 0; i < l; i++ {
		for j := i+1; j < l; j++ {
			if prices[j] - prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

// 动态规划
// 类似于前面提到的子数组最大的和
// 存储： 今天之前买入最小值
// 比较： 今天卖出的最大收益，也是之前的最小值与今天值的比较，如果之前大于今天值，则替换存储，将今天的收益变为0
// 计算:  从所有存储的每天收益中获取最大的收益
func maxProfitV2(prices []int) int {
	profit := 0
	l := len(prices)
	if l == 0 {
		return profit
	}
	min := prices[0]
	for i := 1; i < l; i++ {
		// 是否替换最小值
		if prices[i] < min {
			min = prices[i]
		}

		// 比较当天最大收益与历史最大收益
		if prices[i] - min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}
