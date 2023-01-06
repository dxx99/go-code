package main

import (
	"fmt"
)

func main() {
	fmt.Println(-7%2)
	fmt.Println(baseNeg2(2))
}

// 1017. 负二进制转换
// https://leetcode.cn/problems/convert-to-base-2/
// 思路：
// 	十进制转二进制 -->  除2取余，倒序
//  十进制转负二进制  -->  同理做除，只是如果是-1，当被除数为负数，余数为-1时，计算所得的商应该加一，余数就变为正1
func baseNeg2(n int) string {
	if n == 0 {
		return fmt.Sprintf("%d", n)
	}
	const Base = -2
	arr := make([]byte, 0)
	for n != 1 {
		if n > 0 {
			arr = append(arr, byte('0' + n%Base))
			n /= Base
		} else {
			yu := n % Base
			if yu == -1 {
				n = n/Base + 1
				arr = append(arr, '0'+byte(1))
			} else {
				arr = append(arr, '0'+byte(yu))
				n /= Base
			}
		}
	}

	arr = append(arr, '0'+byte(1))

	// 反转
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	return string(arr)
}
