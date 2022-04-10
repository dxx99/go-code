package main

import "fmt"

// p542 01矩阵
// 给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。
//
//两个相邻元素间的距离为 1 。
//
// 
//
//示例 1：
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[0,0,0]]
//输出：[[0,0,0],[0,1,0],[0,0,0]]
//示例 2：
//
//
//
//输入：mat = [[0,0,0],[0,1,0],[1,1,1]]
//输出：[[0,0,0],[0,1,0],[1,2,1]]
// 
//
//提示：
//
//m == mat.length
//n == mat[i].length
//1 <= m, n <= 104
//1 <= m * n <= 104
//mat[i][j] is either 0 or 1.
//mat 中至少有一个 0 
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/01-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(updateMatrix([][]int{{0,0,0},{0,1,0},{0,0,0}}))
	fmt.Println(updateMatrix([][]int{{0,0,0},{0,1,0},{1,1,1}}))
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}else {
				mat[i][j] = -1		// 没有标记的元素将下标改为-1,方便后面标值，加入队列
			}
		}
	}

	x := 1
	dx := []int{1,0,0,-1}
	dy := []int{0,1,-1,0}
	for len(queue) > 0 {	// BFS
		tmp := queue
		queue = nil	// 清空队列
		for _, p := range tmp {
			// 当前点的上下左右点
			for k := 0; k < 4; k++ {
				px, py := p[0] + dx[k], p[1] + dy[k]
				if px >=0 && px < m && py >= 0 && py<n && mat[px][py] == -1 {
					mat[px][py] = x
					queue = append(queue, []int{px, py})
				}
			}
		}
		x++
	}

	return mat
}