package main

import (
	"fmt"
)

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//
// 
//
//示例:
//
//现有矩阵 matrix 如下：
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//给定 target = 5，返回 true。
//
//给定 target = 20，返回 false。
//
// 
//
//限制：
//
//0 <= n <= 1000
//
//0 <= m <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(findNumberIn2DArray([][]int{{1,   4,  7, 11, 15},{2,   5,  8, 12, 19},{3,   6,  9, 16, 22},{10, 13, 14, 17, 24},{18, 21, 23, 26, 30}}, 5))
	fmt.Println(findNumberIn2DArray([][]int{}, 0))
	fmt.Println(findNumberIn2DArray([][]int{{-5}}, -5))
	fmt.Println(findNumberIn2DArray([][]int{{}}, 1))
	fmt.Println(findNumberIn2DArray([][]int{{-5}}, -10))
	fmt.Println(findNumberIn2DArray([][]int{{1,1,1},{2,3,4}}, 2))
}

// 优化代码的二分查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 右上角对角线的二分查找
	px, py := 0, len(matrix)-1
	for px < len(matrix[0]) && py >= 0 {
		if matrix[px][py] == target {
			return true
		} else if matrix[px][py] > target {
			py--
		}else {
			px++
		}
	}

	return false
}

// 以二维数组的右上角的定点开始查找
func findNumberIn2DArrayV2(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	px, py := 0, n-1
	for px < m && py >= 0 {
		mid := matrix[px][py]
		if mid < target {	// 查找竖行
			left, right := px+1, m-1
			for left <= right {
				pm := (left+right)>>1
				if matrix[pm][py] > target {
					right = pm-1
				}else if matrix[pm][py] < target{
					left = pm+1
				}else {
					return true
				}
			}
		} else if mid > target { //查找横向
			left, right := 0, py-1
			for left <= right {
				pm := (left+right)>>1
				if matrix[px][pm] > target {
					right = pm-1
				}else if matrix[px][pm] < target{
					left = pm+1
				}else {
					return true
				}
			}
		} else {
			return true
		}

		px++
		py--
	}
	return false
}


