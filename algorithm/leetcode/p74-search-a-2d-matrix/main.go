package main

import (
	"fmt"
	"sort"
)

// p74 搜索二维矩阵
// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
// 
//
//示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
// 
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 100
//-104 <= matrix[i][j], target <= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-a-2d-matrix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//fmt.Println(searchMatrix([][]int{{1},{3}}, 3))
	//fmt.Println(searchMatrix([][]int{{1,3}}, 3))
	//fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3))
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l := m * n

	left, right := 0, l-1
	for left <= right {
		mid := (left+right) >> 1

		x, y := mid/n, mid%n

		if matrix[x][y] < target {
			left = mid + 1
		}else if matrix[x][y] > target {
			right = mid-1
		}else {
			return true
		}
	}

	return false
}

// 可以进行两次二分查找
func searchMatrixV2(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {return matrix[i][0] > target}) - 1
	if row < 0 {	// 没有找到
		return false
	}

	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
