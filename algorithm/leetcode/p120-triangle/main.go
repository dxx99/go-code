package main

import "math"

// p120 三角形最小路径和
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
//每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
// 
//
//示例 1：
//
//输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//输出：11
//解释：如下面简图所示：
//   2
//  3 4
// 6 5 7
//4 1 8 3
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//示例 2：
//
//输入：triangle = [[-10]]
//输出：-10
// 
//
//提示：
//
//1 <= triangle.length <= 200
//triangle[0].length == 1
//triangle[i].length == triangle[i - 1].length + 1
//-104 <= triangle[i][j] <= 104
// 
//
//进阶：
//
//你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/triangle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

// 递归 + 记忆 求解
func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle))
	}

	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == len(triangle) {
			return 0
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		dp[i][j] = min(dfs(i+1,j), dfs(i+1, j+1)) + triangle[i][j]
		return dp[i][j]
	}
	return dfs(0, 0)
}

func minimumTotalV2(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle))
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	// 初始化
	dp[len(triangle)-1] = triangle[len(triangle)-1]

	// 从下往上推
	for i := len(triangle)-2; i >= 0 ; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

// f[i][j] = min(f[i-1][j], f[i-1][j-1]) + c[i][j]
func minimumTotalV3(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	// 特殊值
	f[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		// 第一排的值，要特别注意处理
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {

			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
