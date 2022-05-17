package main

import "fmt"

func main() {
	fmt.Println(maximumWhiteTiles([][]int{{1,5},{10,11},{12,18},{20,25},{30,32}}, 10))
	fmt.Println(maximumWhiteTiles([][]int{{10,11},{1,1}}, 2))
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	hash := make(map[int]int)	// 点，长度
	for i := 0; i < len(tiles); i++ {
		left, right := tiles[i][0], tiles[i][1]
		hash[left] = right - left + 1
	}

	res := 0
	// 从每一个点开始
	for j := 0; j < len(tiles); j++ {
		l1, r1 := tiles[j][0], tiles[j][1]
		tmp := hash[l1]	// 初始长度之后
		for i := r1; i <= l1+carpetLen; {
			if v, ok := hash[i]; ok {
				if v <= l1+carpetLen - i {
					tmp += v
				} else {
					tmp += l1+carpetLen - i
				}
				i = i+v
			}else {
				i++
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
