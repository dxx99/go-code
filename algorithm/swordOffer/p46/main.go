package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(506))	 	// 1
	fmt.Println(translateNum(624))		// 2
	fmt.Println(translateNum(12258))	// 5
}

// 动态规划求解
func translateNum(num int) int {
	s := strconv.Itoa(num)
	if len(s) <= 1 {
		return len(s)
	}
	// 1. dp数组的定义
	dp := make([]int, len(s))

	// 2. 初始化
	dp[0] = 1
	dp[1] = 1
	if s[0] != '0' && s[0:2] < "26" {
		dp[1] = 2
	}

	// 3. 遍历
	for i := 2; i < len(s); i++ {
		// 4. 递推公式
		if s[i-1] != '0' && s[i-1:i+1] < "26" {
			dp[i] = dp[i-1]+dp[i-2]
		}else {
			dp[i] = dp[i-1]
		}
	}

	// 5. 输出调式
	//fmt.Println(dp)
	return dp[len(s)-1]
}
