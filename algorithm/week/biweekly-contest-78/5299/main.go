package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(divisorSubstrings(240, 2))
	fmt.Println(divisorSubstrings(430043, 2))
}

func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)

	total := 0
	for i := 0; i < len(s)-k+1; i++ {
		tmp, _ := strconv.Atoi(s[i:i+k])
		if tmp != 0 && num % tmp == 0 {
			total++
		}
	}
	return total
}
