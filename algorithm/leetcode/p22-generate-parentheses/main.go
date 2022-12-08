package main

import "fmt"

// p22 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 
//
//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//示例 2：
//
//输入：n = 1
//输出：["()"]
// 
//
//提示：
//
//1 <= n <= 8
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/generate-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	fmt.Println(generateParenthesis(4))
}

// n=1  1
// n=2  2
// n=3  5
// n=4  14
// 如何解决回溯算法：
// 		1. 选择  	--->  每次最多就两种选择，左括号、右括号，用dfs遍历这颗树，找出所有的解，这个过程叫回溯
//		2. 约束条件	--->  什么情况选择左括号，什么情况选择右括号 (1. 只要左有剩余，就可以选择，2. 当右大于左的数量时候，则可以选择右括号)
//		3. 目标		--->  构建出一个用尽n对括号合法的括号串，当长度为2*n的时候可以终止递归条件
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var dfs func(lNum int, rNum int, path string)
	dfs = func(lNum int, rNum int, path string) {
		// exit条件
		if lNum == 0 && rNum == 0 {
			res = append(res, path)
			return
		}
		
		// 选择左括号
		if lNum > 0 {
			dfs(lNum-1, rNum, path + "(")
		}
		
		// 选择右括号
		if rNum > lNum {
			dfs(lNum, rNum-1, path + ")")
		}
	}

	dfs(n,n, "")
	return res
}
