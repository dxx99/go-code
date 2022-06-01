package main

import (
	"fmt"
	"sort"
)

// 输入一个字符串，打印出该字符串中字符的所有排列。
//
// 
//
//你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
// 
//
//示例:
//
//输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
// 
//
//限制：
//
//1 <= s 的长度 <= 8
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(permutation("dcz"))
	fmt.Println(permutationV2("dcz"))
}

func permutation(s string) []string {
	ans := make([]string, 0)
	track := make([]byte, 0)
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	hashByte := [255]byte{}
	for i := 0; i < len(s); i++ {
		hashByte[s[i]]++
	}

	var backtracking func(m [255]byte)
	backtracking = func(m [255]byte) {
		if len(track) == len(s) {
			ans = append(ans, string(track))
			return
		}
		for i, b := range bs {
			// 跳过同层相同的元素
			if i > 0 && bs[i] == bs[i-1] {
				continue
			}

			if m[b] > 0 {
				track = append(track, b)
				m[b]--
				backtracking( m)

				// 回溯
				m[b]++
				track = track[:len(track)-1]
			}
		}
	}

	backtracking(hashByte)
	return ans
}


func permutationV2(s string) []string {
	ans := make([]string, 0)
	track := make([]byte, 0)

	hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	var backtracking func(m map[byte]int)
	backtracking = func(m map[byte]int) {
		if len(track) == len(s) {
			ans = append(ans, string(track))
			return
		}
		used := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			// 排除到水平重复的元素
			if used[s[i]] == 1 {
				continue
			}

			if v, ok := m[s[i]]; ok && v > 0 {
				track = append(track, s[i])
				used[s[i]]++
				m[s[i]]--
				backtracking( m)

				// 回溯
				m[s[i]]++
				track = track[:len(track)-1]
			}
		}
	}

	backtracking(hash)
	return ans
}
