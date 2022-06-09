package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(strToInt("-2147483649 with words"))
}

func strToInt(str string) int {
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}
	ans := 0
	flag := 1
	for i := 0; i < len(str); i++ {
		if i == 0 {
			if str[0] == '-' {
				flag = -1
				continue
			}
			if str[0] == '+' {
				flag = 1
				continue
			}
			if str[0] > '9' || str[0] < '0' {
				break
			}
		}
		if str[i] >= '0' && str[i] <= '9' {
			ans = 10*ans + int(str[i]-'0')
			if ans > math.MaxInt32 {
				if flag == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
		}else {
			break
		}
	}
	return ans*flag
}
