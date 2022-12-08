package main

import (
	"fmt"
	"strings"
)

// p51 N皇后
// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
// 
//
//示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//示例 2：
//
//输入：n = 1
//输出：[["Q"]]
// 
//
//提示：
//
//1 <= n <= 9
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/n-queens
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	// 存储被其他皇后控制的区域
	columns, skim, driy := make(map[int]bool), make(map[int]bool), make(map[int]bool)

	// 存储每一层皇后存放的位置
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}

	var dfs func(n int, row int)
	dfs = func(n int, row int) {
		// 退出条件
		if row >= n {
			ans = append(ans, genRes(queens))
			return
		}

		for col := 0; col < n; col++ {
			// 不能放皇后的区域
			if columns[col] || skim[row+col] || driy[row-col] {
				continue
			}

			// 每一行皇后的位置
			queens[row] = col

			columns[col] = true
			skim[row+col] = true
			driy[row-col] = true

			// 进入下一层
			dfs(n, row+1)

			// 回溯位置
			queens[row] = -1
			delete(columns, col)
			delete(skim, row+col)
			delete(driy, row-col)
		}
	}

	dfs(n, 0)
	return ans
}

func genRes(queens []int) []string {
	ans := make([]string, 0)
	for _, row := range queens {
		ts := make([]string, len(queens))
		for j := 0; j < len(queens); j++ {
			ts[j] = "."
		}
		ts[row] = "Q"
		ans = append(ans, strings.Join(ts, ""))
	}
	return ans
}
