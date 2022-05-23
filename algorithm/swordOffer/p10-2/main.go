package main

import "fmt"

func main() {
	fmt.Println(numWays(95))
}

// 跟爬楼梯问题同源
func numWays(n int) int {
	const mod = 1e9+7
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)

	// 初始化
	dp[0], dp[1] = 1, 1

	for i := 2; i <=n; i++ {
		dp[i] = (dp[i-1]  + dp[i-2]) % mod
	}

	return dp[n]
}
