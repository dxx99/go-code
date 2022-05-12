package main

import "fmt"

// p417-pacific-atlantic-water-flow 太平洋大西洋水流问题
// 有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
//
//这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上单元格 高于海平面的高度 。
//
//岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
//
//返回网格坐标 result 的 2D 列表 ，其中 result[i] = [ri, ci] 表示雨水从单元格 (ri, ci) 流动 既可流向太平洋也可流向大西洋 。
//
// 
//
//示例 1：
//
//
//
//输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
//输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//示例 2：
//
//输入: heights = [[2,1],[1,2]]
//输出: [[0,0],[0,1],[1,0],[1,1]]
// 
//
//提示：
//
//m == heights.length
//n == heights[r].length
//1 <= m, n <= 200
//0 <= heights[r][c] <= 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/pacific-atlantic-water-flow
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(pacificAtlantic([][]int{{2,1},{1,2}}))
}

func pacificAtlantic(heights [][]int) [][]int {
	// 上下左右偏移量
	xArr := []int{1,0,0,-1}
	yArr := []int{0,1,-1,0}

	m, n := len(heights), len(heights[0])

	pacific := make([][]bool, m)	// 存储太平洋
	atlantic := make([][]bool, m)	// 存储大西洋
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// 定义dfs函数
	var dfs func(int, int, [][]bool)
	dfs = func(x int, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}

		ocean[x][y] = true

		// 处理上下左右
		for i := 0; i < 4; i++ {
			if nx, ny := x + xArr[i], y + yArr[i]; nx >= 0 && nx < m && ny>=0 && ny < n && heights[nx][ny] >= heights[x][y]  {
				dfs(nx, ny, ocean)
			}
		}
	}


	// 处理上下两边
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}

	// 处理左右两边
	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
		dfs(m-1, j, atlantic)
	}

	// 返回结果
	ans := make([][]int, 0)
	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {	// 北冰洋与大西洋都为真才为真
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
