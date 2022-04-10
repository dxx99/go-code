package main

import (
	"fmt"
	"strconv"
)

// p9 回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//例如，121 是回文，而 123 不是。
// 
//
//示例 1：
//
//输入：x = 121
//输出：true
//示例 2：
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//示例 3：
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
// 
//
//提示：
//
//-231 <= x <= 231 - 1
// 
//
//进阶：你能不将整数转为字符串来解决这个问题吗？
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindrome-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(isPalindrome(121))	// true
	fmt.Println(isPalindrome(123))	// false
	fmt.Println(isPalindrome(-121))	// false
	fmt.Println(isPalindrome(112211))	// true
	fmt.Println(isPalindrome(0))	// true

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nArr := make([]int, 0)
	for x > 0 {
		nArr = append(nArr, x % 10)
		x = x / 10
	}

	left, right := 0, len(nArr) - 1
	for left < right {
		if nArr[left] != nArr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 反转一般的数字
func isPalindromeV3(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	// 关键点
	revertedNum := 0
	for x > revertedNum {
		revertedNum = revertedNum * 10 + x % 10
		x = x / 10
	}

	return x == revertedNum || x == revertedNum/10
}
