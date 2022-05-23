package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	px, py := []int{1,0,0,-1}, []int{0,1,-1,0}

	// 用来去重，因为同一个元素不能选择两次
	used := make([][]bool, len(board))
	for i := range board {
		used[i] = make([]bool, len(board[0]))
	}

	// 从矩阵的（x,y）点开始，对应的字符word的下标为k
	var backtracking func(x int, y int, k int) bool
	backtracking = func(x int, y int, k int) bool {
		// 异常终止条件
		if board[x][y] != word[k] {
			return false
		}

		// 所有的元素都找完了，并且都相等
		if k == len(word)-1 {
			return true
		}

		// 还需要对选中的元素去重
		used[x][y] = true

		for i := 0; i < 4; i++ {
			cx, cy := x + px[i], y + py[i]
			if cx >= 0 && cx < len(board) && cy >= 0 && cy < len(board[0]) && !used[cx][cy] {
				if backtracking(cx, cy, k+1) {
					return true
				}
			}
		}

		// 回溯
		used[x][y] = false

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backtracking(i, j, 0) {
				return true
			}
		}
	}

	return false
}


