package main

import "fmt"

func main() {
	fmt.Println(fib(95))
}

func fib(n int) int {
	const mod = 1e9 + 7
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)

	// 初始化
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] % mod + dp[i-2] % mod
	}

	return dp[n] % mod
}
