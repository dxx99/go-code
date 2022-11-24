package main

import "fmt"

// 856. 括号的分数
// 给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
//
// () 得 1 分。
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
// (A) 得 2 * A分，其中 A 是平衡括号字符串。
//
// 示例 1：
//
// 输入： "()"
// 输出： 1
// 示例 2：
//
// 输入： "(())"
// 输出： 2
// 示例 3：
//
// 输入： "()()"
// 输出： 2
// 示例 4：
//
// 输入： "(()(()))"
// 输出： 6
//
// 提示：
//
// S 是平衡括号字符串，且只含有 ( 和 ) 。
// 2 <= S.length <= 50
// 通过次数30,206提交次数45,53
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/score-of-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	//fmt.Println(scoreOfParentheses("()"))
	//fmt.Println(scoreOfParentheses("(())"))
	//fmt.Println(scoreOfParentheses("()()"))
	fmt.Println(scoreOfParentheses("(()(()))"))
}

func scoreOfParentheses(s string) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	stack := make([]int, 0)
	stack = append(stack, 0) // 空字符串的分数
	for _, t := range s {
		if t == '(' {
			stack = append(stack, 0)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*top, 1)
		}
	}
	return stack[0]
}
