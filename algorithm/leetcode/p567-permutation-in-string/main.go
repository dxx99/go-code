package main

import (
	"fmt"
)

// p567 字符串排列
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
//换句话说，s1 的排列之一是 s2 的 子串 。
//
// 
//
//示例 1：
//
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
//示例 2：
//
//输入：s1= "ab" s2 = "eidboaoo"
//输出：false
// 
//
//提示：
//
//1 <= s1.length, s2.length <= 104
//s1 和 s2 仅包含小写字母
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutation-in-string
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(checkInclusion("rvwrk", "lznomzggwrvrkxecjaq"))		// true
	fmt.Println(checkInclusion("islander", "islander"))		// true
	fmt.Println(checkInclusion("ab", "ba"))					// true
	fmt.Println(checkInclusion("ab", "eidbaooo"))			// true
	fmt.Println(checkInclusion("ab", "eidboaoo"))			// false
	fmt.Println(checkInclusion("abc", "bbbca"))				// true
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))	// false
	fmt.Println(checkInclusion("ab", "eidbaooo"))			// true
	fmt.Println(checkInclusion("ab", "abababab"))			// true
	fmt.Println(checkInclusion("ab", "aaabaaab")) 			//true
}

// 【滑动窗口】
// 输入的字符串只包含小写字母
// 两个字符串的长度都在[1,10000]之间
// 出现的字符种类相同
// 各个字符出现的次数也相同
func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}

	//字符数组，指定一个长度为l1的窗口来进行滑动比较
	bs1, bs2 := [26]int{}, [26]int{}
	for i, ch := range s1 {
		bs1[ch - 'a']++
		bs2[s2[i] - 'a']++
	}

	//数组比较，切片没法比较，数组可以
	if bs1 == bs2 {
		return true
	}

	//win滑动,
	for i := l1; i < l2; i++ {
		bs2[s2[i] - 'a']++
		bs2[s2[i-l1] - 'a']--
		if bs1 == bs2 {
			return true
		}
	}
	return false
}

func checkInclusionV2(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}

	//
	cnt1 := [26]int{}
	for _, ch := range s1 {
		cnt1[ch - 'a']--
	}

	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt1[x]++
		for cnt1[x] > 0{
			cnt1[s2[left]-'a']--
			left++
		}

		if right - left + 1 == l1 {
			return true
		}
	}
	return false
}



