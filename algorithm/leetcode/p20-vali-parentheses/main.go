package main

import "fmt"

// p20 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
// 
//
//示例 1：
//
//输入：s = "()"
//输出：true
//示例 2：
//
//输入：s = "()[]{}"
//输出：true
//示例 3：
//
//输入：s = "(]"
//输出：false
//示例 4：
//
//输入：s = "([)]"
//输出：false
//示例 5：
//
//输入：s = "{[]}"
//输出：true
// 
//
//提示：
//
//1 <= s.length <= 104
//s 仅由括号 '()[]{}' 组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(isValid("(){}}{"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[)]"))
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	byteMap := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}

	stack := make([]byte, 0)
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		matchKey, ok := byteMap[s[i]]
		if ok {
			if len(stack)>0 && stack[len(stack)-1] == matchKey {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func isValidV1(s string) bool {
	stack := make([]byte, 0)

	m := map[byte]byte{')':'(', ']':'[', '}':'{'}
	for _, b := range s {
		x, ok := m[byte(b)]
		if !ok {
			stack = append(stack, byte(b))
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != x {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
