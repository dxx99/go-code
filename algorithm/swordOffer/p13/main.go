package main

import "fmt"

func main() {
	fmt.Println(movingCount(2,3,1))
	fmt.Println(movingCount(3,1,0))
}

func movingCount(m int, n int, k int) int {
	ans := 0
	// 用来存放结果，-1 表示不能达到的点，1表示能达到的点，0表示没办法覆盖的点
	grids := make([][]int, m)
	for i := range grids {
		grids[i] = make([]int, n)
	}

	// 用来计算某个数的和
	bitSum := func(x int) int {
		sum := 0
		for x >= 10 {
			sum += x % 10
			x = x / 10
		}
		sum += x
		return sum
	}

	px, py := []int{1,0,0,-1}, []int{0,1,-1,0}
	var backtracking func(x int, y int)
	backtracking = func(x int, y int) {
		if bitSum(x) + bitSum(y) > k {
			grids[x][y] = -1
			return
		}

		grids[x][y] = 1
		ans++

		for i := 0; i < 4; i++ {
			cx, cy := x+px[i], y+py[i]
			if cx >= 0 && cx < len(grids) && cy >= 0 && cy < len(grids[0]) && grids[cx][cy] == 0 {
				backtracking(cx, cy)
			}
		}
		return
	}

	backtracking(0,0)
	return ans
}

