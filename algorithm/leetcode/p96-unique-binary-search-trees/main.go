package main

import "fmt"

// p96 不同的二叉搜索树
// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
// 
//
//示例 1：
//
//
//输入：n = 3
//输出：5
//示例 2：
//
//输入：n = 1
//输出：1
// 
//
//提示：
//
//1 <= n <= 19
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-binary-search-trees
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
}

func numTrees(n int) int {
	// dp数组定义
	dp := make([]int, n+1)

	// 公示推导
	// dp[n] = dp[n-1]*dp[0]  + .... + dp[n-2]*dp[n-2] + dp[0]*dp[n-1]
	// dp[i] += dp[j-1]*dp[i-j]	//j-1为j头节点的左子树数量，i-j为j头节点右子树数量

	// dp数组初始化
	dp[0] = 1

	// 遍历
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {	// j要从1开始
			dp[i] += dp[j-1]*dp[i-j]
		}
	}

	return dp[n]
}
