package main

import "fmt"

// p62 不同路径
// 62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//
//
//
//示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//示例 2：
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//示例 3：
//
//输入：m = 7, n = 3
//输出：28
//示例 4：
//
//输入：m = 3, n = 3
//输出：6
//
//
//提示：
//
//1 <= m, n <= 100
//题目数据保证答案小于等于 2 * 109
func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 3))
}

func uniquePaths(m int, n int) int {
	//dp数组定义
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// dp数组初始化
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {	// dp数组初始化要注意的问题
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]	// 递推公式： 只有可能从上面，或者左边过来
			}
		}
	}

	return dp[m-1][n-1]
}
