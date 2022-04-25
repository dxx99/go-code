package main

import (
	"fmt"
	"math"
)

// p279 完全平方数
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
//完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
// 
//
//示例 1：
//
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//示例 2：
//
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
// 
//提示：
//
//1 <= n <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/perfect-squares
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {
	dp := make([]int, n+1)

	// 初始化
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	// 计算物品数量
	x := int(math.Sqrt(float64(n)))

	// 遍历
	for i := 1; i <= x; i++ {
		for j := i*i; j <= n; j++ {
			if dp[j-i*i] != math.MaxInt64 {
				// 递推公式
				if dp[j] > dp[j-i*i]+1 {
					dp[j] = dp[j-i*i]+1

				}
			}
		}
	}
	// 调式
	//fmt.Println(dp)
	if dp[n] == math.MaxInt64 {
		return 0
	}
	return dp[n]
}
