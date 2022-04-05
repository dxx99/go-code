package main

import "fmt"

// p3 无重复字符的最大子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))	//2
	fmt.Println(lengthOfLongestSubstring("dvdf")) //3
	fmt.Println(lengthOfLongestSubstring("a"))		//1
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //3
	fmt.Println(lengthOfLongestSubstring("bbbbb")) 	//1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))	//3

}

func lengthOfLongestSubstring(s string) int {
	bs := []byte(s)
	maxLen := 0
	left, right := 0, 0	//窗口的前后指针, 用来表示滑动窗口的区间
	widMap := make(map[byte]int, 0)		//缓存所有的子串值，方便更快的找到子串的索引
	for k, b := range bs {
		if w, ok := widMap[b]; ok {
			if k > left && left <= w  {	// 这个条件比较关键，判断什么时候要移动窗口左边的指针。[hash表存在值，且当前元素大于left, left要小于等于hash表中存在的值]
				if right - left > maxLen {
					maxLen = right - left
				}
				// 移到到相同元素的前一位，除掉这个元素
				left = w+1
			}
		}
		widMap[b] = k
		right++
	}

	// 比较最后一个窗口
	if right - left > maxLen {
		maxLen = right - left
	}

	return maxLen
}


