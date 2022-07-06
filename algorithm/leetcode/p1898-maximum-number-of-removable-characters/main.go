package main

import "fmt"

func main() {
	fmt.Println(maximumRemovals("abcbddddd", "abcd", []int{3,2,1,4,5,6}))
}


func maximumRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)
	for left < right {
		mid := int(uint(left+right)>>1)

		// 记录删除元素的数据
		del := make([]bool, len(s))
		for i := 0; i <= mid; i++ {
			del[removable[i]] = true
		}

		// 通过子序列选择区间
		k := 0
		for i := 0; i < len(s); i++ {
			if k < len(p) && !del[i] && s[i] == p[k] {
				k++
			}
		}

		if k >= len(p) {
			left = mid+1
		}else {
			right = mid
		}
	}
	return left
}


// 判断子序列
func isSubsequence(s string, t string) bool {
	k := 0
	for i := 0; i < len(s); i++ {
		if k < len(t) && s[i] == t[k] {
			k++
		}
	}
	return k >= len(t)
}