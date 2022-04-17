package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(digitSum("11111222223", 3))
	fmt.Println(digitSum("00000000", 3))
	fmt.Println(digitSum("11", 3))
}

func digitSum(s string, k int) string {
	l := len(s)
	for l > k {
		newStr := ""
		for i := 0; i < l; i = i + k {
			right := i+k
			if right > l {
				right = l
			}
			newStr += calByteArr(s[i:right])
		}
		s = newStr
		l = len(s)
	}
	return s
}

func calByteArr(s string) string {
	sum := 0
	for _, b := range s {
		sum += int( b - '0')
	}
	return strconv.Itoa(sum)
}
