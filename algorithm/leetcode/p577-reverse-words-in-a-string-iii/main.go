package main

import "fmt"

// p577 反转字符串中的单词III
// 给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
// 
//
//示例 1：
//
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"
//示例 2:
//
//输入： s = "God Ding"
//输出："doG gniD"
// 
//
//提示：
//
//1 <= s.length <= 5 * 104
//s 包含可打印的 ASCII 字符。
//s 不包含任何开头或结尾空格。
//s 里 至少 有一个词。
//s 中的所有单词都用一个空格隔开。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("God Ding"))
	fmt.Println(reverseWords("hehhhhhhe"))
}

func reverseWords(s string) string {
	sArr := []byte(s)
	left, right := 0, 0
	for i := 0; i < len(sArr); i++ {
		if sArr[i] == ' ' {
			if right != 0 {
				right++
			}
			left, right = right, i
			reverseString(sArr[left:right])

		}
	}
	// 处理最后一个元素的反转
	if right != 0 {
		right++
	}
	reverseString(sArr[right:])
	return string(sArr)
}

func reverseString(s []byte)  {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
