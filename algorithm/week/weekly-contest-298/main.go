package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumNumbers(58, 9))
}

// 1
func greatestLetter(s string) string {
	var max byte
	h := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		h[s[i]] = true
		if  s[i] >= 'A' && s[i] <= 'Z' && h[s[i] + 32] {
			if s[i] > max {
				max = s[i]
			}
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' && h[s[i] - 32] {
			if s[i] - 32 > max {
				max = s[i] - 32
			}
			continue
		}
	}
	if max > 0 {
		return string(max)
	}
	return ""
}

// 2
// 核心：减去多个k之后，然后被10整除
// 枚举：n=11的结果与n=1的结果相同，所以枚举到11就可以了
func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}

	for n := 1; n <= 10 && num - n*k > 0; n++ {
		if (num - n * k) % 10 == 0 {
			return n
		}
	}
	return -1
}

// 3
func longestSubsequence(s string, k int) int {
	return 0
}

// 4


