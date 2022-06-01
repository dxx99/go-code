package main

import "fmt"

// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//
// 
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
//提示：
//
//s.length <= 40000
//注意：本题与主站 3 题相同：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//
//通过次数228,060提交次数491,791
//请问您在哪类招聘中遇到此题？
//
//社招
//
//校招
//
//实习
//
//未遇到
//《剑指 Offer（第 2 版）》官方授权
//
//相关企业
//相关标签
//
//题目列表
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))		//3
	//fmt.Println(lengthOfLongestSubstring("auv"))			//3
	//fmt.Println(lengthOfLongestSubstring(" "))			//1
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))	//3
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))		//1
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))		//3
}

// 滑动窗口，左右指针求解
func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	ans, left, right := 0, 0, 0
	for right < len(s){
		if v, ok := hash[s[right]]; ok {	//该元素上一次出现的位置，记录之后left应该加1移动
			if right > left && left <= v {
				if right - left > ans {
					ans = right - left
				}
				left = v+1
			}
		}
		hash[s[right]] = right
		right++
	}
	// 处理最后一个元素
	if right - left > ans {
		ans = right - left
	}
	return ans
}
