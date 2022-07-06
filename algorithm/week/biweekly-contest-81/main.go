package main

import "fmt"

func main() {
	fmt.Println(maximumXOR([]int{3,2,4,6}))
}

func countAsterisks(s string) int {
	xSum := 0
	sNum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			if sNum == 0 {
				sNum = 1
			}else if sNum == 1 {
				sNum = 0
			}
		}
		if sNum == 0 && s[i] == '*' {
			xSum++
		}
	}
	return xSum
}

func countPairs(n int, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	vis := make([]bool, n)
	tot, size := 0, 0

	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}

	fmt.Println(tot,size)


	return 0
}


func maximumXOR(nums []int) int {
	fmt.Println(6&(6^5))
	return 0
}
