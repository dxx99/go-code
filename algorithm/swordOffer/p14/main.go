package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(1))
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(3))
	fmt.Println(cuttingRope(4))
	fmt.Println(cuttingRope(5))
}

// 贪心算法
func cuttingRopeV2(n int) int {

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	ans := 1
	for n > 4 {
		ans *= 3
		n -= 3
	}
	ans *= n

	return ans
}

// 动态规划求解
func cuttingRopeV3(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1

	// 遍历，注意索引的位置
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {

			// dp[i] = dp[i-j]*dp[j] 将数据拆分成4份及以上，这个时候就要注意初始化

			if dp[i-j]*j > dp[i] {
				dp[i] = dp[i-j]*j	// 将数拆分成2份及以上
			}
			if (i-j) * j  > dp[i] {
				dp[i] = (i-j)*j		// 将数拆分成2份
			}
		}
	}

	return dp[n]
}

// 动态规划求解
func cuttingRope(n int) int {
	if n <= 3 {
		return 1*(n-1)
	}
	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 3	// 要注意初始化问题

	// 遍历，注意索引的位置
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// 将数据拆分成4份及以上，这个时候就要注意初始化
			if dp[i] < dp[i-j]*dp[j] {
				dp[i] = dp[i-j]*dp[j]
			}
		}
	}

	return dp[n]
}

