package main

import "fmt"

//p944 腐烂的橘子
// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
//值 0 代表空单元格；
//值 1 代表新鲜橘子；
//值 2 代表腐烂的橘子。
//每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
//返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//
// 
//
//示例 1：
//
//
//
//输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
//输出：4
//示例 2：
//
//输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
//输出：-1
//解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
//示例 3：
//
//输入：grid = [[0,2]]
//输出：0
//解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
// 
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 10
//grid[i][j] 仅为 0、1 或 2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotting-oranges
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(orangesRotting([][]int{{2,1,1},{1,1,0},{0,1,1}}))
	fmt.Println(orangesRotting([][]int{{2,1,1},{0,1,1},{1,0,1}}))
	fmt.Println(orangesRotting([][]int{{1},{2}}))
}

func orangesRotting(grid [][]int) int {
	orangesOkNum := 0
	queue := make([][]int, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
			if grid[i][j] == 1 {
				orangesOkNum += 1
			}
		}
	}

	dx := []int{1,0,0,-1}
	dy := []int{0,1,-1,0}
	usedTime := 0
	for len(queue) > 0 && orangesOkNum > 0 {
		tmp := queue
		queue = nil
		for _, p := range tmp {
			for k := 0; k < 4; k++ {
				px, py := p[0] + dx[k], p[1] + dy[k]
				if px >=0 && px < m && py >=0 && py < n && grid[px][py] == 1 {
					grid[px][py] = 2
					orangesOkNum--
					queue = append(queue, []int{px, py})
				}
			}
		}
		usedTime += 1
	}

	if orangesOkNum > 0 {
		return -1
	}

	return usedTime
}
