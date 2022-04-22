package main

import "fmt"

// p32 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0
// 
//
//提示：
//
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("()(())"))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(")()()))()()()"))
	fmt.Println(longestValidParentheses(""))
}

func longestValidParentheses(s string) int {
	stack := make([]int, 0)	// 用来存储括号索引，这样就可以计算他们的最大长度了
	maxLen := 0

	// 添加一个-1， 方便第一个出栈结果
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}else {
				if i - stack[len(stack)-1] > maxLen {
					maxLen = i - stack[len(stack)-1]
				}
			}
		}
	}

	return maxLen
}
