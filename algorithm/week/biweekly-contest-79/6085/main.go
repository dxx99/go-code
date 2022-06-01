package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumImportance(5, [][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}}))
}

func maximumImportance(n int, roads [][]int) int64 {
	hash := make(map[int]int)
	for i := 0; i < len(roads); i++ {
		for j := 0; j < len(roads[i]); j++ {
			hash[roads[i][j]]++
		}
	}
	fmt.Println(hash)
	arr := make([]int, 0)
	for _, item := range hash {
		arr = append(arr, item)
	}
	sort.Ints(arr)
	fmt.Println(arr)

	var ans int64
	for i := len(arr)-1; i >= 0; i-- {
		ans += int64(arr[i]) * int64(n)

		n--
	}
	return ans
}
