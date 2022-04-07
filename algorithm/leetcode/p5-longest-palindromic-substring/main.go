package main

import "fmt"

// p5 最长回文字符串
// 给你一个字符串 s，找到 s 中最长的回文子串。
//
// 
//
//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"
// 
//
//提示：
//
//1 <= s.length <= 1000
//s 仅由数字和英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(checkPalindrome("abcba"))
}

func longestPalindrome(s string) string {
	return ""
}

// 检测字符串是否是回文字符串
func checkPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	for left <= right{
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
