package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxConsecutiveV2(2,9,[]int{4,6}))
	fmt.Println(maxConsecutiveV2(6,8,[]int{7,6,8}))
}

func maxConsecutiveV2(bottom int, top int, special []int) int {
	max := 0
	sort.Ints(special)

	last := bottom
	for i := 0; i < len(special); i++ {
		tmp := special[i]
		cur := tmp - last
		if cur > max {
			max = cur
		}
		last = tmp + 1
	}

	// 处理最后一个元素
	cur := top - last + 1
	if cur > max {
		max = cur
	}

	return max
}


func maxConsecutive(bottom int, top int, special []int) int {
	max := 0

	hash := make(map[int]int)
	for i := 0; i < len(special); i++ {
		hash[special[i]]++
	}
	last := bottom
	for j := bottom; j <= top; j++ {
		if v, ok := hash[j]; ok && v > 0 {
			cur := j - last
			if cur > max {
				max = cur
			}
			last = j+1
		}
	}

	// 处理最后一个元素
	cur := top - last + 1
	if cur > max {
		max = cur
	}

	return max
}
