package main

import "fmt"

// p733 图像渲染
// 有一幅以 m x n 的二维整数数组表示的图画 image ，其中 image[i][j] 表示该图画的像素值大小。
//
//你也被给予三个整数 sr ,  sc 和 newColor 。你应该从像素 image[sr][sc] 开始对图像进行 上色填充 。
//
//为了完成 上色工作 ，从初始像素开始，记录初始坐标的 上下左右四个方向上 像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应 四个方向上 像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为 newColor 。
//
//最后返回 经过上色渲染后的图像 。
//
// 
//
//示例 1:
//
//
//
//输入: image = [[1,1,1],[1,1,0],[1,0,1]]，sr = 1, sc = 1, newColor = 2
//输出: [[2,2,2],[2,2,0],[2,0,1]]
//解析: 在图像的正中间，(坐标(sr,sc)=(1,1)),在路径上所有符合条件的像素点的颜色都被更改成2。
//注意，右下角的像素没有更改为2，因为它不是在上下左右四个方向上与初始点相连的像素点。
//示例 2:
//
//输入: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
//输出: [[2,2,2],[2,2,2]]
// 
//
//提示:
//
//m == image.length
//n == image[i].length
//1 <= m, n <= 50
//0 <= image[i][j], newColor < 216
//0 <= sr < m
//0 <= sc < n
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/flood-fill
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(floodFill([][]int{{0,0,0},{0,1,1}}, 1 ,1, 1))
	fmt.Println(floodFill([][]int{{0,0,0},{0,0,0}}, 0 ,0, 2))
	fmt.Println(floodFill([][]int{{1,1,1},{1,1,0},{1,0,1}}, 1 ,1, 2))
}


func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// 给定一个hashmap排除重复的元素
	currColor := image[sr][sc]

	// 如果更新的值与初始值相同就会重复更新，应该直接退出
	if currColor == newColor {
		return image
	}

	m, n := len(image), len(image[0])
	dx, dy := []int{1,0,0,-1}, []int{0,1,-1,0}

	// 给一个队列
	queue := make([]point, 0)
	queue = append(queue, point{sr, sc})

	//
	for i := 0; i < len(queue); i++ {
		p := queue[i]

		// 替换值
		image[p.sr][p.sc] = newColor

		// 上下左右处理
		for k := 0; k < 4; k++ {
			mx, ny := p.sr + dx[k], p.sc + dy[k]
			if mx >= 0 && mx < m && ny >= 0 && ny < n && image[mx][ny] == currColor {
				queue = append(queue, point{mx, ny})
			}
		}
	}
	return image
}

type point struct {
	sr int
	sc int
}

