package main

import "fmt"

// p474 一和零
// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
//
//请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
//
//如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
//
// 
//
//示例 1：
//
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
//输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
//其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
//示例 2：
//
//输入：strs = ["10", "0", "1"], m = 1, n = 1
//输出：2
//解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
// 
//
//提示：
//
//1 <= strs.length <= 600
//1 <= strs[i].length <= 100
//strs[i] 仅由 '0' 和 '1' 组成
//1 <= m, n <= 100
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ones-and-zeroes
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5,3))
}

func findMaxForm(str []string, m int, n int) int {
	// m个零，n个一，对应的可能的结果
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化为零即可
	dp[0][0] = 0

	// 递推公式， k0表示当前字符串零的数组，k1表示当前字符串1的数量
	// dp[i][j] = max(dp[i][j], dp[i-num[k0]][j-num[k1]]+1)

	for i := 0; i < len(str); i++ {
		// 计算当前字符串01的数量
		zeroNum, oneNum := 0,0
		for j := 0; j < len(str[i]); j++ {
			if str[i][j] == '0' {
				zeroNum++
			}else {
				oneNum++
			}
		}

		// 对背包容量计算
		for x := m; x >= zeroNum ; x-- {
			for y := n; y >= oneNum; y-- {
				if dp[x-zeroNum][y-oneNum] + 1 > dp[x][y] {
					dp[x][y] = dp[x-zeroNum][y-oneNum] + 1
				}
			}
		}
	}

	return dp[m][n]
}


