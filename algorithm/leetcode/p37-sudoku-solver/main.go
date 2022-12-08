package main

func main() {

}

func solveSudoku(board [][]byte)  {
	var dfs func(board [][]byte, i int, j int) bool
	dfs = func(board [][]byte, i int, j int) bool {
		m, n := 9, 9
		// 穷举到最后一列的话就换到下一行重新开始。
		if j == n {
			return dfs(board, i+1, 0)
		}
		// 找到一个可行解
		if i == m {
			return true
		}

		// 有预设数字，不用穷举
		if board[i][j] != '.' {
			return dfs(board, i, j+1)
		}

		for ch := '1'; ch <= '9'; ch++ {
			if !isValid(board, i, j, byte(ch)) {	// 如果遇到不合法的数字，则跳过
				continue
			}

			board[i][j] = byte(ch)

			if dfs(board, i, j+1) {
				return true
			}

			// 回溯
			board[i][j] = '.'
		}
		// 穷举1-9没有解
		return false
	}

	dfs(board, 0, 0)
}

func isValid(board [][]byte, row int, col int, ch byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == ch || board[i][col] == ch {
			return false
		}

		// 判断3*3是否存在
		if board[(row/3)*3 + i/3][(col/3)*3+i%3] == ch {
			return false
		}
	}
	return true
}