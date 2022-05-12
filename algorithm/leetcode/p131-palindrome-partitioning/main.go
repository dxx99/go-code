package main

import "fmt"

// p131 分隔回文串
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
//回文串 是正着读和反着读都一样的字符串。
//
// 
//
//示例 1：
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//示例 2：
//
//输入：s = "a"
//输出：[["a"]]
// 
//
//提示：
//
//1 <= s.length <= 16
//s 仅由小写英文字母组成
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/palindrome-partitioning
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(partition("aaa"))
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	track := make([]string, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		// 终止条件
		if start >= len(s) {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		//遍历
		for i := start; i < len(s); i++ {
			tmp := s[start:i+1]
			if isPalindrome(tmp) {
				track = append(track, tmp)
			} else {
				continue
			}

			//递归
			backtracking(i+1)
			// 回溯, 不是回文字符串，就会回退元素
			track = track[:len(track)-1]
		}
	}

	backtracking(0)

	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		}else {
			return false
		}
	}
	return true
}