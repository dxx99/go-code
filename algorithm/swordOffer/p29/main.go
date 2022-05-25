package main

func main() {

}

// todo: 需要
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	if len(matrix) == 0 {
		return ans
	}

	m, n := len(matrix), len(matrix[0])
	x, y := 0, 0

	sum := m * n
	for sum > 0 {
		if y == n {
			n--
		}
		if x == m {
			m--
		}

		sum--
	}

	return ans
}
