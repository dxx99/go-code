package main

import "fmt"

// p343 整数拆分
// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
//
//返回 你可以获得的最大乘积 。
//
// 
//
//示例 1:
//
//输入: n = 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。
//示例 2:
//
//输入: n = 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
// 
//cchg&kunming4877
//提示:
//
//2 <= n <= 58
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/integer-break
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//fmt.Println(integerBreak(11))
	fmt.Println(integerBreakV2(2))
	fmt.Println(integerBreakV2(10))
	fmt.Println(integerBreakV2(11))
}

// 动态规划求解
func integerBreakV2(n int) int {
	dp := make([]int, n+1)

	// 推导公式
	// dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))

	//初始化
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 数学公式直接推导求解
func integerBreak(n int) int {
	if n < 4  {
		return n-1
	}
	sum := 1
	for n > 4 {
		n = n-3
		sum *= 3
	}
	sum *= n
	return sum
}


