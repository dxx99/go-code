package main

import "fmt"

// p695 岛屿的最大面积
// 给你一个大小为 m x n 的二进制矩阵 grid 。
//
//岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
//岛屿的面积是岛上值为 1 的单元格的数目。
//
//计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
//
// 
//
//示例 1：
//
//
//输入：grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//输出：6
//解释：答案不应该是 11 ，因为岛屿只能包含水平或垂直这四个方向上的 1 。
//示例 2：
//
//输入：grid = [[0,0,0,0,0,0,0,0]]
//输出：0
// 
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 50
//grid[i][j] 为 0 或 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/max-area-of-island
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	grid := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	fmt.Println(maxAreaOfIsland(grid))

}

type point struct {
	x int
	y int
	val int
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	queue := make([]point, 0)
	maxSum := 0

	//前后左右元素标定
	dx := []int{1, 0, 0, -1}
	dy := []int{0, 1, -1, 0}

	// 添加一个hashmap,标记已扫描的元素避免重复扫描
	hashMap := make(map[string]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {

				if checkMapExist(hashMap, i, j) {
					continue
				}

				queue = append(queue, point{i,j,grid[i][j]}) 	// 加入队列
				cSum := 0
				for len(queue) > 0 {
					// 出列
					p := queue[0]
					cSum += p.val
					if len(queue) == 1 {
						queue = []point{}
					}else {
						queue = queue[1:]
					}

					// 判断这个元素的前后左右元素是否等于1，如果等于就加入到队列
					for k := 0; k < 4; k++ {
						mx, ny := p.x + dx[k], p.y + dy[k]
						if mx >= 0 && mx < m && ny >=0 && ny < n && grid[mx][ny] == 1  {
							if checkMapExist(hashMap, mx, ny) {
								continue
							}
							queue = append(queue, point{mx,ny,grid[mx][ny]})
						}
					}
				}

				if cSum > maxSum {
					maxSum = cSum
				}
			}
		}
	}

	return maxSum
}

func checkMapExist(m map[string]bool, x int, y int) bool {
	key := fmt.Sprintf("%d-%d", x, y)
	if _, ok :=m[key]; ok {
		return true
	}
	m[key] = true
	return false
}
