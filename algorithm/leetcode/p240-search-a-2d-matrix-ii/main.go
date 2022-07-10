package main

func main() {

}

// 左上角的二分查找
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1	// 初始化右上角
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}

		if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

