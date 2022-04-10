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
	fmt.Println(longestPalindrome("abcba"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ab"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println("V2.....................")
	fmt.Println(longestPalindromeV2(""))
	fmt.Println(longestPalindromeV2("abcba"))
	fmt.Println(longestPalindromeV2("a"))
	fmt.Println(longestPalindromeV2("ab"))
	fmt.Println(longestPalindromeV2("babad"))
}
// 中心扩散法
func longestPalindromeV2(s string) string {
	maxStr := ""
	for i := 0; i < len(s); i++ {
		p1Str :=  getCenterStr(s, i, i)
		p2Str :=  getCenterStr(s, i, i+1)
		if len(p1Str) > len(maxStr) {
			maxStr = p1Str
		}
		if len(p2Str) > len(maxStr) {
			maxStr = p2Str
		}
	}

	return maxStr
}

// 通过中心点得到的最大回文串
// 如果left = right 说明这个点
func getCenterStr(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1:right]
}




// 暴利求解【两层循环】
func longestPalindrome(s string) string {
	maxStr := ""
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if checkPalindrome(s[i:j]) {
				if len(s[i:j]) > len(maxStr) {
					maxStr = s[i:j]
				}
			}
		}
	}
	return maxStr
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
