package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	fmt.Println(fb(n))
}

func fb(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+2)
	dp[1], dp[2] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i+1] = dp[i] + dp[i-1]
	}
	return dp[n]
}
